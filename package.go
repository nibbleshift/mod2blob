package main

import (
	"bytes"
	"errors"
	"go/format"
	"html/template"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/42atomys/sprout"
)

var (
	ErrInvalidFunction  = errors.New("invalid function definition")
	ErrFunctionNoReturn = errors.New("function has no return values")
	ErrEmptyString      = errors.New("string is empty")
	ErrInvalidArguments = errors.New("invalid function arguments")
)

type Package struct {
	raw       []byte
	Functions []*Function
	Name      string
	Prefix    string
	Constants []Constant
	// map[method|function][String|Bytes|Int64|Float64|etc]*Function
	Map map[string]map[string][]*Function
}

type Arg struct {
	Name string
	Type string
}

type Constant struct {
	Name  string
	Value string
}

type Function struct {
	Name        string
	Description string
	Args        []Arg
	Return      []Arg
}

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetArgs() []Arg {
	return f.Args
}

func (f *Function) GetReturn() []Arg {
	return f.Return
}

func (p *Package) addToMap(callType string, paramType string, f *Function) error {

	if _, ok := p.Map[callType]; !ok {
		p.Map[callType] = make(map[string][]*Function)
	}

	p.Map[callType][paramType] = append(p.Map[callType][paramType], f)
	return nil
}

func (p *Package) buildMap() error {
	p.Map = make(map[string]map[string][]*Function)

	for _, f := range p.Functions {
		if len(f.Args) == 1 {
			switch f.Args[0].Type {
			case "float64", "float32":
				_ = p.addToMap("method", "Float64", f)
			case "int64", "int32", "int":
				_ = p.addToMap("method", "Int64", f)
			default:
				log.Println("Unsupported method type: " + f.Args[0].Type)
			}
		}

		if len(f.Args) > 0 {
			switch f.Args[0].Type {
			case "float64", "float32":
				_ = p.addToMap("function", "Float64", f)
			case "int64", "int32", "int":
				_ = p.addToMap("function", "Int64", f)
			default:
				log.Println("Unsupported function type: " + f.Args[0].Type)
			}
		}
	}
	return nil
}

func LoadPackage(packageName string, prefix string) (*Package, error) {
	pkg := &Package{}
	cmd := exec.Command("go", "doc", "-all", packageName)

	stdout, err := cmd.Output()

	if err != nil {
		log.Println(packageName + ": " + err.Error())
		return nil, err
	}

	pkg.raw = stdout

	pkg.Name = packageName
	pkg.Prefix = prefix

	err = pkg.parseDoc()

	if err != nil {
		return nil, err
	}

	// build a map of functions, methods etc.
	err = pkg.buildMap()

	if err != nil {
		return nil, err
	}

	return pkg, nil
}

func (pkg *Package) GetName() string {
	return pkg.Name
}

func (pkg *Package) GetPrefix() string {
	return pkg.Prefix
}

func parseReturn(arg string) (*Arg, error) {
	var (
		argType string
	)

	// check for empty arg
	if arg == "" {
		return nil, ErrInvalidFunction
	}

	// remove trailing/leading space
	arg = strings.TrimSpace(arg)

	parts := strings.Split(arg, " ")

	switch len(parts) {
	case 1:
		argType = parts[0]
	case 2:
		// argument type and name provided
		argType = parts[1]

		if argType == "" {
			return nil, ErrInvalidFunction
		}
	default:
		return nil, ErrInvalidArguments
	}

	argObj := Arg{
		Type: argType,
	}

	return &argObj, nil
}

func parseArgument(arg string) (*Arg, error) {
	var (
		argName string
		argType string
	)

	// check for empty arg
	if arg == "" {
		return nil, ErrInvalidFunction
	}

	// remove trailing/leading space
	arg = strings.TrimSpace(arg)

	parts := strings.Split(arg, " ")

	switch len(parts) {
	case 1:
		// this argument only has a name, the type of this argument MUST be resolved
		// outside of this function using the other arguments available
		argName = parts[0]
	case 2:
		// argument type and name provided
		argName = parts[0]
		argType = parts[1]

		if argType == "" || argName == "" {
			return nil, ErrInvalidFunction
		}
	default:
		return nil, ErrInvalidArguments
	}

	argObj := Arg{
		Name: argName,
		Type: argType,
	}

	return &argObj, nil
}

func parseFunctionArguments(args string) ([]Arg, error) {
	var (
		argObjectList []Arg
	)

	// return if we get an empty string
	if args == "" {
		return nil, ErrEmptyString
	}

	argObjectList = make([]Arg, 0)

	if strings.Contains(args, ",") {
		// multiple argument case
		arguments := strings.Split(args, ", ")

		// iterate through arguments and add each to the arg list
		for _, arg := range arguments {
			argObj, err := parseArgument(arg)

			if err != nil {
				return nil, err
			}
			argObjectList = append(argObjectList, *argObj)
		}
	} else {
		// single argument case
		argObj, err := parseArgument(args)

		if err != nil {
			return nil, err
		}

		argObjectList = append(argObjectList, *argObj)
	}

	// check to see if all arg typs have been resolved
	resolve := make([]int, 0)
	resolveToType := ""
	for i, argObject := range argObjectList {
		if argObject.Type == "" {
			resolve = append(resolve, i)
			continue
		}

		if resolveToType == "" {
			resolveToType = argObject.Type
		}
	}

	for i := range resolve {
		argObjectList[i].Type = resolveToType
	}

	return argObjectList, nil
}

func parseReturnArguments(args string) ([]Arg, error) {
	var (
		argObjectList []Arg
	)

	args = strings.ReplaceAll(strings.ReplaceAll(args, "(", ""), ")", "")

	// return if we get an empty string
	if args == "" {
		return nil, ErrEmptyString
	}

	argObjectList = make([]Arg, 0)

	if strings.Contains(args, ",") {
		// multiple argument case
		arguments := strings.Split(args, ", ")

		// iterate through arguments and add each to the arg list
		for _, arg := range arguments {
			argObj, err := parseReturn(arg)

			if err != nil {
				return nil, err
			}
			argObjectList = append(argObjectList, *argObj)
		}
	} else {
		// single argument case
		argObj, err := parseReturn(args)

		if err != nil {
			return nil, err
		}

		argObjectList = append(argObjectList, *argObj)
	}

	// remove "( )" around the args if it is present
	return argObjectList, nil
}

func parseFunction(def string) (*Function, error) {
	var (
		err        error
		funcName   string
		funcArgs   []Arg
		funcReturn []Arg
	)
	pattern := regexp.MustCompile(`^func (?P<funcName>\S+)\((?P<args>.*?)\)(?P<return>.*)`)

	if def == "" {
		return nil, ErrEmptyString
	}
	match := pattern.FindStringSubmatch(def)

	for i, name := range pattern.SubexpNames() {
		// log.Printf("'%8s'\t %d -> %s\n", name, i, match[i])
		// skip matches that are empty strings
		if match[i] == "" {
			continue
		}
		switch name {
		case "funcName":
			funcName = match[i]
		case "args":
			funcArgs, err = parseFunctionArguments(match[i])

			if err != nil {
				log.Printf("%s: err: %s\n", err.Error(), match[i])
			}
		case "return":
			funcReturn, err = parseReturnArguments(match[i])
			if err != nil {
				log.Println(err)
			}
		}
	}

	f := Function{
		Name:   funcName,
		Args:   funcArgs,
		Return: funcReturn,
	}
	return &f, nil
}

func toBenthosType(typeStr string) string {
	switch typeStr {
	case "float", "float32", "float64":
		return "Float64"
	case "int", "int32", "int64", "uint", "uint32", "uint64":
		return "Int64"
	default:
		return typeStr
	}
}

func (pkg *Package) parseDoc() error {
	lines := strings.Split(string(pkg.raw), "\n")

	functions := []*Function{}
	for i := 0; i < len(lines); i += 2 {
		if strings.HasPrefix(lines[i], "func") {
			function, err := parseFunction(lines[i])

			if err != nil {
				log.Println(err)
				continue
			}

			// This only works with single line descriptions,
			// fix to span multiple lines
			function.Description = strings.TrimSpace(lines[i+1])

			functions = append(functions, function)

			log.Printf("Added function %s:%s\n", function.Name, function.Description)
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

					pkg.Constants = append(pkg.Constants, Constant{Name: name, Value: value})
				}
				if strings.HasPrefix(lines[i], ")") {
					break
				}
				i++
			}
		}

	}

	pkg.Functions = functions

	return nil
}

func function(f *Function) Function {
	return *f
}

func (pkg *Package) Generate() error {
	customFuncs := map[string]any{
		"benthosType": toBenthosType,
		"function":    function,
		"getPackage":  pkg.GetName,
		"getPrefix":   pkg.GetPrefix,
	}

	if len(pkg.Map["function"]) > 0 {
		var (
			err    error
			source bytes.Buffer
		)
		// generate functions
		funcTmpl, err := template.New("function").
			Funcs(sprout.FuncMap()).
			Funcs(customFuncs).
			Parse(FunctionTemplate)

		if err != nil {
			panic(err)
		}

		err = funcTmpl.Execute(&source, pkg.Map["function"])

		if err != nil {
			panic(err)
		}

		var formatted []byte

		formatted, err = format.Source(source.Bytes())

		if err != nil {
			panic(err)
		}

		f, err := os.Create(pkg.Name + "_function.go")

		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = f.Write(formatted)

		if err != nil {
			panic(err)
		}
	}

	/*
		if len(pkg.Map["method"]) > 0 {
			// generate methods
			methodTmpl, err := template.New("method").
				Funcs(sprout.FuncMap()).
				Funcs(customFuncs).
				Parse(MethodTemplate)

			if err != nil {
				panic(err)
			}

			f, err := os.Create(pkg.Name + "_method.go")

			if err != nil {
				panic(err)
			}
			defer f.Close()

			err = methodTmpl.Execute(f, pkg.Map["method"])

			if err != nil {
				panic(err)
			}
		}*/

	return nil
}
func (pkg *Package) ListFunctions() []Function {
	return nil
}
