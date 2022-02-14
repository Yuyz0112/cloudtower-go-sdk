// Code generated by go-swagger; DO NOT EDIT.

package snmp_trap_receiver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetSnmpTrapReceiversReader is a Reader for the GetSnmpTrapReceivers structure.
type GetSnmpTrapReceiversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnmpTrapReceiversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnmpTrapReceiversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSnmpTrapReceiversBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnmpTrapReceiversOK creates a GetSnmpTrapReceiversOK with default headers values
func NewGetSnmpTrapReceiversOK() *GetSnmpTrapReceiversOK {
	return &GetSnmpTrapReceiversOK{}
}

/* GetSnmpTrapReceiversOK describes a response with status code 200, with default header values.

Ok
*/
type GetSnmpTrapReceiversOK struct {
	Payload []*models.SnmpTrapReceiver
}

func (o *GetSnmpTrapReceiversOK) Error() string {
	return fmt.Sprintf("[POST /get-snmp-trap-receivers][%d] getSnmpTrapReceiversOK  %+v", 200, o.Payload)
}
func (o *GetSnmpTrapReceiversOK) GetPayload() []*models.SnmpTrapReceiver {
	return o.Payload
}

func (o *GetSnmpTrapReceiversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnmpTrapReceiversBadRequest creates a GetSnmpTrapReceiversBadRequest with default headers values
func NewGetSnmpTrapReceiversBadRequest() *GetSnmpTrapReceiversBadRequest {
	return &GetSnmpTrapReceiversBadRequest{}
}

/* GetSnmpTrapReceiversBadRequest describes a response with status code 400, with default header values.

GetSnmpTrapReceiversBadRequest get snmp trap receivers bad request
*/
type GetSnmpTrapReceiversBadRequest struct {
	Payload string
}

func (o *GetSnmpTrapReceiversBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-snmp-trap-receivers][%d] getSnmpTrapReceiversBadRequest  %+v", 400, o.Payload)
}
func (o *GetSnmpTrapReceiversBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetSnmpTrapReceiversBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
