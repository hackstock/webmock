package parsing

import (
	"encoding/json"
	"io/ioutil"
)

type Endpoint struct {
	URL        string      `json:"url"`
	HTTPMethod string      `json:"method"`
	StatusCode int         `json:"statusCode"`
	Response   interface{} `json:"response"`
}

type Endpoints struct {
	Values []Endpoint `json:"endpoints"`
}

func ParseEndpoints(filename string) (map[string]Endpoint, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var endpoints Endpoints
	if err = json.Unmarshal(fileContent, &endpoints); err != nil {
		return nil, err
	}

	endpointsLookup := make(map[string]Endpoint)
	for _, endpoint := range endpoints.Values {
		endpointsLookup[endpoint.URL] = endpoint
	}

	return endpointsLookup, nil
}
