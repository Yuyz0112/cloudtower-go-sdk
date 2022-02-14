// Code generated by go-swagger; DO NOT EDIT.

package content_library_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetContentLibraryImagesReader is a Reader for the GetContentLibraryImages structure.
type GetContentLibraryImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentLibraryImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentLibraryImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentLibraryImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentLibraryImagesOK creates a GetContentLibraryImagesOK with default headers values
func NewGetContentLibraryImagesOK() *GetContentLibraryImagesOK {
	return &GetContentLibraryImagesOK{}
}

/* GetContentLibraryImagesOK describes a response with status code 200, with default header values.

Ok
*/
type GetContentLibraryImagesOK struct {
	Payload []*models.ContentLibraryImage
}

func (o *GetContentLibraryImagesOK) Error() string {
	return fmt.Sprintf("[POST /get-content-library-images][%d] getContentLibraryImagesOK  %+v", 200, o.Payload)
}
func (o *GetContentLibraryImagesOK) GetPayload() []*models.ContentLibraryImage {
	return o.Payload
}

func (o *GetContentLibraryImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentLibraryImagesBadRequest creates a GetContentLibraryImagesBadRequest with default headers values
func NewGetContentLibraryImagesBadRequest() *GetContentLibraryImagesBadRequest {
	return &GetContentLibraryImagesBadRequest{}
}

/* GetContentLibraryImagesBadRequest describes a response with status code 400, with default header values.

GetContentLibraryImagesBadRequest get content library images bad request
*/
type GetContentLibraryImagesBadRequest struct {
	Payload string
}

func (o *GetContentLibraryImagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-content-library-images][%d] getContentLibraryImagesBadRequest  %+v", 400, o.Payload)
}
func (o *GetContentLibraryImagesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetContentLibraryImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}