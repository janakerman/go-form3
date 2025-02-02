// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package payments

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetPaymentsHealth creates a new GetPaymentsHealthRequest object
// with the default values initialized.
func (c *Client) GetPaymentsHealth() *GetPaymentsHealthRequest {

	return &GetPaymentsHealthRequest{

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentsHealthRequest struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentsHealthRequest) FromJson(j string) *GetPaymentsHealthRequest {

	return o
}

//////////////////
// WithContext adds the context to the get payments health Request
func (o *GetPaymentsHealthRequest) WithContext(ctx context.Context) *GetPaymentsHealthRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payments health Request
func (o *GetPaymentsHealthRequest) WithHTTPClient(client *http.Client) *GetPaymentsHealthRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentsHealthRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
