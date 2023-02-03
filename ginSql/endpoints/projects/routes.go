package projects

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	//////// Projects
	r.GET("/projects")
	r.GET("/projects/:id")
	r.GET("/projects/task/:id")
	r.GET("/projects/employee/:id")
	r.POST("/projects")
	r.PATCH("/projects/:id")
	r.DELETE("/projects/:id")
}
