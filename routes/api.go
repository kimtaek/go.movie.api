package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	r.GET("docs", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/docs/index.html")
	})
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// gorm
	movies := r.Group("movies")
	{
		movies.GET("", services.GetMovies)
		movies.GET(":id/paginate", services.GetMoviesPaginate)
		movies.GET(":id/actors", services.GetMovieActors)
		movies.GET(":id", services.GetMovie)
		movies.POST("", services.PostMovie)
		movies.PUT(":id", services.PutMovie)
		movies.DELETE(":id", services.DeleteMovie)
	}

	actors := r.Group("actors")
	{
		actors.GET("", services.GetActors)
		actors.GET(":id/paginate", services.GetActorsPaginate)
		//actors.GET(":id", services.GetActor)
		//actors.POST("", services.PostActor)
		//actors.PUT(":id", services.PutActor)
		//actors.DELETE(":id", services.DeleteActor)
	}

	// ent orm
	users := r.Group("users")
	{
		users.GET("", services.GetUsers)
		users.GET(":id/paginate", services.GetUsersPaginate)
		users.GET(":id/actors", services.GetUserActors)
		users.GET(":id", services.GetUser)
		users.POST("", services.PostUser)
		users.PUT(":id", services.PutUser)
		users.DELETE(":id", services.DeleteUser)
	}
}
