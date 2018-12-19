package project

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"path"
	"path/filepath"
)

func (p *Project) CreateMainFile() {
	mainTemplate := `package main

import "{{ .importpath }}"

func main() {
	cmd.Execute()
}
`
	data := make(map[string]interface{})
	data["importpath"] = path.Join(p.GetName(), filepath.Base(p.GetCmd()))

	mainScript, err := utils.ExecTemplate(mainTemplate, data)
	logging.IfErr("failed to execute template", err)

	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath(), "main.go"), mainScript)
	logging.IfErr("failed to write file", err)

}
