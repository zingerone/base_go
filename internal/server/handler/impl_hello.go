package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Hello(c *gin.Context) {
	output := h.Service.HelloService.GetHello()
	c.JSON(200, output)
}
