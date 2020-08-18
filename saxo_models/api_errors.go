package saxo_models

import (
	"encoding/json"

	"github.com/gfleury/intensiveTrade/saxo_openapi"
)

func GetStringError(err error) string {
	if errAPI, ok := err.(saxo_openapi.GenericSwaggerError); ok {
		return string(errAPI.Body())
	}
	return ""
}

func GetMapError(err error) map[string]interface{} {
	if errAPI, ok := err.(saxo_openapi.GenericSwaggerError); ok {
		mapError := map[string]interface{}{}
		errMarsh := json.Unmarshal(errAPI.Body(), &mapError)
		if errMarsh == nil {
			return mapError
		}
	}
	return nil
}
