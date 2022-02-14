// Code generated by go-swagger; DO NOT EDIT.

package elf_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// UpdateElfImageReader is a Reader for the UpdateElfImage structure.
type UpdateElfImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateElfImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateElfImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateElfImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateElfImageOK creates a UpdateElfImageOK with default headers values
func NewUpdateElfImageOK() *UpdateElfImageOK {
	return &UpdateElfImageOK{}
}

/* UpdateElfImageOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateElfImageOK struct {
	Payload []*models.WithTaskElfImage
}

func (o *UpdateElfImageOK) Error() string {
	return fmt.Sprintf("[POST /update-elf-image][%d] updateElfImageOK  %+v", 200, o.Payload)
}
func (o *UpdateElfImageOK) GetPayload() []*models.WithTaskElfImage {
	return o.Payload
}

func (o *UpdateElfImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateElfImageBadRequest creates a UpdateElfImageBadRequest with default headers values
func NewUpdateElfImageBadRequest() *UpdateElfImageBadRequest {
	return &UpdateElfImageBadRequest{}
}

/* UpdateElfImageBadRequest describes a response with status code 400, with default header values.

UpdateElfImageBadRequest update elf image bad request
*/
type UpdateElfImageBadRequest struct {
	Payload string
}

func (o *UpdateElfImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-elf-image][%d] updateElfImageBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateElfImageBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateElfImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}