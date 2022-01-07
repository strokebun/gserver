package server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/strokebun/gserver/conf"
	"github.com/strokebun/gserver/iface"
)

// @Description: 传输数据包
// @Author: StrokeBun
// @Date: 2022/1/7 10:32
type DataPack struct {

}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包头的长度
func (d *DataPack) GetHeaderLen() uint32 {
	// Datalen(4 byte) + Id(4 byte)
 	return 8
}

// 封包
func (d *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	// 写入msgLen
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	// 写入Id
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 写入data
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 拆包
func (d *DataPack) Unpack(data []byte) (iface.IMessage, error) {
	buf := bytes.NewReader(data)
	msg := &Message{}
	// 读入msgLen
	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// 读入id
	if err := binary.Read(buf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	// 读入data
	if err := binary.Read(buf, binary.LittleEndian, &msg.Data); err != nil {
		return nil, err
	}

	maxPackageSize := conf.GlobalObject.MaxPackageSize
	if maxPackageSize > 0 && msg.DataLen > maxPackageSize {
		return nil, errors.New("datapack oversize, maxsize is " + string(maxPackageSize) + " bytes")
	}
	return msg, nil
}