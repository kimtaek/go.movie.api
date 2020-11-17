package routes

import (
	"github.com/gin-gonic/gin"
	"movie.api/middleware"
	"movie.api/services"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	r.Use(middleware.CORS())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"code": http.StatusNotFound, "message": "not found"})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": http.StatusOK, "message": "Hi. :D"})
	})

	movies := r.Group("movies")
	{
		movies.GET("", services.GetMovies)
		movies.GET(":id/paginate", services.GetMovies)
		//movies.GET(":id/actors", services.GetMovies)
		//movies.GET(":id/actors/:id/paginate", services.GetMovies)
		//movies.GET(":id/actors/:id", services.GetMovies)
		movies.GET(":id", services.GetMovie)
		movies.POST("", services.PostMovie)
		movies.PUT(":id", services.PutMovie)
		movies.DELETE(":id", services.DeleteMovie)
	}
}
