package account

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBuilderDefaultSettings(t *testing.T){

	//given
	form3AccountClient := NewForm3AccountClientBuilder().Build()

	//then
	assert.Equal(t, "https://form3.tech", form3AccountClient.baseClient.BaseUrl)
	assert.Equal(t, "v1", form3AccountClient.baseClient.Version)
	assert.Equal(t, "organisation/accounts", form3AccountClient.path)
}

func TestBuilderCustomSettings(t *testing.T){

	//given
	expectedBaseUrl := "http://localhost:8080"
	expectedVersion := "v2"
	form3AccountClient := NewForm3AccountClientBuilder().WithBaseUrl(expectedBaseUrl).WithVersion(expectedVersion).Build()

	//then
	assert.Equal(t, expectedBaseUrl, form3AccountClient.baseClient.BaseUrl)
	assert.Equal(t, expectedVersion, form3AccountClient.baseClient.Version)
	assert.Equal(t, "organisation/accounts", form3AccountClient.path)
}
