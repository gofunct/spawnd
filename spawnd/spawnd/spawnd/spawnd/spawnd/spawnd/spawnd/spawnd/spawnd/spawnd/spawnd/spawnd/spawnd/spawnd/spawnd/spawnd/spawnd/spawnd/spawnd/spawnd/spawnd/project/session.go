package project

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"os"
	"path/filepath"
)

func (p *Project) CreateSessionsProto() {
	mainTemplate := session
	serviceTemplate := sessionService
	data := make(map[string]interface{})

	mainScript, err := utils.ExecTemplate(mainTemplate, data)
	logging.IfErr("failed to execute template", err)
	os.MkdirAll("services/sessions", os.ModePerm)
	serviceScript, err := utils.ExecTemplate(serviceTemplate, data)
	logging.IfErr("failed to execute template", err)

	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath()+"/services/session", "session.proto"), mainScript)
	logging.IfErr("failed to write file", err)
	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath()+"/services/session", "service.go"), serviceScript)
	logging.IfErr("failed to write file", err)
}

var session = `syntax = "proto3";

package session;

service SessionService {
rpc Login(LoginRequest) returns (LoginResponse) {}
}

message LoginRequest {
string username = 1;
string password = 2;
}

message LoginResponse {
string token = 1;
string err_msg = 2;
}
`
var sessionService = `package session

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
`
