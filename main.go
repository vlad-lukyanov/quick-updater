// quick-updater project main.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
)

type Config struct {
	Steps []struct {
		Command    string   `json:"command"`
		Args       []string `json:"args"`
		SkipErrors bool     `json:"skip_errors",omitempty`
	} `json:"steps"`
}

var c Config

func main() {
	log.Println("Hello Log!")
	LoadConfig()
	for k, v := range c.Steps {
		log.Printf("Step %d\n", k)
		out, err := exec.Command(v.Command, v.Args...).CombinedOutput()
		if !v.SkipErrors && err != nil {
			log.Fatalln(err)
		}
		log.Println(string(out))
	}
	log.Println("Goodbye Log!")
}

func LoadConfig() {
	file, err := ioutil.ReadFile("updater.json")
	if err != nil {
		log.Panicln(err)
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Panicln(err)
	}
}
