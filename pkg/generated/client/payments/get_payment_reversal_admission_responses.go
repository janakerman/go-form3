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

// GetPaymentReversalAdmissionReader is a Reader for the GetPaymentReversalAdmission structure.
type GetPaymentReversalAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReversalAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentReversalAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReversalAdmissionOK creates a GetPaymentReversalAdmissionOK with default headers values
func NewGetPaymentReversalAdmissionOK() *GetPaymentReversalAdmissionOK {
	return &GetPaymentReversalAdmissionOK{}
}

/*GetPaymentReversalAdmissionOK handles this case with default header values.

Reversal admission details
*/
type GetPaymentReversalAdmissionOK struct {

	//Payload
	*models.ReversalAdmissionDetailsResponse
}

func (o *GetPaymentReversalAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/reversals/{reversalId}/admissions/{admissionId}][%d] getPaymentReversalAdmissionOK  %+v", 200, o)
}

func (o *GetPaymentReversalAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReversalAdmissionDetailsResponse = new(models.ReversalAdmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.ReversalAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
