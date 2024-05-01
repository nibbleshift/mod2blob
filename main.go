package main

import (
	"html/template"
	"log"
	"os"

	"github.com/42atomys/sprout"
	"github.com/nibbleshift/argenv"
)

type Config struct {
	Package    string `default:"" description:"Name of a go module such as 'math' or 'strings'"`
	OutputFile string `default:"" description:"Output file to write generated code to."`
	Template   string `default:"bloblang.tmpl" description:"Template file to use for generated code"`
	Prefix     string `default:"" description:"Prefix to use for function names"`
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

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := &Config{}
	argenv.Init(config)

	customFuncs := map[string]any{
		"benthosType": toBenthosType,
	}

	pkg, err := LoadPackage(config.Package, config.Prefix)

	if err != nil {
		log.Println("LoadPackage: " + err.Error())
		return
	}

	var tmplFile = config.Template

	tmpl, err := template.New(tmplFile).
		Funcs(sprout.FuncMap()).
		Funcs(customFuncs).
		ParseFiles(tmplFile)

	if err != nil {
		panic(err)
	}

	f, err := os.Create(config.OutputFile)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, pkg)

	if err != nil {
		panic(err)
	}
}
