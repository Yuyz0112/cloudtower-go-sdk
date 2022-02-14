// Code generated by go-swagger; DO NOT EDIT.

package vm_nic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vm nic API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vm nic API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetVMNics(params *GetVMNicsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVMNicsOK, error)

	GetVMNicsConnection(params *GetVMNicsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVMNicsConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetVMNics get Vm nics API
*/
func (a *Client) GetVMNics(params *GetVMNicsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVMNicsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMNicsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVmNics",
		Method:             "POST",
		PathPattern:        "/get-vm-nics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVMNicsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVMNicsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVmNics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVMNicsConnection get Vm nics connection API
*/
func (a *Client) GetVMNicsConnection(params *GetVMNicsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVMNicsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMNicsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVmNicsConnection",
		Method:             "POST",
		PathPattern:        "/get-vm-nics-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVMNicsConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVMNicsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVmNicsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
