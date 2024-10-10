package controller

import (
	"fmt"
	"food-recipes-backend/internal/services"
	"food-recipes-backend/internal/vo"
	"food-recipes-backend/pkg/response"
	apierror "food-recipes-backend/pkg/errors"

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

// User registration documentation
// @Summary      User Registration
// @Description  Register a new user
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body vo.UserRegistrationRequest true "User registration request"
// @Success      200  {object} vo.UserRegisterResponse
// @Failure      400  {object} apierror.APIError
// @Failure		 500  {object} apierror.APIError
// @Router       /user/register [post]
func (uc *UserController) Register(c *gin.Context) error {
	var params vo.UserRegistrationRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		return apierror.InvalidRequestData()
	}
	fmt.Printf(`service: %s, %s, %s\n`, params.Name, params.Email, params.Password)
	result, err := uc.userService.Register(c.Request.Context(), params.Name, params.Email, params.Password)
	if err != nil {
		return err
	}
	response.SuccessResponse(c, response.SuccessCode, result)
	return nil
}

// User login documentation
// @Summary      User Login
// @Description  After registration, user can login, receive access token and refresh token
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        payload body vo.UserLoginRequest true "User login request"
// @Success      200  {object} vo.UserRegisterResponse
// @Failure      400  {object} apierror.APIError
// @Failure		 500  {object} apierror.APIError
// @Router       /user/login [post]
func (uc *UserController) Login(c *gin.Context) error {
	var params vo.UserLoginRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		return apierror.InvalidRequestData()
	}
	fmt.Printf(`mail: %s, password: %s `, params.Email, params.Password)
	result, err := uc.userService.Login(c.Request.Context(), params.Email, params.Password)
	if err != nil {
		return err
	}
	response.SuccessResponse(c, response.SuccessCode, result)
	return nil
}

// @Summary      User Logout
// @Description  Logout user, remove refresh token from database
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        X-Client-Id header string true "Client ID"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  apierror.APIError
// @Failure      401  {object}  apierror.APIError
// @Failure      500  {object}  apierror.APIError
// @Security     BearerAuth
// @Router       /user/logout [post]
func (uc *UserController) Logout(c *gin.Context) error {
	userID, exists := c.Get("userId")
	fmt.Println("userID: ", userID)
	if !exists {
		return apierror.NewAPIError(401, "user not authenticated")
	}
	err := uc.userService.Logout(c.Request.Context(), userID.(int))
	if err != nil {
		return err
	}
	response.SuccessResponse(c, response.SuccessCode, nil)
	return nil
}

// @Summary      User Refresh Token
// @Description  Refresh user token, return new access token and refresh token
// @Tags         accounts management
// @Accept       json
// @Produce      json
// @Param        X-Client-Id header string true "Client ID"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  apierror.APIError
// @Failure      401  {object}  apierror.APIError
// @Failure      500  {object}  apierror.APIError
// @Security     BearerAuth
// @Router       /user/refresh-token [post]
func (uc *UserController) RefreshToken(c *gin.Context) error {
	refreshToken, _ := c.Get("refreshToken")
	userId, _ := c.Get("userId")
	fmt.Printf("refresh token: %s, user id: %d  \n", refreshToken, userId)
	userIdInt, ok := userId.(int)
	if !ok {
		return apierror.NewAPIError(400, "invalid user ID type")
	}
	refreshTokenStr, ok := refreshToken.(string)
	if !ok {
		return apierror.NewAPIError(400, "invalid refresh token type")
	}
	data, err := uc.userService.RefreshToken(c.Request.Context(), userIdInt, refreshTokenStr)
	if err != nil {
		return err
	}
	response.SuccessResponse(c, response.SuccessCode, data)
	return nil
}
