// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.ListUserCredentials creates a new ListUserCredentialsRequest object
// with the default values initialized.
func (c *Client) ListUserCredentials() *ListUserCredentialsRequest {
	var ()
	return &ListUserCredentialsRequest{

		UserID: c.Defaults.GetStrfmtUUID("ListUserCredentials", "user_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListUserCredentialsRequest struct {

	/*UserID      User Id      */

	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ListUserCredentialsRequest) FromJson(j string) *ListUserCredentialsRequest {

	return o
}

func (o *ListUserCredentialsRequest) WithUserID(userID strfmt.UUID) *ListUserCredentialsRequest {

	o.UserID = userID

	return o
}

//////////////////
// WithContext adds the context to the list user credentials Request
func (o *ListUserCredentialsRequest) WithContext(ctx context.Context) *ListUserCredentialsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list user credentials Request
func (o *ListUserCredentialsRequest) WithHTTPClient(client *http.Client) *ListUserCredentialsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListUserCredentialsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
