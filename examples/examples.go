package examples

import (
	"encoding/json"
	"fmt"
	"go-hello-world/pkg/form3/client/account"
	"go-hello-world/pkg/form3/client/models"
)

func main(){

	// Create a client with default configuration
	account.NewForm3AccountClientBuilder().Build()


	// Create a client with specific base url and specific version
	form3AccountClient := account.NewForm3AccountClientBuilder().WithBaseUrl("http://localhost:8080").WithVersion("v1").Build()


	// Example: Create account
	createRequest := &models.CreateRequest{AccountData: &models.AccountData{
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
	}}
	createResponse, _ := form3AccountClient.Create(*createRequest)
	createResponseJson, _ := json.Marshal(createResponse)
	fmt.Printf("%s\n", createResponseJson)


	// Example: Fetch account data account
	fetchResponse, _ := form3AccountClient.Fetch("bd27e265-9605-5b4b-a0e5-3003ea9cc419")
	fetchResponseJson, _ := json.Marshal(fetchResponse)
	fmt.Printf("%s\n", fetchResponseJson)


	// Example: Delete account
	deleteResponse, _ := form3AccountClient.Delete(createResponse.ID, *createResponse.Version)
	fmt.Println(deleteResponse)
}

