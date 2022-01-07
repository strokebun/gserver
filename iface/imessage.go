package iface

// @Description: 封装传输的数据
// @Author: StrokeBun
// @Date: 2022/1/7 10:03
type IMessage interface {
	GetMsgId() uint32
	SetMsgId(id uint32)

	GetMsgLen() uint32
	SetMsgLen(msgLen uint32)

	GetData() []byte
	SetData(data []byte)
}