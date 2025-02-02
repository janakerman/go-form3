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

// GetPaymentRecallReversalAdmissionReader is a Reader for the GetPaymentRecallReversalAdmission structure.
type GetPaymentRecallReversalAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRecallReversalAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentRecallReversalAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRecallReversalAdmissionOK creates a GetPaymentRecallReversalAdmissionOK with default headers values
func NewGetPaymentRecallReversalAdmissionOK() *GetPaymentRecallReversalAdmissionOK {
	return &GetPaymentRecallReversalAdmissionOK{}
}

/*GetPaymentRecallReversalAdmissionOK handles this case with default header values.

Reversal admission details
*/
type GetPaymentRecallReversalAdmissionOK struct {

	//Payload
	*models.RecallReversalAdmissionDetailsResponse
}

func (o *GetPaymentRecallReversalAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/reversals/{reversalId}/admissions/{admissionId}][%d] getPaymentRecallReversalAdmissionOK  %+v", 200, o)
}

func (o *GetPaymentRecallReversalAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallReversalAdmissionDetailsResponse = new(models.RecallReversalAdmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.RecallReversalAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
