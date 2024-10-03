//go:build-wireinject
package wire

import (
	"food-recipes-backend/global"
	"food-recipes-backend/internal/controller"
	"food-recipes-backend/internal/repo"
	"food-recipes-backend/internal/services"

	"github.com/google/wire"
)

func initUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		global.ProvideDB,
		repo.NewKeyRepository,
		repo.NewUserRepository,
		services.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}