package tool

import (
	"RequestRelayClient/rlog"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

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

func PrintStackTrace(err interface{})  {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	rlog.Log.Println(buf.String())
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//复写
func WriteFile(fileName,content string) bool {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	defer func() {
		f.Close()
	}()
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		write, e := f.Write([]byte(content))
		if e == nil && write > 0 {
			return true
		}
	}
	return false
}


func ReadFile(fileName string) string {
	exist := CheckFileIsExits(fileName)
	if !exist {
		return ""
	}
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read fail", err)
		return ""
	}
	return string(f)
}