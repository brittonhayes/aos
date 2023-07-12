package service

import (
	"github.com/brittonhayes/warhammer"
	"github.com/brittonhayes/warhammer/api"
)

type WarhammerService struct {
	repo warhammer.WarhammerRepository
}

var _ api.ServerInterface = (*WarhammerService)(nil)

func NewWarhammerService(repository warhammer.WarhammerRepository) api.ServerInterface {
	return &WarhammerService{
		repo: repository,
	}
}
