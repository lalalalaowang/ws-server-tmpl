package main

import (
	"fmt"
	"ws-server/util"
)

func main() {
	fmt.Println(util.Conf.LoginServer.Addr)
	fmt.Println(util.Conf.LoginServer.Port)
}
