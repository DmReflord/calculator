package main

import (
	"calculator/internal/db"
	"calculator/internal/handlers"
	"calculator/internal/service"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database")
	}
	e := echo.New()
	calcRepo := service.NewCalculationRepository(database)
	calcService := service.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)
	e.Start("localhost:8080")
}
