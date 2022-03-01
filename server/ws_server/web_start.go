package ws_server

import (
	"RequestRelayServer/data"
	"github.com/gin-gonic/gin"
	"github.com/ying32/govcl/vcl"
	"os"
)

func WebStart() {
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
	router.POST("/test", func(c *gin.Context) {
		d := make(map[string]interface{})
		var arr []interface{}
		d1 := make(map[string]interface{})
		d2 := make(map[string]interface{})
		d3 := make(map[string]interface{})
		d1["bot_id"] = "1234"
		d1["bot_status"] = 1
		d1["bot_oauth"] = 0
		d1["avator"] = "http://q2.qlogo.cn/headimg_dl?dst_uin=2913083381&spec=640"
		d2["avator"] = "http://q2.qlogo.cn/headimg_dl?dst_uin=2913083381&spec=640"
		d3["avator"] = "http://q2.qlogo.cn/headimg_dl?dst_uin=2913083381&spec=640"
		d2["bot_id"] = "3456"
		d2["bot_status"] = 0
		d3["bot_id"] = "567"
		d2["bot_oauth"] = 1

		arr = append(arr, d1)
		arr = append(arr, d2)
		arr = append(arr, d3)
		//d["bot_id"]=123456
		d["list"] = arr
		c.JSON(200, d)
	})
	router.NoRoute(relayController.HandleRequest)
	//保留一个做web_client
	err := router.Run(":" + data.SettingData.RunPort)
	if err != nil {
		vcl.ShowMessage("运行失败,可能" + data.SettingData.RunPort + "端口被占用")
		os.Exit(0)
	}
}
