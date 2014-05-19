package controllers

// 预定义统一的json消息体最外层格式
type Message struct {
	Message interface{}
	RetCode int
}
