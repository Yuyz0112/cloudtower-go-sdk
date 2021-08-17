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

// DeleteNamespaceGroupReader is a Reader for the DeleteNamespaceGroup structure.
type DeleteNamespaceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNamespaceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNamespaceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNamespaceGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNamespaceGroupOK creates a DeleteNamespaceGroupOK with default headers values
func NewDeleteNamespaceGroupOK() *DeleteNamespaceGroupOK {
	return &DeleteNamespaceGroupOK{}
}

/* DeleteNamespaceGroupOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteNamespaceGroupOK struct {
	Payload []*models.WithTaskDeleteNamespaceGroup
}

func (o *DeleteNamespaceGroupOK) Error() string {
	return fmt.Sprintf("[POST /delete-namespace-group][%d] deleteNamespaceGroupOK  %+v", 200, o.Payload)
}
func (o *DeleteNamespaceGroupOK) GetPayload() []*models.WithTaskDeleteNamespaceGroup {
	return o.Payload
}

func (o *DeleteNamespaceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNamespaceGroupBadRequest creates a DeleteNamespaceGroupBadRequest with default headers values
func NewDeleteNamespaceGroupBadRequest() *DeleteNamespaceGroupBadRequest {
	return &DeleteNamespaceGroupBadRequest{}
}

/* DeleteNamespaceGroupBadRequest describes a response with status code 400, with default header values.

DeleteNamespaceGroupBadRequest delete namespace group bad request
*/
type DeleteNamespaceGroupBadRequest struct {
	Payload string
}

func (o *DeleteNamespaceGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-namespace-group][%d] deleteNamespaceGroupBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteNamespaceGroupBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteNamespaceGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
