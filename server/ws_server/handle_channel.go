package ws_server

import (
	"RequestRelayServer/data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func HandleChannel(wsData data.RequestData,relayId string) chan data.ResponseData {
	ch := make(chan data.ResponseData,1)
	go func() {
		data.ChanMapLock.Lock()
		data.ChanMap[wsData.MessageId] = ch
		data.ChanMapLock.Unlock()
		marshal, err := json.Marshal(wsData)
		if err != nil {
			fmt.Println("json序列化失败:",err)
			ch <- data.ResponseData{Status: -1}
			close(ch)
			return
		}
		var conn *websocket.Conn
		if relayId != "" {
			conn = data.UserConMap[relayId].Con
		}else {
			for _,v := range data.UserConMap {
				conn = v.Con
				break
			}
		}
		if conn == nil {
			fmt.Println("空ws链接")
			ch <- data.ResponseData{Status: -6}
			close(ch)
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			fmt.Println("发送失败:",err)
			ch <- data.ResponseData{Status: -2}
			close(ch)
			return
		}
	}()
	return ch
}
