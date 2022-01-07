package iface

// @Description: 数据包接口，实现封包、拆包
// @Author: StrokeBun
// @Date: 2022/1/7 10:32
type IDataPack interface {
	// 获取包头的长度
	GetHeaderLen() uint32
	// 封包
	Pack(msg IMessage) ([]byte, error)
	// 拆包
	Unpack([]byte) (IMessage, error)
}
