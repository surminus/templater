package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

// Data contains the values that can be accessed inside a template
type Data struct {
	Args  []string
	Env   map[string]string
	Input map[string]string
}

var (
	input   map[string]string
	verbose bool
)

func init() {
	pflag.StringToStringVarP(&input, "input", "i", nil, "key/value pairs of input eg FOO=BAR")
	pflag.BoolVarP(&verbose, "verbose", "v", false, "more output")
	pflag.Parse()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must supply template file as first argument")
		os.Exit(1)
	}

	envvars := make(map[string]string)
	for _, env := range os.Environ() {
		s := strings.Split(env, "=")
		envvars[s[0]] = s[1]
	}

	var args []string
	args = append(args, os.Args[2:]...)

	data := Data{
		Args:  args,
		Env:   envvars,
		Input: input,
	}

	templateData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := data.parseTemplate(string(templateData))
	if err != nil {
		fmt.Println("Unable to parse template. Check syntax.")

		if verbose {
			fmt.Println(err)
		}

		os.Exit(1)
	}

	fmt.Println(out)
}

func (d Data) parseTemplate(t string) (out string, err error) {
	tmpl, err := template.New(time.Now().String()).Parse(t)
	if err != nil {
		return out, err
	}

	b := new(bytes.Buffer)
	err = tmpl.Execute(b, d)

	out = b.String()

	return out, err
}
