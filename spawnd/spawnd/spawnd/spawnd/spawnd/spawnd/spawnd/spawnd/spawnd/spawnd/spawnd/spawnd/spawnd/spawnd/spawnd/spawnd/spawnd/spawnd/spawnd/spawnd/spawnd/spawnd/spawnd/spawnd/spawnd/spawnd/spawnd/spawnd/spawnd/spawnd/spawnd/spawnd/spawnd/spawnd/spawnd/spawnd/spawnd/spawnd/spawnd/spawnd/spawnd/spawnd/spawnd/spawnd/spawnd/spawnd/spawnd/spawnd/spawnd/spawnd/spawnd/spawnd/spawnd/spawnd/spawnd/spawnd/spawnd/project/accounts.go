package project

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"os"
	"path/filepath"
)

func (p *Project) CreateAccountsProto() {
	mainTemplate := accounts
	data := make(map[string]interface{})

	mainScript, err := utils.ExecTemplate(mainTemplate, data)
	logging.IfErr("failed to execute template", err)
	os.MkdirAll("services/account", os.ModePerm)

	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath()+"/services/account", "account.proto"), mainScript)
	logging.IfErr("failed to write file", err)
}

var accounts = `syntax = "proto3";
import "google/protobuf/empty.proto";

package account;

message Account {
string id = 1;
string name = 2;
string email = 3;
string confirm_token = 5;
string password_reset_token = 6;
map<string, string> metadata = 7;
}

message ListAccountsRequest {
int32 page_size = 1;
string page_token = 2;
}

message ListAccountsResponse {
repeated Account accounts = 1;
string next_page_token = 2;
}

message GetByIdRequest {
string id = 1;
}

message GetByEmailRequest {
string email = 1;
}

message AuthenticateByEmailRequest {
string email = 1;
string password = 2;
}

message GeneratePasswordTokenRequest {
string email = 1;
}

message GeneratePasswordTokenResponse {
string token = 1;
}

message ResetPasswordRequest {
string token = 1;
string password = 2;
}

message ConfirmAccountRequest {
string token = 1;
}

message CreateAccountRequest {
Account account = 1;
string password = 2;
}

message UpdateAccountRequest {
string id = 1;
string password = 2;
Account account = 4;
}

message DeleteAccountRequest {
string id = 1;
}

service AccountService {
rpc List (ListAccountsRequest) returns (ListAccountsResponse) {}
rpc GetById (GetByIdRequest) returns (Account) {}
rpc GetByEmail (GetByEmailRequest) returns (Account) {}
rpc AuthenticateByEmail (AuthenticateByEmailRequest) returns (Account) {}
rpc GeneratePasswordToken (GeneratePasswordTokenRequest) returns (GeneratePasswordTokenResponse) {}
rpc ResetPassword (ResetPasswordRequest) returns (Account) {}
rpc ConfirmAccount (ConfirmAccountRequest) returns (Account) {}
rpc Create (CreateAccountRequest) returns (Account) {}
rpc Update (UpdateAccountRequest) returns (Account) {}
rpc Delete (DeleteAccountRequest) returns (google.protobuf.Empty) {}
}
`
