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

// UpdateEntityFilterReader is a Reader for the UpdateEntityFilter structure.
type UpdateEntityFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEntityFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEntityFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEntityFilterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEntityFilterOK creates a UpdateEntityFilterOK with default headers values
func NewUpdateEntityFilterOK() *UpdateEntityFilterOK {
	return &UpdateEntityFilterOK{}
}

/* UpdateEntityFilterOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateEntityFilterOK struct {
	Payload []*models.WithTaskEntityFilter
}

func (o *UpdateEntityFilterOK) Error() string {
	return fmt.Sprintf("[POST /update-entity-filter][%d] updateEntityFilterOK  %+v", 200, o.Payload)
}
func (o *UpdateEntityFilterOK) GetPayload() []*models.WithTaskEntityFilter {
	return o.Payload
}

func (o *UpdateEntityFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEntityFilterBadRequest creates a UpdateEntityFilterBadRequest with default headers values
func NewUpdateEntityFilterBadRequest() *UpdateEntityFilterBadRequest {
	return &UpdateEntityFilterBadRequest{}
}

/* UpdateEntityFilterBadRequest describes a response with status code 400, with default header values.

UpdateEntityFilterBadRequest update entity filter bad request
*/
type UpdateEntityFilterBadRequest struct {
	Payload string
}

func (o *UpdateEntityFilterBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-entity-filter][%d] updateEntityFilterBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateEntityFilterBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateEntityFilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
