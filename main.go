package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kimtaek/gamora/pkg/app"
	"github.com/kimtaek/gamora/pkg/db"
	"movie.api/docs"
	"movie.api/ent"
	"movie.api/routes"
)

func init() {
	app.Setup()
	db.Setup()
	ent.Setup()
}

func main() {
	defer db.CloseDB()
	defer ent.CloseDB()

	// setup swagger host
	docs.SwaggerInfo.Host = app.Config.AppHost

	gin.SetMode(app.Config.Mode)

	r := gin.Default()
	routes.RegisterRouter(r)
	_ = r.Run()
}
