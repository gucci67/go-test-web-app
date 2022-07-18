package api

import (
	"go-test-web-app/api/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	userController := controller.UserController{}

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.GET("/user", userController.GetUserList)
	router.GET("/user/add", func(ctx *gin.Context) {
		ctx.HTML(200, "user-add.html", gin.H{})
	})
	router.GET("/user/edit/:id", userController.DisplayEditUser)

	router.POST("/user/add", userController.CreateUser)
	router.POST("/user/edit", userController.UpdateUser)

	return router
}
