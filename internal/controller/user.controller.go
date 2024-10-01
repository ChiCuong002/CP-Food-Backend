package controller

import (
	"fmt"
	"food-recipes-backend/internal/services"
	"food-recipes-backend/internal/vo"
	"food-recipes-backend/pkg/response"
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
		response.ErrorResponse(c, response.ErrorCodeParamInvalid)
		return
	}
	fmt.Sprintf(`mail: %s, password: %s`, params.Email, params.Password)
	uc.userService.Register(params.Email, params.Password)
	response.SuccessResponse(c, response.SuccessCode, nil)
}
