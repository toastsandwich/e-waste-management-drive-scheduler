package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toastsandwich/ewaste-collection-caller/web/handlers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/assets", "assets")

	e.GET("/", handlers.Home)
	e.GET("/about", handlers.About)
	e.GET("/drives", handlers.UpcomingDrives)
	e.GET("/schedule-drive", handlers.ScheduleDrive)

	e.POST("/submit-drive", handlers.ScheduleDriveForm)

	e.Logger.Fatal(e.Start(":8080"))

}
