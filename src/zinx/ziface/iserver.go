package ziface

//定义服务端接口
type IServer interface {
	//启动
	Start()
	//停止
	Stop()
	//运行
	Serve()
}
