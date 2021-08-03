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

// GetDeploysConnectionReader is a Reader for the GetDeploysConnection structure.
type GetDeploysConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploysConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploysConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploysConnectionOK creates a GetDeploysConnectionOK with default headers values
func NewGetDeploysConnectionOK() *GetDeploysConnectionOK {
	return &GetDeploysConnectionOK{}
}

/* GetDeploysConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetDeploysConnectionOK struct {
	Payload *models.DeployConnection
}

func (o *GetDeploysConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-deploys-connection][%d] getDeploysConnectionOK  %+v", 200, o.Payload)
}
func (o *GetDeploysConnectionOK) GetPayload() *models.DeployConnection {
	return o.Payload
}

func (o *GetDeploysConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeployConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
