// Code generated by go-swagger; DO NOT EDIT.

package nfs_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// CreateNfsExportReader is a Reader for the CreateNfsExport structure.
type CreateNfsExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNfsExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNfsExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNfsExportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNfsExportOK creates a CreateNfsExportOK with default headers values
func NewCreateNfsExportOK() *CreateNfsExportOK {
	return &CreateNfsExportOK{}
}

/* CreateNfsExportOK describes a response with status code 200, with default header values.

Ok
*/
type CreateNfsExportOK struct {
	Payload []*models.WithTaskNfsExport
}

func (o *CreateNfsExportOK) Error() string {
	return fmt.Sprintf("[POST /create-nfs-export][%d] createNfsExportOK  %+v", 200, o.Payload)
}
func (o *CreateNfsExportOK) GetPayload() []*models.WithTaskNfsExport {
	return o.Payload
}

func (o *CreateNfsExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNfsExportBadRequest creates a CreateNfsExportBadRequest with default headers values
func NewCreateNfsExportBadRequest() *CreateNfsExportBadRequest {
	return &CreateNfsExportBadRequest{}
}

/* CreateNfsExportBadRequest describes a response with status code 400, with default header values.

CreateNfsExportBadRequest create nfs export bad request
*/
type CreateNfsExportBadRequest struct {
	Payload string
}

func (o *CreateNfsExportBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-nfs-export][%d] createNfsExportBadRequest  %+v", 400, o.Payload)
}
func (o *CreateNfsExportBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateNfsExportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
