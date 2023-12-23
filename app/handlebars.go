package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aymerick/raymond"
)

type Template struct {
	Value string
}

type Handlebars struct {
	Templates map[string]Template
}

func (hb *Handlebars) Render(template string, data map[string]any) string {
	template_contents := hb.Templates[template]
	result, err := raymond.Render(template_contents.Value, data)
	if err != nil {
		panic("Please report a bug :)")
	}
	return result
}

func (hb *Handlebars) Init(template_dir string) {
	var templates = make(map[string]Template)
	dir, err := os.Open(template_dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range files {
		template_bytes, err := ioutil.ReadFile(template_dir + "/" + v.Name())
		if err != nil {
			log.Fatal("Failed to read config file.")
			return
		}
		templates[v.Name()] = Template{Value: string(template_bytes[:])}
	}
	hb.Templates = templates
}
