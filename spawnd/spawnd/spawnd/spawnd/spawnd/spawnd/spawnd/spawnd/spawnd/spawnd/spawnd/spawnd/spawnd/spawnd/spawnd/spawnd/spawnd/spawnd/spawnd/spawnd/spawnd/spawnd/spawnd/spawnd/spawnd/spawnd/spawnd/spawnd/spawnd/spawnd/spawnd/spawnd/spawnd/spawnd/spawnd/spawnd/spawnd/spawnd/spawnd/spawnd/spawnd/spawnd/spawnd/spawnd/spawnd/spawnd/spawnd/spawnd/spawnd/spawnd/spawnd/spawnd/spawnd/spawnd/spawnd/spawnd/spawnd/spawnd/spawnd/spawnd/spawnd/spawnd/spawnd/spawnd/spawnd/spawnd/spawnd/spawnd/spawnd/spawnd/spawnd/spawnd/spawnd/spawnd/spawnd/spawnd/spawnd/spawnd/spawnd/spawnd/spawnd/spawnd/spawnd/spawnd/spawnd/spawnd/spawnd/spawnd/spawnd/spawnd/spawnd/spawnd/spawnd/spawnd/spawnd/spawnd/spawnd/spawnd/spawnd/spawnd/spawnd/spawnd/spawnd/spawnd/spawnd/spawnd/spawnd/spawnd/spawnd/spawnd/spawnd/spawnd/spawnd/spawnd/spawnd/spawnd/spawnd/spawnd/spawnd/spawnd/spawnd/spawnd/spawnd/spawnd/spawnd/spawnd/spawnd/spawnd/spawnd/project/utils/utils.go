package utils

import (
	"bytes"
	"fmt"
	"github.com/gofunct/grpcgen/logging"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"unicode"
)

var SrcPaths []string

func init() {
	// Initialize SrcPaths.
	envGoPath := os.Getenv("GOPATH")
	goPaths := filepath.SplitList(envGoPath)
	if len(goPaths) == 0 {

		goExecutable := os.Getenv("GEN_GO_EXECUTABLE")
		if len(goExecutable) <= 0 {
			goExecutable = "go"
		}

		out, err := exec.Command(goExecutable, "env", "GOPATH").Output()
		logging.IfErr("failed to execute command", err)

		toolchainGoPath := strings.TrimSpace(string(out))
		goPaths = filepath.SplitList(toolchainGoPath)
		logging.IfErr("$GOPATH is not set", err)

	}
	SrcPaths = make([]string, 0, len(goPaths))
	for _, goPath := range goPaths {
		SrcPaths = append(SrcPaths, filepath.Join(goPath, "src"))
	}
}

func EmptyPath(path string) bool {
	fi, err := os.Stat(path)
	logging.IfErr("failed to check path", err)

	if !fi.IsDir() {
		return fi.Size() == 0
	}

	f, err := os.Open(path)
	logging.IfErr("failed to open path", err)

	defer f.Close()

	names, err := f.Readdirnames(-1)
	if err != nil && err != io.EOF {
		logging.IfErr("failed to read directories", err)

	}

	for _, name := range names {
		if len(name) > 0 && name[0] != '.' && strings.Contains(name, ".yaml") != true {
			return false
		}
	}
	return true
}

func PathExists(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		logging.IfErr("", err)

	}
	return false
}

func ExecTemplate(tmplStr string, data interface{}) (string, error) {
	tmpl, err := template.New("").Funcs(template.FuncMap{"comment": Commentify}).Parse(tmplStr)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	return buf.String(), err
}

func WriteStringToFile(path string, s string) error {
	return WriteToFile(path, strings.NewReader(s))
}

// WriteToFile writes r to file with path only
// if file/directory on given path doesn't exist.
func WriteToFile(path string, r io.Reader) error {
	if PathExists(path) {
		return fmt.Errorf("%v already exists", path)
	}

	dir := filepath.Dir(path)
	if dir != "" {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	return err
}

func Commentify(in string) string {
	var newlines []string
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "//") {
			newlines = append(newlines, line)
		} else {
			if line == "" {
				newlines = append(newlines, "//")
			} else {
				newlines = append(newlines, "// "+line)
			}
		}
	}
	return strings.Join(newlines, "\n")
}

// FindCmdDir checks if base of AbsPath is cmd dir and returns it or
// looks for existing cmd dir in AbsPath.
func FindCmdDir(AbsPath string) string {
	if !PathExists(AbsPath) || EmptyPath(AbsPath) {
		return "cmd"
	}

	if IsCmdDir(AbsPath) {
		return filepath.Base(AbsPath)
	}

	files, _ := filepath.Glob(filepath.Join(AbsPath, "c*"))
	for _, file := range files {
		if IsCmdDir(file) {
			return filepath.Base(file)
		}
	}

	return "cmd"
}

// findPackage returns full path to existing go package in GOPATHs.
func FindPackage(packageName string) string {
	if packageName == "" {
		return ""
	}

	for _, srcPath := range SrcPaths {
		packagePath := filepath.Join(srcPath, packageName)
		if PathExists(packagePath) {
			return packagePath
		}
	}

	return ""
}

// IsCmdDir checks if base of Name is one of cmdDir.
func IsCmdDir(Name string) bool {
	Name = filepath.Base(Name)
	for _, cmdDir := range []string{"cmd", "cmds", "command", "commands"} {
		if Name == cmdDir {
			return true
		}
	}
	return false
}

func FilePathHasPrefix(path string, prefix string) bool {
	if len(path) <= len(prefix) {
		return false
	}
	if runtime.GOOS == "windows" {
		// Paths in windows are case-insensitive.
		return strings.EqualFold(path[0:len(prefix)], prefix)
	}
	return path[0:len(prefix)] == prefix

}

// TrimScrcPath trims at the beginning of AbsPath the SrcPath.
func TrimScrcPath(AbsPath, SrcPath string) string {
	relPath, err := filepath.Rel(SrcPath, AbsPath)
	logging.IfErr("failed to trim source path", err)

	return relPath
}


// validateCmdName returns source without any dashes and underscore.
// If there will be dash or underscore, next letter will be uppered.
// It supports only ASCII (1-byte character) strings.
// https://github.com/spf13/cobra/issues/269
func ValidateCmdName(source string) string {
	i := 0
	l := len(source)
	// The output is initialized on demand, then first dash or underscore
	// occurs.
	var output string

	for i < l {
		if source[i] == '-' || source[i] == '_' {
			if output == "" {
				output = source[:i]
			}

			// If it's last rune and it's dash or underscore,
			// don't add it output and break the loop.
			if i == l-1 {
				break
			}

			// If next character is dash or underscore,
			// just skip the current character.
			if source[i+1] == '-' || source[i+1] == '_' {
				i++
				continue
			}

			// If the current character is dash or underscore,
			// upper next letter and add to output.
			output += string(unicode.ToUpper(rune(source[i+1])))
			// We know, what source[i] is dash or underscore and source[i+1] is
			// uppered character, so make i = i+2.
			i += 2
			continue
		}

		// If the current character isn't dash or underscore,
		// just add it.
		if output != "" {
			output += string(source[i])
		}
		i++
	}

	if output == "" {
		return source // source is initially valid name.
	}
	return output
}
