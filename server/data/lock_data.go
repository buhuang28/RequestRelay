package data

import (
	"sync"
)

var (
	//消息id-channel
	ChanMap = make(map[int64]chan ResponseData)
	ChanMapLock sync.Mutex
)