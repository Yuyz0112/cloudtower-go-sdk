// Code generated by go-swagger; DO NOT EDIT.

package vm_folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// CreateVMFolderReader is a Reader for the CreateVMFolder structure.
type CreateVMFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVMFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVMFolderOK creates a CreateVMFolderOK with default headers values
func NewCreateVMFolderOK() *CreateVMFolderOK {
	return &CreateVMFolderOK{}
}

/* CreateVMFolderOK describes a response with status code 200, with default header values.

Ok
*/
type CreateVMFolderOK struct {
	Payload []*models.WithTaskVMFolder
}

func (o *CreateVMFolderOK) Error() string {
	return fmt.Sprintf("[POST /create-vm-folder][%d] createVmFolderOK  %+v", 200, o.Payload)
}
func (o *CreateVMFolderOK) GetPayload() []*models.WithTaskVMFolder {
	return o.Payload
}

func (o *CreateVMFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVMFolderBadRequest creates a CreateVMFolderBadRequest with default headers values
func NewCreateVMFolderBadRequest() *CreateVMFolderBadRequest {
	return &CreateVMFolderBadRequest{}
}

/* CreateVMFolderBadRequest describes a response with status code 400, with default header values.

CreateVMFolderBadRequest create Vm folder bad request
*/
type CreateVMFolderBadRequest struct {
	Payload string
}

func (o *CreateVMFolderBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vm-folder][%d] createVmFolderBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVMFolderBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateVMFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}