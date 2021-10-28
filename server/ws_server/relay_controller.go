package ws_server

import (
	"RequestRelayServer/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ying32/govcl/vcl"
	"io/ioutil"
	"strconv"
	"sync"
	"time"
)

var (
	timeTemplate = "2006/01/02 15:04:05"
)

type RelayController struct {}

func NewRelayController() RelayController {
	return RelayController{}
}

var (
	MessageId int64 = 0
	MessageIdToListViewIndex = make(map[int64]int32)
	RequestListViewLock sync.Mutex
)

func (r *RelayController) HandleRequest(c *gin.Context)  {
	path := c.Request.RequestURI
	if path == "/favicon.ico" {
		return
	}
	nowTime := time.Now().UnixNano()/1e6

	msgId := MessageId
	MessageId++
	method := c.Request.Method
	requestHeader := c.Request.Header
	header := make(map[string]string)
	for k,v := range requestHeader {
		for _,v2 := range v {
			tempV := header[k]
			if tempV != "" {
				header[k] = tempV + ";" + v2
			}
		}
	}

	//钦点的客户端链接id
	relayId := c.Query("relayId")
	if relayId == "" {
		for k,_ := range data.UserConMap {
			relayId = k
			break
		}
	}
	json := ""
	if method != "GET" {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err == nil {
			json = string(body)
		}
	}
	wsData := data.RequestData{MessageId: msgId, Method: method, Path: path, Header: header,Body: json}
	go func() {
		vcl.ThreadSync(func() {
			item := ServerForm.RequestListView.Items().Add()
			item.SetCaption("")
			item.SubItems().Add(relayId)
			item.SubItems().Add(time.Unix(time.Now().Unix(), 0).Format(timeTemplate))
			item.SubItems().Add(path)
			item.SubItems().Add(method)
			item.SubItems().Add(json)
			item.SubItems().Add("")
			item.SubItems().Add("")
			index := item.Index()
			RequestListViewLock.Lock()
			MessageIdToListViewIndex[msgId] = index
			RequestListViewLock.Unlock()
		})
	}()

	//这里是把接收到的消息丢给客户端
	ch := HandleChannel(wsData,relayId)

	//等待channel返回消息(等待客户端响应)
	var wsResult data.ResponseData
	select {
		case wsResult = <- ch:
			fmt.Println("收到结构体数据")
			go func() {
				vcl.ThreadSync(func() {
					RequestListViewLock.Lock()
					index := MessageIdToListViewIndex[wsResult.MessageId]
					RequestListViewLock.Unlock()
					item := ServerForm.RequestListView.Items().Item(index)
					item.SubItems().SetStrings(5,wsResult.Body)
					nowTime2 := time.Now().UnixNano()/1e6
					useTime := nowTime2 - nowTime
					item.SubItems().SetStrings(6,strconv.FormatInt(useTime,10)+" ms")
					wsResult.UseTime = useTime
				})
			}()
			break
		case <- time.After(time.Second * time.Duration(data.SettingData.OutTime)):
			fmt.Println("超时打断")
			go func() {
				vcl.ThreadSync(func() {
					RequestListViewLock.Lock()
					index := MessageIdToListViewIndex[wsData.MessageId]
					RequestListViewLock.Unlock()
					item := ServerForm.RequestListView.Items().Item(index)
					item.SubItems().SetStrings(5,"null")
					item.SubItems().SetStrings(6,"超时")
				})
			}()
			break
		}
	if wsResult.Status != 1 {
		c.JSON(200, GetStatusExplain(wsResult.Status))
		return
	}
	c.JSON(200,wsResult.Body)
}

func GetStatusExplain(status int64) string {
	switch status {
	case -1:
		return "json序列化失败"
	case -2:
		return "websocket消息发送失败，请检查websocket-client是否开启"
	default:
		return "未知错误"
	}
}