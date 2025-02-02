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

// CreateDirectDebitReturnReader is a Reader for the CreateDirectDebitReturn structure.
type CreateDirectDebitReturnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDirectDebitReturnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateDirectDebitReturnCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDirectDebitReturnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDirectDebitReturnCreated creates a CreateDirectDebitReturnCreated with default headers values
func NewCreateDirectDebitReturnCreated() *CreateDirectDebitReturnCreated {
	return &CreateDirectDebitReturnCreated{}
}

/*CreateDirectDebitReturnCreated handles this case with default header values.

Return creation response
*/
type CreateDirectDebitReturnCreated struct {

	//Payload
	*models.DirectDebitReturnCreationResponse
}

func (o *CreateDirectDebitReturnCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/returns][%d] createDirectDebitReturnCreated  %+v", 201, o)
}

func (o *CreateDirectDebitReturnCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReturnCreationResponse = new(models.DirectDebitReturnCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.DirectDebitReturnCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDirectDebitReturnBadRequest creates a CreateDirectDebitReturnBadRequest with default headers values
func NewCreateDirectDebitReturnBadRequest() *CreateDirectDebitReturnBadRequest {
	return &CreateDirectDebitReturnBadRequest{}
}

/*CreateDirectDebitReturnBadRequest handles this case with default header values.

Return creation error
*/
type CreateDirectDebitReturnBadRequest struct {

	//Payload
	*models.APIError
}

func (o *CreateDirectDebitReturnBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/returns][%d] createDirectDebitReturnBadRequest  %+v", 400, o)
}

func (o *CreateDirectDebitReturnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
