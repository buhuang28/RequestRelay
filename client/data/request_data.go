package data

//向ws-server传递数据的结构体
type RequestData struct {
	//消息类型id  0:消息传递, 1:给client钦点id
	MessageType int64 `json:"message_type"`
	//消息ID
	MessageId int64 `json:"message_id"`
	//请求发起时间
	Time int64 `json:"time"`
	//请求状态    -1:WsData json序列化失败，-2:ws没发出去，一般是ws-client断开
	//-3:WsData json反序列化失败   -4:body json序列化失败    -5:body 反序列化失败
	Status int64 `json:"status"`
	//请求方法
	Method string `json:"method"`
	//请求路径
	Path string `json:"path"`
	//请求头
	Header map[string]string `json:"header"`
	//请求体/返回数据  server -> client:get请求为"",post则获取request.body内容(gin.getRaw()其实也可以)    client -> server:返回响应数据
	Body string `json:"body"`
}
