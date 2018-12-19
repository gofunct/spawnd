package project

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"os"
	"path/filepath"
)

func (p *Project) CreateUsersProto() {
	mainTemplate := users
	serviceTemplate := userService
	data := make(map[string]interface{})

	mainScript, err := utils.ExecTemplate(mainTemplate, data)
	logging.IfErr("failed to execute template", err)
	os.MkdirAll("services/users", os.ModePerm)
	serviceScript, err := utils.ExecTemplate(serviceTemplate, data)
	logging.IfErr("failed to execute template", err)

	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath()+"/services/user", "user.proto"), mainScript)
	logging.IfErr("failed to write file", err)
	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath()+"/services/user", "service.go"), serviceScript)
	logging.IfErr("failed to write file", err)
}

var users = `syntax = "proto3";

package user;

service UserService {
  rpc CreateUser(CreateUserRequest) returns (CreateUserResponse) {}
  rpc GetUser(GetUserRequest) returns (GetUserResponse) {}
}

message CreateUserRequest {
  string name = 1;
}
message CreateUserResponse {
  User user = 1;
  string err_msg = 2;
}

message GetUserRequest {
  string id = 1;
}
message GetUserResponse {
  User user = 1;
  string err_msg = 2;
}

message User {
  string id = 1;
  string name = 2;
}
`
var userService = `package user

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
`
