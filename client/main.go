package main

import (
	"RequestRelayClient/rlog"
	"RequestRelayClient/ws_client"
	"github.com/ying32/govcl/vcl"
)

func init()  {
	rlog.InitLog()
}

func main()  {
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&ws_client.ClientForm)
	vcl.Application.Run()
}