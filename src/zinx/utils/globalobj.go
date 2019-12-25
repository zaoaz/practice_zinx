package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx/ziface"
)

type GlobalObj struct {
	//server
	TcpServer ziface.IServer //当前zinx 全局server对象
	Host      string         //d当前监听的ip
	TcpPort   int            //当前主机监听的端口号
	Name      string         //当前服务器名称
	//zinx
	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

//定义一个全局对外GloabalObj对象
var GlobalObject *GlobalObj

//从zinx.json 加载自定义参数
func (g *GlobalObj) reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("read jso conf/zinx.json err ", err)
	}
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}
}

//import时自动调用init方法
//初始化当前GlobalObj
func init() {
	//默认值
	GlobalObject = &GlobalObj{
		TcpServer:      nil,
		Host:           "0.0.0.0",
		TcpPort:        8999,
		Name:           "zinx server app",
		Version:        "0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	GlobalObject.reload()
}
