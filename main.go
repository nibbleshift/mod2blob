package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/nibbleshift/argenv"
)

type Config struct {
	Package    string `default:"" description:"Name of a go module such as 'math' or 'strings'"`
	OutputFile string `default:"" description:"Output file to write generated code to."`
	Template   string `default:"bloblang.tmpl" description:"Template file to use for generated code"`
	Prefix     string `default:"" description:"Prefix to use for function names"`
}

func main() {
	config := &Config{}
	argenv.Init(config)

	pkg, err := LoadPackage(config.Package, config.Prefix)

	if err != nil {
		fmt.Println("LoadPackage: " + err.Error())
		return
	}

	var tmplFile = config.Template

	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)

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
