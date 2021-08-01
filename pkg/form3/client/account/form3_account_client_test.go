package account

import (
	"encoding/json"
	"github.com/kristo-godari/form3-client-library/pkg/form3/client/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

const expectedResponseAsString = `{
	"data":{
	   "attributes":{
		  "bank_id":"400300",
		  "bank_id_code":"GBDSC",
		  "base_currency":"GBP",
		  "bic":"NWBKGB22",
		  "country":"GB",
		  "name":[
			 "John Doe"
		  ]
	   },
	   "id":"bd27e265-9605-5b4b-a0e5-3003ea9cc419",
	   "organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	   "type":"accounts",
	   "version":0,
	   "created_on":"2021-07-31T11:46:10.55Z",
	   "modified_on":"2021-07-31T11:46:10.55Z"
	},
	"links":{
	   "self":"/v1/organisation/accounts/bd27e265-9605-5b4b-a0e5-3003ea9cc419"
	}
 }`

func TestCreateMethod(t *testing.T){

	// given
	var expectedResponse models.CreateResponse

	json.Unmarshal([]byte(expectedResponseAsString), &expectedResponse)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/v1/organisation/accounts")
		rw.Write([]byte(expectedResponseAsString))
	}))
	defer server.Close()

	form3AccountClient := NewForm3AccountClientBuilder().WithBaseUrl(server.URL).Build()

	expectedAccountData := models.AccountData{
		Type: "accounts",
		ID: "bd27e265-9605-5b4b-a0e5-3003ea9cc419",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Attributes: &models.AccountAttributes{
			Country: "GB",
			BaseCurrency: "GBP",
			BankID: "400300",
			BankIDCode: "GBDSC",
			Bic: "NWBKGB22",
			Name: []string{"John Doe"},
		},
	}
	createRequest := &models.CreateRequest{AccountData: &expectedAccountData}
	
	// when
	createResponse, _ := form3AccountClient.Create(*createRequest)

	// then
	assert.Equal(t, expectedResponse.OrganisationID, createResponse.OrganisationID)
	assert.Equal(t, expectedResponse.AccountData.Attributes, createResponse.AccountData.Attributes)
}


func TestFetchMethod(t *testing.T){
	
	// given
	var expectedResponse models.FetchResponse
	accountId := "bd27e265-9605-5b4b-a0e5-3003ea9cc419"

	json.Unmarshal([]byte(expectedResponseAsString), &expectedResponse)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/v1/organisation/accounts/" + accountId)
		rw.Write([]byte(expectedResponseAsString))
	}))
	defer server.Close()

	form3AccountClient := NewForm3AccountClientBuilder().WithBaseUrl(server.URL).Build()

	// when
	fetchResponse, err := form3AccountClient.Fetch(accountId)

	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse.OrganisationID, fetchResponse.OrganisationID)
	assert.Equal(t, expectedResponse.AccountData.Attributes, fetchResponse.AccountData.Attributes)
}


func TestDeleteMethod(t *testing.T){
	
	// given
	var expectedResponse models.FetchResponse
	id := "bd27e265-9605-5b4b-a0e5-3003ea9cc419"
	version := 0

	json.Unmarshal([]byte(expectedResponseAsString), &expectedResponse)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/v1/organisation/accounts/" + id + "?version=0")
		rw.WriteHeader(http.StatusNoContent)
		rw.Write([]byte(``))
	}))
	defer server.Close()

	form3AccountClient := NewForm3AccountClientBuilder().WithBaseUrl(server.URL).Build()

	// when
	deleteResponse, err := form3AccountClient.Delete(id, int64(version))

	// then
	assert.Nil(t, err)
	assert.True(t, deleteResponse)
}