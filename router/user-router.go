package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/controllers"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/helpers"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/middlewares"
)

var (
	userHelpers    helpers.UserHelper         = helpers.NewUserHelper(userRepo)
	userController controllers.UserController = controllers.NewUserController(userHelpers, jwtHelpers)
)

func UserRouter() {
	router := gin.Default()

	userRoutes := router.Group("api/users", middlewares.AuthJWT(jwtHelper))
	{
		userRoutes.GET("/:userId", userController.Profile)
		userRoutes.PUT("/:userId", userController.Update)
	}
	router.Run()
}
