package rlog

import (
	"RequestRelayServer/tool"
	"log"
	"os"
	"time"
)

var (
	Log *log.Logger
)

func InitLog() {
	logFileNmae := `./rlog/`+time.Now().Format("20060102")+"-server.log"
	logFileAllPath := logFileNmae
	_,err :=os.Stat(logFileAllPath)

	exits := tool.CheckFileIsExits(`rlog`)
	if !exits {
		_ = os.Mkdir("./rlog", os.ModePerm)
	}

	var f *os.File
	if  err != nil {
		f, _= os.Create(logFileAllPath)
	}else{
		//如果存在文件则 追加log
		f ,_= os.OpenFile(logFileAllPath,os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}
	Log = log.New(f, "", log.LstdFlags)
	Log.SetFlags(log.LstdFlags | log.Lshortfile)
}
