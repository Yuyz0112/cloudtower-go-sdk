// Code generated by go-swagger; DO NOT EDIT.

package snapshot_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetSnapshotGroupsConnectionReader is a Reader for the GetSnapshotGroupsConnection structure.
type GetSnapshotGroupsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnapshotGroupsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnapshotGroupsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSnapshotGroupsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnapshotGroupsConnectionOK creates a GetSnapshotGroupsConnectionOK with default headers values
func NewGetSnapshotGroupsConnectionOK() *GetSnapshotGroupsConnectionOK {
	return &GetSnapshotGroupsConnectionOK{}
}

/* GetSnapshotGroupsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetSnapshotGroupsConnectionOK struct {
	Payload *models.SnapshotGroupConnection
}

func (o *GetSnapshotGroupsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-groups-connection][%d] getSnapshotGroupsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetSnapshotGroupsConnectionOK) GetPayload() *models.SnapshotGroupConnection {
	return o.Payload
}

func (o *GetSnapshotGroupsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotGroupConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnapshotGroupsConnectionBadRequest creates a GetSnapshotGroupsConnectionBadRequest with default headers values
func NewGetSnapshotGroupsConnectionBadRequest() *GetSnapshotGroupsConnectionBadRequest {
	return &GetSnapshotGroupsConnectionBadRequest{}
}

/* GetSnapshotGroupsConnectionBadRequest describes a response with status code 400, with default header values.

GetSnapshotGroupsConnectionBadRequest get snapshot groups connection bad request
*/
type GetSnapshotGroupsConnectionBadRequest struct {
	Payload string
}

func (o *GetSnapshotGroupsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-groups-connection][%d] getSnapshotGroupsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetSnapshotGroupsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetSnapshotGroupsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
