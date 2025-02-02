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

// Client.GetUserAces creates a new GetUserAcesRequest object
// with the default values initialized.
func (c *Client) GetUserAces() *GetUserAcesRequest {
	var ()
	return &GetUserAcesRequest{

		FilterAction: c.Defaults.GetStringPtr("GetUserAces", "filter[action]"),

		FilterRecordType: c.Defaults.GetStringPtr("GetUserAces", "filter[record_type]"),

		UserID: c.Defaults.GetStrfmtUUID("GetUserAces", "user_id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetUserAcesRequest struct {

	/*FilterAction      Access action      */

	FilterAction *string

	/*FilterRecordType      Record type      */

	FilterRecordType *string

	/*UserID      User Id      */

	UserID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetUserAcesRequest) FromJson(j string) *GetUserAcesRequest {

	return o
}

func (o *GetUserAcesRequest) WithFilterAction(filterAction string) *GetUserAcesRequest {

	o.FilterAction = &filterAction

	return o
}

func (o *GetUserAcesRequest) WithoutFilterAction() *GetUserAcesRequest {

	o.FilterAction = nil

	return o
}

func (o *GetUserAcesRequest) WithFilterRecordType(filterRecordType string) *GetUserAcesRequest {

	o.FilterRecordType = &filterRecordType

	return o
}

func (o *GetUserAcesRequest) WithoutFilterRecordType() *GetUserAcesRequest {

	o.FilterRecordType = nil

	return o
}

func (o *GetUserAcesRequest) WithUserID(userID strfmt.UUID) *GetUserAcesRequest {

	o.UserID = userID

	return o
}

//////////////////
// WithContext adds the context to the get user aces Request
func (o *GetUserAcesRequest) WithContext(ctx context.Context) *GetUserAcesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get user aces Request
func (o *GetUserAcesRequest) WithHTTPClient(client *http.Client) *GetUserAcesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetUserAcesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterAction != nil {

		// query param filter[action]
		var qrFilterAction string
		if o.FilterAction != nil {
			qrFilterAction = *o.FilterAction
		}
		qFilterAction := qrFilterAction
		if qFilterAction != "" {
			if err := r.SetQueryParam("filter[action]", qFilterAction); err != nil {
				return err
			}
		}

	}

	if o.FilterRecordType != nil {

		// query param filter[record_type]
		var qrFilterRecordType string
		if o.FilterRecordType != nil {
			qrFilterRecordType = *o.FilterRecordType
		}
		qFilterRecordType := qrFilterRecordType
		if qFilterRecordType != "" {
			if err := r.SetQueryParam("filter[record_type]", qFilterRecordType); err != nil {
				return err
			}
		}

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
