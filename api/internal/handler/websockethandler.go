/*
 * @Author: LEESON
 * @Date: 2024-11-06 23:05:42
 */
package handler

import (
	"AiChatPartner/api/internal/svc"
	"net/http"
	"sync"

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
	// ReadMessage(uint32) ([]byte, error)
}

var wg sync.WaitGroup

var gcm ConnectionManager

var uid uint32 = 1

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

func (n *Session) SendMessage() {
	for {
		select {
		case message, ok := <-n.Message:
			if !ok {
				return
			}
			err := n.WsConn.WriteMessage(1, []byte(message))
			if err != nil {
				logx.Error("[SendMessage] write message error. ", err)
				n.Close()
				return
			}
		}
	}
}

func (n *Session) RecvMessage() {
	logx.Info("[RecvMessage] start ...")
	for {
		_, message, err := n.WsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Error("[RecvMessage] read message error: ", err)
			}
			n.Close() // 关闭通道，通知其他监听者
			return
		}
		logx.Info("[RecvMessage] recv message: ", string(message))
		n.Message <- string(message)
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
	uid++
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
		node := Session{
			WsConn:  conn,
			Message: make(chan string),
			ID:      UserID(uid),
		}
		defer node.Close()

		gcm.connections = make(map[UserID]*Session)
		gcm.Add(&node)
		defer gcm.Remove(uid)
		logx.Infof("[WebsocketHandler] User %d connected", uid)

		wg.Add(2)
		go node.SendMessage()
		go node.RecvMessage()

		wg.Wait()
	}
}
