package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"AiChatPartner/api/websocket/internal/svc"
	"AiChatPartner/rpc/db/db"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type Session struct {
	WsConn    *websocket.Conn
	Message   chan string
	closeOnce sync.Once
	ID        UserID
}

type UserID uint64

type ConnectionManager struct {
	connections map[UserID]*Session
	mutex       sync.RWMutex
}

type IConnectionManager interface {
	Add(*Session)
	Remove(UserID)
	// RemoveWithCode(uint32, int, string)
	Get(UserID) (*Session, bool)
	SendMessage(UserID) error
	ReadMessage(UserID) ([]byte, error)
}

var (
	wg       sync.WaitGroup
	ic       IConnectionManager = NewConnectionManager()
	closeMsg chan struct{}      // 关闭信号
)

func upgrade(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext) (*websocket.Conn, error) {
	ws := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewConnectionManager() IConnectionManager {
	return &ConnectionManager{
		connections: make(map[UserID]*Session),
		mutex:       sync.RWMutex{},
	}
}

func (cm *ConnectionManager) SendMessage(userId UserID) error {
	for {
		select {
		case message, ok := <-cm.connections[userId].Message:
			if !ok {
				return nil
			}
			for uin, s := range cm.connections {
				if uin == userId {
					continue
				}
				err := cm.connections[s.ID].WsConn.WriteMessage(1, []byte(message))
				if err != nil {
					logx.Error("[SendMessage] write message error. ", err)
					// cm.Remove(s.ID)
					return err
				}
			}
		}
	}
}

func (cm *ConnectionManager) readMessage(userID UserID) ([]byte, error) {
	_, msg, err := cm.connections[UserID(userID)].WsConn.ReadMessage()
	if err != nil {
		logx.Error("[RecvMessage] read message error. ", err)
		return msg, err
	}

	logx.Infof("[RecvMessage] User %d received message: %s", userID, string(msg))
	cm.connections[UserID(userID)].Message <- string(msg)
	return msg, nil
}

func (cm *ConnectionManager) ReadMessage(userID UserID) ([]byte, error) {
	for {
		msg, err := cm.readMessage(userID)
		if err != nil {
			return msg, err
		}

		//TODO: 业务逻辑处理

	}
}

func (n *Session) Close() {
	n.closeOnce.Do(func() {
		if n.WsConn != nil {
			n.WsConn.Close()
		}
		close(n.Message)
	})
}

func (cm *ConnectionManager) Add(s *Session) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.connections[s.ID] = s
}

func (cm *ConnectionManager) Remove(userID UserID) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	if s, ok := cm.connections[UserID(userID)]; ok {
		s.WsConn.Close()
		delete(cm.connections, UserID(userID))
	}
}

func (cm *ConnectionManager) Get(userID UserID) (*Session, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	s, ok := cm.connections[UserID(userID)]
	return s, ok
}

func parseJwtToken(tokenString, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是正确的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}

func (s *Server) WebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("[WebsocketHandler] start ...")

		// 获取uid
		username := r.URL.Query().Get("username")
		dbrsp, err := s.svc.DbServer.Read(context.Background(), &db.ReadRequest{
			TableName: "ac_user",
			Key:       username,
			KeyType:   2,
		})
		if err != nil {
			logx.Error("[WebsocketHandler] get user by username:[%s] error: %s", username, err)
			return
		}
		sUin := dbrsp.Data["uin"]
		uin, err := strconv.Atoi(sUin)
		if uin == -1 {
			logx.Error("[WebsocketHandler] get uin error. username:", username, " err:", err)
			return
		}
		userId := UserID(uin)
		// userIdStr := strconv.FormatUint(uint64(userId), 10)

		// TODO: 检查token
		// token, err := redis.GetRedisClient().Hget(userIdStr, "token")
		// if err != nil {
		// 	logx.Errorf("[WebsocketHandler] redis get token error. key:[%s] err:[%s]", userIdStr, err)
		// 	return
		// }
		// _, err = parseJwtToken(token, s.svc.Config.Auth.AccessSecret)
		// if err != nil {
		// 	logx.Errorf("[WebsocketHandler] parse token:%s error:%s ", token, err)
		// 	return
		// }

		// 升级为websocket
		conn, err := upgrade(w, r, svcCtx)
		if err != nil {
			logx.Error("[WebsocketHandler] upgrade error. ", err)
			return
		}
		node := Session{
			WsConn:  conn,
			Message: make(chan string),
			ID:      userId,
		}
		defer node.Close()

		ic.Add(&node)
		defer ic.Remove(userId)
		logx.Infof("[WebsocketHandler] User %d connected", userId)

		// wg.Add(2)

		go func() {
			err := ic.SendMessage(userId)
			if err != nil {
				logx.Error("[WebsocketHandler] send message error. ", err)
				ic.Remove(userId)
				closeMsg <- struct{}{}
				return
			}
		}()

		go func() {
			_, err := ic.ReadMessage(userId)
			if err != nil {
				logx.Error("[WebsocketHandler] read message error. ", err)
				ic.Remove(userId)
				closeMsg <- struct{}{}
				return
			}
		}()

		// wg.Wait()

		for {
			<-closeMsg
			logx.Info("[WebsocketHandler] closeMsg received ...")
			return
		}

	}
}
