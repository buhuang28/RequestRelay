package ws_server

import (
	"RequestRelayServer/data"
	"github.com/gin-gonic/gin"
	"os"
)

func WebStart()  {
	router := gin.Default()
	relayController := NewRelayController()
	router.GET("/test", func(c *gin.Context) {
		c.String(200,"success")
	})
	router.NoRoute(relayController.HandleRequest)
	//保留一个做web_client
	err := router.Run(":" + data.SettingData.RunPort)
	if err != nil {
		//vcl.ShowMessage("运行失败,可能"+data.SettingData.RunPort+"端口被占用")
		os.Exit(0)
	}
}
