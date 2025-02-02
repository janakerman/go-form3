// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// CreatePaymentRecallDecisionSubmissionReader is a Reader for the CreatePaymentRecallDecisionSubmission structure.
type CreatePaymentRecallDecisionSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentRecallDecisionSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentRecallDecisionSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentRecallDecisionSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentRecallDecisionSubmissionCreated creates a CreatePaymentRecallDecisionSubmissionCreated with default headers values
func NewCreatePaymentRecallDecisionSubmissionCreated() *CreatePaymentRecallDecisionSubmissionCreated {
	return &CreatePaymentRecallDecisionSubmissionCreated{}
}

/*CreatePaymentRecallDecisionSubmissionCreated handles this case with default header values.

Recall decision submission creation response
*/
type CreatePaymentRecallDecisionSubmissionCreated struct {

	//Payload
	*models.RecallDecisionSubmissionCreationResponse
}

func (o *CreatePaymentRecallDecisionSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/submissions][%d] createPaymentRecallDecisionSubmissionCreated  %+v", 201, o)
}

func (o *CreatePaymentRecallDecisionSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallDecisionSubmissionCreationResponse = new(models.RecallDecisionSubmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.RecallDecisionSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentRecallDecisionSubmissionBadRequest creates a CreatePaymentRecallDecisionSubmissionBadRequest with default headers values
func NewCreatePaymentRecallDecisionSubmissionBadRequest() *CreatePaymentRecallDecisionSubmissionBadRequest {
	return &CreatePaymentRecallDecisionSubmissionBadRequest{}
}

/*CreatePaymentRecallDecisionSubmissionBadRequest handles this case with default header values.

Recall decision submission creation error
*/
type CreatePaymentRecallDecisionSubmissionBadRequest struct {

	//Payload
	*models.APIError
}

func (o *CreatePaymentRecallDecisionSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/submissions][%d] createPaymentRecallDecisionSubmissionBadRequest  %+v", 400, o)
}

func (o *CreatePaymentRecallDecisionSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
