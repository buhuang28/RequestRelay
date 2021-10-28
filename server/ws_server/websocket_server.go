package ws_server

import (
	"RequestRelayServer/data"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"net/http"
	"os"
)

func WSStart()  {
	http.HandleFunc("/echo", Echo)
	fmt.Println("localhost:"+data.SettingData.WSPort)
	err := http.ListenAndServe(":"+data.SettingData.WSPort, nil)
	if err != nil {
		fmt.Println("运行2")
		vcl.ShowMessage("运行失败，可能"+data.SettingData.WSPort+"端口被占用")
		os.Exit(0)
	}else {
		fmt.Println("运行成功")
	}
}
