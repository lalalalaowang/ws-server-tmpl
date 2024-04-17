package net

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	Addr   string
	router *router
}

func NewServer(addr string) *Server {
	return &Server{Addr: addr}
}

func (s *Server) StartServer() error {
	http.HandleFunc("/", s.wsHandler)
	err := http.ListenAndServe(s.Addr, nil)
	return err
}

func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func(wsConn *websocket.Conn) {
		err := wsConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(wsConn)

	// 每当有websocket连接进来，启动一个处理服务
	wsServer := NewWsServer(wsConn)
	wsServer.Router(s.router)
}
