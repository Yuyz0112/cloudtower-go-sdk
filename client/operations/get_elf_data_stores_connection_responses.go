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

// GetElfDataStoresConnectionReader is a Reader for the GetElfDataStoresConnection structure.
type GetElfDataStoresConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetElfDataStoresConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetElfDataStoresConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetElfDataStoresConnectionOK creates a GetElfDataStoresConnectionOK with default headers values
func NewGetElfDataStoresConnectionOK() *GetElfDataStoresConnectionOK {
	return &GetElfDataStoresConnectionOK{}
}

/* GetElfDataStoresConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetElfDataStoresConnectionOK struct {
	Payload *models.ElfDataStoreConnection
}

func (o *GetElfDataStoresConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-elf-data-stores-connection][%d] getElfDataStoresConnectionOK  %+v", 200, o.Payload)
}
func (o *GetElfDataStoresConnectionOK) GetPayload() *models.ElfDataStoreConnection {
	return o.Payload
}

func (o *GetElfDataStoresConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElfDataStoreConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
