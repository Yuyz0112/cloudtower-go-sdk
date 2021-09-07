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

// UpdateClusterLicenseReader is a Reader for the UpdateClusterLicense structure.
type UpdateClusterLicenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterLicenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterLicenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateClusterLicenseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateClusterLicenseOK creates a UpdateClusterLicenseOK with default headers values
func NewUpdateClusterLicenseOK() *UpdateClusterLicenseOK {
	return &UpdateClusterLicenseOK{}
}

/* UpdateClusterLicenseOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateClusterLicenseOK struct {
	Payload []*models.WithTaskCluster
}

func (o *UpdateClusterLicenseOK) Error() string {
	return fmt.Sprintf("[POST /update-cluster-license][%d] updateClusterLicenseOK  %+v", 200, o.Payload)
}
func (o *UpdateClusterLicenseOK) GetPayload() []*models.WithTaskCluster {
	return o.Payload
}

func (o *UpdateClusterLicenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterLicenseBadRequest creates a UpdateClusterLicenseBadRequest with default headers values
func NewUpdateClusterLicenseBadRequest() *UpdateClusterLicenseBadRequest {
	return &UpdateClusterLicenseBadRequest{}
}

/* UpdateClusterLicenseBadRequest describes a response with status code 400, with default header values.

UpdateClusterLicenseBadRequest update cluster license bad request
*/
type UpdateClusterLicenseBadRequest struct {
	Payload string
}

func (o *UpdateClusterLicenseBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-cluster-license][%d] updateClusterLicenseBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateClusterLicenseBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateClusterLicenseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}