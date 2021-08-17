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

// DeleteVdsReader is a Reader for the DeleteVds structure.
type DeleteVdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVdsOK creates a DeleteVdsOK with default headers values
func NewDeleteVdsOK() *DeleteVdsOK {
	return &DeleteVdsOK{}
}

/* DeleteVdsOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVdsOK struct {
	Payload []*models.WithTaskDeleteVds
}

func (o *DeleteVdsOK) Error() string {
	return fmt.Sprintf("[POST /delete-vds][%d] deleteVdsOK  %+v", 200, o.Payload)
}
func (o *DeleteVdsOK) GetPayload() []*models.WithTaskDeleteVds {
	return o.Payload
}

func (o *DeleteVdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVdsBadRequest creates a DeleteVdsBadRequest with default headers values
func NewDeleteVdsBadRequest() *DeleteVdsBadRequest {
	return &DeleteVdsBadRequest{}
}

/* DeleteVdsBadRequest describes a response with status code 400, with default header values.

DeleteVdsBadRequest delete vds bad request
*/
type DeleteVdsBadRequest struct {
	Payload string
}

func (o *DeleteVdsBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vds][%d] deleteVdsBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVdsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteVdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
