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

// GetNvmfNamespaceSnapshotsReader is a Reader for the GetNvmfNamespaceSnapshots structure.
type GetNvmfNamespaceSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNvmfNamespaceSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNvmfNamespaceSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNvmfNamespaceSnapshotsOK creates a GetNvmfNamespaceSnapshotsOK with default headers values
func NewGetNvmfNamespaceSnapshotsOK() *GetNvmfNamespaceSnapshotsOK {
	return &GetNvmfNamespaceSnapshotsOK{}
}

/* GetNvmfNamespaceSnapshotsOK describes a response with status code 200, with default header values.

Ok
*/
type GetNvmfNamespaceSnapshotsOK struct {
	Payload []*models.NvmfNamespaceSnapshot
}

func (o *GetNvmfNamespaceSnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /get-nvmf-namespace-snapshots][%d] getNvmfNamespaceSnapshotsOK  %+v", 200, o.Payload)
}
func (o *GetNvmfNamespaceSnapshotsOK) GetPayload() []*models.NvmfNamespaceSnapshot {
	return o.Payload
}

func (o *GetNvmfNamespaceSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
