package handler

import (
	"net/http"
	"tama_foundation/helper"
	"tama_foundation/users"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandlerService(userService users.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	/*
		- tangkap input dari user
		- map input user dari user ke struct RegisterUserInput
		- struct di atas bakal di passing ke service
		- dan nanti service bakal nge save ke repository
		- dan datanya akan muncul di database
	*/
	var input = users.RegisterUserInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed ", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return //agar eksekusi tidak lanjut ke bawah
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account Failed ", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return

	}

	formaterr := users.FormatUser(newUser, "token")

	response := helper.APIResponse("Account has been Registered", http.StatusOK, "success", formaterr)
	c.JSON(http.StatusOK, response)
}
