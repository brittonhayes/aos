package service

import (
	"github.com/brittonhayes/warhammer"
	"github.com/brittonhayes/warhammer/api"
)

type WarhammerService struct {
	repo warhammer.WarhammerRepository
}

var _ api.ServerInterface = (*WarhammerService)(nil)

const (
	ErrUnitNotFound   = "Unit not found with this ID"
	ErrUnitsNotFound  = "No units found"
	ErrArmyNotFound   = "Army not found with this ID"
	ErrArmiesNotFound = "No armies found"
)

func NewWarhammerService(repository warhammer.WarhammerRepository) api.ServerInterface {
	return &WarhammerService{
		repo: repository,
	}
}
