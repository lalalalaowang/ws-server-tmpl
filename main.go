package main

import (
	"ws-server/net"
	"ws-server/util"
)

func main() {
	server := net.NewServer(util.Conf.LoginServer.Host + ":" + util.Conf.LoginServer.Port)
	err := server.StartServer()
	if err != nil {
		return
	}
}
