package net

type ReqBody struct {
	Seq   int         `json:"seq"`
	Name  string      `json:"name"`
	Msg   interface{} `json:"msg"`
	Proxy string      `json:"proxy"`
}

type RespBody struct {
	Seq  int         `json:"seq"`
	Name string      `json:"name"`
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

type WSReqMsg struct {
	Body *ReqBody `json:"body"`
	Conn WSConn
}

type WSRespMsg struct {
	Body *RespBody `json:"body"`
}

type WSConn interface {
	SetProperty(key string, value string)
	GetProperty(key string) string
	RemoveProperty(key string)
	Addr() string
	Push(name string, data interface{})
}
