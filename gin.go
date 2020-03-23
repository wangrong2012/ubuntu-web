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

	r.GET("/delay", func(context *gin.Context) {
		time.Sleep(time.Second * 65)
		context.JSON(http.StatusOK, gin.H{
			"message": "Delayed, Resp time:" + GetCurrentTime(),
		})
	})

	fmt.Printf("===============port: %v============\n", port)
	go r.Run(":" + port) // listen and serve on 0.0.0.0:8080

	//阻塞程序
	select {}
}

