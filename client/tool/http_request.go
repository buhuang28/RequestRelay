package tool

import (
	"RequestRelayClient/rlog"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GetRequest(u string,header map[string]string) (bool,[]byte) {
	defer func() {
		e := recover()
		if e != nil {
			PrintStackTrace(e)
		}
	}()
	if u == "" {
		rlog.Log.Println("请求空地址")
		return false,[]byte("请求空地址")
	}
	if !strings.Contains(u,"https://") && !strings.Contains(u,"http://"){
		rlog.Log.Println("请求非法地址")
		return false,[]byte("请求非法地址")
	}

	request,_ := http.NewRequest("GET", u, nil)
	if header != nil && len(header) != 0 {
		for k,v := range header {
			request.Header.Set(k, v)
		}
	}
	//加入get参数
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()
	timeout := time.Duration(6 * time.Second)
	client := http.Client {
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp == nil {
		rlog.Log.Println("Get请求访问错误:",err,resp)
		return false,[]byte("Get请求访问错误")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil || data == nil {
		rlog.Log.Println("Get请求返回数据读取错误:",err,string(data))
		return false,[]byte("Get请求返回数据读取错误")
	}
	return true,data
}

func PostRequest(u string,header map[string]string,data interface{}) (bool,[]byte)  {
	defer func() {
		e := recover()
		if e != nil {
			PrintStackTrace(e)
		}
	}()
	bytesData,_ := json.Marshal(data)
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", u, reader)
	if err != nil {
		rlog.Log.Println("Post请求失败")
		return false,[]byte("Post请求失败")
	}
	if header != nil && len(header) != 0 {
		for k,v := range header {
			request.Header.Set(k,v)
		}
	}
	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		rlog.Log.Println("请求失败")
		return false,[]byte("请求失败")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rlog.Log.Println("数据读取失败")
		return false,[]byte("数据读取失败")
	}
	return true,respBytes
}
