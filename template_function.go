package main

var FunctionTemplate string = `
package bloblang

{{ $moduleName := "math" -}}
import (
	"{{getPackage}}"
	"github.com/benthosdev/benthos/v4/public/bloblang"
)

func init() {
	var (
		err error
	)
{{ range $k, $v := . -}}
        {{ range $v }}
	{{- $nArgs := len .Args -}}
	{{- if gt $nArgs 0 -}}
	{{- $funcName := .Name }}
	object{{.Name}}Spec := bloblang.NewPluginSpec().
		{{- range $i, $el := .Args -}}
			{{if $i}}.{{end}}Param(bloblang.New{{ benthosType .Type}}Param("{{$el.Name}}"))
		{{- end }}
	// {{.Description}}
	err = bloblang.RegisterFunctionV2("{{ getPrefix }}{{ lower .Name}}", object{{.Name}}Spec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			{{- $argStr := "" -}}
			{{- range .Args -}}
			{{- $bType := benthosType .Type }}
			{{.Name}}, err := args.Get{{ $bType }}("{{ .Name }}")
			if err != nil {
				return nil, err
			}

			{{ .Name }}a := {{.Type}}({{.Name}})


			{{- if eq $argStr "" -}}
			{{ $argStr = (printf "%sa" .Name) }}
			{{ else }}
			{{ $argStr = (printf "%s, %sa" $argStr .Name) }}
			{{- end -}}
			{{ end -}}

			return func() (interface{}, error) {
				return {{$moduleName}}.{{$funcName}}({{ $argStr }}), nil
			}, nil
	})

	if err != nil {
		panic(err)
	}
	{{ end }}
{{- end }}
{{- end }}
}`
