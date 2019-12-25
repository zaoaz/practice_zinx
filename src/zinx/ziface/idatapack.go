package ziface

//封包拆包 模块
//直接面向tcp连接中的数据流 用于处理tcp粘包问题

type IDataPack interface {
	//获取包长度方法
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	UnPack([]byte) (IMessage, error)
}
