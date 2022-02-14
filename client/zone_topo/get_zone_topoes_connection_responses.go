// Code generated by go-swagger; DO NOT EDIT.

package zone_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetZoneTopoesConnectionReader is a Reader for the GetZoneTopoesConnection structure.
type GetZoneTopoesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZoneTopoesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZoneTopoesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetZoneTopoesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetZoneTopoesConnectionOK creates a GetZoneTopoesConnectionOK with default headers values
func NewGetZoneTopoesConnectionOK() *GetZoneTopoesConnectionOK {
	return &GetZoneTopoesConnectionOK{}
}

/* GetZoneTopoesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetZoneTopoesConnectionOK struct {
	Payload *models.ZoneTopoConnection
}

func (o *GetZoneTopoesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes-connection][%d] getZoneTopoesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetZoneTopoesConnectionOK) GetPayload() *models.ZoneTopoConnection {
	return o.Payload
}

func (o *GetZoneTopoesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZoneTopoConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZoneTopoesConnectionBadRequest creates a GetZoneTopoesConnectionBadRequest with default headers values
func NewGetZoneTopoesConnectionBadRequest() *GetZoneTopoesConnectionBadRequest {
	return &GetZoneTopoesConnectionBadRequest{}
}

/* GetZoneTopoesConnectionBadRequest describes a response with status code 400, with default header values.

GetZoneTopoesConnectionBadRequest get zone topoes connection bad request
*/
type GetZoneTopoesConnectionBadRequest struct {
	Payload string
}

func (o *GetZoneTopoesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes-connection][%d] getZoneTopoesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetZoneTopoesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetZoneTopoesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
