package project

import (
	"errors"
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"os"
	"path/filepath"
	"strings"
)

type Project struct {
	AbsPath string
	CmdPath string
	SrcPath string
	Name    string
}

func InitializeProject(p *Project) {
	CheckPathBeforeProject(p)
	p.CreateMainFile()
	p.CreatePrototoolfile()
	p.CreateGitIgnore()
	p.CreateDockerfile()
	p.CreateMakeFile()
	p.CreateRootCmdFile()
	p.CreateSessionsProto()
	p.CreateUsersProto()
	p.CreateAccountsProto()
	p.CreateGoKitServerCmdFile()
}

func NewGokitServerCmd(p *Project) {
	p.CreateGoKitServerCmdFile()
}

func NewProjectFromCurrentPath() *Project {
	wd, err := os.Getwd()
	logging.IfErr("failed to get working directory", err)
	return NewProjectFromPath(wd)
}

func CheckPathBeforeProject(p *Project) {
	var err error
	if !utils.PathExists(p.GetAbsPath()) { // If path doesn't yet exist, create it
		if err = os.MkdirAll(p.GetAbsPath(), os.ModePerm); err != nil {
			logging.IfErr("path error", err)
		}

	} else if !utils.EmptyPath(p.GetAbsPath()) { // If path exists and is not empty don't use it
		logging.IfErr("path error", errors.New("cannot create a new project in a non empty directory: "+p.GetAbsPath()))
	}
}

// NewProject returns Project with specified project Name.
func NewProject(projectName string) *Project {
	if projectName == "" {
		logging.Exit("can't create project with blank Name")
	}

	p := new(Project)
	p.Name = projectName

	// 1. Find already created protect.
	p.AbsPath = utils.FindPackage(projectName)

	// 2. If there are no created project with this path, and user is in GOPATH,
	// then use GOPATH/src/projectName.
	if p.AbsPath == "" {
		wd, err := os.Getwd()
		logging.IfErr("failed to get working directory", err)
		for _, SrcPath := range utils.SrcPaths {
			goPath := filepath.Dir(SrcPath)
			if utils.FilePathHasPrefix(wd, goPath) {
				p.AbsPath = filepath.Join(SrcPath, projectName)
				break
			}
		}
	}

	// 3. If user is not in GOPATH, then use (first GOPATH)/src/projectName.
	if p.AbsPath == "" {
		p.AbsPath = filepath.Join(utils.SrcPaths[0], projectName)
	}

	return p
}

// NewProjectFromPath returns Project with specified absolute path to
// package.
func NewProjectFromPath(AbsPath string) *Project {
	if AbsPath == "" {
		logging.Exit("can't create project: AbsPath can't be blank")
	}
	if !filepath.IsAbs(AbsPath) {
		logging.Exit("can't create project: AbsPath is not absolute")
	}

	// If AbsPath is symlink, use its destination.
	fi, err := os.Lstat(AbsPath)
	logging.IfErr("can't read path info: ", err)

	if fi.Mode()&os.ModeSymlink != 0 {
		path, err := os.Readlink(AbsPath)
		logging.IfErr("can't read the destination of symlink: ", err)
		AbsPath = path
	}

	p := new(Project)
	p.AbsPath = strings.TrimSuffix(AbsPath, utils.FindCmdDir(AbsPath))
	p.Name = filepath.ToSlash(utils.TrimScrcPath(p.AbsPath, p.GetSource()))
	return p
}

// Name returns the Name of project, e.g. "github.com/spf13/cobra"
func (p *Project) GetName() string {
	return p.Name
}

// CmdPath returns absolute path to directory, where all commands are located.
func (p *Project) GetCmd() string {
	if p.AbsPath == "" {
		return ""
	}
	if p.CmdPath == "" {
		p.CmdPath = filepath.Join(p.AbsPath, utils.FindCmdDir(p.AbsPath))
	}
	return p.CmdPath
}

// AbsPath returns absolute path of project.
func (p Project) GetAbsPath() string {
	return p.AbsPath
}

// AbsPath returns absolute path of project.
func (p Project) Absolute() string {
	return p.AbsPath
}

// SrcPath returns absolute path to $GOPATH/src where project is located.
func (p *Project) GetSource() string {
	if p.SrcPath != "" {
		return p.SrcPath
	}
	if p.AbsPath == "" {
		p.SrcPath = utils.SrcPaths[0]
		return p.SrcPath
	}

	for _, SrcPath := range utils.SrcPaths {
		if utils.FilePathHasPrefix(p.AbsPath, SrcPath) {
			p.SrcPath = SrcPath
			break
		}
	}

	return p.SrcPath
}
