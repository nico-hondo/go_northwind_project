package server

import (
	"codeid.northwind/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *controllers.ControllerManager) *gin.Engine {
	router := gin.Default()

	categoryRoute := router.Group("/category")
	{
		//router endpoint
		categoryRoute.GET("", controllerMgr.CategoryController.GetListCategory)
		categoryRoute.GET("/:id", controllerMgr.CategoryController.GetCategory)
		categoryRoute.POST("/", controllerMgr.CategoryController.CreateCategory)

		categoryRoute.POST("/withproduct", controllerMgr.CategoryController.CreateCategoryWithProduct)

		categoryRoute.PUT("/:id", controllerMgr.CategoryController.UpdateCategory)
		categoryRoute.DELETE("/:id", controllerMgr.CategoryController.DeleteCategory)
	}

	userRoute := router.Group("/user")
	{
		//router for signUp
		userRoute.POST("/signup", controllerMgr.UsersController.Signup)
		userRoute.POST("/signin", controllerMgr.UsersController.Signin)
		userRoute.POST("/logout", controllerMgr.UsersController.Logout)
	}
	return router
}
