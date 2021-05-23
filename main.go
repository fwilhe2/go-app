package main

import (
	"github.com/goccy/go-yaml"
	"io/ioutil"
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
		Run  string
		Uses string
	}

	type Job struct {
		RunsOn string `yaml:"runs-on"`
		Steps  []Step `yaml:"steps"`
	}

	var workflow struct {
		Name string
		On   map[string]interface{}
		Jobs map[string]Job
	}
	if err := yaml.Unmarshal(dat, &workflow); err != nil {
		panic("foo")
	}
	println(workflow.Name)
	for trigger, _ := range workflow.On {
		println("Workflow runs on trigger " + trigger)
	}

	for jobName, job := range workflow.Jobs {
		println("Job: " + jobName)
		println("Runs On: " + job.RunsOn)
		for x := range job.Steps {
			println("  Name: " + job.Steps[x].Name)
			println("  Uses: " + job.Steps[x].Uses)
			println("  Run:  " + job.Steps[x].Run)
			println()
		}
	}
}
