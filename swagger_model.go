package main

type Parameter struct {
	Name        string `json:"name" yaml:",omitempty"`
	In          string `json:"in" yaml:",omitempty"`
	Description string `json:"description" yaml:",omitempty"`
	Required    bool   `json:"required" yaml:",omitempty"`
	Style       string `json:"style" yaml:",omitempty"`
	Explode     bool   `json:"explode" yaml:",omitempty"`
	Schema      struct {
		Type  string `json:"type" yaml:",omitempty"`
		Items struct {
			Type    string   `json:"type" yaml:",omitempty"`
			Default string   `json:"default" yaml:",omitempty"`
			Enum    []string `json:"enum" yaml:",omitempty"`
		} `json:"items" yaml:",omitempty"`
	} `json:"schema" yaml:",omitempty"`
}

type Method struct {
	Tags        []string    `json:"tags" yaml:",omitempty"`
	Summary     string      `json:"summary" yaml:",omitempty"`
	OperationID string      `json:"operationId" yaml:"operationId,omitempty"`
	Parameters  []Parameter `json:"parameters" yaml:",omitempty"`
	RequestBody struct {
		Description string `json:"description" yaml:",omitempty"`
		Content     struct {
			ApplicationJSON struct {
				Schema struct {
					Ref string `json:"$ref" yaml:",omitempty"`
				} `json:"schema" yaml:",omitempty"`
			} `json:"application/json" yaml:",omitempty"`
			ApplicationXML struct {
				Schema struct {
					Ref string `json:"$ref" yaml:",omitempty"`
				} `json:"schema" yaml:",omitempty"`
			} `json:"application/xml" yaml:",omitempty"`
		} `json:"content" yaml:",omitempty"`
		Required bool `json:"required" yaml:",omitempty"`
	} `json:"requestBody" yaml:",omitempty"`
	Responses struct {
		Num200 struct {
			Description string `json:"description"`
			Content     struct {
			} `json:"content"`
		} `json:"200" yaml:"200"`
	} `json:"responses"`
	Security []struct {
		PetstoreAuth []string `json:"petstore_auth" yaml:",omitempty"`
	} `json:"security" yaml:",omitempty"`
	XCodegenRequestBodyName string `json:"x-codegen-request-body-name" yaml:",omitempty"`
}

type Endpoint struct {
	Methods map[string]*Method `json:"methods" yaml:"methods,omitempty,inline"`
}

type Paths struct {
	Endpoints map[string]*Endpoint `json:"paths" yaml:"paths,omitempty"`
}
