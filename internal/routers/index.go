package routers

import "food-recipes-backend/internal/routers/user"

type RouterGroup struct {
	User user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)