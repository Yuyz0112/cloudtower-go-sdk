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

// GetElfDataStoresReader is a Reader for the GetElfDataStores structure.
type GetElfDataStoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetElfDataStoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetElfDataStoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetElfDataStoresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetElfDataStoresOK creates a GetElfDataStoresOK with default headers values
func NewGetElfDataStoresOK() *GetElfDataStoresOK {
	return &GetElfDataStoresOK{}
}

/* GetElfDataStoresOK describes a response with status code 200, with default header values.

Ok
*/
type GetElfDataStoresOK struct {
	Payload []*models.ElfDataStore
}

func (o *GetElfDataStoresOK) Error() string {
	return fmt.Sprintf("[POST /get-elf-data-stores][%d] getElfDataStoresOK  %+v", 200, o.Payload)
}
func (o *GetElfDataStoresOK) GetPayload() []*models.ElfDataStore {
	return o.Payload
}

func (o *GetElfDataStoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetElfDataStoresBadRequest creates a GetElfDataStoresBadRequest with default headers values
func NewGetElfDataStoresBadRequest() *GetElfDataStoresBadRequest {
	return &GetElfDataStoresBadRequest{}
}

/* GetElfDataStoresBadRequest describes a response with status code 400, with default header values.

GetElfDataStoresBadRequest get elf data stores bad request
*/
type GetElfDataStoresBadRequest struct {
	Payload string
}

func (o *GetElfDataStoresBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-elf-data-stores][%d] getElfDataStoresBadRequest  %+v", 400, o.Payload)
}
func (o *GetElfDataStoresBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetElfDataStoresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
