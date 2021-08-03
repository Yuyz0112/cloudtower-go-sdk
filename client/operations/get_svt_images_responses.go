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

// GetSvtImagesReader is a Reader for the GetSvtImages structure.
type GetSvtImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSvtImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSvtImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSvtImagesOK creates a GetSvtImagesOK with default headers values
func NewGetSvtImagesOK() *GetSvtImagesOK {
	return &GetSvtImagesOK{}
}

/* GetSvtImagesOK describes a response with status code 200, with default header values.

Ok
*/
type GetSvtImagesOK struct {
	Payload []*models.SvtImage
}

func (o *GetSvtImagesOK) Error() string {
	return fmt.Sprintf("[POST /get-svt-images][%d] getSvtImagesOK  %+v", 200, o.Payload)
}
func (o *GetSvtImagesOK) GetPayload() []*models.SvtImage {
	return o.Payload
}

func (o *GetSvtImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
