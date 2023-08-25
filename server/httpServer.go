package server

import (
	"database/sql"
	"log"

	"codeid.northwind/controllers"
	"codeid.northwind/repositories"
	"codeid.northwind/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config     *viper.Viper
	router     *gin.Engine
	controller *controllers.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	// categoryRepository := repositories.NewCategoryRepository(dbHandler)
	// categoryService := services.NewCategoryService(categoryRepository)
	// categoryController := controllers.NewCategoryController(categoryService)

	repositoryManager := repositories.NewRepositoryManager(dbHandler)

	serviceManager := services.NewServiceManager(repositoryManager)

	controllerManager := controllers.NewControllerManager(serviceManager)

	router := InitRouter(controllerManager)

	return HttpServer{
		config:     config,
		router:     router,
		controller: controllerManager,
	}
}

// running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP SERVER : %v", err)
	}
}
