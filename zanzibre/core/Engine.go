package core

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Engine struct {
	Template *Model
}

func NewEngine(templatePath string) (*Engine, error) {
	jsonFile, err := os.Open(templatePath)

	if err != nil {
		return nil, fmt.Errorf("failed to open template at path %s with error message %s", templatePath, err.Error())
	}

	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read json file %s with error message %s", templatePath, err.Error())
	}

	m := &Model{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, fmt.Errorf("failed to map json file template to model with error message %s", err.Error())
	}

	e := &Engine{
		Template: m,
	}

	return e, nil
}

func (e *Engine) PrintTemplate() {
	t := e.Template
	fmt.Printf("Namespace: %s\n", t.Namespace)
	for i, r := range t.Relations {
		fmt.Printf("\tRelation %d: %s\n", i, r.Name)
		for j, info := range r.AdditionalInfo {
			fmt.Printf("\t\tInfo %d: %s\n", j, info.Type)
			for k, c := range info.Children {
				fmt.Printf("\t\t\tChild %d: %s\n", k, c.RelationName)
			}
		}
	}
}
