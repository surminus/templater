package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTemplates(t *testing.T) {
	d := Data{
		Env:   map[string]string{"CAT": "fluffy"},
		Args:  []string{"melon", "orange"},
		Input: map[string]string{"nightmare": "freddy"},
	}

	actual, err := d.parseTemplate(`Test data below this line:
My cat is very {{ .Env.CAT }}.
{{ range .Args }}
I love {{ . }}!
{{ end }}

{{- if (eq .Input.nightmare "freddy") }}
I'm coming to your dreams tonight!
{{- end }}
`)
	expected := `Test data below this line:
My cat is very fluffy.

I love melon!

I love orange!

I'm coming to your dreams tonight!
`

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
