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

// GetVMFoldersConnectionReader is a Reader for the GetVMFoldersConnection structure.
type GetVMFoldersConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMFoldersConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMFoldersConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMFoldersConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMFoldersConnectionOK creates a GetVMFoldersConnectionOK with default headers values
func NewGetVMFoldersConnectionOK() *GetVMFoldersConnectionOK {
	return &GetVMFoldersConnectionOK{}
}

/* GetVMFoldersConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMFoldersConnectionOK struct {
	Payload *models.VMFolderConnection
}

func (o *GetVMFoldersConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-folders-connection][%d] getVmFoldersConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVMFoldersConnectionOK) GetPayload() *models.VMFolderConnection {
	return o.Payload
}

func (o *GetVMFoldersConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMFolderConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMFoldersConnectionBadRequest creates a GetVMFoldersConnectionBadRequest with default headers values
func NewGetVMFoldersConnectionBadRequest() *GetVMFoldersConnectionBadRequest {
	return &GetVMFoldersConnectionBadRequest{}
}

/* GetVMFoldersConnectionBadRequest describes a response with status code 400, with default header values.

GetVMFoldersConnectionBadRequest get Vm folders connection bad request
*/
type GetVMFoldersConnectionBadRequest struct {
	Payload string
}

func (o *GetVMFoldersConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-folders-connection][%d] getVmFoldersConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMFoldersConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVMFoldersConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
