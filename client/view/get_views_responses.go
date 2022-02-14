// Code generated by go-swagger; DO NOT EDIT.

package view

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetViewsReader is a Reader for the GetViews structure.
type GetViewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetViewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetViewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetViewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetViewsOK creates a GetViewsOK with default headers values
func NewGetViewsOK() *GetViewsOK {
	return &GetViewsOK{}
}

/* GetViewsOK describes a response with status code 200, with default header values.

Ok
*/
type GetViewsOK struct {
	Payload []*models.View
}

func (o *GetViewsOK) Error() string {
	return fmt.Sprintf("[POST /get-views][%d] getViewsOK  %+v", 200, o.Payload)
}
func (o *GetViewsOK) GetPayload() []*models.View {
	return o.Payload
}

func (o *GetViewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetViewsBadRequest creates a GetViewsBadRequest with default headers values
func NewGetViewsBadRequest() *GetViewsBadRequest {
	return &GetViewsBadRequest{}
}

/* GetViewsBadRequest describes a response with status code 400, with default header values.

GetViewsBadRequest get views bad request
*/
type GetViewsBadRequest struct {
	Payload string
}

func (o *GetViewsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-views][%d] getViewsBadRequest  %+v", 400, o.Payload)
}
func (o *GetViewsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetViewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
