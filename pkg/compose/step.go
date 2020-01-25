// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    stepType, err := UnmarshalStepType(bytes)
//    bytes, err = stepType.Marshal()

package compose

import "encoding/json"

func UnmarshalStepType(data []byte) (StepType, error) {
	var r StepType
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StepType) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StepType struct {
	Version  string   `json:"version"`
	Kind     string   `json:"kind"`
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
}

type Metadata struct {
	Name        string       `json:"name"`
	Version     string       `json:"version"`
	IsPublic    bool         `json:"isPublic"`
	Description string       `json:"description"`
	Sources     []string     `json:"sources"`
	Stage       string       `json:"stage"`
	Maintainers []Maintainer `json:"maintainers"`
	Categories  []string     `json:"categories"`
	Official    bool         `json:"official"`
	Tags        []string     `json:"tags"`
	Icon        Icon         `json:"icon"`
	Examples    []Example    `json:"examples"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Latest      bool         `json:"latest"`
	ID          string       `json:"id"`
}

type Example struct {
	Description string                 `json:"description"`
	Workflow    map[string]interface{} `json:"workflow"`
}

type Icon struct {
	Type       string `json:"type"`
	URL        string `json:"url"`
	Background string `json:"background"`
}

type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Spec struct {
	Arguments     string     `json:"arguments"`
	Returns       string     `json:"returns"`
	Delimiters    Delimiters `json:"delimiters"`
	StepsTemplate string     `json:"stepsTemplate"`
}

type Delimiters struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}
