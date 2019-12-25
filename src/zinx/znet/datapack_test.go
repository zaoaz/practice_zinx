package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// datapack拆包封包单元测试
func TestDataPack(t *testing.T) {
	//模拟的服务器
	// 创建socketTcp Server
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	// 从客户端读取数据进行拆包
	if err != nil {
		fmt.Println("server listen err ", err)
		return
	}

	//创建一个go 负责从处理客户端业务

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept err ", err)
			}

			//处理客户端请求
			go func(conn net.Conn) {
				//处理客户端请求
				//拆包
				//定义拆包对象dp
				dp := NewDataPack()
				for {
					//第一次 从conne 读取到包的head
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read  head err ", err)
						break
					}
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("server unpack head err ", err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						// msg 是有数据的 需要进行第二次读取
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msgHead.GetMsgLen())
						//根据datalen长度再次从io流中读取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data err ", err)
							return
						}
						//完整消息已经读取完毕

						fmt.Println("--->Recv MsgId ", msg.Id, " datalen =", msg.GetMsgLen(), " data = ", string(msg.GetData()))
					}

				}

			}(conn)
		}
	}()

	//模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client DialTCP  err ", err)
		return
	}
	//创建封包对象dp
	dp := NewDataPack()

	//模拟粘包过程 封装两个msg一同发送
	//封装第一个msg1包
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1  err ", err)
		return
	}
	//封装第二个msg2包
	msg2 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'d', 'w', 'e', 'a', 'b'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack msg2  err ", err)
		return
	}
	//将两个包粘在一起
	sendData1 = append(sendData1, sendData2...)
	//一次性发送给服务端
	conn.Write(sendData1)
	//客户端阻塞
	select {}
}
