// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// DeleteClusterReader is a Reader for the DeleteCluster structure.
type DeleteClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteClusterOK creates a DeleteClusterOK with default headers values
func NewDeleteClusterOK() *DeleteClusterOK {
	return &DeleteClusterOK{}
}

/* DeleteClusterOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteClusterOK struct {
	Payload []*models.WithTaskDeleteCluster
}

func (o *DeleteClusterOK) Error() string {
	return fmt.Sprintf("[POST /delete-cluster][%d] deleteClusterOK  %+v", 200, o.Payload)
}
func (o *DeleteClusterOK) GetPayload() []*models.WithTaskDeleteCluster {
	return o.Payload
}

func (o *DeleteClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterBadRequest creates a DeleteClusterBadRequest with default headers values
func NewDeleteClusterBadRequest() *DeleteClusterBadRequest {
	return &DeleteClusterBadRequest{}
}

/* DeleteClusterBadRequest describes a response with status code 400, with default header values.

DeleteClusterBadRequest delete cluster bad request
*/
type DeleteClusterBadRequest struct {
	Payload string
}

func (o *DeleteClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-cluster][%d] deleteClusterBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteClusterBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
