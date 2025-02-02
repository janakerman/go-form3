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

// Client.DeleteUserCredential creates a new DeleteUserCredentialRequest object
// with the default values initialized.
func (c *Client) DeleteUserCredential() *DeleteUserCredentialRequest {
	var ()
	return &DeleteUserCredentialRequest{

		ClientID: c.Defaults.GetString("DeleteUserCredential", "client_id"),

		UserID: c.Defaults.GetStrfmtUUID("DeleteUserCredential", "user_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type DeleteUserCredentialRequest struct {

	/*ClientID      client id      */

	ClientID string

	/*UserID      User Id      */

	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *DeleteUserCredentialRequest) FromJson(j string) *DeleteUserCredentialRequest {

	return o
}

func (o *DeleteUserCredentialRequest) WithClientID(clientID string) *DeleteUserCredentialRequest {

	o.ClientID = clientID

	return o
}

func (o *DeleteUserCredentialRequest) WithUserID(userID strfmt.UUID) *DeleteUserCredentialRequest {

	o.UserID = userID

	return o
}

//////////////////
// WithContext adds the context to the delete user credential Request
func (o *DeleteUserCredentialRequest) WithContext(ctx context.Context) *DeleteUserCredentialRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the delete user credential Request
func (o *DeleteUserCredentialRequest) WithHTTPClient(client *http.Client) *DeleteUserCredentialRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *DeleteUserCredentialRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
