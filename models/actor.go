package models

import "github.com/kimtaek/gamora/pkg/db"

type Actor struct {
	db.Model
	Name    string    `json:"name"`
}

// Index
func GetActors(w *Actor) []*Actor {
	var m []*Actor
	if err := db.Connection().Where(w).Find(&m).Error; err == nil {
		return m
	}
	return nil
}

// Show
func GetActor(w *Actor) *Actor {
	var m Actor
	if err := db.Connection().Where(w).First(&m).Error; err == nil {
		return &m
	}
	return nil
}

// Create
func CreateActor(d *Actor) *Actor {
	if err := db.Connection().Create(d).Error; err == nil {
		return d
	}
	return nil
}

// Update
func UpdateActor(w *Actor, d *Actor) *Actor {
	var m Actor
	if err := db.Connection().Model(&m).Where(w).Update(d).Error; err == nil {
		return &m
	}
	return nil
}

// Delete
func DeleteActor(w *Actor) bool {
	var m Actor
	if err := db.Connection().Where(w).Delete(&m).Error; err == nil {
		return true
	}
	return false
}
