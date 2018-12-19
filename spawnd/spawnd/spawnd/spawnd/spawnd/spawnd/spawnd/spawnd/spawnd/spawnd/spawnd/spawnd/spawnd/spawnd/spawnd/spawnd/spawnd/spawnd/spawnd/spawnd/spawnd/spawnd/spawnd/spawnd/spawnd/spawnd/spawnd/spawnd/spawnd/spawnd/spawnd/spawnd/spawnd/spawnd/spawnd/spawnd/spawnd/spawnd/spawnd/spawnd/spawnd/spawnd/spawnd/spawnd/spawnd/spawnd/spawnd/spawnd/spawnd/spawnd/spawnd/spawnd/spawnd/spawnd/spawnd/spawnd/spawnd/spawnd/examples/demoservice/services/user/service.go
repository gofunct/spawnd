package user

import (
	"fmt"

	"golang.org/x/net/context"
)

type Service struct{}

func New() UserServiceServer { return &Service{} }

func (svc *Service) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) GetUser(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
