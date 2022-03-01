package main

import (
	"RequestRelayServer/logs"
	"RequestRelayServer/ws_server"
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func init() {
	logs.InitLog()
}

func main() {
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&ws_server.ServerForm)
	vcl.Application.Run()
}
