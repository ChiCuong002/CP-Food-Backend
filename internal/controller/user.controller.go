package controller

import (
	"fmt"
	"food-recipes-backend/internal/services"
	"food-recipes-backend/internal/vo"
	"food-recipes-backend/pkg/response"
	//"net/http"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(
	userService services.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistrationRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrorCodeParamInvalid, err.Error())
		return
	}
	fmt.Printf(`service: %s, %s, %s\n`, params.Name, params.Email, params.Password)
	result, err := uc.userService.Register(c.Request.Context(), params.Name, params.Email, params.Password)
	if err != nil {
		response.ErrorResponse(c, response.ErrorInternalServer, err.Error())
		return
	}
	response.SuccessResponse(c, response.SuccessCode, result)
}

func (uc *UserController) Login(c *gin.Context) {
	var params vo.UserLoginRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrorInternalServer, err.Error())
		return
	}
	fmt.Printf(`mail: %s, password: %s `, params.Email, params.Password)
	result, err := uc.userService.Login(c.Request.Context(), params.Email, params.Password)
	if err != nil {
		response.ErrorResponse(c, response.ErrorInternalServer, err.Error())
		return
	}
	response.SuccessResponse(c, response.SuccessCode, result)
}