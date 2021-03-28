package main

//client
import (
	"fmt"
	"net"
	"os"
	"runtime"
)

const ip string = "localhost:13487"
const Maxsize = 256

// TCP client

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	// 1.连接服务端
	fmt.Println("start tcp")
	conn_service, err := net.Dial("tcp", ip)
	fmt.Println("tcpok")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2.给服务端发生消息
	// fmt.Printf("args[%v]=[%v]\n", k, v)
	v := os.Args[1]
	message := []byte(v)
	_, err = conn_service.Write(message)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Print("Send success!")
	}

	// 3.接收服务端消息
	var msg [Maxsize]byte // 声明一个接收信息的变量
	n, err := conn_service.Read(msg[:])
	if err != nil {
		fmt.Println("Read error of:", err)
		return
	}
	myfun(msg, n)
	// 4.关闭连接
	conn_service.Close()
}

//对数据进行处理
func myfun(msg [Maxsize]byte, n int) {
	fmt.Println("The service msg:", string(msg[:n]))
}

