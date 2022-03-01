package ws_server

import (
	"RequestRelayServer/data"
	"github.com/gin-gonic/gin"
	"github.com/ying32/govcl/vcl"
	"os"
)

func WebRun() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	relayController := NewRelayController()
	router.NoRoute(relayController.HandleRequest)
	//保留一个做web_client
	err := router.Run(":" + data.SettingData.RunPort)
	if err != nil {
		vcl.ShowMessage("运行失败,可能" + data.SettingData.RunPort + "端口被占用")
		os.Exit(0)
	}
}
