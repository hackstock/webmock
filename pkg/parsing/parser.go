package parsing

import (
	"encoding/json"
	"io/ioutil"
)

type Endpoint struct {
	URL        string                 `json:"url"`
	HTTPMethod string                 `json:"method"`
	StatusCode int                    `json:"statusCode"`
	Response   map[string]interface{} `json:"response"`
}

type Endpoints struct {
	Values []Endpoint `json:"endpoints"`
}

func ParseEndpoints(filename string) ([]Endpoint, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var endpoints Endpoints
	if err = json.Unmarshal(fileContent, &endpoints); err != nil {
		return nil, err
	}
	return endpoints.Values, nil
}
