package mockcrash

import (
	"bytes"
	"context"
	"text/template"
)

type mockKey string

var MockKey mockKey = "mockKey"

func Execute(ctx context.Context, tplText string, vals map[string]any) (string, error) {
	wp := ctx.Value(MockKey).(WithPtr)
	tpl, err := template.New("tpl").Funcs(template.FuncMap{
		"doit": func() (Value, error) {
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
