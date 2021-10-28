package data

type Setting struct {
	//服务端地址--端口
	ServerAddr string `json:"server_addr"`

	//链接状态
	ConStatus int64 `json:"con_status"`

	//本地websocket的端口
	WSPort string `json:"ws_port"`

	//本地服务端口
	ServicePort string `json:"service_port"`

	//链接id
	WsId string `json:"ws_id"`

	//链接备注
	Note string `json:"note"`
}

var (
	//全局设置
	SettingData Setting
)

const (
	SETTINGFILE = "client_setting.json"
)