// Code generated by go-swagger; DO NOT EDIT.

package rack_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetRackTopoesReader is a Reader for the GetRackTopoes structure.
type GetRackTopoesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRackTopoesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRackTopoesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRackTopoesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRackTopoesOK creates a GetRackTopoesOK with default headers values
func NewGetRackTopoesOK() *GetRackTopoesOK {
	return &GetRackTopoesOK{}
}

/* GetRackTopoesOK describes a response with status code 200, with default header values.

Ok
*/
type GetRackTopoesOK struct {
	Payload []*models.RackTopo
}

func (o *GetRackTopoesOK) Error() string {
	return fmt.Sprintf("[POST /get-rack-topoes][%d] getRackTopoesOK  %+v", 200, o.Payload)
}
func (o *GetRackTopoesOK) GetPayload() []*models.RackTopo {
	return o.Payload
}

func (o *GetRackTopoesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRackTopoesBadRequest creates a GetRackTopoesBadRequest with default headers values
func NewGetRackTopoesBadRequest() *GetRackTopoesBadRequest {
	return &GetRackTopoesBadRequest{}
}

/* GetRackTopoesBadRequest describes a response with status code 400, with default header values.

GetRackTopoesBadRequest get rack topoes bad request
*/
type GetRackTopoesBadRequest struct {
	Payload string
}

func (o *GetRackTopoesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-rack-topoes][%d] getRackTopoesBadRequest  %+v", 400, o.Payload)
}
func (o *GetRackTopoesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetRackTopoesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
