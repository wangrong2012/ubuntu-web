package main

import (
	"fmt"
	"log"
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

var delaySecond int = 65

func main() {
	var port = "8080" //os.Args[1]
	fmt.Printf("--------------port: %v\n", port)
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "resp host:" + GetHostName() ,
		})
	})
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "time:" + GetCurrentTime(),
		})
	})

	fmt.Printf("===============port: %v============\n", port)
	go r.Run(":" + port) // listen and serve on 0.0.0.0:8080



	//延迟测试
	var port1 = "9090"

	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", myHandler1)
	fmt.Printf("===============port: %v========Resp delay %vs====\n", port1, delaySecond)
	go http.ListenAndServe(":9090", mux1)


	//阻塞程序
	select {}
}

func myHandler1(res http.ResponseWriter, req *http.Request)  {

	time.Sleep(time.Second * time.Duration(delaySecond))
	//fmt.Println(req.URL, req.Host)
	fmt.Println("time:" + GetCurrentTime())
}