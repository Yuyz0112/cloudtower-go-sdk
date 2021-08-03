// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"cloudtower-go-sdk/models"
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
	Payload []*models.ClusterSettings2
}

func (o *GetClusterSettingsesOK) Error() string {
	return fmt.Sprintf("[POST /get-cluster-settingses][%d] getClusterSettingsesOK  %+v", 200, o.Payload)
}
func (o *GetClusterSettingsesOK) GetPayload() []*models.ClusterSettings2 {
	return o.Payload
}

func (o *GetClusterSettingsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}