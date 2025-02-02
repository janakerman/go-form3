// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package a_c_e

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.ListAces creates a new ListAcesRequest object
// with the default values initialized.
func (c *Client) ListAces() *ListAcesRequest {
	var ()
	return &ListAcesRequest{

		RoleID: c.Defaults.GetStrfmtUUID("ListAces", "role_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListAcesRequest struct {

	/*RoleID      Role Id      */

	RoleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ListAcesRequest) FromJson(j string) *ListAcesRequest {

	return o
}

func (o *ListAcesRequest) WithRoleID(roleID strfmt.UUID) *ListAcesRequest {

	o.RoleID = roleID

	return o
}

//////////////////
// WithContext adds the context to the list aces Request
func (o *ListAcesRequest) WithContext(ctx context.Context) *ListAcesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list aces Request
func (o *ListAcesRequest) WithHTTPClient(client *http.Client) *ListAcesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListAcesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
