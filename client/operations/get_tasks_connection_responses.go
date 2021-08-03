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

// GetTasksConnectionReader is a Reader for the GetTasksConnection structure.
type GetTasksConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTasksConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTasksConnectionOK creates a GetTasksConnectionOK with default headers values
func NewGetTasksConnectionOK() *GetTasksConnectionOK {
	return &GetTasksConnectionOK{}
}

/* GetTasksConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetTasksConnectionOK struct {
	Payload *models.TaskConnection
}

func (o *GetTasksConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-tasks-connection][%d] getTasksConnectionOK  %+v", 200, o.Payload)
}
func (o *GetTasksConnectionOK) GetPayload() *models.TaskConnection {
	return o.Payload
}

func (o *GetTasksConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
