package models

import "github.com/kimtaek/gamora/pkg/db"

type Actor struct {
	Model
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

// GetActors
func GetActors(w *Actor) []*Actor {
	var m []*Actor
	if err := db.Connection().Where(w).Find(&m).Error; err == nil {
		return m
	}
	return nil
}

// GetActorsPaginate
func GetActorsPaginate(w *Actor, pp *db.PaginationParam) *db.Pagination {
	var m []*Actor
	d := db.Connection().Where(w).Order("`id` DESC").Limit(pp.Limit).Find(&m)
	pp.DB = d
	return db.Paginate(pp, &m)
}
