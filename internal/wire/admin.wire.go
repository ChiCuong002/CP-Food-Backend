package wire

import (
	"food-recipes-backend/global"
	"food-recipes-backend/internal/controller"
	"food-recipes-backend/internal/repo"
	"food-recipes-backend/internal/services"

	"github.com/google/wire"
)

func initAdminRouterHandler() (*controller.AdminController, error) {
	wire.Build(
		global.ProvideRedis,
		global.ProvideDB,
		repo.NewKeyRepository,
		repo.NewUserRepository,
		services.NewUserService,
		controller.NewAdminController,
	)
	return new(controller.AdminController), nil
}