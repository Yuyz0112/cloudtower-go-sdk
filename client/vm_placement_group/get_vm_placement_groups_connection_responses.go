// Code generated by go-swagger; DO NOT EDIT.

package vm_placement_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetVMPlacementGroupsConnectionReader is a Reader for the GetVMPlacementGroupsConnection structure.
type GetVMPlacementGroupsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMPlacementGroupsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMPlacementGroupsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMPlacementGroupsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMPlacementGroupsConnectionOK creates a GetVMPlacementGroupsConnectionOK with default headers values
func NewGetVMPlacementGroupsConnectionOK() *GetVMPlacementGroupsConnectionOK {
	return &GetVMPlacementGroupsConnectionOK{}
}

/* GetVMPlacementGroupsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMPlacementGroupsConnectionOK struct {
	Payload *models.VMPlacementGroupConnection
}

func (o *GetVMPlacementGroupsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-placement-groups-connection][%d] getVmPlacementGroupsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVMPlacementGroupsConnectionOK) GetPayload() *models.VMPlacementGroupConnection {
	return o.Payload
}

func (o *GetVMPlacementGroupsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMPlacementGroupConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMPlacementGroupsConnectionBadRequest creates a GetVMPlacementGroupsConnectionBadRequest with default headers values
func NewGetVMPlacementGroupsConnectionBadRequest() *GetVMPlacementGroupsConnectionBadRequest {
	return &GetVMPlacementGroupsConnectionBadRequest{}
}

/* GetVMPlacementGroupsConnectionBadRequest describes a response with status code 400, with default header values.

GetVMPlacementGroupsConnectionBadRequest get Vm placement groups connection bad request
*/
type GetVMPlacementGroupsConnectionBadRequest struct {
	Payload string
}

func (o *GetVMPlacementGroupsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-placement-groups-connection][%d] getVmPlacementGroupsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMPlacementGroupsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVMPlacementGroupsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
