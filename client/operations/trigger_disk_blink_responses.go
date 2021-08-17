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

// TriggerDiskBlinkReader is a Reader for the TriggerDiskBlink structure.
type TriggerDiskBlinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerDiskBlinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriggerDiskBlinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTriggerDiskBlinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTriggerDiskBlinkOK creates a TriggerDiskBlinkOK with default headers values
func NewTriggerDiskBlinkOK() *TriggerDiskBlinkOK {
	return &TriggerDiskBlinkOK{}
}

/* TriggerDiskBlinkOK describes a response with status code 200, with default header values.

Ok
*/
type TriggerDiskBlinkOK struct {
	Payload []*models.WithTaskHost
}

func (o *TriggerDiskBlinkOK) Error() string {
	return fmt.Sprintf("[POST /trigger-disk-blink][%d] triggerDiskBlinkOK  %+v", 200, o.Payload)
}
func (o *TriggerDiskBlinkOK) GetPayload() []*models.WithTaskHost {
	return o.Payload
}

func (o *TriggerDiskBlinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggerDiskBlinkBadRequest creates a TriggerDiskBlinkBadRequest with default headers values
func NewTriggerDiskBlinkBadRequest() *TriggerDiskBlinkBadRequest {
	return &TriggerDiskBlinkBadRequest{}
}

/* TriggerDiskBlinkBadRequest describes a response with status code 400, with default header values.

TriggerDiskBlinkBadRequest trigger disk blink bad request
*/
type TriggerDiskBlinkBadRequest struct {
	Payload string
}

func (o *TriggerDiskBlinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /trigger-disk-blink][%d] triggerDiskBlinkBadRequest  %+v", 400, o.Payload)
}
func (o *TriggerDiskBlinkBadRequest) GetPayload() string {
	return o.Payload
}

func (o *TriggerDiskBlinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
