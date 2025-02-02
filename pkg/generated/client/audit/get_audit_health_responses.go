// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetAuditHealthReader is a Reader for the GetAuditHealth structure.
type GetAuditHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuditHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuditHealthOK creates a GetAuditHealthOK with default headers values
func NewGetAuditHealthOK() *GetAuditHealthOK {
	return &GetAuditHealthOK{}
}

/*GetAuditHealthOK handles this case with default header values.

Audit service health
*/
type GetAuditHealthOK struct {

	//Payload
	*models.Health
}

func (o *GetAuditHealthOK) Error() string {
	return fmt.Sprintf("[GET /audit/entries/health][%d] getAuditHealthOK  %+v", 200, o)
}

func (o *GetAuditHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Health = new(models.Health)

	// response payload
	if err := consumer.Consume(response.Body(), o.Health); err != nil && err != io.EOF {
		return err
	}

	return nil
}
