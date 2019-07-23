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

// Client.GetPaymentSubmission creates a new GetPaymentSubmissionRequest object
// with the default values initialized.
func (c *Client) GetPaymentSubmission() *GetPaymentSubmissionRequest {
	var ()
	return &GetPaymentSubmissionRequest{

		ID: c.Defaults.GetStrfmtUUID("GetPaymentSubmission", "id"),

		SubmissionID: c.Defaults.GetStrfmtUUID("GetPaymentSubmission", "submissionId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentSubmissionRequest struct {

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*SubmissionID      Submission Id      */

	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentSubmissionRequest) FromJson(j string) *GetPaymentSubmissionRequest {

	return o
}

func (o *GetPaymentSubmissionRequest) WithID(id strfmt.UUID) *GetPaymentSubmissionRequest {

	o.ID = id

	return o
}

func (o *GetPaymentSubmissionRequest) WithSubmissionID(submissionID strfmt.UUID) *GetPaymentSubmissionRequest {

	o.SubmissionID = submissionID

	return o
}

//////////////////
// WithContext adds the context to the get payment submission Request
func (o *GetPaymentSubmissionRequest) WithContext(ctx context.Context) *GetPaymentSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment submission Request
func (o *GetPaymentSubmissionRequest) WithHTTPClient(client *http.Client) *GetPaymentSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}