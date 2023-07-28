package service

import (
	"github.com/brittonhayes/aos"
	"github.com/brittonhayes/aos/api"
)

type aosService struct {
	repo aos.HammerRepository
}

var _ api.ServerInterface = (*aosService)(nil)

func NewaosService(repository aos.HammerRepository) api.ServerInterface {
	return &aosService{
		repo: repository,
	}
}
