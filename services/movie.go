package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kimtaek/gamora/pkg/helper"
	"github.com/kimtaek/gamora/pkg/respond"
	"movie.api/models"
)

// @Tags movies
// @Summary 영화 리스트
// @Router /movies [get]
func GetMovies(c *gin.Context) {
	var source, w = respond.Source{}, models.Movie{}
	source.Data = models.GetMoviesWithActors(&w)
	respond.Data(c, &source)
}

// @Tags movies
// @Summary 영화 추가
// @Router /movies [post]
// @Param   movie body models.Movie true "Movie info"
func PostMovie(c *gin.Context) {
	var source, request = respond.Source{}, models.Movie{}
	_ = c.Bind(&request)
	// need's check validate

	source.Data = models.CreateMovie(&request)
	respond.Data(c, &source)
}

// @Tags movies
// @Summary 영화 가져 오기
// @Router /movies/{id} [get]
// @Param   id path int true "Movie.Id"
func GetMovie(c *gin.Context) {
	var source = respond.Source{}
	// need's check validate

	source.Data = models.GetMovieByID(helper.String2Int64(c.Param("id")))
	respond.Data(c, &source)
}

// @Tags movies
// @Summary 단일 영화 수정
// @Router /movies/{id} [put]
// @Param   id path int true "Movie.Id"
func PutMovie(c *gin.Context) {
	var source, where, request = respond.Source{}, models.Movie{}, models.Movie{}
	_ = c.Bind(&request)
	// need's check validate

	where.ID = helper.String2Int64(c.Param("id"))
	source.Data = models.UpdateMovie(&where, &request)
	respond.Data(c, &source)
}

// @Tags movies
// @Summary 영화 삭제
// @Router /movies/{id} [delete]
// @Param   id path int true "Movie.Id"
func DeleteMovie(c *gin.Context) {
	var source, where = respond.Source{}, models.Movie{}
	// need's check validate

	where.ID = helper.String2Int64(c.Param("id"))
	source.Data = models.DeleteMovie(&where)
	respond.Data(c, &source)
}
