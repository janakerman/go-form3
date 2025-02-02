// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create account API
*/
func (a *CreateAccountRequest) Do() (*CreateAccountCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateAccount",
		Method:             "POST",
		PathPattern:        "/organisation/accounts",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateAccountReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAccountCreated), nil

}

func (a *CreateAccountRequest) MustDo() *CreateAccountCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
delete account API
*/
func (a *DeleteAccountRequest) Do() (*DeleteAccountNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAccount",
		Method:             "DELETE",
		PathPattern:        "/organisation/accounts/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &DeleteAccountReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAccountNoContent), nil

}

func (a *DeleteAccountRequest) MustDo() *DeleteAccountNoContent {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get account API
*/
func (a *GetAccountRequest) Do() (*GetAccountOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccount",
		Method:             "GET",
		PathPattern:        "/organisation/accounts/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetAccountReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountOK), nil

}

func (a *GetAccountRequest) MustDo() *GetAccountOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
list accounts API
*/
func (a *ListAccountsRequest) Do() (*ListAccountsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAccounts",
		Method:             "GET",
		PathPattern:        "/organisation/accounts",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListAccountsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAccountsOK), nil

}

func (a *ListAccountsRequest) MustDo() *ListAccountsOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
modify account API
*/
func (a *ModifyAccountRequest) Do() (*ModifyAccountOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyAccount",
		Method:             "PATCH",
		PathPattern:        "/organisation/accounts/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ModifyAccountReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyAccountOK), nil

}

func (a *ModifyAccountRequest) MustDo() *ModifyAccountOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
