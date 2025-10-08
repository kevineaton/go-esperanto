package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
)

// TestEndpoint allows an easy way to test HTTP end points in unit testing
func TestEndpoint(method string, endpoint string, data io.Reader, handler http.HandlerFunc, valid bool) (code int, body *bytes.Buffer, err error) {
	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		return 500, nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	rr := httptest.NewRecorder()

	chi := SetupRouter()
	if valid {
		c := context.WithValue(req.Context(), appContextAuthenticationFound, valid)
		req.Header.Add("X-API-TOKEN", config.AuthenticationToken)
		req = req.WithContext(c)
	}
	chi.ServeHTTP(rr, req)

	return rr.Code, rr.Body, nil
}

func testEndpointResultToStruct(bu *bytes.Buffer, target any) error {
	m := &apiReturn{}
	err := json.Unmarshal(bu.Bytes(), m)
	if err != nil {
		return err
	}
	jsonb, err := json.Marshal(m.Data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonb, target)
	return err
}
