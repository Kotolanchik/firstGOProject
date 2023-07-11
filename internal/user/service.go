package user

import (
	"context"
	"firstGOProject/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, request CreateUserRequest) (user User, err error) {
	return
}
