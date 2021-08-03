// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cloudtower-go-sdk/models"
)

// StartVMReader is a Reader for the StartVM structure.
type StartVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartVMOK creates a StartVMOK with default headers values
func NewStartVMOK() *StartVMOK {
	return &StartVMOK{}
}

/* StartVMOK describes a response with status code 200, with default header values.

Ok
*/
type StartVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *StartVMOK) Error() string {
	return fmt.Sprintf("[POST /start-vm][%d] startVmOK  %+v", 200, o.Payload)
}
func (o *StartVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *StartVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
