package main

import (
	"RequestRelayClient/logs"
	"RequestRelayClient/ws_client"
	"github.com/ying32/govcl/vcl"
)

func init() {
	logs.InitLog()
}

func main() {
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&ws_client.ClientForm)
	vcl.Application.Run()
}
