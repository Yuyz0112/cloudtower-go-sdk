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
	case 400:
		result := NewGetConsistencyGroupsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
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

// NewGetConsistencyGroupsConnectionBadRequest creates a GetConsistencyGroupsConnectionBadRequest with default headers values
func NewGetConsistencyGroupsConnectionBadRequest() *GetConsistencyGroupsConnectionBadRequest {
	return &GetConsistencyGroupsConnectionBadRequest{}
}

/* GetConsistencyGroupsConnectionBadRequest describes a response with status code 400, with default header values.

GetConsistencyGroupsConnectionBadRequest get consistency groups connection bad request
*/
type GetConsistencyGroupsConnectionBadRequest struct {
	Payload string
}

func (o *GetConsistencyGroupsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-consistency-groups-connection][%d] getConsistencyGroupsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetConsistencyGroupsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetConsistencyGroupsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
