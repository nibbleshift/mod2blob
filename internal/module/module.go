package module

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path"
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

func LoadModule(packageName string, prefix string) (*Module, error) {
	mod := &Module{}

	// if module does not have slash, assume it is a runtime mod
	if strings.Count(packageName, "/") > 1 {
		goPath := os.Getenv("GOPATH")

		if goPath == "" {
			log.Println("GOPATH must be set")
			return nil, ErrEmptyString
		}

		if !checkIfDownloaded(packageName) {
			packageDir := path.Join(goPath, "src", packageName)

			_ = os.MkdirAll(packageDir, 0o755)

			gitURL := "https://" + packageName

			_, err := git.PlainClone(packageDir, false, &git.CloneOptions{
				URL:               gitURL,
				RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			})
			if err != nil {
				log.Println("Fetching module failed: " + err.Error())
				return nil, ErrCloneFailed
			}
		}
	}

	if strings.Contains(packageName, "@") {
		parts := strings.Split(packageName, "@")

		if len(parts) > 0 {
			packageName = parts[0]
		}
	}

	godoc := exec.Command("go", "doc", "-all", packageName)
	godoc.Env = os.Environ()
	godoc.Env = append(godoc.Env, "GO111MODULE=off")

	stdout, err := godoc.Output()
	if err != nil {
		log.Println(packageName + ": " + err.Error())
		return nil, err
	}

	mod.raw = stdout

	mod.Name = packageName
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
		"benthosType": toBenthosType,
		"function":    derefFunction,
		"getModule":   mod.GetName,
		"getPrefix":   mod.GetPrefix,
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
