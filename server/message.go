package server

// @Description: 消息传输模块
// @Author: StrokeBun
// @Date: 2022/1/7 10:03
type Message struct {
	// 消息id
	Id uint32
	// 消息长度
	DataLen uint32
	// 消息的内容
	Data    []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message)  SetMsgLen(msgLen uint32){
	m.DataLen = msgLen
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}



