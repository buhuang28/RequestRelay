package rlog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	Log *log.Logger
)

func InitLog() {
	logFileNmae := `./rlog/` + time.Now().Format("20060102") + "-client.log"
	logFileAllPath := logFileNmae
	_, err := os.Stat(logFileAllPath)

	exits := CheckFileIsExits(`rlog`)
	if !exits {
		_ = os.Mkdir("./rlog", os.ModePerm)
	}

	var f *os.File
	if err != nil {
		f, _ = os.Create(logFileAllPath)
	} else {
		//如果存在文件则 追加log
		f, _ = os.OpenFile(logFileAllPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}
	Log = log.New(f, "", log.LstdFlags)
	Log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func CheckFileIsExits(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func PrintStackTrace(err interface{}) {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	Log.Println(buf.String())
}
