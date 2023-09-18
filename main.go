package main

import (
	"fmt"
	"github.com/jmespath/go-jmespath"
	"os"
	"strings"
	"text/template"
)

var Functions = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	"Title": strings.ToTitle,
	"Lowercase": func(evaluation interface{}) string {
		if v, ok := evaluation.(string); ok {
			return strings.ToLower(v)
		}

		// @todo: Report Warning
		return "{{ ! }}" // Invalid evaluation, incorrect data-type.
	},
	"Index": func(meta interface{}, literal string) interface{} {
		parser := jmespath.MustCompile(fmt.Sprintf("%s", literal))

		value, e := parser.Search(meta)
		if e != nil || value == nil {
			// @todo: Report Warning
			return "{{ ! }}"
		}

		return value
	},
}

func main() {
	// Run the template to verify the output.
	if e := Template.Execute(os.Stdout, map[string]interface{}{
		"Package": map[string]interface{}{
			"Name": "test",
		},
	}); e != nil {
		panic(e)
	}
}

var Template = template.Must(template.New("root.go.template").Funcs(Functions).Parse(`
{{- define "root.go.template" -}}
package {{ Index $ "Package.Name" }}
{{- end -}}
`))
