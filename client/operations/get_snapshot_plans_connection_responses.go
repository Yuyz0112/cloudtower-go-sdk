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

// GetSnapshotPlansConnectionReader is a Reader for the GetSnapshotPlansConnection structure.
type GetSnapshotPlansConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnapshotPlansConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnapshotPlansConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSnapshotPlansConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnapshotPlansConnectionOK creates a GetSnapshotPlansConnectionOK with default headers values
func NewGetSnapshotPlansConnectionOK() *GetSnapshotPlansConnectionOK {
	return &GetSnapshotPlansConnectionOK{}
}

/* GetSnapshotPlansConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetSnapshotPlansConnectionOK struct {
	Payload *models.SnapshotPlanConnection
}

func (o *GetSnapshotPlansConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-plans-connection][%d] getSnapshotPlansConnectionOK  %+v", 200, o.Payload)
}
func (o *GetSnapshotPlansConnectionOK) GetPayload() *models.SnapshotPlanConnection {
	return o.Payload
}

func (o *GetSnapshotPlansConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotPlanConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnapshotPlansConnectionBadRequest creates a GetSnapshotPlansConnectionBadRequest with default headers values
func NewGetSnapshotPlansConnectionBadRequest() *GetSnapshotPlansConnectionBadRequest {
	return &GetSnapshotPlansConnectionBadRequest{}
}

/* GetSnapshotPlansConnectionBadRequest describes a response with status code 400, with default header values.

GetSnapshotPlansConnectionBadRequest get snapshot plans connection bad request
*/
type GetSnapshotPlansConnectionBadRequest struct {
	Payload string
}

func (o *GetSnapshotPlansConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-plans-connection][%d] getSnapshotPlansConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetSnapshotPlansConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetSnapshotPlansConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
