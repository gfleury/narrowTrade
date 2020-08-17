package saxo_models

import (
	"encoding/json"

	"github.com/gfleury/intensiveTrade/saxo_openapi"
)

func GetStringError(err error) string {
	if errAPI, ok := err.(saxo_openapi.GenericSwaggerError); ok {
		errJson, err := json.Marshal(errAPI.Model())
		if err == nil {
			return string(errJson)
		}
	}
	return ""
}
