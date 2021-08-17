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

// GetHostsReader is a Reader for the GetHosts structure.
type GetHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHostsOK creates a GetHostsOK with default headers values
func NewGetHostsOK() *GetHostsOK {
	return &GetHostsOK{}
}

/* GetHostsOK describes a response with status code 200, with default header values.

Ok
*/
type GetHostsOK struct {
	Payload []*models.Host
}

func (o *GetHostsOK) Error() string {
	return fmt.Sprintf("[POST /get-hosts][%d] getHostsOK  %+v", 200, o.Payload)
}
func (o *GetHostsOK) GetPayload() []*models.Host {
	return o.Payload
}

func (o *GetHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostsBadRequest creates a GetHostsBadRequest with default headers values
func NewGetHostsBadRequest() *GetHostsBadRequest {
	return &GetHostsBadRequest{}
}

/* GetHostsBadRequest describes a response with status code 400, with default header values.

GetHostsBadRequest get hosts bad request
*/
type GetHostsBadRequest struct {
	Payload string
}

func (o *GetHostsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-hosts][%d] getHostsBadRequest  %+v", 400, o.Payload)
}
func (o *GetHostsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
