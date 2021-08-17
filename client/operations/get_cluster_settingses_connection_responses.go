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

// GetClusterSettingsesConnectionReader is a Reader for the GetClusterSettingsesConnection structure.
type GetClusterSettingsesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterSettingsesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterSettingsesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterSettingsesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterSettingsesConnectionOK creates a GetClusterSettingsesConnectionOK with default headers values
func NewGetClusterSettingsesConnectionOK() *GetClusterSettingsesConnectionOK {
	return &GetClusterSettingsesConnectionOK{}
}

/* GetClusterSettingsesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterSettingsesConnectionOK struct {
	Payload *models.ClusterSettingsConnection
}

func (o *GetClusterSettingsesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-cluster-settingses-connection][%d] getClusterSettingsesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetClusterSettingsesConnectionOK) GetPayload() *models.ClusterSettingsConnection {
	return o.Payload
}

func (o *GetClusterSettingsesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterSettingsConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterSettingsesConnectionBadRequest creates a GetClusterSettingsesConnectionBadRequest with default headers values
func NewGetClusterSettingsesConnectionBadRequest() *GetClusterSettingsesConnectionBadRequest {
	return &GetClusterSettingsesConnectionBadRequest{}
}

/* GetClusterSettingsesConnectionBadRequest describes a response with status code 400, with default header values.

GetClusterSettingsesConnectionBadRequest get cluster settingses connection bad request
*/
type GetClusterSettingsesConnectionBadRequest struct {
	Payload string
}

func (o *GetClusterSettingsesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cluster-settingses-connection][%d] getClusterSettingsesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetClusterSettingsesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetClusterSettingsesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
