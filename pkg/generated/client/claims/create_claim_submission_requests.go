// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package claims

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// Client.CreateClaimSubmission creates a new CreateClaimSubmissionRequest object
// with the default values initialized.
func (c *Client) CreateClaimSubmission() *CreateClaimSubmissionRequest {
	var ()
	return &CreateClaimSubmissionRequest{

		ClaimSubmissionCreation: models.ClaimSubmissionCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateClaimSubmission", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateClaimSubmissionRequest struct {

	/*ClaimSubmissionCreationRequest*/

	*models.ClaimSubmissionCreation

	/*ID      Claim Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateClaimSubmissionRequest) FromJson(j string) *CreateClaimSubmissionRequest {

	var m models.ClaimSubmissionCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.ClaimSubmissionCreation = &m

	return o
}

func (o *CreateClaimSubmissionRequest) WithClaimSubmissionCreationRequest(claimSubmissionCreationRequest models.ClaimSubmissionCreation) *CreateClaimSubmissionRequest {

	o.ClaimSubmissionCreation = &claimSubmissionCreationRequest

	return o
}

func (o *CreateClaimSubmissionRequest) WithoutClaimSubmissionCreationRequest() *CreateClaimSubmissionRequest {

	o.ClaimSubmissionCreation = &models.ClaimSubmissionCreation{}

	return o
}

func (o *CreateClaimSubmissionRequest) WithID(id strfmt.UUID) *CreateClaimSubmissionRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create claim submission Request
func (o *CreateClaimSubmissionRequest) WithContext(ctx context.Context) *CreateClaimSubmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create claim submission Request
func (o *CreateClaimSubmissionRequest) WithHTTPClient(client *http.Client) *CreateClaimSubmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateClaimSubmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ClaimSubmissionCreation != nil {
		if err := r.SetBodyParam(o.ClaimSubmissionCreation); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
