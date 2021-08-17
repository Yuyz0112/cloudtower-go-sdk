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

// UnmountUsbDeviceReader is a Reader for the UnmountUsbDevice structure.
type UnmountUsbDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnmountUsbDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnmountUsbDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnmountUsbDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnmountUsbDeviceOK creates a UnmountUsbDeviceOK with default headers values
func NewUnmountUsbDeviceOK() *UnmountUsbDeviceOK {
	return &UnmountUsbDeviceOK{}
}

/* UnmountUsbDeviceOK describes a response with status code 200, with default header values.

Ok
*/
type UnmountUsbDeviceOK struct {
	Payload []*models.WithTaskUsbDevice
}

func (o *UnmountUsbDeviceOK) Error() string {
	return fmt.Sprintf("[POST /unmount-usb-device][%d] unmountUsbDeviceOK  %+v", 200, o.Payload)
}
func (o *UnmountUsbDeviceOK) GetPayload() []*models.WithTaskUsbDevice {
	return o.Payload
}

func (o *UnmountUsbDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnmountUsbDeviceBadRequest creates a UnmountUsbDeviceBadRequest with default headers values
func NewUnmountUsbDeviceBadRequest() *UnmountUsbDeviceBadRequest {
	return &UnmountUsbDeviceBadRequest{}
}

/* UnmountUsbDeviceBadRequest describes a response with status code 400, with default header values.

UnmountUsbDeviceBadRequest unmount usb device bad request
*/
type UnmountUsbDeviceBadRequest struct {
	Payload string
}

func (o *UnmountUsbDeviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /unmount-usb-device][%d] unmountUsbDeviceBadRequest  %+v", 400, o.Payload)
}
func (o *UnmountUsbDeviceBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UnmountUsbDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
