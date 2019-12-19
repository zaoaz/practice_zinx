package znet

import "zinx/ziface"

type Request struct {
	//已和客户端建立好的连接 conn
	conn ziface.IConnection
	//客户端请求数据
	data []byte
}

//得到当前连接
func (r *Request) GetConn() ziface.IConnection {
	return r.conn
}

//得到请求数据
func (r *Request) GetData() []byte {
	return r.data
}
