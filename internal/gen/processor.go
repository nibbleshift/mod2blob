package gen

var Processor string = `---
processor_resources:
  - label: execute_map
    mapping: |
      root = {}
      {{- range . }}
      {{ $argStr := "" }}
      {{- range .Args }}
      {{- $randValue := randInt 1 1000 -}}
      {{- if eq $argStr "" -}}
      {{- $argStr = (printf "%d"  $randValue) -}}
      {{- else -}}
      {{ $argStr = (printf "%s, %d" $argStr $randValue) }}
      {{- end -}}
      {{- end -}}
      root.{{lower .Name}} = {{lower .Name}}({{$argStr}})
      {{- end }}`
