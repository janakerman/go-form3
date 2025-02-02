// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get report API
*/
func (a *GetReportRequest) Do() (*GetReportOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReport",
		Method:             "GET",
		PathPattern:        "/notification/reports/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetReportReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReportOK), nil

}

func (a *GetReportRequest) MustDo() *GetReportOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get report admission API
*/
func (a *GetReportAdmissionRequest) Do() (*GetReportAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReportAdmission",
		Method:             "GET",
		PathPattern:        "/notification/reports/{id}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetReportAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReportAdmissionOK), nil

}

func (a *GetReportAdmissionRequest) MustDo() *GetReportAdmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
list reports API
*/
func (a *ListReportsRequest) Do() (*ListReportsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListReports",
		Method:             "GET",
		PathPattern:        "/notification/reports",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListReportsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReportsOK), nil

}

func (a *ListReportsRequest) MustDo() *ListReportsOK {
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
