package routers

import (
	"food-recipes-backend/internal/routers/admin"
	"food-recipes-backend/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)