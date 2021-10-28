package ws_server

import (
	"RequestRelayServer/data"
	"RequestRelayServer/rlog"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	upgrader = websocket.Upgrader{} // use default options
)

func Echo(w http.ResponseWriter, r *http.Request) {
	var e error
	WsCon, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		rlog.Log.Println("upgrade",e)
		log.Print("upgrade:", e)
		return
	}
	if r.Header == nil {
		fmt.Println("请求头不能为空")
		return
	}
	wsId := r.Header.Get("ws_id")
	if wsId == "" {
		fmt.Println("链接id为空")
		return
	}

	var userCon data.UserCon
	userCon.Time = time.Now().Unix()
	userCon.Con = WsCon
	userCon.Id = wsId

	runes := []rune(wsId)
	id := ""
	note := ""
	id = string(runes[:32])
	if len(runes) > 32 {
		note = string(runes[33:])
	}
	item := ServerForm.ClientListView.Items().Add()
	item.SetCaption("")
	item.SubItems().Add(id)
	item.SubItems().Add(note)
	userCon.Note = note

	data.WsConLock.Lock()
	data.UserConMap[wsId] = userCon
	data.WsConLock.Unlock()

	defer func() {
		fmt.Println("defer执行了")
		data.WsConLock.Lock()
		delete(data.UserConMap,wsId)
		var i int32 = 0
		for i = 0;i < ServerForm.ClientListView.Items().Count();i++ {
			fmt.Println(wsId)
			if strings.Contains(wsId,"|") {
				split := strings.Split(wsId, "|")
				wsId = split[0]
			}
			if wsId == ServerForm.ClientListView.Items().Item(i).SubItems().Strings(0) {
				ServerForm.ClientListView.Items().Delete(i)
				break
			}
		}
		data.WsConLock.Unlock()
		WsCon.Close()
	}()
	for {
		//收到client的消息
		_, message, err := WsCon.ReadMessage()
		if err != nil {
			rlog.Log.Println("read:",err)
			log.Println("read:", err)
			break
		}
		responseData := data.ResponseData{}
		err = json.Unmarshal(message, &responseData)
		if err != nil {
			break
		}
		//0是需要转发的消息，1是客户端对获取到的备注
		switch responseData.MessageType {
		case 0:
			id := responseData.MessageId
			data.ChanMapLock.Lock()
			c := data.ChanMap[id]
			data.ChanMapLock.Unlock()
			c <- responseData
		case 1:
			//body := responseData.Body
			//if body == "" {
			//	rlog.Log.Println("接收到的body为空")
			//	break
			//}
			////body:32位MD5+|+备注
			//runes := []rune(body)
			//id := ""
			//note := ""
			//id = string(runes[:32])
			//if len(runes) > 32 {
			//	note = string(runes[33:])
			//}
			//item := vcl_ui.ServerForm.ClientListView.Items().Add()
			//item.SetCaption("")
			//item.SubItems().Add(id)
			//item.SubItems().Add(note)
			//userCon.Note = note
			//data.WsConLock.Lock()
			//data.UserConMap[wsId] = userCon
			//data.WsConLock.Unlock()
		}
	}
	return
}

