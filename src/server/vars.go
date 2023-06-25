package server

import "github.com/gin-gonic/gin"

var ServerEngine *gin.Engine

var (
	getHandlers  = make(map[string]gin.HandlerFunc)
	postHandlers = make(map[string]gin.HandlerFunc)
)
