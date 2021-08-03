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

// GetClusterTopoesReader is a Reader for the GetClusterTopoes structure.
type GetClusterTopoesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterTopoesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterTopoesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterTopoesOK creates a GetClusterTopoesOK with default headers values
func NewGetClusterTopoesOK() *GetClusterTopoesOK {
	return &GetClusterTopoesOK{}
}

/* GetClusterTopoesOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterTopoesOK struct {
	Payload []*models.ClusterTopo
}

func (o *GetClusterTopoesOK) Error() string {
	return fmt.Sprintf("[POST /get-cluster-topoes][%d] getClusterTopoesOK  %+v", 200, o.Payload)
}
func (o *GetClusterTopoesOK) GetPayload() []*models.ClusterTopo {
	return o.Payload
}

func (o *GetClusterTopoesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
