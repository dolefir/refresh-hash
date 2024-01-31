package services

import (
	"context"

	"github.com/dolefir/refresh-hash/models"
)

// Hash is the service interface that
// describes business logic for working with hash
type Hash interface {
	Get(ctx context.Context) (*models.Hash, error)
	Refresh(ctx context.Context) error
}
