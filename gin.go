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
/*	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})*/
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "resp host:" + GetHostName() + "\n",
		})
	})
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "time:" + GetCurrentTime() + "\n",
		})

		// Handle all requests using net/http
		http.Handle("/", r)
	})

	fmt.Printf("===============port: %v============\n", port)
	r.Run(":"+port) // listen and serve on 0.0.0.0:8080
}