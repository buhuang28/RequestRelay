package data

import (
	"sync"
)

var (
	//消息id-channel
	ChanMap     = make(map[int64]chan ResponseData)
	ChanMapLock sync.Mutex
)

func GetResponseDataChan(messageId int64) chan ResponseData {
	ChanMapLock.Lock()
	defer ChanMapLock.Unlock()
	data := ChanMap[messageId]
	return data
}

func DeleteResponseDataChan(messageId int64) {
	ChanMapLock.Lock()
	defer ChanMapLock.Unlock()
	delete(ChanMap, messageId)
}

func SaveResponseDataChan(messageId int64, data chan ResponseData) {
	ChanMapLock.Lock()
	defer ChanMapLock.Unlock()
	ChanMap[messageId] = data
}
