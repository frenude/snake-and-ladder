package utils

type BaseRsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Body interface{} `json:"body,omitempty"`
}
