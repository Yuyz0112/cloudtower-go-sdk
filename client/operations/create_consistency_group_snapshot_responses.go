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

// CreateConsistencyGroupSnapshotReader is a Reader for the CreateConsistencyGroupSnapshot structure.
type CreateConsistencyGroupSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConsistencyGroupSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateConsistencyGroupSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateConsistencyGroupSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConsistencyGroupSnapshotOK creates a CreateConsistencyGroupSnapshotOK with default headers values
func NewCreateConsistencyGroupSnapshotOK() *CreateConsistencyGroupSnapshotOK {
	return &CreateConsistencyGroupSnapshotOK{}
}

/* CreateConsistencyGroupSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type CreateConsistencyGroupSnapshotOK struct {
	Payload []*models.WithTaskConsistencyGroupSnapshot
}

func (o *CreateConsistencyGroupSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /create-consistency-snapshot-group][%d] createConsistencyGroupSnapshotOK  %+v", 200, o.Payload)
}
func (o *CreateConsistencyGroupSnapshotOK) GetPayload() []*models.WithTaskConsistencyGroupSnapshot {
	return o.Payload
}

func (o *CreateConsistencyGroupSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsistencyGroupSnapshotBadRequest creates a CreateConsistencyGroupSnapshotBadRequest with default headers values
func NewCreateConsistencyGroupSnapshotBadRequest() *CreateConsistencyGroupSnapshotBadRequest {
	return &CreateConsistencyGroupSnapshotBadRequest{}
}

/* CreateConsistencyGroupSnapshotBadRequest describes a response with status code 400, with default header values.

CreateConsistencyGroupSnapshotBadRequest create consistency group snapshot bad request
*/
type CreateConsistencyGroupSnapshotBadRequest struct {
	Payload string
}

func (o *CreateConsistencyGroupSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-consistency-snapshot-group][%d] createConsistencyGroupSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CreateConsistencyGroupSnapshotBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateConsistencyGroupSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
