package model

import "time"

type BaseModel struct {
	ID 		  int64 		`gorm:"primary_key" json:"id"`
	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
	DeletedAt *time.Time `	json:"deleted_at"`
}
