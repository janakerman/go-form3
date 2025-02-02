// Code generated by go-swagger; DO NOT EDIT.

package balances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetBalancedReader is a Reader for the GetBalanced structure.
type GetBalancedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBalancedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBalancedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBalancedOK creates a GetBalancedOK with default headers values
func NewGetBalancedOK() *GetBalancedOK {
	return &GetBalancedOK{}
}

/*GetBalancedOK handles this case with default header values.

Associations details
*/
type GetBalancedOK struct {

	//Payload
	*models.BalanceDetailsListResponse
}

func (o *GetBalancedOK) Error() string {
	return fmt.Sprintf("[GET /organisation/balances][%d] getBalancedOK  %+v", 200, o)
}

func (o *GetBalancedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.BalanceDetailsListResponse = new(models.BalanceDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.BalanceDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
