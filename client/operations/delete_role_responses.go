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

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoleOK creates a DeleteRoleOK with default headers values
func NewDeleteRoleOK() *DeleteRoleOK {
	return &DeleteRoleOK{}
}

/* DeleteRoleOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteRoleOK struct {
	Payload []*models.WithTaskDeleteRole
}

func (o *DeleteRoleOK) Error() string {
	return fmt.Sprintf("[POST /delete-role][%d] deleteRoleOK  %+v", 200, o.Payload)
}
func (o *DeleteRoleOK) GetPayload() []*models.WithTaskDeleteRole {
	return o.Payload
}

func (o *DeleteRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleBadRequest creates a DeleteRoleBadRequest with default headers values
func NewDeleteRoleBadRequest() *DeleteRoleBadRequest {
	return &DeleteRoleBadRequest{}
}

/* DeleteRoleBadRequest describes a response with status code 400, with default header values.

DeleteRoleBadRequest delete role bad request
*/
type DeleteRoleBadRequest struct {
	Payload string
}

func (o *DeleteRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-role][%d] deleteRoleBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteRoleBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
