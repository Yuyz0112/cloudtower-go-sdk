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

// GetNfsInodesConnectionReader is a Reader for the GetNfsInodesConnection structure.
type GetNfsInodesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsInodesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNfsInodesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNfsInodesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNfsInodesConnectionOK creates a GetNfsInodesConnectionOK with default headers values
func NewGetNfsInodesConnectionOK() *GetNfsInodesConnectionOK {
	return &GetNfsInodesConnectionOK{}
}

/* GetNfsInodesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetNfsInodesConnectionOK struct {
	Payload *models.NfsInodeConnection
}

func (o *GetNfsInodesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-nfs-inodes-connection][%d] getNfsInodesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNfsInodesConnectionOK) GetPayload() *models.NfsInodeConnection {
	return o.Payload
}

func (o *GetNfsInodesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsInodeConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsInodesConnectionBadRequest creates a GetNfsInodesConnectionBadRequest with default headers values
func NewGetNfsInodesConnectionBadRequest() *GetNfsInodesConnectionBadRequest {
	return &GetNfsInodesConnectionBadRequest{}
}

/* GetNfsInodesConnectionBadRequest describes a response with status code 400, with default header values.

GetNfsInodesConnectionBadRequest get nfs inodes connection bad request
*/
type GetNfsInodesConnectionBadRequest struct {
	Payload string
}

func (o *GetNfsInodesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nfs-inodes-connection][%d] getNfsInodesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNfsInodesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNfsInodesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
