// 代码由简易GoVCL IDE自动生成。
// 不要更改此文件名
// 在这里写你的事件。

package ws_client

import (
    "RequestRelayClient/data"
    "RequestRelayClient/tool"
    "encoding/json"
    "fmt"
    "github.com/ying32/govcl/vcl"
    "strconv"
    "strings"
)

//::private::
type TClientFormFields struct {
}

func (f *TClientForm) OnFormCreate(sender vcl.IObject) {
    f.ScreenCenter()

    readFile := tool.ReadFile(data.SETTINGFILE)
    if readFile == "" {
        return
    }
    var setting data.Setting
    err := json.Unmarshal([]byte(readFile), &setting)
    if err != nil {
        vcl.ShowMessage("设置数据反序列化失败")
        return
    }
    if setting.ServicePort != "" {
        f.LocalPort.SetText(setting.ServicePort)
    }

    if setting.Note != "" {
        f.WSNote.SetText(setting.Note)
    }

    if setting.WsId != "" {
        f.WSId.SetCaption(setting.WsId)
    }

    if setting.ServerAddr != "" {
        f.ServerAddr.SetText(setting.ServerAddr)
    }

    if setting.WSPort != "" {
        f.WSClientPort.SetText(setting.WSPort)
    }

    data.SettingData = setting
    fmt.Println(data.SettingData.WsId)
}

func (f *TClientForm) OnRunButtonClick (sender vcl.IObject)  {
    ServerAddr := f.ServerAddr.Text()
    ServerAddr = strings.TrimSpace(ServerAddr)
    if ServerAddr == "" {
        vcl.ShowMessage("服务端地址不可为空")
        return
    }

    clientPort := f.WSClientPort.Text()
    clientPort = strings.TrimSpace(clientPort)
    if clientPort == "" {
        vcl.ShowMessage("本地链接端口不可为空")
        return
    }

    _, err := strconv.ParseInt(clientPort, 10, 64)
    if err != nil {
        vcl.ShowMessage("本地链接端口错误")
        return
    }

    note := f.WSNote.Text()
    note = strings.TrimSpace(note)

    servicePort := f.LocalPort.Text()
    servicePort = strings.TrimSpace(servicePort)
    if servicePort == "" {
        vcl.ShowMessage("本地服务端口不可为空")
        return
    }

    _, err = strconv.ParseInt(servicePort, 10, 64)
    if err != nil {
        vcl.ShowMessage("本地服务端口错误")
        return
    }

    data.SettingData.ServicePort = servicePort
    data.SettingData.Note = note
    data.SettingData.ServerAddr = ServerAddr
    data.SettingData.WSPort = clientPort

    marshal, e := json.Marshal(data.SettingData)
    if e != nil {
        vcl.ShowMessage("数据序列化失败")
        return
    }
    tool.WriteFile(data.SETTINGFILE,string(marshal))
    go func() {
       WsDailCall()
    }()
}


