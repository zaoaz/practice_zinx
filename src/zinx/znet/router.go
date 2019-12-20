package znet

import "zinx/ziface"

//发现router时 嵌入router基类  然后根据需要对基类进行重写
type BaseRouter struct {
}

//baserouter的方法都为空
//因为有的router不希望有prehandle posthandle
//所以router全部继承baserouter的好处就是，只需要实现想覆盖的handle
//处理connection 之前的hook
func (r *BaseRouter) PreHandle(request ziface.IRquest) {
}

//处理conn业务主方法hook
func (r *BaseRouter) Handle(request ziface.IRquest) {

}

//处理conn业务之后hook
func (r *BaseRouter) PostHandle(request ziface.IRquest) {

}
