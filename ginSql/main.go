package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/endpoints/ping"
)

func main() {
	r := gin.Default()
	describeEndpoints(r)
	r.Run()
}

func describeEndpoints(r *gin.Engine) {
	r.GET("/ping", ping.Ping)
}
