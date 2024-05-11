package gen

var Function string = `
package bloblang

import (
	"{{getModulePath}}"
	"github.com/benthosdev/benthos/v4/public/bloblang"
)

func init() {
	var (
		err error
	)
{{ range $k, $v := . -}}
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
			{{- $returnVal := "" -}}
			{{- range .Args -}}
			{{- $bType := benthosType .Type }}
			{{- $getType := $bType }}
			{{- if eq $getType "Any" }}
			{{ $getType = "" }}
			{{ end }}
			{{.Name}}, err := args.Get{{ $getType }}("{{ .Name }}")
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

			{{- range .Return -}}
			{{- if eq $returnVal "" -}}
			{{ $returnVal = (printf "%s" .Name) }}
			{{ else }}
			{{ $returnVal = (printf "%s, %s" $returnVal .Name) }}
			{{- end -}}
			{{- end -}}

			return func() (any, error) {
				{{ $nReturn := len .Return }}
				{{ if gt $nReturn 1 }}
					{{$returnVal}} := {{getModuleName}}.{{$funcName}}({{ $argStr }})
					obj := map[string]any{}
					{{- range .Return }}
					obj["{{.Name}}"] = {{.Name}}
					{{- end }}
					return obj, nil
				{{ else }}
				return {{getModuleName}}.{{$funcName}}({{ $argStr }}), nil
				{{- end -}}
			}, nil
	})

	if err != nil {
		panic(err)
	}
	{{ end }}
{{- end }}
}`
