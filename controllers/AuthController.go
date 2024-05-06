package controller

import (
	"darkr/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService services.AuthenticationService
}

func (a *AuthController) Index(c echo.Context) error {
	// First, determine if user is authenticated.
	if a.authService.IsUserAuthenticated(c) {
		return c.String(http.StatusOK, "You are authenticated")
	} else {
		return c.String(http.StatusUnauthorized, "You are not authenticated")
	}

}

func (a *AuthController) Login() {
	// Register a new user
}

func (a *AuthController) Callback() {
	// Register a new user
}
