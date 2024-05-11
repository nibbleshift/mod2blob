package module

import (
	"os"
	"path"
	"regexp"
	"slices"
	"strings"
)

var native = []string{
	"int", "uint", "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64",
	"float", "float32", "float64",
	"string", "byte", "rune", "bool",
	"[]int", "[]uint", "[]int8", "[]uint8", "[]int16", "[]uint16", "[]int32", "[]uint32", "[]int64", "[]uint64",
	"float", "[]float32", "[]float64",
	"string", "[]byte", "[]rune", "[]bool",
	"error",
}

// Check to see if function accepts and returns
// only primitive types
func checkValidFunction(f *Function) bool {
	if f == nil {
		return false
	}

	if len(f.Args) == 0 {
		return false
	}

	/*
		for _, r := range f.Return {
			if !slices.Contains(native, r.Type) {
				return false
			}
		}*/

	for _, a := range f.Args {
		if !slices.Contains(native, a.Type) {
			return false
		}
	}
	return true
}

func checkIfDownloaded(packageName string) bool {
	gopath := os.Getenv("GOPATH")

	// if there are no "/", then it should be a module in the runtime
	if !strings.Contains(gopath, "/") {
		return true
	}

	packageDir := path.Join(gopath, "src", packageName)

	_, err := os.Stat(packageDir)

	if err == nil {
		return true
	}

	if os.IsExist(err) {
		return true
	}

	return false
}

func parseReturn(arg string) (*Arg, error) {
	var (
		argType string
		argName string
	)

	// check for empty arg
	if arg == "" {
		return nil, ErrInvalidFunction
	}

	// remove trailing/leading space
	arg = strings.TrimSpace(arg)

	if strings.Contains(arg, " ") {
		parts := strings.Split(arg, " ")

		switch len(parts) {
		case 1:
			argType = parts[0]
		case 2:
			// argument type and name provided
			argName = parts[0]
			argType = parts[1]

			if argType == "" {
				return nil, ErrInvalidFunction
			}
		default:
			return nil, ErrInvalidArguments
		}
	} else {
		argType = arg
	}

	argObj := Arg{
		Type: argType,
		Name: argName,
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
	var argObjectList []Arg

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
	var argObjectList []Arg

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

	// in some cases we need to resolve return types a bit,
	// esp when there are multiple args but only the last
	// specifies a return type
	nArgs := len(argObjectList)

	if nArgs > 1 {
		lastType := ""

		for i := nArgs - 1; i >= 0; i-- {
			if i == nArgs-1 {
				if argObjectList[i].Name == "" {
					break
				}
				lastType = argObjectList[i].Type
			} else if argObjectList[i].Name == "" && argObjectList[i].Type != "" {
				// fix the case: func Test() (one, two float64)
				argObjectList[i].Name = argObjectList[i].Type
				argObjectList[i].Type = lastType
			}

		}
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

	if len(match) == 0 {
		return nil, ErrInvalidFunction
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
				break
			}
		case "return":
			funcReturn, err = parseReturnArguments(match[i])
			if err != nil {
				break
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
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "Int64"
	case "string":
		return "String"
	case "[]byte", "[]string":
		return "Any"
	case "[]int", "[]int8", "[]int16", "[]int32", "[]int64", "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64":
		return "Any"
	case "[]float", "[]float32", "[]float64":
		return "Any"
	default:
		return typeStr
	}
}

func getFileName(name string) string {
	if strings.Contains(name, "/") {
		tmp := strings.Split(name, "/")
		name = tmp[len(tmp)-1]
	}

	return strings.ReplaceAll(name, "-", "_") + ".go"
}
