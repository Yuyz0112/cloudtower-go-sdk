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

// GetVmsConnectionReader is a Reader for the GetVmsConnection structure.
type GetVmsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVmsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVmsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVmsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVmsConnectionOK creates a GetVmsConnectionOK with default headers values
func NewGetVmsConnectionOK() *GetVmsConnectionOK {
	return &GetVmsConnectionOK{}
}

/* GetVmsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVmsConnectionOK struct {
	Payload *models.VMConnection
}

func (o *GetVmsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vms-connection][%d] getVmsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVmsConnectionOK) GetPayload() *models.VMConnection {
	return o.Payload
}

func (o *GetVmsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmsConnectionBadRequest creates a GetVmsConnectionBadRequest with default headers values
func NewGetVmsConnectionBadRequest() *GetVmsConnectionBadRequest {
	return &GetVmsConnectionBadRequest{}
}

/* GetVmsConnectionBadRequest describes a response with status code 400, with default header values.

GetVmsConnectionBadRequest get vms connection bad request
*/
type GetVmsConnectionBadRequest struct {
	Payload string
}

func (o *GetVmsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vms-connection][%d] getVmsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVmsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVmsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
