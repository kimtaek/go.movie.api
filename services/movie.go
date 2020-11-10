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
	source.Data = models.GetMovies(&w)
	respond.Data(c, &source)
}

// @Tags movies
// @Summary 영화 추가
// @Router /movies [post]
// @Param   movie body models.Movie true "Movie info"
func PostMovie(c *gin.Context) {
	var source, request = respond.Source{}, models.Movie{}
	_ = c.Bind(&request)
	// validate
	source.Data = models.CreateMovie(&request)
	respond.Data(c, &source)
}

// @Tags movies
// @Summary 대표 영화 가져오기
// @Router /movies/{id} [get]
// @Param   id path string true "*"
func GetMovie(c *gin.Context) {
	var source = respond.Source{}
	source.Data = models.GetMovie(&models.Movie{
	})
	respond.Data(c, &source)
}

// @Tags movies
// @Summary 단일 영화 수정
// @Router /movies/{id} [put]
// @Param   id path int true "Movie.Id"
func PutMovie(c *gin.Context) {
	var source, request = respond.Source{}, models.Movie{}
	models.UpdateMovie(&request, &models.Movie{})

	_ = c.Bind(&request)
	request.ID = helper.String2Int64(c.Param("id"))
	respond.Data(c, &source)
}

// @Tags movies
// @Summary 영화 삭제
// @Router /movies/{id} [delete]
// @Param   id path int true "Movie.Id"
func DeleteMovie(c *gin.Context) {
	var source, where = respond.Source{}, models.Movie{}

	where.ID = helper.String2Int64(c.Param("id"))
	respond.Data(c, &source)
}
