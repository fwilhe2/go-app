package main

import (
	"github.com/goccy/go-yaml"
)

func main() {
	yml := `
%YAML 1.2
---
a: 1
b: c
`
	var v struct {
		A int
		B string
	}
	if err := yaml.Unmarshal([]byte(yml), &v); err != nil {
		panic("foo")
	}
	println(&v)
}