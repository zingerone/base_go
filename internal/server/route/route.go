package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zingerone/base_go/internal/server/handler"
	"github.com/zingerone/base_go/internal/service"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	h := handler.Handler{
		Service: service.NewService(),
	}

	r.GET("/hello", h.Hello)

	return r
}
