package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tessyelots/ambulance-webapi/api"
	"github.com/tessyelots/ambulance-webapi/internal/ambulance_wl"
	"log"
	"os"
	"strings"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("AMBULANCE_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") { // case insensitive comparison
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	// request routings
	handleFunctions := &ambulance_wl.ApiHandleFunctions{
		AmbulanceConditionsAPI:  ambulance_wl.NewAmbulanceConditionsApi(),
		AmbulanceWaitingListAPI: ambulance_wl.NewAmbulanceWaitingListApi(),
	}
	ambulance_wl.NewRouterWithGinEngine(engine, *handleFunctions)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
