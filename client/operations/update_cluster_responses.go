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

// UpdateClusterReader is a Reader for the UpdateCluster structure.
type UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateClusterOK creates a UpdateClusterOK with default headers values
func NewUpdateClusterOK() *UpdateClusterOK {
	return &UpdateClusterOK{}
}

/* UpdateClusterOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateClusterOK struct {
	Payload []*models.WithTaskCluster
}

func (o *UpdateClusterOK) Error() string {
	return fmt.Sprintf("[POST /update-cluster][%d] updateClusterOK  %+v", 200, o.Payload)
}
func (o *UpdateClusterOK) GetPayload() []*models.WithTaskCluster {
	return o.Payload
}

func (o *UpdateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterBadRequest creates a UpdateClusterBadRequest with default headers values
func NewUpdateClusterBadRequest() *UpdateClusterBadRequest {
	return &UpdateClusterBadRequest{}
}

/* UpdateClusterBadRequest describes a response with status code 400, with default header values.

UpdateClusterBadRequest update cluster bad request
*/
type UpdateClusterBadRequest struct {
	Payload string
}

func (o *UpdateClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-cluster][%d] updateClusterBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateClusterBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
