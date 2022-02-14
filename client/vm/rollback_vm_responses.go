// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// RollbackVMReader is a Reader for the RollbackVM structure.
type RollbackVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RollbackVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRollbackVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRollbackVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRollbackVMOK creates a RollbackVMOK with default headers values
func NewRollbackVMOK() *RollbackVMOK {
	return &RollbackVMOK{}
}

/* RollbackVMOK describes a response with status code 200, with default header values.

Ok
*/
type RollbackVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *RollbackVMOK) Error() string {
	return fmt.Sprintf("[POST /rollback-vm][%d] rollbackVmOK  %+v", 200, o.Payload)
}
func (o *RollbackVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *RollbackVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRollbackVMBadRequest creates a RollbackVMBadRequest with default headers values
func NewRollbackVMBadRequest() *RollbackVMBadRequest {
	return &RollbackVMBadRequest{}
}

/* RollbackVMBadRequest describes a response with status code 400, with default header values.

RollbackVMBadRequest rollback Vm bad request
*/
type RollbackVMBadRequest struct {
	Payload string
}

func (o *RollbackVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /rollback-vm][%d] rollbackVmBadRequest  %+v", 400, o.Payload)
}
func (o *RollbackVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *RollbackVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
