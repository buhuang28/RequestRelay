package ws_client

import (
	"RequestRelayClient/data"
	"RequestRelayClient/tool"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	//wsServerAddr string = "ws://127.0.0.1:8080/echo"
	LocalServerAddr = "http://127.0.0.1:"
	WsCon           *websocket.Conn
	ConSuccess      bool = false
	WSRLock         sync.Mutex
	WSWLock         sync.Mutex
)

const (
	SUCCESS = "链接成功"
	FAIL    = "链接失败"
)

func WsDailCall() {
	data.SettingData = data.SettingData
	var err error
	var header http.Header = make(map[string][]string)
	header.Add("origin", data.SettingData.ServerAddr)
	if data.SettingData.WsId == "" {
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
		if ConSuccess {
			break
		}
		fmt.Println(data.SettingData.ServerAddr)
		WsCon, _, err = websocket.DefaultDialer.Dial(data.SettingData.ServerAddr, header)
		if err != nil || WsCon == nil {
			fmt.Println(err)
			ClientForm.WSStatus.SetCaption(FAIL)
			log.Info("dial:", err)
		} else {
			ClientForm.WSStatus.SetCaption(SUCCESS)
			ConSuccess = true
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
		if WsCon == nil || !ConSuccess {
			time.Sleep(time.Second * 2)
			continue
		}
		WSRLock.Lock()
		_, message, e := WsCon.ReadMessage()
		WSRLock.Unlock()
		fmt.Println("收到消息:", string(message))
		if e != nil {
			log.Info(e)
			time.Sleep(time.Second * 2)
			go func() {
				ConSuccess = false
				fmt.Println("ws-server掉线，正在重连")
				WsDailCall()
			}()
			return
		}
		var requestData data.RequestData
		e = json.Unmarshal(message, &requestData)
		if e != nil {
			log.Info(e)
			requestData.Status = -3
			SendRequestMessage(requestData)
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
		SendResponseMessage(responseData)
	}
}

func SendResponseMessage(data data.ResponseData) error {
	WSWLock.Lock()
	defer WSWLock.Unlock()
	marshal, _ := json.Marshal(data)
	err := WsCon.WriteMessage(websocket.TextMessage, marshal)
	return err
}

func SendRequestMessage(data data.RequestData) error {
	WSWLock.Lock()
	defer WSWLock.Unlock()
	marshal, _ := json.Marshal(data)
	err := WsCon.WriteMessage(websocket.TextMessage, marshal)
	return err
}
