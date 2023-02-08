package user

import "Bugs-Bunny/pkg/models"

type Service interface {
	GetUserByEmail(email string) (*models.User, error)
	ReplaceApiKey(UserId *uuid.UUID) (*models.ApiKey, error)
}

type userSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &userSvc{repo: r}
}
