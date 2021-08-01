package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const Form3BaseUrl = "https://form3.tech"
const Form3Version = "v1"

type Form3BaseClient struct {
	BaseUrl    string
	Version    string
	HttpClient http.Client
}

func (form3BaseClient *Form3BaseClient) WithBaseUrl(baseUrl string) {
	form3BaseClient.BaseUrl = baseUrl
}

func (form3BaseClient *Form3BaseClient) WithVersion(version string) {
	form3BaseClient.Version = version
}

func (form3BaseClient *Form3BaseClient) GetBaseUrl()  string{
	return form3BaseClient.BaseUrl
}

func (form3BaseClient *Form3BaseClient) GetVersion()  string{
	return form3BaseClient.Version
}

func (form3BaseClient *Form3BaseClient) GetBaseUrlAndVersion()  string{
	return fmt.Sprintf("%s/%s", form3BaseClient.BaseUrl, form3BaseClient.Version)
}

func (form3BaseClient *Form3BaseClient) ValidateAndReadResponse(expectedStatusCode int, resp *http.Response, err error) ([]byte, error) {

	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatusCode {
		statusCodeError := errors.New( fmt.Sprintf("Status code does not match, expected %d, got %d.", expectedStatusCode, resp.StatusCode))
		if resp.Body != nil{
			responseInBytes, _ := form3BaseClient.readResponse(resp)
			return responseInBytes, statusCodeError
		}
		return nil, statusCodeError
	}

	return form3BaseClient.readResponse(resp)
}


func (form3BaseClient *Form3BaseClient) readResponse( resp *http.Response) ([]byte, error) {

	responseInBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return responseInBytes, nil
}