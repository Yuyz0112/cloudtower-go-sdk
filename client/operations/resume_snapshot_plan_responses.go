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

// ResumeSnapshotPlanReader is a Reader for the ResumeSnapshotPlan structure.
type ResumeSnapshotPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeSnapshotPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeSnapshotPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResumeSnapshotPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResumeSnapshotPlanOK creates a ResumeSnapshotPlanOK with default headers values
func NewResumeSnapshotPlanOK() *ResumeSnapshotPlanOK {
	return &ResumeSnapshotPlanOK{}
}

/* ResumeSnapshotPlanOK describes a response with status code 200, with default header values.

Ok
*/
type ResumeSnapshotPlanOK struct {
	Payload []*models.WithTaskSnapshotPlan
}

func (o *ResumeSnapshotPlanOK) Error() string {
	return fmt.Sprintf("[POST /resume-snapshot-plan][%d] resumeSnapshotPlanOK  %+v", 200, o.Payload)
}
func (o *ResumeSnapshotPlanOK) GetPayload() []*models.WithTaskSnapshotPlan {
	return o.Payload
}

func (o *ResumeSnapshotPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeSnapshotPlanBadRequest creates a ResumeSnapshotPlanBadRequest with default headers values
func NewResumeSnapshotPlanBadRequest() *ResumeSnapshotPlanBadRequest {
	return &ResumeSnapshotPlanBadRequest{}
}

/* ResumeSnapshotPlanBadRequest describes a response with status code 400, with default header values.

ResumeSnapshotPlanBadRequest resume snapshot plan bad request
*/
type ResumeSnapshotPlanBadRequest struct {
	Payload string
}

func (o *ResumeSnapshotPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /resume-snapshot-plan][%d] resumeSnapshotPlanBadRequest  %+v", 400, o.Payload)
}
func (o *ResumeSnapshotPlanBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ResumeSnapshotPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
