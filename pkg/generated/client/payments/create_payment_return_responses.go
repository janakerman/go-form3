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

// CreatePaymentReturnReader is a Reader for the CreatePaymentReturn structure.
type CreatePaymentReturnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentReturnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentReturnCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentReturnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentReturnCreated creates a CreatePaymentReturnCreated with default headers values
func NewCreatePaymentReturnCreated() *CreatePaymentReturnCreated {
	return &CreatePaymentReturnCreated{}
}

/*CreatePaymentReturnCreated handles this case with default header values.

Return creation response
*/
type CreatePaymentReturnCreated struct {

	//Payload
	*models.ReturnCreationResponse
}

func (o *CreatePaymentReturnCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns][%d] createPaymentReturnCreated  %+v", 201, o)
}

func (o *CreatePaymentReturnCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnCreationResponse = new(models.ReturnCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.ReturnCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentReturnBadRequest creates a CreatePaymentReturnBadRequest with default headers values
func NewCreatePaymentReturnBadRequest() *CreatePaymentReturnBadRequest {
	return &CreatePaymentReturnBadRequest{}
}

/*CreatePaymentReturnBadRequest handles this case with default header values.

Return creation error
*/
type CreatePaymentReturnBadRequest struct {

	//Payload
	*models.APIError
}

func (o *CreatePaymentReturnBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/payments/{id}/returns][%d] createPaymentReturnBadRequest  %+v", 400, o)
}

func (o *CreatePaymentReturnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
