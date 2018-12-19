package session

import (
	"fmt"

	"golang.org/x/net/context"
)

type Service struct{}

func New() SessionServiceServer {
	return &Service{}
}

func (svc *Service) Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
