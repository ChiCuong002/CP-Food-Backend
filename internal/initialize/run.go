package initialize

import "fmt"

import "food-recipes-backend/global"

func Run() {
	LoadConfig()
	InitLogger()
	InitDatabase()
	r := Initialize()
	port := global.Config.Server.Port
	fmt.Println("port " + port)
	r.Run(fmt.Sprintf(":%s", port))
}
