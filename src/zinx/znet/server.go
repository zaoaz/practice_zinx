package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

//定义server模块
type Server struct {
	//name
	Name string
	//ip version
	IPVersion string
	//ip
	IP string
	//port
	Port int
	//当前server添加一个router server链接对应的处理业务
	Router ziface.IRouter
}

//---------------------
//定义当前客户端绑定的handler api（目前写死，后续修改成成从demo定义）
//func CalllBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
//	fmt.Printf("[conn handler] CalllBackToClient \n")
//	if _, err := conn.Write(data[:cnt]); err != nil {
//		fmt.Printf("write back buf err ", err)
//		return errors.New("CalllBackToClient error \n")
//	}
//	return nil
//}
//--------------------

//iserver 接口实现
func (s *Server) Start() {
	fmt.Printf("[start ] server listen at IP %s port %d is starting \n", s.IP, s.Port)
	go func() {
		//获取tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error :", err)
			return
		}
		//监听地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listener  :", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start zinx server success ", s.Name, " listening ......")
		////阻塞连接，等待客户端连接
		var cid uint32
		cid = 0
		for {
			//如有客户端访问 则返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP  err:", err)
				continue
			}
			//连接建立成功 ，回显客户端发送的信息(长度限制512字节)
			//将处理新链接业务方法和链接绑定，得到链接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			go dealConn.Start()
			//go func() {
			//	for{
			//		//读取客户端发送的数据
			//		buf:=make([]byte,512)
			//		cnt,err := conn.Read(buf)
			//		if err !=nil{
			//			fmt.Println("receive buff err ",err)
			//			continue;
			//		}
			//		fmt.Printf("recieve buf %s,cnt=%d \n",buf[:cnt],cnt)
			//		//回发到客户端
			//		if _,err:=conn.Write(buf[:cnt]);err !=nil{
			//			fmt.Println("write buf err  ",err)
			//			continue;
			//		}
			//	}
			//}()

		}

	}()

}

func (s *Server) Stop() {
	//TODO
}

func (s *Server) Serve() {
	//启动server服务功能
	s.Start()

	//TODO 做启动服务器之后额外功能

	//阻塞状态
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("add router success")
}

//初始化 server模块
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
