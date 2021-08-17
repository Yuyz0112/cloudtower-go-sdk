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

// GetLabelsReader is a Reader for the GetLabels structure.
type GetLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLabelsOK creates a GetLabelsOK with default headers values
func NewGetLabelsOK() *GetLabelsOK {
	return &GetLabelsOK{}
}

/* GetLabelsOK describes a response with status code 200, with default header values.

Ok
*/
type GetLabelsOK struct {
	Payload []*models.Label
}

func (o *GetLabelsOK) Error() string {
	return fmt.Sprintf("[POST /get-labels][%d] getLabelsOK  %+v", 200, o.Payload)
}
func (o *GetLabelsOK) GetPayload() []*models.Label {
	return o.Payload
}

func (o *GetLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsBadRequest creates a GetLabelsBadRequest with default headers values
func NewGetLabelsBadRequest() *GetLabelsBadRequest {
	return &GetLabelsBadRequest{}
}

/* GetLabelsBadRequest describes a response with status code 400, with default header values.

GetLabelsBadRequest get labels bad request
*/
type GetLabelsBadRequest struct {
	Payload string
}

func (o *GetLabelsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-labels][%d] getLabelsBadRequest  %+v", 400, o.Payload)
}
func (o *GetLabelsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
