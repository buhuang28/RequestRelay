package data

type ResponseData struct {
	//消息类型id  0:消息传递, 1:给client钦点id
	MessageType int64 `json:"message_type"`
	//消息ID
	MessageId int64 `json:"message_id"`
	//请求状态    -1:WsData json序列化失败，-2:ws没发出去，一般是ws-client断开
	//-3:WsData json反序列化失败   -4:body json序列化失败    -5:body 反序列化失败
	Status int64 `json:"status"`
	//请求体
	Body string `json:"body"`
}