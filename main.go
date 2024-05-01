package main

import (
	"log"

	"github.com/nibbleshift/argenv"
)

type Config struct {
	Package    string `default:"" description:"Name of a go module such as 'math' or 'strings'"`
	OutputFile string `default:"" description:"Output file to write generated code to."`
	Prefix     string `default:"" description:"Prefix to use for function names"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := &Config{}
	argenv.Init(config)

	pkg, err := LoadPackage(config.Package, config.Prefix)

	if err != nil {
		log.Println("LoadPackage: " + err.Error())
		return
	}

	err = pkg.Generate()

	if err != nil {
		log.Println("Generate: " + err.Error())
		return
	}
}
