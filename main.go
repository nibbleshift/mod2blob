package main

import (
	"log"

	"github.com/nibbleshift/argenv"
)

type Config struct {
	Package string `default:"" description:"Name of a go module such as 'math' or 'strings'"`
	Prefix  string `default:"" description:"Prefix to use for function names. Format: [a-Z0-9]"`
	Debug   bool   `default:"false" description:"Enable debug logging"`
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
