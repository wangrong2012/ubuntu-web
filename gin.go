package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)


func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return hostname
}

func GetCurrentTime() string {
	timer := time.Now()
	return timer.String()
}


func main() {
	var port = "8080" //os.Args[1]
	fmt.Printf("--------------port: %v\n", port)
	r0 := gin.Default()

	r0.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "resp host:" + GetHostName() ,
		})
	})
	r0.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "time:" + GetCurrentTime(),
		})
	})

	r0.GET("/delay", func(context *gin.Context) {
		time.Sleep(time.Second * 65)
		context.JSON(http.StatusOK, gin.H{
			"message": "Delayed, Resp time:" + GetCurrentTime(),
		})
	})

	r0.GET("/getHeader", HandleGetAllData)



	fmt.Printf("===============port: %v============\n", port)
	go r0.Run(":" + port) // listen and serve on 0.0.0.0:8080



	port2 := "9090"
	Start(port2)


	//阻塞程序
	select {}
}


func HandleGetAllData(c *gin.Context)  {
	//log.Print("handle log")
	body,_ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("---body/--- \r\n " + string(body))

	fmt.Printf("---header/--- \r\n")
	for k,v :=range c.Request.Header {
		fmt.Println(k,v)
	}
	//fmt.Println("header \r\n",c.Request.Header)

	c.JSON(200,gin.H{
		"receive":c.Request.Header,
	})

}

// 启动服务器
func Start(port string) {
	host := ":" + port

	// 获取tcp地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		log.Printf("resolve tcp addr failed: %v\n", err)
		return
	}

	// 监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Printf("listen tcp port failed: %v\n", err)
		return
	}

	// 建立连接池，用于广播消息
	conns := make(map[string]net.Conn)

	// 消息通道
	//messageChan := make(chan string, 10)

	// 广播消息
	//go BroadMessages(&conns, messageChan)

	// 启动
	for {
		fmt.Printf("listening port %s ...\n", port)
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Accept failed:%v\n", err)
			continue
		}

		// 把每个客户端连接扔进连接池
		conns[conn.RemoteAddr().String()] = conn
		fmt.Println(conns)

		// 处理消息
		//go Handler(conn, &conns, messageChan)
	}
}

// 向所有连接上的乡亲们发广播
func BroadMessages(conns *map[string]net.Conn, messages chan string) {
	for {

		// 不断从通道里读取消息
		msg := <-messages
		fmt.Println(msg)

		// 向所有的乡亲们发消息
		for key, conn := range *conns {
			fmt.Println("connection is connected from ", key)
			_, err := conn.Write([]byte(msg))
			if err != nil {
				log.Printf("broad message to %s failed: %v\n", key, err)
				delete(*conns, key)
			}
		}
	}
}




// 处理客户端发到服务端的消息，将其扔到通道中
func Handler(conn net.Conn, conns *map[string]net.Conn, messages chan string) {
	fmt.Println("connect from client ", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("read client message failed:%v\n", err)
			delete(*conns, conn.RemoteAddr().String())
			conn.Close()
			break
		}

		// 把收到的消息写到通道中
		recvStr := string(buf[0:length])
		messages <- recvStr
	}
}
