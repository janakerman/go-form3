// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package direct_debits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetDirectDebit creates a new GetDirectDebitRequest object
// with the default values initialized.
func (c *Client) GetDirectDebit() *GetDirectDebitRequest {
	var ()
	return &GetDirectDebitRequest{

		ID: c.Defaults.GetStrfmtUUID("GetDirectDebit", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetDirectDebitRequest struct {

	/*ID      Direct Debit Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetDirectDebitRequest) FromJson(j string) *GetDirectDebitRequest {

	return o
}

func (o *GetDirectDebitRequest) WithID(id strfmt.UUID) *GetDirectDebitRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the get direct debit Request
func (o *GetDirectDebitRequest) WithContext(ctx context.Context) *GetDirectDebitRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get direct debit Request
func (o *GetDirectDebitRequest) WithHTTPClient(client *http.Client) *GetDirectDebitRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetDirectDebitRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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