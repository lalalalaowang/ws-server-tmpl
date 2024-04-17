package net

import (
	"github.com/gorilla/websocket"
	"sync"
)

type wsServer struct {
	conn         *websocket.Conn
	router       *router
	outChan      chan *WSRespMsg
	Seq          int
	property     map[string]string
	propertyLock sync.RWMutex
}

func NewWsServer(conn *websocket.Conn) *wsServer {
	return &wsServer{
		conn:     conn,
		outChan:  make(chan *WSRespMsg, 1000),
		property: make(map[string]string),
		Seq:      0,
	}
}
func (w *wsServer) Router(router *router) {
	w.router = router
}

func (w *wsServer) Start() {
	go w.readMsgLoop()
	go w.writeMsgLoop()
}

func (w *wsServer) SetProperty(key string, value string) {
	w.propertyLock.Lock()
	defer w.propertyLock.Unlock()
	w.property[key] = value
}

func (w *wsServer) GetProperty(key string) string {
	w.propertyLock.RLock()
	defer w.propertyLock.RUnlock()
	return w.property[key]
}

func (w *wsServer) RemoveProperty(key string) {
	w.propertyLock.Lock()
	defer w.propertyLock.Unlock()
	delete(w.property, key)
}

func (w *wsServer) Addr() string {
	return w.conn.RemoteAddr().String()
}

func (w *wsServer) Push(name string, data interface{}) {
	resp := &WSRespMsg{
		Body: &RespBody{
			Seq:  0,
			Name: name,
			Msg:  data,
		},
	}

	w.outChan <- resp
}

func (w *wsServer) readMsgLoop() {

}

func (w *wsServer) writeMsgLoop() {

}
