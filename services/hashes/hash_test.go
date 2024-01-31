package hashes

import (
	"context"
	"reflect"
	"testing"

	"github.com/dolefir/refresh-hash/config"
	"github.com/dolefir/refresh-hash/logger"
	"github.com/dolefir/refresh-hash/models"
	"github.com/dolefir/refresh-hash/repository"
	"github.com/dolefir/refresh-hash/services/mock"
)

func TestService_Get(t *testing.T) {
	mockInmem := mock.NewInmemMock()
	type fields struct {
		hashRepo repository.Inmem
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Hash
		wantErr bool
	}{
		{
			name: "should returns hash",
			fields: fields{
				hashRepo: mockInmem,
			},
			args: args{
				ctx: context.Background(),
			},
			want: &models.Hash{
				ID: "996f2357-31af-4b1a-9889-a075be3de0a9",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewConfig("")
			log := logger.NewLogger((*logger.CFGLogger)(&cfg.Logger), nil)
			s := Service{
				hashRepo: tt.fields.hashRepo,
				log:      log,
			}
			got, err := s.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetWithError(t *testing.T) {
	mockInmem := mock.NewInmemErrMock()
	type fields struct {
		hashRepo repository.Inmem
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Hash
		wantErr bool
	}{
		{
			name: "should returns error",
			fields: fields{
				hashRepo: mockInmem,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewConfig("")
			log := logger.NewLogger((*logger.CFGLogger)(&cfg.Logger), nil)
			s := Service{
				hashRepo: tt.fields.hashRepo,
				log:      log,
			}
			got, err := s.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Refresh(t *testing.T) {
	mockInmem := mock.NewInmemMock()
	type fields struct {
		hashRepo repository.Inmem
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should returns ok",
			fields: fields{
				hashRepo: mockInmem,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewConfig("")
			log := logger.NewLogger((*logger.CFGLogger)(&cfg.Logger), nil)
			s := Service{
				hashRepo: tt.fields.hashRepo,
				log:      log,
			}
			if err := s.Refresh(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Service.Refresh() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_RefreshWithError(t *testing.T) {
	mockInmem := mock.NewInmemErrMock()
	type fields struct {
		hashRepo repository.Inmem
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should returns error",
			fields: fields{
				hashRepo: mockInmem,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewConfig("")
			log := logger.NewLogger((*logger.CFGLogger)(&cfg.Logger), nil)
			s := Service{
				hashRepo: tt.fields.hashRepo,
				log:      log,
			}
			if err := s.Refresh(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Service.Refresh() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
