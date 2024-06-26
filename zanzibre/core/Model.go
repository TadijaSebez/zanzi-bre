package core

type Model struct {
	Namespace string      `json:"namespace"`
	Relations []*Relation `json:"relations"`
}

type Relation struct {
	Name           string  `json:"name"`
	AdditionalInfo []*Info `json:"additional"`
}

type Info struct {
	Type     string   `json:"type"`
	Children []*Child `json:"children"`
}

type Child struct {
	RelationName string `json:"child"`
}

type Grammar struct {
	Namespace string `json:"namespace"`
}
