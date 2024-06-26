package core

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Engine struct {
	Template []*Model
}

func NewEngine(templatePath string) (*Engine, error) {
	e := &Engine{
		Template: make([]*Model, 0),
	}
	m, err := e.ParseTemplate(templatePath)

	if err != nil {
		return nil, err
	}

	e.Template = append(e.Template, m)
	return e, nil
}

func (e *Engine) PrintTemplate() {
	for _, t := range e.Template {
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
}

func (e *Engine) ParseTemplate(templatePath string) (*Model, error) {
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

	return m, nil
}

func (e *Engine) AddTemplate(template string) error {
	m := &Model{}
	err := json.Unmarshal([]byte(template), &m)
	if err != nil {
		return err
	}

	found := false
	for _, t := range e.Template {
		if t.Namespace == m.Namespace {
			found = true
			t.Relations = m.Relations
			break
		}
	}

	if !found {
		e.Template = append(e.Template, m)
	}
	return nil
}
