// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package mandates

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetMandateReturn creates a new GetMandateReturnRequest object
// with the default values initialized.
func (c *Client) GetMandateReturn() *GetMandateReturnRequest {
	var ()
	return &GetMandateReturnRequest{

		ID: c.Defaults.GetStrfmtUUID("GetMandateReturn", "id"),

		ReturnID: c.Defaults.GetStrfmtUUID("GetMandateReturn", "returnId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetMandateReturnRequest struct {

	/*ID      Mandate Id      */

	ID strfmt.UUID

	/*ReturnID      Return Id      */

	ReturnID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetMandateReturnRequest) FromJson(j string) *GetMandateReturnRequest {

	return o
}

func (o *GetMandateReturnRequest) WithID(id strfmt.UUID) *GetMandateReturnRequest {

	o.ID = id

	return o
}

func (o *GetMandateReturnRequest) WithReturnID(returnID strfmt.UUID) *GetMandateReturnRequest {

	o.ReturnID = returnID

	return o
}

//////////////////
// WithContext adds the context to the get mandate return Request
func (o *GetMandateReturnRequest) WithContext(ctx context.Context) *GetMandateReturnRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get mandate return Request
func (o *GetMandateReturnRequest) WithHTTPClient(client *http.Client) *GetMandateReturnRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetMandateReturnRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
