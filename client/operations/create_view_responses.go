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

// CreateViewReader is a Reader for the CreateView structure.
type CreateViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateViewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateViewOK creates a CreateViewOK with default headers values
func NewCreateViewOK() *CreateViewOK {
	return &CreateViewOK{}
}

/* CreateViewOK describes a response with status code 200, with default header values.

Ok
*/
type CreateViewOK struct {
	Payload []*models.WithTaskView
}

func (o *CreateViewOK) Error() string {
	return fmt.Sprintf("[POST /create-view][%d] createViewOK  %+v", 200, o.Payload)
}
func (o *CreateViewOK) GetPayload() []*models.WithTaskView {
	return o.Payload
}

func (o *CreateViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateViewBadRequest creates a CreateViewBadRequest with default headers values
func NewCreateViewBadRequest() *CreateViewBadRequest {
	return &CreateViewBadRequest{}
}

/* CreateViewBadRequest describes a response with status code 400, with default header values.

CreateViewBadRequest create view bad request
*/
type CreateViewBadRequest struct {
	Payload string
}

func (o *CreateViewBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-view][%d] createViewBadRequest  %+v", 400, o.Payload)
}
func (o *CreateViewBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateViewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
