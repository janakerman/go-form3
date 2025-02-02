// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetDirectDebitAdmissionReader is a Reader for the GetDirectDebitAdmission structure.
type GetDirectDebitAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectDebitAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectDebitAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectDebitAdmissionOK creates a GetDirectDebitAdmissionOK with default headers values
func NewGetDirectDebitAdmissionOK() *GetDirectDebitAdmissionOK {
	return &GetDirectDebitAdmissionOK{}
}

/*GetDirectDebitAdmissionOK handles this case with default header values.

Direct Debit Admission details
*/
type GetDirectDebitAdmissionOK struct {

	//Payload
	*models.DirectDebitAdmissionDetailsResponse
}

func (o *GetDirectDebitAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/admissions/{admissionId}][%d] getDirectDebitAdmissionOK  %+v", 200, o)
}

func (o *GetDirectDebitAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitAdmissionDetailsResponse = new(models.DirectDebitAdmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.DirectDebitAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
