package main

import (
	"github.com/kristo-godari/form3-client-library/pkg/form3/client/account"
	"github.com/kristo-godari/form3-client-library/pkg/form3/client/models"
	"testing"
	"github.com/stretchr/testify/assert"
)

/**
 * The hostname form3-client-library_accountapi_1 is the container name, update it acordingly.
 * Change "http://form3-client-library_accountapi_1:8080" to "http://localhost:8080", if you are running this outside the docker network.
 */
func TestFullIntegrationFlow(t *testing.T) {

	// given
	form3AccountClient := account.NewForm3AccountClientBuilder().WithBaseUrl("http://form3-client-library_accountapi_1:8080").WithVersion("v1").Build()

	expectedOrganisationId := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	expectedId := "bd27e265-9605-5b4b-a0e5-3003ea9cc419"

	expectedAccountAttributes := &models.AccountAttributes{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "400300",
		BankIDCode:   "GBDSC",
		Bic:          "NWBKGB22",
		Name:         []string{"John Doe"},
	}

	createRequest := &models.CreateRequest{AccountData: &models.AccountData{
		Type:           "accounts",
		ID:             expectedId,
		OrganisationID: expectedOrganisationId,
		Attributes:     expectedAccountAttributes,
	}}


	// when
	createResponse, createError := form3AccountClient.Create(*createRequest)

	// then
	assert.Nil(t, createError)
	assert.NotNil(t, createResponse)
	assert.Equal(t, expectedOrganisationId, createResponse.OrganisationID)
	assert.Equal(t, expectedId, createResponse.ID)
	assert.Equal(t, expectedAccountAttributes, createResponse.Attributes)


	// when
	fetchResponse, fetchError := form3AccountClient.Fetch(expectedId)

	// then
	assert.Nil(t, fetchError)
	assert.Equal(t, expectedOrganisationId, fetchResponse.OrganisationID)
	assert.Equal(t, expectedId, fetchResponse.ID)
	assert.Equal(t, expectedAccountAttributes, fetchResponse.Attributes)


	// when
	deleteResponse, deleteError := form3AccountClient.Delete(expectedId, *createResponse.Version)
	
	// then
	assert.Nil(t, deleteError)
	assert.True(t, deleteResponse)
}
