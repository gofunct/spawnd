package project

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"path"
	"path/filepath"
)

func (p *Project) CreateMakeFile() {
	mainTemplate := `SOURCES :=	$(shell find . -name "*.proto" -not -path ./vendor/\*)
TEMPLATES := $(shell find ~/go/src/github.com/gofunct/grpcgen/templates -name "*.tmpl" -not -path ./vendor/\*)
TARGETS_GO :=	$(foreach source, $(SOURCES), $(source)_go)
TARGETS_TMPL :=	$(foreach source, $(SOURCES), $(source)_tmpl)
import_path := {{ .importpath }}
app_name = {{ .appName }}
service_name =	$(word 2,$(subst /, ,$1))

.PHONY: setup
setup: ## download dependencies and tls certificates
	brew install protobuf
	brew install prototool
	go get -u \
		google.golang.org/grpc \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/gogo/protobuf/protoc-gen-gogo \
		github.com/gogo/protobuf/protoc-gen-gogofast \
		github.com/ckaznocha/protoc-gen-lint \
		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
		github.com/golang/protobuf/{proto,protoc-gen-go} \
		moul.io/protoc-gen-gotemplate


.PHONY: list
list: list-templates list-protos

.PHONY: list-templates
list-templates:
	@echo "using templates:"
	@find ~/go/src/github.com/gofunct/grpcgen/templates -name "*.tmpl" -not -path ./vendor/\*

.PHONY: list-protos
list-protos:
	@echo "using protos:"
	@find . -name "*.proto" -not -path ./vendor/\*

.PHONY: session
session: services/session/session.pb.go

services/session/session.pb.go:	services/session/session.proto
	@protoc --gotemplate_out=destination_dir=services/session,template_dir=$(GOPATH)/src/github.com/gofunct/grpcgen/templates:services/session services/session/session.proto
	gofmt -w services/session
	@protoc --gogo_out=plugins=grpc:. services/session/session.proto

.PHONY: user
user: services/user/user.pb.go

services/user/user.pb.go:	services/user/user.proto
	@protoc --gotemplate_out=destination_dir=services/user,template_dir=$(GOPATH)/src/github.com/gofunct/grpcgen/templates:services/user services/user/user.proto
	gofmt -w services/user
	@protoc --gogo_out=plugins=grpc:. services/user/user.proto

.PHONY: account
account: services/account/account.pb.go

services/account/account.pb.go:	services/account/account.proto
	@protoc --gotemplate_out=destination_dir=services/account,template_dir=$(GOPATH)/src/github.com/gofunct/grpcgen/templates:services/account services/account/account.proto
	gofmt -w services/account
	@protoc --gogo_out=plugins=grpc:. services/account/account.proto

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort
`
	data := make(map[string]interface{})
	data["importpath"] = path.Join(p.GetName(), filepath.Base(p.GetCmd()))
	data["appName"] = path.Base(p.GetName())
	mainScript, err := utils.ExecTemplate(mainTemplate, data)
	logging.IfErr("failed to execute template", err)

	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath(), "Makefile"), mainScript)
	logging.IfErr("failed to write file", err)

}
