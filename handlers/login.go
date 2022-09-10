package handlers

import (
	"net/http"
	"test-case-backend-majoo/middlewares"
	"test-case-backend-majoo/models"
	"test-case-backend-majoo/repository"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var login models.LoginRequest
	if err := c.Bind(&login); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	checkUser, err := repository.GetUserByUsername(login.Username, login.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if checkUser.ID == 0 {
		return c.String(http.StatusNotFound, "Username or Password Not Found!!!")
	}

	token, err := middlewares.CreateToken(checkUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error Generate Token")
	}

	resp := models.LoginResponse{
		Token: token,
	}

	return c.JSON(http.StatusCreated, resp)
}
