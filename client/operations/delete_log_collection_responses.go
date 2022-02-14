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

// DeleteLogCollectionReader is a Reader for the DeleteLogCollection structure.
type DeleteLogCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLogCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLogCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLogCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLogCollectionOK creates a DeleteLogCollectionOK with default headers values
func NewDeleteLogCollectionOK() *DeleteLogCollectionOK {
	return &DeleteLogCollectionOK{}
}

/* DeleteLogCollectionOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteLogCollectionOK struct {
	Payload []*models.WithTaskDeleteLogCollection
}

func (o *DeleteLogCollectionOK) Error() string {
	return fmt.Sprintf("[POST /delete-log-collection][%d] deleteLogCollectionOK  %+v", 200, o.Payload)
}
func (o *DeleteLogCollectionOK) GetPayload() []*models.WithTaskDeleteLogCollection {
	return o.Payload
}

func (o *DeleteLogCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLogCollectionBadRequest creates a DeleteLogCollectionBadRequest with default headers values
func NewDeleteLogCollectionBadRequest() *DeleteLogCollectionBadRequest {
	return &DeleteLogCollectionBadRequest{}
}

/* DeleteLogCollectionBadRequest describes a response with status code 400, with default header values.

DeleteLogCollectionBadRequest delete log collection bad request
*/
type DeleteLogCollectionBadRequest struct {
	Payload string
}

func (o *DeleteLogCollectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-log-collection][%d] deleteLogCollectionBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteLogCollectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteLogCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
