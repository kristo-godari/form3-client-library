Form3 Account Client
===

Form3 Account Client is a simple, fast, compact package for interacting with Form3 Account API.


## Installation

Using this package requires a working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html).

Go Modules are required when using this package. [See the go blog guide on using Go Modules](https://blog.golang.org/using-go-modules).

### Using the client

```
$ GO111MODULE=on go get github.com/kristo.godari/form3-client-api
```

```go
import (
  "ithub.com/kristo.godari/form3-client-api" // imports as package "cli"
)
...
```

```go
// Create a client with default configuration
account.NewForm3AccountClientBuilder().Build()


// Create a client with specific base url and specific version
form3AccountClient := account.NewForm3AccountClientBuilder().WithBaseUrl("http://my-new-base-url.com").WithVersion("v2").Build()
```


```go
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
  }
}
createResponse, _ := form3AccountClient.Create(*createRequest)
```

```go
// Example: Fetch account data account
fetchResponse, _ := form3AccountClient.Fetch("bd27e265-9605-5b4b-a0e5-3003ea9cc419")
```
```go
// Example: Delete account
deleteResponse, _ := form3AccountClient.Delete(createResponse.ID, 1)
```

### GOPATH

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:
```
export PATH=$PATH:$GOPATH/bin
```