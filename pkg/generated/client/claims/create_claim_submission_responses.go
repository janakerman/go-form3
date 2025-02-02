// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// CreateClaimSubmissionReader is a Reader for the CreateClaimSubmission structure.
type CreateClaimSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClaimSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateClaimSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateClaimSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateClaimSubmissionCreated creates a CreateClaimSubmissionCreated with default headers values
func NewCreateClaimSubmissionCreated() *CreateClaimSubmissionCreated {
	return &CreateClaimSubmissionCreated{}
}

/*CreateClaimSubmissionCreated handles this case with default header values.

Claim Submission creation response
*/
type CreateClaimSubmissionCreated struct {

	//Payload
	*models.ClaimSubmissionDetailsResponse
}

func (o *CreateClaimSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/claims/{id}/submissions][%d] createClaimSubmissionCreated  %+v", 201, o)
}

func (o *CreateClaimSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimSubmissionDetailsResponse = new(models.ClaimSubmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.ClaimSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClaimSubmissionBadRequest creates a CreateClaimSubmissionBadRequest with default headers values
func NewCreateClaimSubmissionBadRequest() *CreateClaimSubmissionBadRequest {
	return &CreateClaimSubmissionBadRequest{}
}

/*CreateClaimSubmissionBadRequest handles this case with default header values.

Claim Submission creation error
*/
type CreateClaimSubmissionBadRequest struct {

	//Payload
	*models.APIError
}

func (o *CreateClaimSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/claims/{id}/submissions][%d] createClaimSubmissionBadRequest  %+v", 400, o)
}

func (o *CreateClaimSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
