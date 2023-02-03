package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/endpoints/employees"
	"github.com/golangLearning/ginSQL/endpoints/ping"
	"github.com/golangLearning/ginSQL/endpoints/projects"
	"github.com/golangLearning/ginSQL/endpoints/tasks"
)

func main() {
	r := gin.Default()
	describeEndpoints(r)
	r.Run()
}

func describeEndpoints(r *gin.Engine) {
	r.GET("/ping", ping.Ping)

	// Defining the routes inside each package, to be easy to read and to mantain
	employees.Routes(r)
	projects.Routes(r)
	tasks.Routes(r)
}
