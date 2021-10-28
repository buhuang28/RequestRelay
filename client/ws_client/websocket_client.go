package ws_client

import (
	"RequestRelayClient/data"
	"RequestRelayClient/rlog"
	"RequestRelayClient/tool"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
)

var (
	//wsServerAddr string = "ws://127.0.0.1:8080/echo"
	LocalServerAddr = "http://127.0.0.1:"
	WsCon           *websocket.Conn
	ConSucess       bool = false
)

const (
	SUCCESS = "链接成功"
	FAIL    = "链接失败"
)

func WsDailCall() {
	data.SettingData = data.SettingData
	fmt.Println(data.SettingData)
	var err error
	var header http.Header = make(map[string][]string)
	header.Add("origin", data.SettingData.ServerAddr)

	if data.SettingData.WsId == "" {
		fmt.Println("卧槽尼玛：", data.SettingData.WsId)
		nano := time.Now().UnixNano()
		wsId := strconv.FormatInt(nano, 10) + uuid.NewV4().String()
		wsId = tool.Md5(wsId)
		data.SettingData.WsId = wsId
		marshal, _ := json.Marshal(data.SettingData)
		tool.WriteFile(data.SETTINGFILE, string(marshal))
	}

	wsId := data.SettingData.WsId

	if data.SettingData.Note != "" {
		wsId += "|" + data.SettingData.Note
	}

	//请求头这里带上id
	header.Add("ws_id", wsId)

	//Label:
	for {
		if ConSucess {
			break
		}
		fmt.Println(data.SettingData.ServerAddr)
		WsCon, _, err = websocket.DefaultDialer.Dial(data.SettingData.ServerAddr, header)
		if err != nil || WsCon == nil {
			fmt.Println(err)
			ClientForm.WSStatus.SetCaption(FAIL)
			rlog.Log.Println("dial:", err)
		} else {
			ClientForm.WSStatus.SetCaption(SUCCESS)
			ConSucess = true
			go func() {
				HandleWsMsg()
			}()
			return
		}
		time.Sleep(time.Second * 2)
	}
}

func HandleWsMsg() {
	for {
		_, message, e := WsCon.ReadMessage()
		fmt.Println("收到消息:", string(message))
		if e != nil {
			fmt.Println("出错了")
			rlog.Log.Println(e)
			time.Sleep(time.Second * 2)
			go func() {
				ConSucess = false
				fmt.Println("ws-server掉线，正在重连")
				WsDailCall()
			}()
			return
		}
		var requestData data.RequestData
		e = json.Unmarshal(message, &requestData)
		if e != nil {
			rlog.Log.Println(e)
			requestData.Status = -3
			marshal, _ := json.Marshal(requestData)
			WsCon.WriteMessage(websocket.TextMessage, marshal)
			continue
		}

		//请求后返回给服务端的数据
		var responseData data.ResponseData

		method := requestData.Method
		var result []byte
		var requestStatus bool
		switch method {
		case "GET":
			requestStatus, result = tool.GetRequest(LocalServerAddr+data.SettingData.ServicePort+requestData.Path, requestData.Header)
		case "POST":
			requestStatus, result = tool.PostRequest(LocalServerAddr+data.SettingData.ServicePort+requestData.Path, requestData.Header, requestData.Body)
		}
		if !requestStatus {
			responseData.Status = -7
		} else {
			responseData.Status = 1
		}
		//应该在这里修改GUI
		responseData.MessageId = requestData.MessageId
		responseData.Body = string(result)
		marshal, _ := json.Marshal(responseData)
		WsCon.WriteMessage(websocket.TextMessage, marshal)
	}
}
