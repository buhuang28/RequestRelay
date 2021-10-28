// 代码由简易GoVCL IDE自动生成。
// 不要更改此文件名
// 在这里写你的事件。

package ws_server

import (
	"RequestRelayServer/data"
	"RequestRelayServer/tool"
	"encoding/json"
	"github.com/ying32/govcl/vcl"
	"strconv"
	"strings"
)

//::private::
type TServerFormFields struct {
}

func (f *TServerForm) OnFormCreate(sender vcl.IObject) {

	f.ScreenCenter()

	readFile := tool.ReadFile(data.SETTINGFILE)
	if readFile == "" {
		return
	}
	err := json.Unmarshal([]byte(readFile), &data.SettingData)
	if err != nil {
		vcl.ShowMessage("设置数据反序列化失败")
		return
	}

	if data.SettingData.OutTime == 0 {
		data.SettingData.OutTime = 6
	}else {
		f.OutTime.SetText(strconv.FormatInt(data.SettingData.OutTime,10))
	}

	if data.SettingData.SecretKey != "" {
		f.Secret.SetText(data.SettingData.SecretKey)
	}

	if data.SettingData.RunPort != "" {
		f.Port.SetText(data.SettingData.RunPort)
	}

	if data.SettingData.WSPort != "" {
		f.WSPort.SetText(data.SettingData.WSPort)
	}

}

func (f *TServerForm) OnRunButtonClick(sender vcl.IObject) {
	runPort := f.Port.Text()
	runPort = strings.TrimSpace(runPort)
	if runPort == "" {
		vcl.ShowMessage("运行端口不可为空")
		return
	}

	_, err := strconv.ParseInt(runPort, 10, 64)
	if err != nil {
		vcl.ShowMessage("运行端口不正确")
		return
	}

	secretKey := f.Secret.Text()
	secretKey = strings.TrimSpace(secretKey)

	wsPort := f.WSPort.Text()
	wsPort = strings.TrimSpace(wsPort)
	if wsPort == "" {
		vcl.ShowMessage("WebSocket端口不可为空")
		return
	}

	_, err = strconv.ParseInt(wsPort, 10, 64)
	if err != nil {
		vcl.ShowMessage("WebSocket端口错误")
		return
	}

	var setting data.Setting

	outTime := f.OutTime.Text()
	outTime = strings.TrimSpace(outTime)
	if outTime != "" {
		ot, e := strconv.ParseInt(outTime, 10, 64)
		if e != nil || ot == 0 {
			vcl.ShowMessage("超时时间错误")
			return
		}
		setting.OutTime = ot
	}

	setting.RunPort = runPort
	setting.SecretKey = secretKey
	setting.WSPort = wsPort
	marshal, _ := json.Marshal(setting)
	tool.WriteFile(data.SETTINGFILE,string(marshal))

	data.SettingData = setting
	go func() {
		WSStart()
	}()
	go func() {
		WebStart()
	}()

}