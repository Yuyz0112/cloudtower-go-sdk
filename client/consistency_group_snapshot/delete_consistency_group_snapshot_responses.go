// Code generated by go-swagger; DO NOT EDIT.

package consistency_group_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// DeleteConsistencyGroupSnapshotReader is a Reader for the DeleteConsistencyGroupSnapshot structure.
type DeleteConsistencyGroupSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConsistencyGroupSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConsistencyGroupSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConsistencyGroupSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConsistencyGroupSnapshotOK creates a DeleteConsistencyGroupSnapshotOK with default headers values
func NewDeleteConsistencyGroupSnapshotOK() *DeleteConsistencyGroupSnapshotOK {
	return &DeleteConsistencyGroupSnapshotOK{}
}

/* DeleteConsistencyGroupSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteConsistencyGroupSnapshotOK struct {
	Payload []*models.WithTaskDeleteConsistencyGroupSnapshot
}

func (o *DeleteConsistencyGroupSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /delete-consistency-snapshot-group][%d] deleteConsistencyGroupSnapshotOK  %+v", 200, o.Payload)
}
func (o *DeleteConsistencyGroupSnapshotOK) GetPayload() []*models.WithTaskDeleteConsistencyGroupSnapshot {
	return o.Payload
}

func (o *DeleteConsistencyGroupSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConsistencyGroupSnapshotBadRequest creates a DeleteConsistencyGroupSnapshotBadRequest with default headers values
func NewDeleteConsistencyGroupSnapshotBadRequest() *DeleteConsistencyGroupSnapshotBadRequest {
	return &DeleteConsistencyGroupSnapshotBadRequest{}
}

/* DeleteConsistencyGroupSnapshotBadRequest describes a response with status code 400, with default header values.

DeleteConsistencyGroupSnapshotBadRequest delete consistency group snapshot bad request
*/
type DeleteConsistencyGroupSnapshotBadRequest struct {
	Payload string
}

func (o *DeleteConsistencyGroupSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-consistency-snapshot-group][%d] deleteConsistencyGroupSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteConsistencyGroupSnapshotBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteConsistencyGroupSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
