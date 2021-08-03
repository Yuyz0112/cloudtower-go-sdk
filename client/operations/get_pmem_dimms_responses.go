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

// GetPmemDimmsReader is a Reader for the GetPmemDimms structure.
type GetPmemDimmsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPmemDimmsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPmemDimmsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPmemDimmsOK creates a GetPmemDimmsOK with default headers values
func NewGetPmemDimmsOK() *GetPmemDimmsOK {
	return &GetPmemDimmsOK{}
}

/* GetPmemDimmsOK describes a response with status code 200, with default header values.

Ok
*/
type GetPmemDimmsOK struct {
	Payload []*models.PmemDimm
}

func (o *GetPmemDimmsOK) Error() string {
	return fmt.Sprintf("[POST /get-pmem-dimms][%d] getPmemDimmsOK  %+v", 200, o.Payload)
}
func (o *GetPmemDimmsOK) GetPayload() []*models.PmemDimm {
	return o.Payload
}

func (o *GetPmemDimmsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
