package tasks

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	/////// Tasks
	r.GET("/tasks")
	r.POST("/tasks")
	r.GET("/tasks/:id")
	r.GET("/tasks/employee/:id")
	r.GET("/tasks/project/:id")
	r.PATCH("/tasks/:id")
	r.DELETE("/tasks/:id")
}
