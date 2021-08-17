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

// UpdateSnmpTransportReader is a Reader for the UpdateSnmpTransport structure.
type UpdateSnmpTransportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnmpTransportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSnmpTransportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSnmpTransportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSnmpTransportOK creates a UpdateSnmpTransportOK with default headers values
func NewUpdateSnmpTransportOK() *UpdateSnmpTransportOK {
	return &UpdateSnmpTransportOK{}
}

/* UpdateSnmpTransportOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateSnmpTransportOK struct {
	Payload []*models.WithTaskSnmpTransport
}

func (o *UpdateSnmpTransportOK) Error() string {
	return fmt.Sprintf("[POST /update-snmp-transport][%d] updateSnmpTransportOK  %+v", 200, o.Payload)
}
func (o *UpdateSnmpTransportOK) GetPayload() []*models.WithTaskSnmpTransport {
	return o.Payload
}

func (o *UpdateSnmpTransportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSnmpTransportBadRequest creates a UpdateSnmpTransportBadRequest with default headers values
func NewUpdateSnmpTransportBadRequest() *UpdateSnmpTransportBadRequest {
	return &UpdateSnmpTransportBadRequest{}
}

/* UpdateSnmpTransportBadRequest describes a response with status code 400, with default header values.

UpdateSnmpTransportBadRequest update snmp transport bad request
*/
type UpdateSnmpTransportBadRequest struct {
	Payload string
}

func (o *UpdateSnmpTransportBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-snmp-transport][%d] updateSnmpTransportBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateSnmpTransportBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateSnmpTransportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
