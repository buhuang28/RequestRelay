package data

type Setting struct {
	//gin的端口
	RunPort string `json:"run_port"`
	SecretKey string `json:"secret_key"`
	OutTime int64 `json:"out_time"`
	WSPort string `json:"ws_port"`
}

var (
	SettingData Setting
)

const (
	SETTINGFILE = "server_setting.json"
)