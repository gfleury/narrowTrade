package saxo_oauth2

import (
	"encoding/json"

	"golang.org/x/oauth2"
)

func LoadConfiguration(data []byte) (*oauth2.Config, error) {
	oa := &oauth2.Config{}

	err := json.Unmarshal(data, oa)
	if err != nil {
		return nil, err
	}

	return oa, nil
}
