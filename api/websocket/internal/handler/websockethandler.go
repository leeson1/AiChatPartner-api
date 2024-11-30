package handler

import (
	"net/http"
	"strconv"
	"sync"

	"AiChatPartner/api/websocket/internal/svc"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type Session struct {
	WsConn    *websocket.Conn
	Message   chan string
	closeOnce sync.Once
	ID        UserID
}

type UserID uint32

type ConnectionManager struct {
	connections map[UserID]*Session
	mutex       sync.RWMutex
}

type IConnectionManager interface {
	Add(*Session)
	Remove(uint32)
	// RemoveWithCode(uint32, int, string)
	// Get(uint32) (*Session, bool)
	// SendMessage(uint32, []byte) error
	ReadMessage(uint32) ([]byte, error)
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

// func (cm *ConnectionManager) SendMessage(UserID uint32, message []byte) error {
// 	for {
// 		select {
// 		case message, ok := <-n.Message:
// 			if !ok {
// 				return
// 			}
// 			err := n.WsConn.WriteMessage(1, []byte(message))
// 			if err != nil {
// 				logx.Error("[SendMessage] write message error. ", err)
// 				n.Close()
// 				return
// 			}
// 		}
// 	}
// }

func (cm *ConnectionManager) HandlerReadMessage(userID uint32) ([]byte, error) {
	_, msg, err := cm.connections[UserID(userID)].WsConn.ReadMessage()
	if err != nil {
		logx.Error("[RecvMessage] read message error. ", err)
		return msg, err
	}

	logx.Infof("[RecvMessage] User %d received message: %s", userID, string(msg))
	return msg, nil
}

func (cm *ConnectionManager) ReadMessage(userID uint32) ([]byte, error) {
	for {
		msg, err := cm.HandlerReadMessage(userID)
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

func (cm *ConnectionManager) Remove(userID uint32) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	if s, ok := cm.connections[UserID(userID)]; ok {
		s.WsConn.Close()
		delete(cm.connections, UserID(userID))
	}
}

func WebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("[WebsocketHandler] start ...")
		conn, err := upgrade(w, r, svcCtx)
		if err != nil {
			logx.Error("[WebsocketHandler] upgrade error. ", err)
			return
		}

		sUid := r.URL.Query().Get("uid")
		if sUid == "" {
			logx.Error("[WebsocketHandler] uid is empty.")
			conn.Close()
			return
		}
		uid, err := strconv.ParseUint(sUid, 10, 32)
		if err != nil {
			logx.Error("[WebsocketHandler] uid convert error: ", err)
			conn.Close()
			return
		}
		userId := uint32(uid)

		node := Session{
			WsConn:  conn,
			Message: make(chan string),
			ID:      UserID(uid),
		}
		defer node.Close()

		ic.Add(&node)
		defer ic.Remove(userId)
		logx.Infof("[WebsocketHandler] User %d connected", userId)

		// wg.Add(2)
		// go ic.SendMessage()
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
