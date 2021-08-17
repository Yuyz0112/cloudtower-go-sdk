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

// GetIscsiTargetsReader is a Reader for the GetIscsiTargets structure.
type GetIscsiTargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIscsiTargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIscsiTargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIscsiTargetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIscsiTargetsOK creates a GetIscsiTargetsOK with default headers values
func NewGetIscsiTargetsOK() *GetIscsiTargetsOK {
	return &GetIscsiTargetsOK{}
}

/* GetIscsiTargetsOK describes a response with status code 200, with default header values.

Ok
*/
type GetIscsiTargetsOK struct {
	Payload []*models.IscsiTarget
}

func (o *GetIscsiTargetsOK) Error() string {
	return fmt.Sprintf("[POST /get-iscsi-targets][%d] getIscsiTargetsOK  %+v", 200, o.Payload)
}
func (o *GetIscsiTargetsOK) GetPayload() []*models.IscsiTarget {
	return o.Payload
}

func (o *GetIscsiTargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIscsiTargetsBadRequest creates a GetIscsiTargetsBadRequest with default headers values
func NewGetIscsiTargetsBadRequest() *GetIscsiTargetsBadRequest {
	return &GetIscsiTargetsBadRequest{}
}

/* GetIscsiTargetsBadRequest describes a response with status code 400, with default header values.

GetIscsiTargetsBadRequest get iscsi targets bad request
*/
type GetIscsiTargetsBadRequest struct {
	Payload string
}

func (o *GetIscsiTargetsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-iscsi-targets][%d] getIscsiTargetsBadRequest  %+v", 400, o.Payload)
}
func (o *GetIscsiTargetsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetIscsiTargetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
