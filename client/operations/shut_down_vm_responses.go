// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// ShutDownVMReader is a Reader for the ShutDownVM structure.
type ShutDownVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutDownVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShutDownVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShutDownVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShutDownVMOK creates a ShutDownVMOK with default headers values
func NewShutDownVMOK() *ShutDownVMOK {
	return &ShutDownVMOK{}
}

/* ShutDownVMOK describes a response with status code 200, with default header values.

Ok
*/
type ShutDownVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *ShutDownVMOK) Error() string {
	return fmt.Sprintf("[POST /shut-down-vm][%d] shutDownVmOK  %+v", 200, o.Payload)
}
func (o *ShutDownVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *ShutDownVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutDownVMBadRequest creates a ShutDownVMBadRequest with default headers values
func NewShutDownVMBadRequest() *ShutDownVMBadRequest {
	return &ShutDownVMBadRequest{}
}

/* ShutDownVMBadRequest describes a response with status code 400, with default header values.

ShutDownVMBadRequest shut down Vm bad request
*/
type ShutDownVMBadRequest struct {
	Payload string
}

func (o *ShutDownVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /shut-down-vm][%d] shutDownVmBadRequest  %+v", 400, o.Payload)
}
func (o *ShutDownVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ShutDownVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
