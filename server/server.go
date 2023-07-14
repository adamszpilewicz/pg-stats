package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"stats-pg/controllers"
	"stats-pg/repository"
)

type HttpServer struct {
	router     *gin.Engine
	controller *controllers.Controller
}

func InitHttpServer(dbHandler repository.Repository) HttpServer {
	controller := controllers.NewController(dbHandler)

	router := gin.Default()
	router.GET("/stats", controller.GetStats)
	router.GET("/stats/:dbName", controller.GetStatsByName)

	return HttpServer{
		router:     router,
		controller: controller,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(":8888")
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
