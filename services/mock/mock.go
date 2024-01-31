package mock

import (
	"errors"

	"github.com/dolefir/refresh-hash/models"
	"github.com/dolefir/refresh-hash/repository"
)

type InmemMock struct {
	repository.Inmem
}

func NewInmemMock() *InmemMock {
	return &InmemMock{}
}

func (s InmemMock) Get() (*models.Hash, error) {
	return &models.Hash{ID: "996f2357-31af-4b1a-9889-a075be3de0a9"}, nil
}

func (s InmemMock) Set(h *models.Hash) error {
	return nil
}

type InmemErrMock struct {
	repository.Inmem
}

func NewInmemErrMock() *InmemErrMock {
	return &InmemErrMock{}
}

func (r InmemErrMock) Get() (*models.Hash, error) {
	return nil, errors.New("error")
}

func (s InmemErrMock) Set(h *models.Hash) error {
	return errors.New("error")
}
