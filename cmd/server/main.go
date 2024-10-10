package main

import (
	"fmt"
	"food-recipes-backend/global"
	"food-recipes-backend/internal/initialize"

	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "food-recipes-backend/cmd/swag/docs"
)
// @title           API Food Recipes
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/ChiCuong002/CP-Food-Backend

// @contact.name   Bui Chi Cuong
// @contact.url    https://github.com/ChiCuong002/CP-Food-Backend
// @contact.email  buichicuong6110@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1/api
// @schemes   http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true)))
	port := global.Config.Server.Port
	r.Run(fmt.Sprintf(":%s", port))
}