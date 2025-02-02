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

// ListPaymentsReader is a Reader for the ListPayments structure.
type ListPaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPaymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPaymentsOK creates a ListPaymentsOK with default headers values
func NewListPaymentsOK() *ListPaymentsOK {
	return &ListPaymentsOK{}
}

/*ListPaymentsOK handles this case with default header values.

List of payment details
*/
type ListPaymentsOK struct {

	//Payload
	*models.PaymentDetailsListResponse
}

func (o *ListPaymentsOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments][%d] listPaymentsOK  %+v", 200, o)
}

func (o *ListPaymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentDetailsListResponse = new(models.PaymentDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.PaymentDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
