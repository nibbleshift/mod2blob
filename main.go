package main

import (
	"log"

	"github.com/nibbleshift/argenv"
	"github.com/nibbleshift/mod2blob/internal/module"
)

type Config struct {
	Module    string `default:"" description:"Name of a go module such as 'math' or 'strings'"`
	Prefix    string `default:"" description:"Prefix to use for function names. Format: [a-Z0-9]"`
	Debug     bool   `default:"false" description:"Enable debug logging"`
	OutputDir string `default:"." description:"Directory to write generated code to"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := &Config{}
	argenv.Init(config)

	pkg, err := module.LoadModule(config.Module, config.Prefix)
	if err != nil {
		log.Println(err)
		return
	}

	err = pkg.Generate(config.OutputDir)
	if err != nil {
		log.Println("Generate: " + err.Error())
		return
	}
}
