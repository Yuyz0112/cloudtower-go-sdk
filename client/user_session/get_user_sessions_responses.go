// Code generated by go-swagger; DO NOT EDIT.

package user_session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetUserSessionsReader is a Reader for the GetUserSessions structure.
type GetUserSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserSessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserSessionsOK creates a GetUserSessionsOK with default headers values
func NewGetUserSessionsOK() *GetUserSessionsOK {
	return &GetUserSessionsOK{}
}

/* GetUserSessionsOK describes a response with status code 200, with default header values.

Ok
*/
type GetUserSessionsOK struct {
	Payload []*models.UserSession
}

func (o *GetUserSessionsOK) Error() string {
	return fmt.Sprintf("[POST /get-user-sessions][%d] getUserSessionsOK  %+v", 200, o.Payload)
}
func (o *GetUserSessionsOK) GetPayload() []*models.UserSession {
	return o.Payload
}

func (o *GetUserSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserSessionsBadRequest creates a GetUserSessionsBadRequest with default headers values
func NewGetUserSessionsBadRequest() *GetUserSessionsBadRequest {
	return &GetUserSessionsBadRequest{}
}

/* GetUserSessionsBadRequest describes a response with status code 400, with default header values.

GetUserSessionsBadRequest get user sessions bad request
*/
type GetUserSessionsBadRequest struct {
	Payload string
}

func (o *GetUserSessionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-user-sessions][%d] getUserSessionsBadRequest  %+v", 400, o.Payload)
}
func (o *GetUserSessionsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetUserSessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
