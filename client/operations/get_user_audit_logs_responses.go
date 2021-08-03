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

// GetUserAuditLogsReader is a Reader for the GetUserAuditLogs structure.
type GetUserAuditLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserAuditLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserAuditLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserAuditLogsOK creates a GetUserAuditLogsOK with default headers values
func NewGetUserAuditLogsOK() *GetUserAuditLogsOK {
	return &GetUserAuditLogsOK{}
}

/* GetUserAuditLogsOK describes a response with status code 200, with default header values.

Ok
*/
type GetUserAuditLogsOK struct {
	Payload []*models.UserAuditLog
}

func (o *GetUserAuditLogsOK) Error() string {
	return fmt.Sprintf("[POST /get-user-audit-logs][%d] getUserAuditLogsOK  %+v", 200, o.Payload)
}
func (o *GetUserAuditLogsOK) GetPayload() []*models.UserAuditLog {
	return o.Payload
}

func (o *GetUserAuditLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
