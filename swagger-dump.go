package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	s := map[string]interface{}{}
	file, err := ioutil.ReadFile("openapi.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(file, &s)
	if err != nil {
		fmt.Println(err)
		return
	}

	paths := Paths{
		Endpoints: map[string]*Endpoint{},
	}

	operationsIDs := map[string]int{}

	for _, obj := range s["servicegroup"].([]interface{}) {
		key := obj.(map[string]interface{})
		tag := key["Key"].(string)

		for _, obj := range key["Services"].([]interface{}) {
			key := obj.(map[string]interface{})

			for _, obj := range key["Endpoints"].([]interface{}) {
				key := obj.(map[string]interface{})

				urlx := key["Url"].(string)
				u, err := url.Parse(urlx)
				if err != nil {
					log.Fatal(err)
				}
				methodPath := strings.TrimPrefix(u.Path, "/sim")

				method := strings.ToLower(key["HttpMethod"].(string))

				if _, ok := paths.Endpoints[methodPath]; !ok {
					paths.Endpoints[methodPath] = &Endpoint{
						Methods: map[string]*Method{},
					}
				}
				e := paths.Endpoints[methodPath]

				if _, ok := e.Methods[method]; !ok {
					e.Methods[method] = &Method{
						Tags: []string{tag},
					}
				}
				m := e.Methods[method]

				m.OperationID = key["Key"].(string)

				operationsIDs[m.OperationID]++
				if operationsIDs[m.OperationID] == 2 {
					m.OperationID = fmt.Sprintf("%s_%s", m.OperationID, tag)
				} else if operationsIDs[m.OperationID] > 2 {
					m.OperationID = fmt.Sprintf("%s_%s%d", m.OperationID, tag, operationsIDs[m.OperationID])
				}

				for _, oobj := range key["Parameters"].([]interface{}) {
					p := Parameter{}
					param := oobj.(map[string]interface{})
					in := strings.ToLower(param["Origin"].(string))
					dataType := strings.ToLower(param["DataType"].(string))
					if in != "body" {
						switch in {
						case "query-string":
							p.In = "query"
						case "route":
							p.In = "path"
						default:
							p.In = in
						}
						DataTypeMap(dataType, &p.Schema)
						p.Name = param["Name"].(string)
						p.Required = param["IsRequired"].(bool)

						m.Parameters = append(m.Parameters, p)
					} else {
						m.RequestBody.Content.ApplicationJSON.Schema.Type = "object"
						if m.RequestBody.Content.ApplicationJSON.Schema.Properties == nil {
							m.RequestBody.Content.ApplicationJSON.Schema.Properties = map[string]Schema{}
						}
						o := Schema{}
						DataTypeMap(dataType, &o)
						m.RequestBody.Content.ApplicationJSON.Schema.Properties[param["Name"].(string)] = o
					}
				}
			}
		}
	}

	d, err := yaml.Marshal(&paths)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(`
openapi: 3.0.1
info:
	title: Swagger Petstore
	description: xxxx
	version: 1.0.0
servers:
  - url: 'https://gateway.saxobank.com/sim/openapi'
	- url: 'https://gateway.saxobank.com/openapi'`)

	fmt.Println(string(d))
}
