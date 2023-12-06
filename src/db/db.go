package db

import (
	"gorm.io/gorm"
)

func New(db *gorm.DB) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db *gorm.DB
}
