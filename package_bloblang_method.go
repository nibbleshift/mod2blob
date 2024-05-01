package main

var MethodTemplate string = `
package bloblang

{{ $moduleName := .Name -}}
import (
	"log"

	"{{.Name}}"
	"github.com/benthosdev/benthos/v4/public/bloblang"
)

func init() {
	var (
		err error
	)
{{ range .Functions }}
	{{- $nArgs := len .Args -}}
	{{- if gt $nArgs 0 -}}
	{{- $funcName := .Name }}
	object{{.Name}}Spec := bloblang.NewPluginSpec().
		{{- range $i, $el := .Args -}}
			{{if $i}}.{{end}}Param(bloblang.New{{ benthosType .Type}}Param("{{$el.Name}}"))
		{{- end }}

	err = bloblang.RegisterFunctionV2("{{.Name}}", object{{.Name}}Spec,
		func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			{{- $argStr := "" -}}
			{{- range .Args -}}
			{{- if eq $argStr "" -}}
			{{ $argStr = (printf "%s" .Name) }}
			{{ else }}
			{{ $argStr = (printf "%s, %s" $argStr .Name) }}
			{{- end -}}
			{{.Name}}, err := args.Get{{ benthosType .Type }}("{{ .Name }}")
			if err != nil {
				return nil, err
			}
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
}`
