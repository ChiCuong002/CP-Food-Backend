package controller

import (
	"encoding/json"
	"fmt"
	"food-recipes-backend/global"
	"food-recipes-backend/internal/queries"
	"food-recipes-backend/internal/services"
	apierror "food-recipes-backend/pkg/errors"
	pkg "food-recipes-backend/pkg/json"
	"food-recipes-backend/pkg/pagination"
	"food-recipes-backend/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type AdminController struct {
	userService services.IUserService
	redisClient *redis.Client
}

func NewAdminController(
	userService services.IUserService,
	redisClient *redis.Client,
) *AdminController {
	return &AdminController{
		userService: userService,
		redisClient: redisClient,
	}
}

// @Summary      Detail User
// @Description  Detail infomation of a user
// @Tags         users management
// @Accept       json
// @Produce      json
// @Param        X-Client-Id header string true "Client ID"
// @Param        id path string true "user ID"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  apierror.APIError
// @Failure      500  {object}  apierror.APIError
// @Security     BearerAuth
// @Router       /admin/user/{id} [get]
func (ac *AdminController) DetailUser(c *gin.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return apierror.NewAPIError(400, "Invalid user ID")
	}
	cachedKey := fmt.Sprintf("user:%d", id)
	cachedData, err := ac.redisClient.Get(c, cachedKey).Result()
	if err == redis.Nil {
		// Key does not exist
		fmt.Println("key does not exist")
		user, err := ac.userService.GetUserByID(c, int32(id))
		if err != nil {
			return apierror.NewAPIError(500, "Failed to get user")
		}
		// Cache user data
		jsonData, err := pkg.JSONMarshal(user)
		if err != nil {
			fmt.Println("failed to marshal user data", err)
			global.Logger.Error("failed to marshal user data", zap.Error(err))
			return apierror.NewAPIError(500, "Server error")
		}
		fmt.Println("caching user data")
		err = ac.redisClient.Set(c, cachedKey, jsonData, time.Minute*10).Err()
		if err != nil {
			fmt.Println("failed to cached user data", err)
			global.Logger.Error("failed to cached user data", zap.Error(err))
		}
		response.SuccessResponse(c, response.SuccessCode, user)
		return nil
	} else if err != nil {
		fmt.Println("server err: ", err)
		global.Logger.Error("server err", zap.Error(err))
		return apierror.NewAPIError(500, "Server error")
	}
	// Key exists
	fmt.Println("key does exist")
	var result queries.DetailUserRow
	err = json.Unmarshal([]byte(cachedData), &result)
	if err != nil {
		fmt.Println("server err: ", err)
		return apierror.NewAPIError(500, "Server error")
	}
	response.SuccessResponse(c, response.SuccessCode, result)
	return nil
}

// @Summary      List users
// @Description  Get list of all users
// @Tags         users management
// @Accept       json
// @Produce      json
// @Param        X-Client-Id header string true "Client ID"
// @Param        search query string false "Search by name"
// @Param        status query string false "Filter by status"
// @Param		 sort query string false "Sort by id"
// @Param        limit query int false "Limit"
// @Param        page query int false "Page"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  apierror.APIError
// @Failure      401  {object}  apierror.APIError
// @Failure      500  {object}  apierror.APIError
// @Security     BearerAuth
// @Router       /admin/users [get]
func (ac *AdminController) ListUsers(c *gin.Context) error {
	search := c.Query("search")
	status := c.Query("status")
	sort := c.Query("sort")
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // default limit
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil && page < 1 {
		page = 1
	}
	offset := (page - 1) * limit
	users, err := ac.userService.ListUser(c, queries.ListUsersParams{
		Search: search,
		Status: status,
		Sort:   sort,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return apierror.NewAPIError(500, "Failed to list users")
	}
	if len(users) == 0 {
		return apierror.NewAPIError(404, "No users found")
	}
	// Calculate total pages
	totalRows := users[0].TotalCount
	totalPages := (int32(totalRows) + int32(limit) - 1) / int32(limit)

	result := pagination.Pagination{
		Search:     search,
		Limit:      limit,
		Page:       offset,
		Sort:       sort,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Rows:       users,
	}
	response.SuccessResponse(c, response.SuccessCode, result)
	return nil
}
