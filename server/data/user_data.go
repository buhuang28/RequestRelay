package data

import (
	"github.com/gorilla/websocket"
	"sync"
)

var (
	WsConLock   sync.Mutex
	UserConLock sync.Mutex
	UserConMap  = make(map[string]UserCon)
)

type UserCon struct {
	Id   string          `json:"id"`   //id
	Note string          `json:"note"` //备注
	Time int64           `json:"time"` //链接时间
	Con  *websocket.Conn //websocket链接
}

func GetUserCon(conId string) UserCon {
	UserConLock.Lock()
	defer UserConLock.Unlock()
	var userCon UserCon
	if conId == "" {
		for _, v := range UserConMap {
			userCon = v
			break
		}
	} else {
		userCon = UserConMap[conId]
	}
	return userCon
}
