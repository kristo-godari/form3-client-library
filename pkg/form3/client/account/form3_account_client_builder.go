package account

import (
	"github.com/kristo-godari/form3-client-library/internal/form3/client"
	"net/http"
)

const accountPath = "organisation/accounts"

type Form3AccountClientBuilder struct {
	form3AccountClient Form3AccountClient
}

func (form3AccountClientBuilder *Form3AccountClientBuilder) WithBaseUrl(baseUrl string) *Form3AccountClientBuilder {
	form3AccountClientBuilder.form3AccountClient.baseClient.WithBaseUrl(baseUrl)
	return form3AccountClientBuilder
}

func (form3AccountClientBuilder *Form3AccountClientBuilder) WithVersion(version string)  *Form3AccountClientBuilder {
	form3AccountClientBuilder.form3AccountClient.baseClient.WithVersion(version)
	return form3AccountClientBuilder
}

func (form3AccountClientBuilder *Form3AccountClientBuilder) Build() Form3AccountClient {
	return  form3AccountClientBuilder.form3AccountClient
}

func (form3AccountClientBuilder *Form3AccountClientBuilder) setAccountClient(form3AccountClient Form3AccountClient) {
	form3AccountClientBuilder.form3AccountClient = form3AccountClient
}

func NewForm3AccountClientBuilder()  *Form3AccountClientBuilder {

	form3AccountClient := Form3AccountClient{
		baseClient: &client.Form3BaseClient{
			BaseUrl:    client.Form3BaseUrl,
			Version:    client.Form3Version,
			HttpClient: http.Client{},
		},
		path: accountPath,
	}

	return  &Form3AccountClientBuilder{form3AccountClient: form3AccountClient}
}
