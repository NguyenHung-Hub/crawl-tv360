package db

import (
	"context"
	"tv360/src/models"
)

type Querier interface {
	IsExistsContent360(ctx context.Context, id string, tablName string) bool
	CreateContent(ctx context.Context, data models.Content2) error
	CreateContentTV(ctx context.Context, data models.ContentTV) error
}

var _ Querier = (*Queries)(nil)
