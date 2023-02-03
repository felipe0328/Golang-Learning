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

	///////// Employees
	r.GET("/employees") // Get List of employees
	r.POST("/employees")
	r.GET("/employees/:id")
	r.PATCH("/employees/:id")
	r.PUT("/employees/:id")
	r.DELETE("/employees/:id")

	//////// Projects
	r.GET("/projects")
	r.GET("/projects/:id")
	r.GET("/projects/task/:id")
	r.GET("/projects/employee/:id")
	r.POST("/projects")
	r.PATCH("/projects/:id")
	r.DELETE("/projects/:id")

	/////// Tasks
	r.GET("/tasks")
	r.POST("/tasks")
	r.GET("/tasks/:id")
	r.GET("/tasks/employee/:id")
	r.GET("/tasks/project/:id")
	r.PATCH("/tasks/:id")
	r.DELETE("/tasks/:id")
}
