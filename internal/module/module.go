package module

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/42atomys/sprout"
	"github.com/go-git/go-git/v5"
	"github.com/nibbleshift/mod2blob/internal/gen"
	"mvdan.cc/gofumpt/format"
)

func (mod *Module) addToMap(callType string, f *Function) error {
	if _, ok := mod.Map[callType]; !ok {
		mod.Map = make(map[string][]*Function)
	}

	mod.Map[callType] = append(mod.Map[callType], f)

	return nil
}

func (mod *Module) buildMap() error {
	mod.Map = make(map[string][]*Function)

	for _, f := range mod.Functions {
		if !checkValidFunction(f) {
			log.Printf("Skipped function %+v Args:%v Return:%v\n", f.Name, f.Args, f.Return)
			continue
		}

		if len(f.Args) > 0 {
			_ = mod.addToMap("function", f)
			log.Printf("Added function %+v Args:%v Return:%v\n", f.Name, f.Args, f.Return)
		}
	}
	return nil
}

func gitClone(moduleUrl string) error {
	packageDir, err := getModuleSrcPath(moduleUrl)

	if err != nil {
		return err
	}

	_ = os.MkdirAll(packageDir, 0755)

	gitURL := "https://" + moduleUrl

	_, err = git.PlainClone(packageDir, false, &git.CloneOptions{
		URL:               gitURL,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		log.Println("Fetching module failed: " + err.Error())
		return ErrCloneFailed
	}
	return nil
}

func runGoDoc(modulePath string) ([]byte, error) {
	godoc := exec.Command("go", "doc", "-all", modulePath)
	godoc.Env = os.Environ()
	godoc.Env = append(godoc.Env, "GO111MODULE=off")

	docStr, err := godoc.Output()

	if err != nil {
		log.Println(modulePath + ": " + err.Error())
		return nil, err
	}
	return docStr, nil
}

func getModuleSrcPath(moduleUrl string) (string, error) {
	goPath := os.Getenv("GOPATH")

	if goPath == "" {
		log.Println("GOPATH must be set")
		return "", ErrEmptyString
	}
	return path.Join(goPath, "src", moduleUrl), nil
}

func getModuleName(moduleUrl string) (string, error) {
	var moduleName string

	// This should be better, if it is a runtime package
	// then we dont need to look for the module name
	if strings.Count(moduleUrl, "/") < 2 {
		return moduleUrl, nil
	}

	modulePath, err := getModuleSrcPath(moduleUrl)

	if err != nil {
		return "", err
	}

	files, err := filepath.Glob(modulePath + "/*.go")

	if err != nil {
		return "", err
	}

	for _, f := range files {
		f, err := os.Open(f)

		if err != nil {
			continue
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		// read the first line
		scanner.Scan()

		// extract text for first line
		line := scanner.Text()

		// attempt to get the package name from file
		if strings.HasPrefix(line, "package ") {
			parts := strings.Split(line, " ")

			if len(parts) != 2 {
				continue
			}
			moduleName = strings.TrimSpace(parts[1])
			break
		}
	}

	if moduleName == "" {
		return "", ErrEmptyString
	}

	return moduleName, nil
}
func LoadModule(modulePath string, prefix string) (*Module, error) {
	var (
		err    error
		docStr []byte
	)

	mod := &Module{}

	// if module does not have slash, assume it is a runtime mod
	if strings.Count(modulePath, "/") > 1 {
		if !checkIfDownloaded(modulePath) {
			err = gitClone(modulePath)

			if err != nil {
				return nil, err
			}
		}
	}

	docStr, err = runGoDoc(modulePath)

	if err != nil {
		return nil, err
	}

	moduleName, err := getModuleName(modulePath)

	if err != nil {
		return nil, err
	}

	mod.raw = docStr
	mod.Name = moduleName
	mod.Path = modulePath
	mod.Prefix = prefix

	err = mod.parseDoc()
	if err != nil {
		return nil, err
	}

	// build a map of functions, methods etc.
	err = mod.buildMap()
	if err != nil {
		return nil, err
	}

	return mod, nil
}

func (mod *Module) GetName() string {
	return mod.Name
}

func (mod *Module) GetPath() string {
	return mod.Path
}

func (mod *Module) GetPrefix() string {
	return mod.Prefix
}

func (mod *Module) parseDoc() error {
	lines := strings.Split(string(mod.raw), "\n")

	functions := []*Function{}
	for i := 0; i < len(lines); i += 2 {
		if strings.HasPrefix(lines[i], "func") {
			function, err := parseFunction(lines[i])
			if err != nil {
				continue
			}

			// This only works with single line descriptions,
			// fix to span multiple lines
			function.Description = strings.TrimSpace(lines[i+1])

			functions = append(functions, function)

		} else if strings.HasPrefix(lines[i], "const (") {
			i++ // skip passed the const ( line
			for i < len(lines) {
				// if "//" exists within a string this will cause a problem
				if strings.Contains(lines[i], "//") {
					lines[i] = strings.Split(lines[i], "//")[0]
				}
				pattern := regexp.MustCompile(`\s+(?P<name>\S+)\s+=\s+(?P<value>.*)`)
				match := pattern.FindStringSubmatch(lines[i])

				if len(match) > 0 {
					var (
						name  string
						value string
					)

					for j, expName := range pattern.SubexpNames() {
						// log.Printf("'%8s'\t %d -> %s\n", name, i, match[i])
						// skip matches that are empty strings
						if match[j] == "" {
							continue
						}
						switch expName {
						case "name":
							name = match[j]
						case "value":
							value = match[j]
						}
					}

					mod.Constants = append(mod.Constants, Constant{Name: name, Value: value})
				}
				if strings.HasPrefix(lines[i], ")") {
					break
				}
				i++
			}
		}
	}

	mod.Functions = functions

	return nil
}

func derefFunction(f *Function) Function {
	return *f
}

func (mod *Module) Generate() error {
	customFuncs := map[string]any{
		"benthosType":   toBenthosType,
		"function":      derefFunction,
		"getModulePath": mod.GetPath,
		"getModuleName": mod.GetName,
		"getPrefix":     mod.GetPrefix,
	}

	if len(mod.Map["function"]) > 0 {
		var (
			err    error
			source bytes.Buffer
		)
		// generate functions
		funcTmpl, err := template.New("function").
			Funcs(sprout.FuncMap()).
			Funcs(customFuncs).
			Parse(gen.Function)
		if err != nil {
			panic(err)
		}

		err = funcTmpl.Execute(&source, mod.Map["function"])
		if err != nil {
			panic(err)
		}

		var formatted []byte

		formatted, err = format.Source(source.Bytes(), format.Options{ExtraRules: true})
		if err != nil {
			panic(err)
		}

		f, err := os.Create(getFileName(mod.Name))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = f.Write(formatted)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func (mod *Module) ListFunctions() []Function {
	return nil
}
