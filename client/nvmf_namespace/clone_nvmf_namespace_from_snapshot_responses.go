// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// CloneNvmfNamespaceFromSnapshotReader is a Reader for the CloneNvmfNamespaceFromSnapshot structure.
type CloneNvmfNamespaceFromSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneNvmfNamespaceFromSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneNvmfNamespaceFromSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloneNvmfNamespaceFromSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneNvmfNamespaceFromSnapshotOK creates a CloneNvmfNamespaceFromSnapshotOK with default headers values
func NewCloneNvmfNamespaceFromSnapshotOK() *CloneNvmfNamespaceFromSnapshotOK {
	return &CloneNvmfNamespaceFromSnapshotOK{}
}

/* CloneNvmfNamespaceFromSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type CloneNvmfNamespaceFromSnapshotOK struct {
	Payload []*models.WithTaskNvmfNamespace
}

func (o *CloneNvmfNamespaceFromSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /clone-nvmf-namespace-from-snapshot][%d] cloneNvmfNamespaceFromSnapshotOK  %+v", 200, o.Payload)
}
func (o *CloneNvmfNamespaceFromSnapshotOK) GetPayload() []*models.WithTaskNvmfNamespace {
	return o.Payload
}

func (o *CloneNvmfNamespaceFromSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneNvmfNamespaceFromSnapshotBadRequest creates a CloneNvmfNamespaceFromSnapshotBadRequest with default headers values
func NewCloneNvmfNamespaceFromSnapshotBadRequest() *CloneNvmfNamespaceFromSnapshotBadRequest {
	return &CloneNvmfNamespaceFromSnapshotBadRequest{}
}

/* CloneNvmfNamespaceFromSnapshotBadRequest describes a response with status code 400, with default header values.

CloneNvmfNamespaceFromSnapshotBadRequest clone nvmf namespace from snapshot bad request
*/
type CloneNvmfNamespaceFromSnapshotBadRequest struct {
	Payload string
}

func (o *CloneNvmfNamespaceFromSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /clone-nvmf-namespace-from-snapshot][%d] cloneNvmfNamespaceFromSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CloneNvmfNamespaceFromSnapshotBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CloneNvmfNamespaceFromSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
