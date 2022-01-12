package service

import "github.com/korpgoodness/services.git/pkg/repository"

type Authorization interface {
}

type Servcie struct {
	Authorization
}

func NewService(repos *repository.Repository) *Servcie {
	return &Servcie{}
}
