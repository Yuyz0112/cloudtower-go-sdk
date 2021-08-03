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

// GetLabelsConnectionReader is a Reader for the GetLabelsConnection structure.
type GetLabelsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLabelsConnectionOK creates a GetLabelsConnectionOK with default headers values
func NewGetLabelsConnectionOK() *GetLabelsConnectionOK {
	return &GetLabelsConnectionOK{}
}

/* GetLabelsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetLabelsConnectionOK struct {
	Payload *models.LabelConnection
}

func (o *GetLabelsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-labels-connection][%d] getLabelsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetLabelsConnectionOK) GetPayload() *models.LabelConnection {
	return o.Payload
}

func (o *GetLabelsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
