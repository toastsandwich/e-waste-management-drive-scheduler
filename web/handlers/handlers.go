package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/ewaste-collection-caller/db"
	"github.com/toastsandwich/ewaste-collection-caller/utils"
	"github.com/toastsandwich/ewaste-collection-caller/views/pages"
)

func Home(c echo.Context) error {
	return utils.Render(c, http.StatusOK, pages.HomePage())
}

func About(c echo.Context) error {
	return utils.Render(c, http.StatusOK, pages.AboutPage())
}

func UpcomingDrives(c echo.Context) error {
	return utils.Render(c, http.StatusOK, pages.UpcomingDrives(db.ScheduledDrives))
}

func ScheduleDrive(c echo.Context) error {
	return utils.Render(c, http.StatusOK, pages.ScheduleDrive())
}

func ScheduleDriveForm(c echo.Context) error {
	driveName := c.FormValue("drive-name")
	address := c.FormValue("address")
	name := c.FormValue("name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")

	drive := db.Drive{
		Name:    driveName,
		Address: address,
		Customer: db.Customer{
			Name:  name,
			Email: email,
			Phone: phone,
		},
	}
	db.ScheduledDrives = append(db.ScheduledDrives, drive)
	return utils.Render(c, http.StatusOK, pages.UpcomingDrives(db.ScheduledDrives))
}
