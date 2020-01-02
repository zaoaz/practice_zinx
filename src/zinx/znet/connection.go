package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

//链接模块
type Connection struct {
	//当前链接socket Tcp
	Conn *net.TCPConn
	//当前链接id
	ConnId uint32
	//当前连接状态
	isClose bool
	//----------------
	////当前链接绑定的处理业务方法
	//handelAPI ziface.HandleFunc
	//------------------
	//告知当前链接已经停止的channel
	ExitChan chan bool

	//当前链接处理的方法router
	Router ziface.IRouter
}

//启动连接(让当前连接开始工作)
func (c *Connection) Start() {
	fmt.Printf("conn start connId = %d", c.ConnId)
	//启动当前链接读数据的业务
	go c.StartReader()
	//TODO 启动当前链接写数据业务
}
func (c *Connection) StartReader() {
	fmt.Printf("reader goroutine is running\n")
	defer fmt.Printf("connId  = %d is stop reader is exit ,remote add is %s \n", c.ConnId, c.RemoteAddr())
	defer c.Stop()

	for {
		//读取客户端数据到buf中
		//------------------v0.4
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("receive buff err ", err)
		//	continue
		//}
		//------------------------
		//--------------- v0.3
		//fmt.Printf("receive buff  %s \n", buf[:cnt])
		////调用当前链接绑定的handleAPI
		//if err := c.handelAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("ConnId    ", c.ConnId, " handle is error ", err)
		//	break
		//}
		//----------------------

		//创建拆包解包对象
		//v0.5
		dp := NewDataPack()
		//读取客户端的msg head 二进制流 8字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTcpConnection(), headData); err != nil {
			fmt.Println("ConnId    ", c.ConnId, " read head error ", err)
			break
		}
		//得到msg进行拆包  获得id 消息长度 放到msg对象种
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack msg  error ", err)
		}
		//根据datalen 再次读取data，放在msg.data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				fmt.Println("ConnId    ", c.ConnId, " read msg data error ", err)
				break
			}

		}
		msg.SetData(data)

		//得到当前链接request
		req := Request{
			conn: c,
			msg:  msg,
		}
		//执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}

//停止连接 结束当前连接
func (c *Connection) Stop() {
	fmt.Printf("conn stop connId = %d", c.ConnId)
	if c.isClose == true {
		return
	}
	c.isClose = true
	//关闭socketl链接
	c.Conn.Close()
	//h回收资源
	close(c.ExitChan)

}

//获取当前链接的socket conn
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接模块的id
func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

//获取远程客户端 tcp状态
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//提供sendmsg方法 把发往客户端的数据进行封包处理 再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClose == true {
		return errors.New("connection cloese when send msg")
	}
	//将data进行封包 msg datalen/msgid
	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err == nil {
		return errors.New("connection cloese when package msg")
	}
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("write msg id", msgId, " error ", err)
	}

	return nil
}
func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) ziface.IConnection {
	c := &Connection{
		Conn:     conn,
		ConnId:   connId,
		isClose:  false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}
