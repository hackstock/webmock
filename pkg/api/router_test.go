package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hackstock/webmock/pkg/parsing"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_Router(t *testing.T) {
	endpoints, err := parsing.ParseEndpoints("./testdata/test_endpoints.json")
	assert.NoError(t, err)

	logger := zap.NewExample()

	router := InitRoutes(endpoints, logger)

	for _, endpoint := range endpoints {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(endpoint.HTTPMethod, endpoint.URL, nil)
		assert.NoError(t, err)

		router.ServeHTTP(w, req)

		assert.Equal(t, endpoint.StatusCode, w.Code)

		resBody, err := json.Marshal(endpoint.Response)
		assert.NoError(t, err)
		assert.Equal(t, resBody, w.Body.Bytes())
	}
}
