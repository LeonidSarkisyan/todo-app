package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Message string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	switch message {
	case "EOF":
		c.AbortWithStatusJSON(statusCode, Error{"отсутствует тело запроса"})
		return
	case "sql: no rows in result set":
		c.AbortWithStatusJSON(statusCode, Error{"не найдено"})
	default:
		c.AbortWithStatusJSON(statusCode, Error{message})
	}
}
