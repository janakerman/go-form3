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

// CreatePaymentReturnReversalReader is a Reader for the CreatePaymentReturnReversal structure.
type CreatePaymentReturnReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentReturnReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentReturnReversalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentReturnReversalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentReturnReversalCreated creates a CreatePaymentReturnReversalCreated with default headers values
func NewCreatePaymentReturnReversalCreated() *CreatePaymentReturnReversalCreated {
	return &CreatePaymentReturnReversalCreated{}
}

/*CreatePaymentReturnReversalCreated handles this case with default header values.

Reversal creation response
*/
type CreatePaymentReturnReversalCreated struct {

	//Payload
	*models.ReturnReversalCreationResponse
}

func (o *CreatePaymentReturnReversalCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalCreated  %+v", 201, o)
}

func (o *CreatePaymentReturnReversalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnReversalCreationResponse = new(models.ReturnReversalCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.ReturnReversalCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnReversalBadRequest creates a CreatePaymentReturnReversalBadRequest with default headers values
func NewCreatePaymentReturnReversalBadRequest() *CreatePaymentReturnReversalBadRequest {
	return &CreatePaymentReturnReversalBadRequest{}
}

/*CreatePaymentReturnReversalBadRequest handles this case with default header values.

Reversal creation error
*/
type CreatePaymentReturnReversalBadRequest struct {

	//Payload
	*models.APIError
}

func (o *CreatePaymentReturnReversalBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns/{returnId}/reversals][%d] createPaymentReturnReversalBadRequest  %+v", 400, o)
}

func (o *CreatePaymentReturnReversalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
