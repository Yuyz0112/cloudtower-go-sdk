// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// DeleteDatacenterReader is a Reader for the DeleteDatacenter structure.
type DeleteDatacenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDatacenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDatacenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDatacenterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDatacenterOK creates a DeleteDatacenterOK with default headers values
func NewDeleteDatacenterOK() *DeleteDatacenterOK {
	return &DeleteDatacenterOK{}
}

/* DeleteDatacenterOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteDatacenterOK struct {
	Payload []*models.WithTaskDeleteDatacenter
}

func (o *DeleteDatacenterOK) Error() string {
	return fmt.Sprintf("[POST /delete-datacenter][%d] deleteDatacenterOK  %+v", 200, o.Payload)
}
func (o *DeleteDatacenterOK) GetPayload() []*models.WithTaskDeleteDatacenter {
	return o.Payload
}

func (o *DeleteDatacenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDatacenterBadRequest creates a DeleteDatacenterBadRequest with default headers values
func NewDeleteDatacenterBadRequest() *DeleteDatacenterBadRequest {
	return &DeleteDatacenterBadRequest{}
}

/* DeleteDatacenterBadRequest describes a response with status code 400, with default header values.

DeleteDatacenterBadRequest delete datacenter bad request
*/
type DeleteDatacenterBadRequest struct {
	Payload string
}

func (o *DeleteDatacenterBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-datacenter][%d] deleteDatacenterBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDatacenterBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteDatacenterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
