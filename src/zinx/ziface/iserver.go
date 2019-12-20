package ziface

//定义服务端接口
type IServer interface {
	//启动
	Start()
	//停止
	Stop()
	//运行
	Serve()
	//路由功能；给当前服务注册路由方法，供客户端链接处理时使用
	AddRouter(router IRouter)
}
