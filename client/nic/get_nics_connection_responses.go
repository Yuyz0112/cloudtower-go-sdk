// Code generated by go-swagger; DO NOT EDIT.

package nic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetNicsConnectionReader is a Reader for the GetNicsConnection structure.
type GetNicsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNicsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNicsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNicsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNicsConnectionOK creates a GetNicsConnectionOK with default headers values
func NewGetNicsConnectionOK() *GetNicsConnectionOK {
	return &GetNicsConnectionOK{}
}

/* GetNicsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetNicsConnectionOK struct {
	Payload *models.NicConnection
}

func (o *GetNicsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-nics-connection][%d] getNicsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNicsConnectionOK) GetPayload() *models.NicConnection {
	return o.Payload
}

func (o *GetNicsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NicConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNicsConnectionBadRequest creates a GetNicsConnectionBadRequest with default headers values
func NewGetNicsConnectionBadRequest() *GetNicsConnectionBadRequest {
	return &GetNicsConnectionBadRequest{}
}

/* GetNicsConnectionBadRequest describes a response with status code 400, with default header values.

GetNicsConnectionBadRequest get nics connection bad request
*/
type GetNicsConnectionBadRequest struct {
	Payload string
}

func (o *GetNicsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nics-connection][%d] getNicsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNicsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNicsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
