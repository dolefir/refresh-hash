package inmem

import (
	"sync"

	"github.com/dolefir/refresh-hash/models"
)

// Repository holds methods for works with hash data inmem.
type Repository struct {
	hash models.Hash
	*sync.RWMutex
}

// NewRepository returns new hash Repository.
func NewRepository() *Repository {
	return &Repository{
		RWMutex: new(sync.RWMutex),
	}
}

// Set the information record.
func (r *Repository) Set(h *models.Hash) error {
	r.RWMutex.Lock()
	r.hash.ID = h.ID
	r.hash.Datatime = h.Datatime
	r.RWMutex.Unlock()

	return nil
}

// Get the read information.
// It is possible not to use Repository as
// a pointer it is only for the read.
func (r *Repository) Get() (*models.Hash, error) {
	r.RWMutex.RLock()
	defer r.RWMutex.RUnlock()

	return &r.hash, nil
}
