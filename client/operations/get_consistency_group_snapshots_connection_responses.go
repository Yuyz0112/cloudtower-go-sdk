// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cloudtower-go-sdk/models"
)

// GetConsistencyGroupSnapshotsConnectionReader is a Reader for the GetConsistencyGroupSnapshotsConnection structure.
type GetConsistencyGroupSnapshotsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsistencyGroupSnapshotsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConsistencyGroupSnapshotsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConsistencyGroupSnapshotsConnectionOK creates a GetConsistencyGroupSnapshotsConnectionOK with default headers values
func NewGetConsistencyGroupSnapshotsConnectionOK() *GetConsistencyGroupSnapshotsConnectionOK {
	return &GetConsistencyGroupSnapshotsConnectionOK{}
}

/* GetConsistencyGroupSnapshotsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetConsistencyGroupSnapshotsConnectionOK struct {
	Payload *models.ConsistencyGroupSnapshotConnection
}

func (o *GetConsistencyGroupSnapshotsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-consistency-group-snapshots-connection][%d] getConsistencyGroupSnapshotsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetConsistencyGroupSnapshotsConnectionOK) GetPayload() *models.ConsistencyGroupSnapshotConnection {
	return o.Payload
}

func (o *GetConsistencyGroupSnapshotsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsistencyGroupSnapshotConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
