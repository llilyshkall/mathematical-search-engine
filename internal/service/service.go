package service

import repo "github.com/llilyshkall/mathematical-search-engine/internal/repository/postgres"

type Service struct {
	db repo.ExpressionRepository
}

func NewService(
	db repo.ExpressionRepository,
) *Service {
	return &Service{db: db}
}
