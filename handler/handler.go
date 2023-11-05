package handler

import "github.com/gin-gonic/gin"

func SetUpServer() *gin.Engine {
	router := gin.Default()
	router.GET("/profile/:id", Profile)
	return router
}
