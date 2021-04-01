package models

import "time"

// Model global basic model struct
type Model struct {
	ID        uint64     `form:"id" json:"id" gorm:"primary_key"`
	CreatedBy uint64     `json:"-"`
	UpdatedBy uint64     `json:"-"`
	DeletedBy uint64     `json:"-"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
