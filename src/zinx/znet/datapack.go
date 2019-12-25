package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

//封包 拆包 具体模块
type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	//dataLen unint32 (4字节) id uint32 (4字节)
	return 8
}

//封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建存放byte字节流的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	//将datalen 写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//将msgid 写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//将data数据 写入databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

//拆包方法  将包的head信息读取出来  根据head信息中的 data的长度在进行读取
func (dp *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个输入二进制数据的ioreader
	dataBuff := bytes.NewReader(binaryData)
	//只解压head信息，得到datalen和msgid

	msg := &Message{}
	//读取datalen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//msgid
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	//判断datalen是否超出允许包长度
	if utils.GlobalObject.MaxPackageSize > 0 && utils.GlobalObject.MaxPackageSize < msg.DataLen {
		return nil, errors.New("too large msg data recv!")
	}
	//读取msgid

	return msg, nil
}
