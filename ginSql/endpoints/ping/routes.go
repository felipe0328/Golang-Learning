package ping

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.GET("/ping", Ping)
}
