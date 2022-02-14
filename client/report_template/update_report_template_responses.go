// Code generated by go-swagger; DO NOT EDIT.

package report_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// UpdateReportTemplateReader is a Reader for the UpdateReportTemplate structure.
type UpdateReportTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReportTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateReportTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateReportTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateReportTemplateOK creates a UpdateReportTemplateOK with default headers values
func NewUpdateReportTemplateOK() *UpdateReportTemplateOK {
	return &UpdateReportTemplateOK{}
}

/* UpdateReportTemplateOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateReportTemplateOK struct {
	Payload []*models.WithTaskReportTemplate
}

func (o *UpdateReportTemplateOK) Error() string {
	return fmt.Sprintf("[POST /update-report-template][%d] updateReportTemplateOK  %+v", 200, o.Payload)
}
func (o *UpdateReportTemplateOK) GetPayload() []*models.WithTaskReportTemplate {
	return o.Payload
}

func (o *UpdateReportTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReportTemplateBadRequest creates a UpdateReportTemplateBadRequest with default headers values
func NewUpdateReportTemplateBadRequest() *UpdateReportTemplateBadRequest {
	return &UpdateReportTemplateBadRequest{}
}

/* UpdateReportTemplateBadRequest describes a response with status code 400, with default header values.

UpdateReportTemplateBadRequest update report template bad request
*/
type UpdateReportTemplateBadRequest struct {
	Payload string
}

func (o *UpdateReportTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-report-template][%d] updateReportTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateReportTemplateBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateReportTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
