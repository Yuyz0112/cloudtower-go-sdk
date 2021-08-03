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

// GetNvmfNamespacesReader is a Reader for the GetNvmfNamespaces structure.
type GetNvmfNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNvmfNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNvmfNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNvmfNamespacesOK creates a GetNvmfNamespacesOK with default headers values
func NewGetNvmfNamespacesOK() *GetNvmfNamespacesOK {
	return &GetNvmfNamespacesOK{}
}

/* GetNvmfNamespacesOK describes a response with status code 200, with default header values.

Ok
*/
type GetNvmfNamespacesOK struct {
	Payload []*models.NvmfNamespace
}

func (o *GetNvmfNamespacesOK) Error() string {
	return fmt.Sprintf("[POST /get-nvmf-namespaces][%d] getNvmfNamespacesOK  %+v", 200, o.Payload)
}
func (o *GetNvmfNamespacesOK) GetPayload() []*models.NvmfNamespace {
	return o.Payload
}

func (o *GetNvmfNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
