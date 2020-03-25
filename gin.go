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

	fmt.Printf("===============port: %v============\n", port)
	go r0.Run(":" + port) // listen and serve on 0.0.0.0:8080


    // 9090 for delay response
	var port1 = "9090" //os.Args[1]
	fmt.Printf("--------------port: %v\n", port1)
	r1 := gin.Default()
	r1.GET("/", func(context *gin.Context) {
		time.Sleep(time.Second * 110)
		context.JSON(http.StatusOK, gin.H{
			"message": "Delayed, Resp time:" + GetCurrentTime(),
		})
	})
	fmt.Printf("===============port: %v============\n", port1)
	go r1.Run(":" + port1) // listen and serve on 0.0.0.0:9090


	// 9090 for delay response
	var port2 = "9999" //os.Args[1]
	fmt.Printf("--------------port: %v\n", port2)
	r2 := gin.Default()
/*	r2.GET("/", func(context *gin.Context) {

	})*/
	fmt.Printf("===============port: %v============\n", port2)
	go r2.Run(":" + port2) // listen and serve on 0.0.0.0:9999


	//阻塞程序
	select {}
}

