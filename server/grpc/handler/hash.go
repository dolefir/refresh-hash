package handler

import (
	"context"

	gen "github.com/dolefir/refresh-hash/gen/proto"
	"github.com/dolefir/refresh-hash/services"
)

type HashService struct {
	gen.UnimplementedHashServiceServer
	hashSrv services.Hash
}

func NewHashService(hashSrv services.Hash) *HashService {
	return &HashService{hashSrv: hashSrv}
}

func (hs HashService) GetHash(ctx context.Context, in *gen.GetHashRequest) (*gen.GetHashResponse, error) {
	resp, err := hs.hashSrv.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &gen.GetHashResponse{Uid: resp.ID}, nil
}
