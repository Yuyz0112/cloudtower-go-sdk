// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetApplicationsConnectionReader is a Reader for the GetApplicationsConnection structure.
type GetApplicationsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApplicationsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApplicationsConnectionOK creates a GetApplicationsConnectionOK with default headers values
func NewGetApplicationsConnectionOK() *GetApplicationsConnectionOK {
	return &GetApplicationsConnectionOK{}
}

/* GetApplicationsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetApplicationsConnectionOK struct {
	Payload *models.ApplicationConnection
}

func (o *GetApplicationsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-applications-connection][%d] getApplicationsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetApplicationsConnectionOK) GetPayload() *models.ApplicationConnection {
	return o.Payload
}

func (o *GetApplicationsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationsConnectionBadRequest creates a GetApplicationsConnectionBadRequest with default headers values
func NewGetApplicationsConnectionBadRequest() *GetApplicationsConnectionBadRequest {
	return &GetApplicationsConnectionBadRequest{}
}

/* GetApplicationsConnectionBadRequest describes a response with status code 400, with default header values.

GetApplicationsConnectionBadRequest get applications connection bad request
*/
type GetApplicationsConnectionBadRequest struct {
	Payload string
}

func (o *GetApplicationsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-applications-connection][%d] getApplicationsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetApplicationsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetApplicationsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}