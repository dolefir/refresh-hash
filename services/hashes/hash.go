package hashes

import (
	"context"
	"time"

	"github.com/dolefir/refresh-hash/logger"
	"github.com/dolefir/refresh-hash/models"
	"github.com/dolefir/refresh-hash/repository"
	"github.com/google/uuid"
)

// Service handle common for hash operations.
type Service struct {
	hashRepo repository.Inmem
	log      logger.Logger
}

// NewService creates new hash service.
func NewService(hashRepo repository.Inmem, log logger.Logger) *Service {
	return &Service{
		hashRepo: hashRepo,
		log:      log,
	}
}

// Get returns hash.
func (s Service) Get(ctx context.Context) (*models.Hash, error) {
	s.log.Debug("service.Hash.Get: get hash")
	hash, err := s.hashRepo.Get()
	if err != nil {
		s.log.Errorf("service.Hash.Get: %s", err)
		return nil, err
	}
	if hash.ID == "" {
		// Create a new hash if not exist.
		if err := s.Refresh(ctx); err != nil {
			return nil, err
		}
	}

	s.log.Debugf("service.Hash.Get: hash exist %s", hash)

	return hash, nil
}

// Refresh to create/update hash inmem.
func (s Service) Refresh(ctx context.Context) error {
	s.log.Debug("service.Hash.Refresh: refresh hash")

	hash := &models.Hash{
		ID:       uuid.New().String(),
		Datatime: time.Now(),
	}

	if err := s.hashRepo.Set(hash); err != nil {
		s.log.Errorf("service.Hash.Refresh: %s", err)
		return err
	}

	s.log.Debug("service.Hash.Refresh: hash updated")

	return nil
}
