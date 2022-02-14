// Code generated by go-swagger; DO NOT EDIT.

package brick_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetBrickTopoesConnectionReader is a Reader for the GetBrickTopoesConnection structure.
type GetBrickTopoesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBrickTopoesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBrickTopoesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBrickTopoesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBrickTopoesConnectionOK creates a GetBrickTopoesConnectionOK with default headers values
func NewGetBrickTopoesConnectionOK() *GetBrickTopoesConnectionOK {
	return &GetBrickTopoesConnectionOK{}
}

/* GetBrickTopoesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetBrickTopoesConnectionOK struct {
	Payload *models.BrickTopoConnection
}

func (o *GetBrickTopoesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-brick-topoes-connection][%d] getBrickTopoesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetBrickTopoesConnectionOK) GetPayload() *models.BrickTopoConnection {
	return o.Payload
}

func (o *GetBrickTopoesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BrickTopoConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBrickTopoesConnectionBadRequest creates a GetBrickTopoesConnectionBadRequest with default headers values
func NewGetBrickTopoesConnectionBadRequest() *GetBrickTopoesConnectionBadRequest {
	return &GetBrickTopoesConnectionBadRequest{}
}

/* GetBrickTopoesConnectionBadRequest describes a response with status code 400, with default header values.

GetBrickTopoesConnectionBadRequest get brick topoes connection bad request
*/
type GetBrickTopoesConnectionBadRequest struct {
	Payload string
}

func (o *GetBrickTopoesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-brick-topoes-connection][%d] getBrickTopoesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetBrickTopoesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetBrickTopoesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}