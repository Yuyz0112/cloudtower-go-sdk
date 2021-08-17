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

// GetClustersReader is a Reader for the GetClusters structure.
type GetClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClustersOK creates a GetClustersOK with default headers values
func NewGetClustersOK() *GetClustersOK {
	return &GetClustersOK{}
}

/* GetClustersOK describes a response with status code 200, with default header values.

Ok
*/
type GetClustersOK struct {
	Payload []*models.Cluster
}

func (o *GetClustersOK) Error() string {
	return fmt.Sprintf("[POST /get-clusters][%d] getClustersOK  %+v", 200, o.Payload)
}
func (o *GetClustersOK) GetPayload() []*models.Cluster {
	return o.Payload
}

func (o *GetClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersBadRequest creates a GetClustersBadRequest with default headers values
func NewGetClustersBadRequest() *GetClustersBadRequest {
	return &GetClustersBadRequest{}
}

/* GetClustersBadRequest describes a response with status code 400, with default header values.

GetClustersBadRequest get clusters bad request
*/
type GetClustersBadRequest struct {
	Payload string
}

func (o *GetClustersBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-clusters][%d] getClustersBadRequest  %+v", 400, o.Payload)
}
func (o *GetClustersBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
