package app

import (
	"jakmall/app/utils"
	"jakmall/config"
	"jakmall/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(utils.MsgGetEnvErr, err)
	}

	ginMode := os.Getenv("GIN_MODE")

	// Router
	router := config.InitRouter(ginMode)
	v1 := router.Group("/v1")

	// Routes
	routes.ProductRoute(v1)
	routes.ReviewRoute(v1)

	err = router.Run()
	if err != nil {
		log.Fatalf(utils.MsgRunServerErr)
	}
}
