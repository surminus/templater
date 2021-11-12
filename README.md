# Templater

Output a template using [Go template](https://pkg.go.dev/text/template) syntax.

## How to use

Create a test file:

```
cat << EOF > test.tmpl
My name is {{ .Env.USER }}.

I like to eat {{ index .Args 0 }}.
EOF
```

Build the tool:
```
make build
```

Run the tool:
```
bin/templater test.tmpl melons
```

Redirect to where you need it:
```
bin/templater test.tmpl melons > /my/important/file
```
