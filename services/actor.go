package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kimtaek/gamora/pkg/db"
	"github.com/kimtaek/gamora/pkg/respond"
	"movie.api/models"
)

// @Tags actors
// @Summary 배우 리스트
// @Router /actors [get]
func GetActors(c *gin.Context) {
	var source, w = respond.Source{}, models.Actor{}
	source.Data = models.GetActors(&w)
	respond.Data(c, &source)
}

// @Tags actors
// @Summary 배우 리스트
// @Router /actors/{id}/paginate [get]
// @Param   id path string true "*"
func GetActorsPaginate(c *gin.Context) {
	var source, w, paging = respond.Source{}, models.Actor{}, db.PaginationParam{}
	_ = c.Bind(&paging)
	source.Data = models.GetActorsPaginate(&w, &paging)
	respond.Data(c, &source)
}
