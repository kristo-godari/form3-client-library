package client

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetBaseUrlAndVersion(t *testing.T) {

	// given
	expectedBaseUrl := "testUrl"
	expectedVersion := "v1"
	expectedBaseUrlAndVersion := "testUrl/v1"

	// when
	form3Client := Form3BaseClient{
		BaseUrl:    expectedBaseUrl,
		Version:    expectedVersion,
		HttpClient: http.Client{},
	}

	// then
	assert.Equal(t, expectedBaseUrl, form3Client.GetBaseUrl(), "Base Url should be equal.")
	assert.Equal(t, expectedVersion, form3Client.GetVersion(), "Versions should be equal.")
	assert.Equal(t, expectedBaseUrlAndVersion, form3Client.GetBaseUrlAndVersion(), "Base Url and Version should be equal.")
}


func TestVladiateAndReadResponseReturnsError(t *testing.T){

	// given
	form3Client := Form3BaseClient{
		BaseUrl:   "nil",
		Version:   "nil",
		HttpClient: http.Client{},
	}
	expectedError := errors.New("Connection error.")
	stringReadCloser := io.NopCloser(strings.NewReader("{body}"))

	httpResponse := http.Response{}
	httpResponse.StatusCode = 200
	httpResponse.Body = stringReadCloser

	// when
	responseBytes, err:=form3Client.ValidateAndReadResponse(200, &httpResponse, expectedError)

	// then
	assert.NotNil(t, err)
	assert.Nil(t, responseBytes)
	assert.Equal(t, expectedError, err)
}

func TestVladiateAndReadResponse(t *testing.T){

	// given
	form3Client := Form3BaseClient{
		BaseUrl:   "nil",
		Version:   "nil",
		HttpClient: http.Client{},
	}

	expectedBody := "{\"success\": \"true\"}"
	stringReadCloser := io.NopCloser(strings.NewReader(expectedBody))

	httpResponse := http.Response{}
	httpResponse.StatusCode = 200
	httpResponse.Body = stringReadCloser

	// when
	responseBytes, err:=form3Client.ValidateAndReadResponse(200, &httpResponse, nil)

	// then
	assert.Nil(t, err)
	assert.NotNil(t, responseBytes)
	assert.Equal(t, expectedBody, string(responseBytes))
}