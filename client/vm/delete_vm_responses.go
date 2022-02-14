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

// DeleteVMReader is a Reader for the DeleteVM structure.
type DeleteVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMOK creates a DeleteVMOK with default headers values
func NewDeleteVMOK() *DeleteVMOK {
	return &DeleteVMOK{}
}

/* DeleteVMOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVMOK struct {
	Payload []*models.WithTaskDeleteVM
}

func (o *DeleteVMOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm][%d] deleteVmOK  %+v", 200, o.Payload)
}
func (o *DeleteVMOK) GetPayload() []*models.WithTaskDeleteVM {
	return o.Payload
}

func (o *DeleteVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMBadRequest creates a DeleteVMBadRequest with default headers values
func NewDeleteVMBadRequest() *DeleteVMBadRequest {
	return &DeleteVMBadRequest{}
}

/* DeleteVMBadRequest describes a response with status code 400, with default header values.

DeleteVMBadRequest delete Vm bad request
*/
type DeleteVMBadRequest struct {
	Payload string
}

func (o *DeleteVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm][%d] deleteVmBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
