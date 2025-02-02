// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// CreateUserCredentialsReader is a Reader for the CreateUserCredentials structure.
type CreateUserCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateUserCredentialsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserCredentialsCreated creates a CreateUserCredentialsCreated with default headers values
func NewCreateUserCredentialsCreated() *CreateUserCredentialsCreated {
	return &CreateUserCredentialsCreated{}
}

/*CreateUserCredentialsCreated handles this case with default header values.

Credential creation response
*/
type CreateUserCredentialsCreated struct {

	//Payload
	*models.CredentialCreationResponse
}

func (o *CreateUserCredentialsCreated) Error() string {
	return fmt.Sprintf("[POST /security/users/{user_id}/credentials][%d] createUserCredentialsCreated  %+v", 201, o)
}

func (o *CreateUserCredentialsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.CredentialCreationResponse = new(models.CredentialCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.CredentialCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
