package rest

import (
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (h *Handler) loggingMiddleware(c *gin.Context) {
	logrus.Infof("[%s] - %s", c.Request.Method, c.Request.RequestURI)
}