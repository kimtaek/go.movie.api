package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kimtaek/gamora/pkg/app"
	"github.com/kimtaek/gamora/pkg/db"
	"movie.api/routes"
)

func init() {
	app.Setup()
	db.Setup()
}

func main() {
	defer db.CloseDB()

	gin.SetMode(app.Config.Mode)

	r := gin.Default()
	routes.RegisterRouter(r)
	_ = r.Run()
}
