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

// GetWitnessesConnectionReader is a Reader for the GetWitnessesConnection structure.
type GetWitnessesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWitnessesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWitnessesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWitnessesConnectionOK creates a GetWitnessesConnectionOK with default headers values
func NewGetWitnessesConnectionOK() *GetWitnessesConnectionOK {
	return &GetWitnessesConnectionOK{}
}

/* GetWitnessesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetWitnessesConnectionOK struct {
	Payload *models.WitnessConnection
}

func (o *GetWitnessesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-witnesses-connection][%d] getWitnessesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetWitnessesConnectionOK) GetPayload() *models.WitnessConnection {
	return o.Payload
}

func (o *GetWitnessesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WitnessConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
