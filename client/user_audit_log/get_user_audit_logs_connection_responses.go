// Code generated by go-swagger; DO NOT EDIT.

package user_audit_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetUserAuditLogsConnectionReader is a Reader for the GetUserAuditLogsConnection structure.
type GetUserAuditLogsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserAuditLogsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserAuditLogsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserAuditLogsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserAuditLogsConnectionOK creates a GetUserAuditLogsConnectionOK with default headers values
func NewGetUserAuditLogsConnectionOK() *GetUserAuditLogsConnectionOK {
	return &GetUserAuditLogsConnectionOK{}
}

/* GetUserAuditLogsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetUserAuditLogsConnectionOK struct {
	Payload *models.UserAuditLogConnection
}

func (o *GetUserAuditLogsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-user-audit-logs-connection][%d] getUserAuditLogsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetUserAuditLogsConnectionOK) GetPayload() *models.UserAuditLogConnection {
	return o.Payload
}

func (o *GetUserAuditLogsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAuditLogConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserAuditLogsConnectionBadRequest creates a GetUserAuditLogsConnectionBadRequest with default headers values
func NewGetUserAuditLogsConnectionBadRequest() *GetUserAuditLogsConnectionBadRequest {
	return &GetUserAuditLogsConnectionBadRequest{}
}

/* GetUserAuditLogsConnectionBadRequest describes a response with status code 400, with default header values.

GetUserAuditLogsConnectionBadRequest get user audit logs connection bad request
*/
type GetUserAuditLogsConnectionBadRequest struct {
	Payload string
}

func (o *GetUserAuditLogsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-user-audit-logs-connection][%d] getUserAuditLogsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetUserAuditLogsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetUserAuditLogsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
