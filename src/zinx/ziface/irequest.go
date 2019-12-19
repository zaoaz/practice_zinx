package ziface

//irequest 把客户端请求的数据包装到request中
type IRquest interface {
	//得到当前连接
	GetConn() IConnection
	//得到请求数据
	GetData() []byte
}
