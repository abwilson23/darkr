package services

import "github.com/labstack/echo/v4"

type AuthenticationService interface {
	IsUserAuthenticated(c echo.Context) bool
}

func isAuthenticated() bool {
	return false
}
