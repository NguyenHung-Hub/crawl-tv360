package db

import (
	"context"
	"errors"
	"tv360/src/models"

	"gorm.io/gorm"
)

func (q *Queries) IsExistsContent360(ctx context.Context, id string, tblName string) bool {
	var data models.Content
	err := q.db.Table(tblName).Where("contentId = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		} else {
			return true
		}

	}

	if data.ContentID != "" && data.Name != "" {
		return true
	}

	return false
}

func (q *Queries) CreateContent(ctx context.Context, data models.Content2) error {
	err := q.db.Create(data).Error
	return err
}
func (q *Queries) CreateContentTV(ctx context.Context, data models.ContentTV) error {
	err := q.db.Create(data).Error
	return err
}
