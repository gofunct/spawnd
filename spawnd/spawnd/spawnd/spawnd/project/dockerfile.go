package project

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"path"
	"path/filepath"
)

func (p *Project) CreateDockerfile() {
	mainTemplate := `FROM golang
COPY    . "{{ .importpath }}"
WORKDIR "{{ .importpath }}"
CMD     ["{{ .appName }}"]
EXPOSE  8000 9000
RUN     make install
`
	data := make(map[string]interface{})
	data["importpath"] = path.Join(p.GetName(), filepath.Base(p.GetCmd()))
	data["appName"] = path.Base(p.GetName())

	mainScript, err := utils.ExecTemplate(mainTemplate, data)
	logging.IfErr("failed to execute template", err)

	err = utils.WriteStringToFile(filepath.Join(p.GetAbsPath(), "Dockerfile"), mainScript)
	logging.IfErr("failed to write file", err)

}
