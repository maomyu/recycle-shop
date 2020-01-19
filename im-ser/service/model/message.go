package model

// 需要显示的消息结构体
type Message struct {
	// 消息的唯一id
	ID string `json:"id"`
	// 发送者的id
	SendID string `json:"sendid"`
	// 接收者的id
	ReceiveID string `json:"receiveid"`
	// 消息的内容
	Content string `json:"content"`
	// 消息的类型
	Type int `json:"type"`
	// 消息显示的状态
	Status int `json:"status"`
}
