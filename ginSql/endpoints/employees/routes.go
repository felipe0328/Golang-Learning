package employees

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	// /////// Employees
	r.GET("/employees", GetEmployees)
	r.POST("/employees")
	r.GET("/employees/:id")
	r.PATCH("/employees/:id")
	r.PUT("/employees/:id")
	r.DELETE("/employees/:id")
}
