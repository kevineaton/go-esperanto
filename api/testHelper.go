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

// UnmarshalSliceFromTestRoute unmarshals a response that is an array in the data field
func unmarshalSliceFromTestRoute(body *bytes.Buffer) ([]interface{}, error) {
	response := apiReturn{}
	ret := []interface{}{}
	retBuf := new(bytes.Buffer)
	retBuf.ReadFrom(body)
	err := json.Unmarshal(retBuf.Bytes(), &response)
	if err != nil {
		return []interface{}{}, err
	}
	retBody, ok := response.Data.([]interface{})
	if ok {
		ret = retBody
	}
	return ret, nil
}

func unmarshalMapFromTestRoute(body *bytes.Buffer) (map[string]interface{}, error) {
	response := apiReturn{}
	ret := map[string]interface{}{}
	retBuf := new(bytes.Buffer)
	retBuf.ReadFrom(body)
	err := json.Unmarshal(retBuf.Bytes(), &response)
	if err != nil {
		return map[string]interface{}{}, err
	}
	retBody, ok := response.Data.(map[string]interface{})
	if ok {
		ret = retBody
	}

	return ret, nil
}
