// Code generated by go-swagger; DO NOT EDIT.

package cluster_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetClusterSettingsesReader is a Reader for the GetClusterSettingses structure.
type GetClusterSettingsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterSettingsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterSettingsesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterSettingsesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterSettingsesOK creates a GetClusterSettingsesOK with default headers values
func NewGetClusterSettingsesOK() *GetClusterSettingsesOK {
	return &GetClusterSettingsesOK{}
}

/* GetClusterSettingsesOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterSettingsesOK struct {
	Payload []*models.ClusterSettings
}

func (o *GetClusterSettingsesOK) Error() string {
	return fmt.Sprintf("[POST /get-cluster-settingses][%d] getClusterSettingsesOK  %+v", 200, o.Payload)
}
func (o *GetClusterSettingsesOK) GetPayload() []*models.ClusterSettings {
	return o.Payload
}

func (o *GetClusterSettingsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterSettingsesBadRequest creates a GetClusterSettingsesBadRequest with default headers values
func NewGetClusterSettingsesBadRequest() *GetClusterSettingsesBadRequest {
	return &GetClusterSettingsesBadRequest{}
}

/* GetClusterSettingsesBadRequest describes a response with status code 400, with default header values.

GetClusterSettingsesBadRequest get cluster settingses bad request
*/
type GetClusterSettingsesBadRequest struct {
	Payload string
}

func (o *GetClusterSettingsesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cluster-settingses][%d] getClusterSettingsesBadRequest  %+v", 400, o.Payload)
}
func (o *GetClusterSettingsesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetClusterSettingsesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}