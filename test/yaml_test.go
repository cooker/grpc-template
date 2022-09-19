package test

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

var data = `
a: Easy!
`

type Cc struct {
	A string `yaml:"a" json:"a"`
}

func TestYaml(t *testing.T) {
	c := Cc{}

	err := yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	str, err := json.Marshal(c)
	if err == nil {
		fmt.Printf("--- t:\n%s\n\n", str)
	}

}
