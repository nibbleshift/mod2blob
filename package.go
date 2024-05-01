package main

import (
	"errors"
	"log"
	"os/exec"
	"regexp"
	"strings"
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
}

type Arg struct {
	Name string
	Type string
}

type Function struct {
	Name   string
	Args   []Arg
	Return []Arg
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

func LoadPackage(packageName string, prefix string) (*Package, error) {
	pkg := &Package{}
	cmd := exec.Command("go", "doc", packageName)

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

	return pkg, nil
}

func (pkg *Package) GetName() string {
	return pkg.Name
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

	match := pattern.FindStringSubmatch(def)

	if def == "" {
		return nil, ErrEmptyString
	}

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

func (pkg *Package) parseDoc() error {
	lines := strings.Split(string(pkg.raw), "\n")

	functions := []*Function{}
	for _, line := range lines {
		if strings.HasPrefix(line, "func") {
			function, err := parseFunction(line)

			if err != nil {
				log.Println(err)
				continue
			}
			functions = append(functions, function)
		}
	}

	pkg.Functions = functions

	return nil
}

func (pkg *Package) ListFunctions() []Function {
	return nil
}
