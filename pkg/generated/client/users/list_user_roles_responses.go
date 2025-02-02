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

// ListUserRolesReader is a Reader for the ListUserRoles structure.
type ListUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListUserRolesOK creates a ListUserRolesOK with default headers values
func NewListUserRolesOK() *ListUserRolesOK {
	return &ListUserRolesOK{}
}

/*ListUserRolesOK handles this case with default header values.

List of roles for user
*/
type ListUserRolesOK struct {

	//Payload
	*models.UserRoleListResponse
}

func (o *ListUserRolesOK) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}/roles][%d] listUserRolesOK  %+v", 200, o)
}

func (o *ListUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.UserRoleListResponse = new(models.UserRoleListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.UserRoleListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
