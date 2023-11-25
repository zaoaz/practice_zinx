package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

func main() {
	//模拟客户端
	fmt.Println("client start ")
	//连接远程服务器
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("connect err ", err)
		return
	}
	fmt.Println("connect success")
	for {
		//发送封包的msg消息
		dp :=znet.NewDataPack()
		binaryMsg,err:=dp.Pack(znet.NewMsgPackage(0,[]byte("hi zinx v0.5 ")))
		if err!=nil{
			fmt.Println("pack err ",err)
			return
		}


		//连接调用write 写入数据
		if _, err := conn.Write(binaryMsg);err!=nil{
			fmt.Println("write err ", err)
			return
		}


		binaryHead := make([]byte, dp.GetHeadLen())
		//buf := make([]byte, 512)
		cnt, err := conn.Read(binaryHead)
		if err != nil {
			fmt.Println("read head error  ", err," ",cnt)
			return
		}

		//读取流中的head部分 得到id 和datalen
		msgHead,err := dp.UnPack(binaryHead)
		if err != nil {
			fmt.Println("unpack head error  ", err)
			return
		}
		if msgHead.GetMsgLen()>0{
			//根据datalen进行第二次读取，获得data
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _,err := io.ReadFull(conn,msg.Data);err!=nil{
				fmt.Println("unpack msg data error  ", err)
				return
			}
			fmt.Printf("recv id %d data %s \n",msg.GetMsgId(), string(msg.GetData()))
		}




		//阻塞 避免循环过快
		time.Sleep(1 * time.Second)
	}

}
