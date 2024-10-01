package initialize

import (
	"fmt"
	"food-recipes-backend/global"
	"food-recipes-backend/internal/routers"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDatabase()
	r := routers.InitializeRoutes()
	port := global.Config.Server.Port
	fmt.Println("port " + port)
	r.Run(fmt.Sprintf(":%s", port))
}
