// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetSnapshotPlanTasksReader is a Reader for the GetSnapshotPlanTasks structure.
type GetSnapshotPlanTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnapshotPlanTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnapshotPlanTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSnapshotPlanTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnapshotPlanTasksOK creates a GetSnapshotPlanTasksOK with default headers values
func NewGetSnapshotPlanTasksOK() *GetSnapshotPlanTasksOK {
	return &GetSnapshotPlanTasksOK{}
}

/* GetSnapshotPlanTasksOK describes a response with status code 200, with default header values.

Ok
*/
type GetSnapshotPlanTasksOK struct {
	Payload []*models.SnapshotPlanTask
}

func (o *GetSnapshotPlanTasksOK) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-plan-tasks][%d] getSnapshotPlanTasksOK  %+v", 200, o.Payload)
}
func (o *GetSnapshotPlanTasksOK) GetPayload() []*models.SnapshotPlanTask {
	return o.Payload
}

func (o *GetSnapshotPlanTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnapshotPlanTasksBadRequest creates a GetSnapshotPlanTasksBadRequest with default headers values
func NewGetSnapshotPlanTasksBadRequest() *GetSnapshotPlanTasksBadRequest {
	return &GetSnapshotPlanTasksBadRequest{}
}

/* GetSnapshotPlanTasksBadRequest describes a response with status code 400, with default header values.

GetSnapshotPlanTasksBadRequest get snapshot plan tasks bad request
*/
type GetSnapshotPlanTasksBadRequest struct {
	Payload string
}

func (o *GetSnapshotPlanTasksBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-plan-tasks][%d] getSnapshotPlanTasksBadRequest  %+v", 400, o.Payload)
}
func (o *GetSnapshotPlanTasksBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetSnapshotPlanTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
