// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package direct_debits

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

// Client.CreateDirectDebitReversal creates a new CreateDirectDebitReversalRequest object
// with the default values initialized.
func (c *Client) CreateDirectDebitReversal() *CreateDirectDebitReversalRequest {
	var ()
	return &CreateDirectDebitReversalRequest{

		DirectDebitReversalCreation: models.DirectDebitReversalCreationWithDefaults(c.Defaults),

		ID: c.Defaults.GetStrfmtUUID("CreateDirectDebitReversal", "id"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateDirectDebitReversalRequest struct {

	/*ReversalCreationRequest*/

	*models.DirectDebitReversalCreation

	/*ID      DirectDebit Id      */

	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateDirectDebitReversalRequest) FromJson(j string) *CreateDirectDebitReversalRequest {

	var m models.DirectDebitReversalCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		log.Fatal(err)
	}

	o.DirectDebitReversalCreation = &m

	return o
}

func (o *CreateDirectDebitReversalRequest) WithReversalCreationRequest(reversalCreationRequest models.DirectDebitReversalCreation) *CreateDirectDebitReversalRequest {

	o.DirectDebitReversalCreation = &reversalCreationRequest

	return o
}

func (o *CreateDirectDebitReversalRequest) WithoutReversalCreationRequest() *CreateDirectDebitReversalRequest {

	o.DirectDebitReversalCreation = &models.DirectDebitReversalCreation{}

	return o
}

func (o *CreateDirectDebitReversalRequest) WithID(id strfmt.UUID) *CreateDirectDebitReversalRequest {

	o.ID = id

	return o
}

//////////////////
// WithContext adds the context to the create direct debit reversal Request
func (o *CreateDirectDebitReversalRequest) WithContext(ctx context.Context) *CreateDirectDebitReversalRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create direct debit reversal Request
func (o *CreateDirectDebitReversalRequest) WithHTTPClient(client *http.Client) *CreateDirectDebitReversalRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateDirectDebitReversalRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.DirectDebitReversalCreation != nil {
		if err := r.SetBodyParam(o.DirectDebitReversalCreation); err != nil {
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
