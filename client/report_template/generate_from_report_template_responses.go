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

// GenerateFromReportTemplateReader is a Reader for the GenerateFromReportTemplate structure.
type GenerateFromReportTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateFromReportTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateFromReportTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateFromReportTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateFromReportTemplateOK creates a GenerateFromReportTemplateOK with default headers values
func NewGenerateFromReportTemplateOK() *GenerateFromReportTemplateOK {
	return &GenerateFromReportTemplateOK{}
}

/* GenerateFromReportTemplateOK describes a response with status code 200, with default header values.

Ok
*/
type GenerateFromReportTemplateOK struct {
	Payload []*models.WithTaskReportTask
}

func (o *GenerateFromReportTemplateOK) Error() string {
	return fmt.Sprintf("[POST /generate-from-report-template][%d] generateFromReportTemplateOK  %+v", 200, o.Payload)
}
func (o *GenerateFromReportTemplateOK) GetPayload() []*models.WithTaskReportTask {
	return o.Payload
}

func (o *GenerateFromReportTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateFromReportTemplateBadRequest creates a GenerateFromReportTemplateBadRequest with default headers values
func NewGenerateFromReportTemplateBadRequest() *GenerateFromReportTemplateBadRequest {
	return &GenerateFromReportTemplateBadRequest{}
}

/* GenerateFromReportTemplateBadRequest describes a response with status code 400, with default header values.

GenerateFromReportTemplateBadRequest generate from report template bad request
*/
type GenerateFromReportTemplateBadRequest struct {
	Payload string
}

func (o *GenerateFromReportTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /generate-from-report-template][%d] generateFromReportTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *GenerateFromReportTemplateBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GenerateFromReportTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
