package main

import (
	"io/ioutil"
	"github.com/goccy/go-yaml"
)

func noError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile(".github/workflows/ci.yml")
	noError(err)

	type Step struct {
		Name string
		Run string
		Uses string
	}

	type Job struct {
		RunsOn string `yaml:"runs-on"`
		Steps []Step `yaml:"steps"`
	}

	var v struct {
		Name string
		On map[string]interface{}
		Jobs map[string] struct {
			RunsOn string `yaml:"runs-on"`
			Steps []Step `yaml:"steps"`
		}
	}
	if err := yaml.Unmarshal(dat, &v); err != nil {
		panic("foo")
	}
	println(v.Name)
	println(v.On["push"])
	println(v.Jobs["build"].RunsOn)
	for i := 0; i < len(v.Jobs["build"].Steps); i++ {
		println(v.Jobs["build"].Steps[i].Name)
		println(v.Jobs["build"].Steps[i].Uses)
		println(v.Jobs["build"].Steps[i].Run)
	}
}