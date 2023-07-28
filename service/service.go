package service

import (
	"github.com/brittonhayes/aos"
	"github.com/brittonhayes/aos/api"
)

type Service struct {
	repo aos.Repository
}

var _ api.ServerInterface = (*Service)(nil)

func New(repository aos.Repository) api.ServerInterface {
	return &Service{
		repo: repository,
	}
}
