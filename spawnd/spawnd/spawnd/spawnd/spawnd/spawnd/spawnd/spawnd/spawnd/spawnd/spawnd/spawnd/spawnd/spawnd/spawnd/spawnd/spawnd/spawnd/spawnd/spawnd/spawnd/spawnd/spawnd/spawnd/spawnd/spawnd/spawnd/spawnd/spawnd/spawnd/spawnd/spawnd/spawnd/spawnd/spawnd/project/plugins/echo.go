package plugins

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"os"
	"path/filepath"
)

func CreateEchoProto() {
	mainTemplate := `syntax = "proto3";
package echo;

import "google/api/annotations.proto";

message EchoMessage {
 string value = 1;
}

service EchoService {
  rpc Echo(EchoMessage) returns (EchoMessage) {
    option (google.api.http) = {
      post: "/v1/echo"
      body: "*"
    };
  }
}
`
	data := make(map[string]interface{})

	mainScript, err := utils.ExecTemplate(mainTemplate, data)
	logging.IfErr("failed to execute template", err)
	os.MkdirAll("services/echo", os.ModePerm)
	err = utils.WriteStringToFile(filepath.Join("/services/echo", "echo.proto"), mainScript)
	logging.IfErr("failed to write file", err)

}

