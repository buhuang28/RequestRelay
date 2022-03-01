package ws_server

import (
	"RequestRelayServer/data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func HandleChannel(wsData data.RequestData, relayId string) chan data.ResponseData {
	ch := make(chan data.ResponseData, 1)
	go func() {
		data.SaveResponseDataChan(wsData.MessageId, ch)
		userCon := data.GetUserCon(relayId)
		if userCon.Con == nil {
			fmt.Println("空ws链接")
			ch <- data.ResponseData{Status: -6}
			close(ch)
			return
		}
		err := SendWSMessage(userCon.Con, wsData)
		if err != nil {
			fmt.Println("发送失败:", err)
			ch <- data.ResponseData{Status: -2}
			close(ch)
			return
		}
	}()
	return ch
}

func SendWSMessage(conn *websocket.Conn, data data.RequestData) error {
	marshal, _ := json.Marshal(data)
	WSWLock.Lock()
	defer WSWLock.Unlock()
	err := conn.WriteMessage(websocket.TextMessage, marshal)
	return err
}
