// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteRoleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRoleNoContent creates a DeleteRoleNoContent with default headers values
func NewDeleteRoleNoContent() *DeleteRoleNoContent {
	return &DeleteRoleNoContent{}
}

/*DeleteRoleNoContent handles this case with default header values.

Role deleted
*/
type DeleteRoleNoContent struct {
}

func (o *DeleteRoleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{role_id}][%d] deleteRoleNoContent ", 204)
}

func (o *DeleteRoleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
