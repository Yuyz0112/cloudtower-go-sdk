// Code generated by go-swagger; DO NOT EDIT.

package vm_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetVMSnapshotsConnectionReader is a Reader for the GetVMSnapshotsConnection structure.
type GetVMSnapshotsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMSnapshotsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMSnapshotsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMSnapshotsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMSnapshotsConnectionOK creates a GetVMSnapshotsConnectionOK with default headers values
func NewGetVMSnapshotsConnectionOK() *GetVMSnapshotsConnectionOK {
	return &GetVMSnapshotsConnectionOK{}
}

/* GetVMSnapshotsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMSnapshotsConnectionOK struct {
	Payload *models.VMSnapshotConnection
}

func (o *GetVMSnapshotsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-snapshots-connection][%d] getVmSnapshotsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVMSnapshotsConnectionOK) GetPayload() *models.VMSnapshotConnection {
	return o.Payload
}

func (o *GetVMSnapshotsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMSnapshotConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMSnapshotsConnectionBadRequest creates a GetVMSnapshotsConnectionBadRequest with default headers values
func NewGetVMSnapshotsConnectionBadRequest() *GetVMSnapshotsConnectionBadRequest {
	return &GetVMSnapshotsConnectionBadRequest{}
}

/* GetVMSnapshotsConnectionBadRequest describes a response with status code 400, with default header values.

GetVMSnapshotsConnectionBadRequest get Vm snapshots connection bad request
*/
type GetVMSnapshotsConnectionBadRequest struct {
	Payload string
}

func (o *GetVMSnapshotsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-snapshots-connection][%d] getVmSnapshotsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMSnapshotsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVMSnapshotsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}