package project

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"path/filepath"
)

func (p *Project) CreateGitIgnore() {
	mainTemplate := `vendor/
gen/
`
	data := make(map[string]interface{})

	mainScript, err := utils.ExecTemplate(mainTemplate, data)
	logging.IfErr("failed to execute template", err)

	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath(), ".gitignore"), mainScript)
	logging.IfErr("failed to write file", err)
}
