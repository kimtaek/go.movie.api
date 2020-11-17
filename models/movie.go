package models

import "github.com/kimtaek/gamora/pkg/db"

type Movie struct {
	db.Model
	Title       string   `form:"title" json:"title"`
	Description string   `form:"description" json:"description"`
	Actors      []*Actor `json:"actors" gorm:"many2many:movie_actors;PRELOAD:false"`
}

// GetMovies get movies
func GetMovies(w *Movie) []*Movie {
	var m []*Movie
	if err := db.Connection().Where(w).Order("`id` DESC").Find(&m).Error; err == nil {
		return m
	}
	return nil
}

// GetMoviesWithActors
func GetMoviesWithActors(w *Movie) []*Movie {
	var m []*Movie
	if err := db.Connection().Where(w).Preload("Actors").Order("`id` DESC").Find(&m).Error; err == nil {
		return m
	}
	return m
}

// GetMoviesWithActorsPaginate
func GetMoviesWithActorsPaginate(w *Movie, pp *db.PaginationParam) *db.Pagination {
	var m []*Movie
	d := db.Connection().Where(w).Preload("Actors").Order("`id` DESC").Limit(pp.Limit).Find(&m)
	pp.DB = d
	return db.Paginate(pp, &m)
}

// GetMovieByID get single movie by id
func GetMovieByID(id uint64) *Movie {
	var m Movie
	if err := db.Connection().Preload("Actors").Where("id = ?", id).First(&m).Error; err == nil {
		return &m
	}
	return nil
}

// GetMovie get single movie by Movie struct
func GetMovie(w *Movie) *Movie {
	var m Movie
	if err := db.Connection().Where(w).First(&m).Error; err == nil {
		return &m
	}
	return nil
}

// CreateMovie create single movie
func CreateMovie(d *Movie) *Movie {
	if err := db.Connection().Create(d).Error; err == nil {
		return d
	}
	if d.ID == 0 {
		return nil
	}
	return nil
}

// UpdateMovie update single movie by Movie struct
func UpdateMovie(w *Movie, d *Movie) *Movie {
	var m Movie
	if err := db.Connection().Model(&m).Where(w).Update(d).Error; err == nil {
		return &m
	}
	return nil
}

// DeleteMovie delete single movie
func DeleteMovie(w *Movie) bool {
	var m Movie
	if err := db.Connection().Where(w).Delete(&m).Error; err == nil {
		return true
	}
	return false
}
