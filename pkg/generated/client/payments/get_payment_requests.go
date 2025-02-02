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

// Client.GetPayment creates a new GetPaymentRequest object
// with the default values initialized.
func (c *Client) GetPayment() *GetPaymentRequest {
	var ()
	return &GetPaymentRequest{

		ID: c.Defaults.GetStrfmtUUID("GetPayment", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentRequest struct {

	/*ID      Payment Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentRequest) FromJson(j string) *GetPaymentRequest {

	return o
}

func (o *GetPaymentRequest) WithID(id strfmt.UUID) *GetPaymentRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get payment Request
func (o *GetPaymentRequest) WithContext(ctx context.Context) *GetPaymentRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment Request
func (o *GetPaymentRequest) WithHTTPClient(client *http.Client) *GetPaymentRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
