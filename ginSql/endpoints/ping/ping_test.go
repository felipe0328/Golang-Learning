package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingEndpoint(t *testing.T) {
	router := gin.Default()
	Routes(router)

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(writer, req)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Equal(t, "Pong", writer.Body.String())
}
