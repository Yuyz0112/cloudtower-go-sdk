// Code generated by go-swagger; DO NOT EDIT.

package report_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetReportTasksConnectionReader is a Reader for the GetReportTasksConnection structure.
type GetReportTasksConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportTasksConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportTasksConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReportTasksConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportTasksConnectionOK creates a GetReportTasksConnectionOK with default headers values
func NewGetReportTasksConnectionOK() *GetReportTasksConnectionOK {
	return &GetReportTasksConnectionOK{}
}

/* GetReportTasksConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetReportTasksConnectionOK struct {
	Payload *models.ReportTaskConnection
}

func (o *GetReportTasksConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-report-tasks-connection][%d] getReportTasksConnectionOK  %+v", 200, o.Payload)
}
func (o *GetReportTasksConnectionOK) GetPayload() *models.ReportTaskConnection {
	return o.Payload
}

func (o *GetReportTasksConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportTaskConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportTasksConnectionBadRequest creates a GetReportTasksConnectionBadRequest with default headers values
func NewGetReportTasksConnectionBadRequest() *GetReportTasksConnectionBadRequest {
	return &GetReportTasksConnectionBadRequest{}
}

/* GetReportTasksConnectionBadRequest describes a response with status code 400, with default header values.

GetReportTasksConnectionBadRequest get report tasks connection bad request
*/
type GetReportTasksConnectionBadRequest struct {
	Payload string
}

func (o *GetReportTasksConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-report-tasks-connection][%d] getReportTasksConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetReportTasksConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetReportTasksConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
