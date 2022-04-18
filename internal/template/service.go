package template

import (
	"context"
	"template/pkg/logging"
)

type service struct {
	storage Storage
	logger  logging.Logger
}

func NewService(userStorage Storage, logger logging.Logger) (Service, error) {
	return &service{
		storage: userStorage,
		logger:  logger,
	}, nil
}

type Service interface {
	Create(ctx context.Context, dto CreateUserHashedDTO) error
}

func (s service) Create(ctx context.Context, dto CreateUserHashedDTO) error {
	var user *User
	user = dto.NewUser()
	err := s.storage.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
