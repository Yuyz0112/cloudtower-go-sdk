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

// UpdateVMNicReader is a Reader for the UpdateVMNic structure.
type UpdateVMNicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMNicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMNicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVMNicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMNicOK creates a UpdateVMNicOK with default headers values
func NewUpdateVMNicOK() *UpdateVMNicOK {
	return &UpdateVMNicOK{}
}

/* UpdateVMNicOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateVMNicOK struct {
	Payload []*models.WithTaskVM
}

func (o *UpdateVMNicOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic][%d] updateVmNicOK  %+v", 200, o.Payload)
}
func (o *UpdateVMNicOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *UpdateVMNicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMNicBadRequest creates a UpdateVMNicBadRequest with default headers values
func NewUpdateVMNicBadRequest() *UpdateVMNicBadRequest {
	return &UpdateVMNicBadRequest{}
}

/* UpdateVMNicBadRequest describes a response with status code 400, with default header values.

UpdateVMNicBadRequest update Vm nic bad request
*/
type UpdateVMNicBadRequest struct {
	Payload string
}

func (o *UpdateVMNicBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-nic][%d] updateVmNicBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMNicBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateVMNicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
