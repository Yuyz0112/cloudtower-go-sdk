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

// GetOrganizationsReader is a Reader for the GetOrganizations structure.
type GetOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationsOK creates a GetOrganizationsOK with default headers values
func NewGetOrganizationsOK() *GetOrganizationsOK {
	return &GetOrganizationsOK{}
}

/* GetOrganizationsOK describes a response with status code 200, with default header values.

Ok
*/
type GetOrganizationsOK struct {
	Payload []*models.Organization
}

func (o *GetOrganizationsOK) Error() string {
	return fmt.Sprintf("[POST /get-organizations][%d] getOrganizationsOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationsOK) GetPayload() []*models.Organization {
	return o.Payload
}

func (o *GetOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}