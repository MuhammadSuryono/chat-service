package server

import (
	"chat-service/models/response"
	"chat-service/system"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var config *gin.Engine

func ConfigServer() *gin.RouterGroup {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed, response.FailureResponse(
			response.METHOD_NOT_ALLOWED_STATUS,
			response.METHOD_NOT_ALLOWED_MESSAGE,
			nil,
		))
		return
	})

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, response.FailureResponse(
			response.NOT_FOUND_STATUS,
			response.NOT_FOUND_MESSAGE,
			nil,
		))
		return
	})

	config = r

	api := r.Group("api/v1/", func(context *gin.Context) {
		system.Context = context
	})
	api.GET("/", func(context *gin.Context) {
		context.JSON(200, response.SuccessResponse(false, "CHAT GROUP SERVICE API v.0.0.1", nil))
	})

	return api
}

func RunServer() {
	_ = config.Run("0.0.0.0:8081")
}
