package repository

import "github.com/dolefir/refresh-hash/models"

// Inmem is the interface that wraps works in-memory with hash.
type Inmem interface {
	Set(h *models.Hash) error
	Get() (*models.Hash, error)
}
