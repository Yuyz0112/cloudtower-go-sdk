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

// GetConsistencyGroupsConnectionReader is a Reader for the GetConsistencyGroupsConnection structure.
type GetConsistencyGroupsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsistencyGroupsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConsistencyGroupsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConsistencyGroupsConnectionOK creates a GetConsistencyGroupsConnectionOK with default headers values
func NewGetConsistencyGroupsConnectionOK() *GetConsistencyGroupsConnectionOK {
	return &GetConsistencyGroupsConnectionOK{}
}

/* GetConsistencyGroupsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetConsistencyGroupsConnectionOK struct {
	Payload *models.ConsistencyGroupConnection
}

func (o *GetConsistencyGroupsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-consistency-groups-connection][%d] getConsistencyGroupsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetConsistencyGroupsConnectionOK) GetPayload() *models.ConsistencyGroupConnection {
	return o.Payload
}

func (o *GetConsistencyGroupsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsistencyGroupConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
