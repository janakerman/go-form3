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

// GetClaimReversalSubmissionReader is a Reader for the GetClaimReversalSubmission structure.
type GetClaimReversalSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClaimReversalSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClaimReversalSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetClaimReversalSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClaimReversalSubmissionOK creates a GetClaimReversalSubmissionOK with default headers values
func NewGetClaimReversalSubmissionOK() *GetClaimReversalSubmissionOK {
	return &GetClaimReversalSubmissionOK{}
}

/*GetClaimReversalSubmissionOK handles this case with default header values.

Claim Reversal Submission details
*/
type GetClaimReversalSubmissionOK struct {

	//Payload
	*models.ClaimReversalSubmissionDetailsResponse
}

func (o *GetClaimReversalSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}/submissions/{submissionId}][%d] getClaimReversalSubmissionOK  %+v", 200, o)
}

func (o *GetClaimReversalSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ClaimReversalSubmissionDetailsResponse = new(models.ClaimReversalSubmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.ClaimReversalSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClaimReversalSubmissionBadRequest creates a GetClaimReversalSubmissionBadRequest with default headers values
func NewGetClaimReversalSubmissionBadRequest() *GetClaimReversalSubmissionBadRequest {
	return &GetClaimReversalSubmissionBadRequest{}
}

/*GetClaimReversalSubmissionBadRequest handles this case with default header values.

Error
*/
type GetClaimReversalSubmissionBadRequest struct {

	//Payload
	*models.APIError
}

func (o *GetClaimReversalSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/claims/{id}/reversals/{reversalId}/submissions/{submissionId}][%d] getClaimReversalSubmissionBadRequest  %+v", 400, o)
}

func (o *GetClaimReversalSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
