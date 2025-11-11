package mockcrash

import (
	"bytes"
	"text/template"
)

func Execute(tplText string, vals map[string]any) (string, error) {
	tpl, err := template.New("tpl").Funcs(template.FuncMap{
		"doit": func(wp WithPtr) (Value, error) {
			return wp.Do(&Value{})
		},
	}).Option("missingkey=zero").Parse(tplText)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err = tpl.Execute(&buf, vals); err != nil {
		return "", err
	}
	return buf.String(), nil
}
