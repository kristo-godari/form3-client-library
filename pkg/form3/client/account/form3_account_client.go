package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-hello-world/internal/form3/client"
	"go-hello-world/pkg/form3/client/models"
	"net/http"
)

const contentType = "application/vnd.form3+json"

type Form3AccountClient struct {
	baseClient *client.Form3BaseClient
	path       string
}

func (form3AccountClient *Form3AccountClient) Create(createRequest models.CreateRequest) (*models.CreateResponse, error) {

	var expectedStatusCode = 201
	var createResponse models.CreateResponse

	bodyEncoded, _ := json.Marshal(createRequest)
	ioReaderBody := bytes.NewBuffer(bodyEncoded)

	url := fmt.Sprintf("%s/%s", form3AccountClient.baseClient.GetBaseUrlAndVersion(), form3AccountClient.path)
	resp, err := http.Post(url, contentType, ioReaderBody)

	responseInBytes, err := form3AccountClient.baseClient.ValidateAndReadResponse(expectedStatusCode, resp, err)

	if err != nil {
		json.Unmarshal(responseInBytes, &createResponse)
		if responseInBytes != nil{
			return  &createResponse, err
		}
		return nil, err
	}

	json.Unmarshal(responseInBytes, &createResponse)

	return &createResponse, nil
}

func (form3AccountClient *Form3AccountClient) Fetch(accountId string) (*models.FetchResponse, error) {
	var expectedStatusCode = 200
	var fetchResponse models.FetchResponse

	url := fmt.Sprintf("%s/%s/%s", form3AccountClient.baseClient.GetBaseUrlAndVersion(), form3AccountClient.path, accountId)
	response, err := http.Get(url)

	responseInBytes, err := form3AccountClient.baseClient.ValidateAndReadResponse(expectedStatusCode, response, err)

	if err != nil {
		json.Unmarshal(responseInBytes, &fetchResponse)
		if responseInBytes != nil{
			return  &fetchResponse, err
		}
		return nil, err
	}

	json.Unmarshal(responseInBytes, &fetchResponse)

	return &fetchResponse, nil
}

func (form3AccountClient *Form3AccountClient) Delete(accountId string, version int64) (bool,error) {
	var expectedStatusCode = 204

	url := fmt.Sprintf("%s/%s/%s?version=%d", form3AccountClient.baseClient.GetBaseUrlAndVersion(), form3AccountClient.path, accountId, version)
	response, err := http.NewRequest("DELETE", url, nil)

	resp, err := form3AccountClient.baseClient.HttpClient.Do(response)

	_, err = form3AccountClient.baseClient.ValidateAndReadResponse(expectedStatusCode, resp, err)

	if err != nil {
		return false, err
	}

	return true, nil
}
