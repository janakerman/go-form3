// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package scheme_messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new scheme messages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for scheme messages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get scheme message API
*/
func (a *GetSchemeMessageRequest) Do() (*GetSchemeMessageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSchemeMessage",
		Method:             "GET",
		PathPattern:        "/notification/schememessages/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetSchemeMessageReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSchemeMessageOK), nil

}

func (a *GetSchemeMessageRequest) MustDo() *GetSchemeMessageOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get scheme message admission API
*/
func (a *GetSchemeMessageAdmissionRequest) Do() (*GetSchemeMessageAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSchemeMessageAdmission",
		Method:             "GET",
		PathPattern:        "/notification/schememessages/{id}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetSchemeMessageAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSchemeMessageAdmissionOK), nil

}

func (a *GetSchemeMessageAdmissionRequest) MustDo() *GetSchemeMessageAdmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
list scheme messages API
*/
func (a *ListSchemeMessagesRequest) Do() (*ListSchemeMessagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListSchemeMessages",
		Method:             "GET",
		PathPattern:        "/notification/schememessages",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListSchemeMessagesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSchemeMessagesOK), nil

}

func (a *ListSchemeMessagesRequest) MustDo() *ListSchemeMessagesOK {
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
