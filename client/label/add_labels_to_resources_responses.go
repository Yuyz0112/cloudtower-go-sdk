// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// AddLabelsToResourcesReader is a Reader for the AddLabelsToResources structure.
type AddLabelsToResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddLabelsToResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddLabelsToResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddLabelsToResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddLabelsToResourcesOK creates a AddLabelsToResourcesOK with default headers values
func NewAddLabelsToResourcesOK() *AddLabelsToResourcesOK {
	return &AddLabelsToResourcesOK{}
}

/* AddLabelsToResourcesOK describes a response with status code 200, with default header values.

Ok
*/
type AddLabelsToResourcesOK struct {
	Payload []*models.WithTaskLabel
}

func (o *AddLabelsToResourcesOK) Error() string {
	return fmt.Sprintf("[POST /add-labels-to-resources][%d] addLabelsToResourcesOK  %+v", 200, o.Payload)
}
func (o *AddLabelsToResourcesOK) GetPayload() []*models.WithTaskLabel {
	return o.Payload
}

func (o *AddLabelsToResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLabelsToResourcesBadRequest creates a AddLabelsToResourcesBadRequest with default headers values
func NewAddLabelsToResourcesBadRequest() *AddLabelsToResourcesBadRequest {
	return &AddLabelsToResourcesBadRequest{}
}

/* AddLabelsToResourcesBadRequest describes a response with status code 400, with default header values.

AddLabelsToResourcesBadRequest add labels to resources bad request
*/
type AddLabelsToResourcesBadRequest struct {
	Payload string
}

func (o *AddLabelsToResourcesBadRequest) Error() string {
	return fmt.Sprintf("[POST /add-labels-to-resources][%d] addLabelsToResourcesBadRequest  %+v", 400, o.Payload)
}
func (o *AddLabelsToResourcesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *AddLabelsToResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
