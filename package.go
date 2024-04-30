package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	ErrInvalidFunction = errors.New("invalid function definition")
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

func isValidFunction(definition string) bool {
	if !strings.HasPrefix(definition, "func ") {
		return false
	}

	if !strings.Contains(definition, "(") {
		return false
	}

	if !strings.Contains(definition, ")") {
		return false
	}

	return true
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
		fmt.Println(packageName + ": " + err.Error())
		return nil, err
	}

	pkg.raw = stdout

	pkg.Name = packageName
	pkg.Prefix = prefix
	pkg.parseDoc()

	return pkg, nil
}

func (pkg *Package) GetName() string {
	return pkg.Name
}

func parseFunction(def string) (*Function, error) {
	if !isValidFunction(def) {
		return nil, ErrInvalidFunction
	}

	f := Function{}

	function := def[5:]

	fields := strings.Split(function, "(")

	argsStrList := strings.Split(strings.Join(fields[1:], ""), ")")[:1]

	argObjList := make([]Arg, 0)
	for _, argsStr := range argsStrList {
		argStrList := strings.Split(string(argsStr), ",")

		for _, argStr := range argStrList {
			argStr := strings.Trim(argStr, " ")
			argObj := Arg{}

			arg := strings.Split(argStr, " ")
			argObj.Name = arg[0]

			if len(arg) > 1 {
				argObj.Type = strings.Join(arg[1:], "")
			}

			argObjList = append(argObjList, argObj)
		}

		nArgs := len(argStrList)
		argType := ""

		for i := nArgs - 1; i != 0; i-- {
			if argType == "" && argObjList[i].Type != "" {
				argType = argObjList[i].Type
				continue
			}
			argObjList[i].Type = argType
		}

	}

	returnString := strings.Join(strings.Split(strings.Join(fields[1:], ""), ")")[1:], "")
	if len(returnString) != 0 {
		if strings.Contains(returnString, "(") && strings.Contains(returnString, ")") {
			_ = Arg{Type: strings.Trim(argsStrList[1], " ")}
		} else {
			f.Return = []Arg{}
			f.Return = append(f.Return, Arg{Type: strings.Trim(returnString, " ")})
		}
	}
	if len(argObjList) > 0 {
		f.Args = argObjList
	}
	f.Name = fields[0]

	return &f, nil
}

func (pkg *Package) parseDoc() error {
	lines := strings.Split(string(pkg.raw), "\n")

	functions := []*Function{}
	for _, line := range lines {
		if strings.HasPrefix(line, "func") {
			function, err := parseFunction(line)

			if err != nil {
				fmt.Println(err)
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
