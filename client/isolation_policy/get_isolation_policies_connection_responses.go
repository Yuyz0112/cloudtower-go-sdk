// Code generated by go-swagger; DO NOT EDIT.

package isolation_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// GetIsolationPoliciesConnectionReader is a Reader for the GetIsolationPoliciesConnection structure.
type GetIsolationPoliciesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIsolationPoliciesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIsolationPoliciesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIsolationPoliciesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIsolationPoliciesConnectionOK creates a GetIsolationPoliciesConnectionOK with default headers values
func NewGetIsolationPoliciesConnectionOK() *GetIsolationPoliciesConnectionOK {
	return &GetIsolationPoliciesConnectionOK{}
}

/* GetIsolationPoliciesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetIsolationPoliciesConnectionOK struct {
	Payload *models.IsolationPolicyConnection
}

func (o *GetIsolationPoliciesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-isolation-policies-connection][%d] getIsolationPoliciesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetIsolationPoliciesConnectionOK) GetPayload() *models.IsolationPolicyConnection {
	return o.Payload
}

func (o *GetIsolationPoliciesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IsolationPolicyConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIsolationPoliciesConnectionBadRequest creates a GetIsolationPoliciesConnectionBadRequest with default headers values
func NewGetIsolationPoliciesConnectionBadRequest() *GetIsolationPoliciesConnectionBadRequest {
	return &GetIsolationPoliciesConnectionBadRequest{}
}

/* GetIsolationPoliciesConnectionBadRequest describes a response with status code 400, with default header values.

GetIsolationPoliciesConnectionBadRequest get isolation policies connection bad request
*/
type GetIsolationPoliciesConnectionBadRequest struct {
	Payload string
}

func (o *GetIsolationPoliciesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-isolation-policies-connection][%d] getIsolationPoliciesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetIsolationPoliciesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetIsolationPoliciesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
