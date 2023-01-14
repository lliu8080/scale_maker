package main

import (
	"log"

	"nuc.lliu.ca/gitea/app/scale_maker/internal/api"
	"nuc.lliu.ca/gitea/app/scale_maker/internal/config"
)

func main() {
	app := api.InitialSetup()
	port := config.AppConfig.Port
	log.Fatal(app.Listen(port)) // go run app.go -port=:3000
}
