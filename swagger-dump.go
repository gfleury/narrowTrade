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
						switch dataType {
						case "bool":
							p.Schema.Type = "boolean"
						case "int":
							p.Schema.Type = "integer"
						case "clientkey":
							p.Schema.Type = "string"
						case "guid":
							p.Schema.Type = "string"
						case "accountkey":
							p.Schema.Type = "string"
						case "datetime":
							p.Schema.Type = "string"
						case "utcdatetime":
							p.Schema.Type = "string"
						case "date":
							p.Schema.Type = "string"
						case "standardperiodtype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"AllTime", "Month", "Quarter", "Year"}
						case "assettype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"Bond", "Cash", "CfdIndexOption",
								"CfdOnFutures", "CfdOnIndex", "CfdOnStock", "ContractFutures",
								"FuturesOption", "FuturesStrategy", "FxBinaryOption", "FxForwards",
								"FxKnockInOption", "FxKnockOutOption", "FxNoTouchOption", "FxOneTouchOption",
								"FxSpot", "FxVanillaOption", "IpoOnStock", "ManagedFund", "MutualFund",
								"Stock", "StockIndex", "StockIndexOption", "StockOption"}
						case "chartfieldgroupspec []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"BlockInfo", "ChartInfo", "Data", "DisplayAndFormat"}
						case "chartrequestmode":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"From", "UpTo"}
						case "signupflowdocumenttype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"AccountViewToIb",
								"EsaContract", "FeePaymentAuthorization", "GeneralBusinessTerms",
								"PensionTransferRequest", "PowerOfAttorney", "PowerOfAttorneyToIb",
								"ProofOfIdentity", "ProofOfResidency", "SourceOfFundsDocument",
								"TermsAndConditions", "TermsAndConditionsAldersopsparingPrivate",
								"TermsAndConditionsKapitalPensionPrivate", "TermsAndConditionsRatePensionPrivate"}
						default:
							p.Schema.Type = dataType
						}
						p.Name = param["Name"].(string)
						p.Required = param["IsRequired"].(bool)

						m.Parameters = append(m.Parameters, p)
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
