// Code generated by go-swagger; DO NOT EDIT.

package vm_entity_filter_result

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vm entity filter result API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vm entity filter result API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetVMEntityFilterResults(params *GetVMEntityFilterResultsParams, opts ...ClientOption) (*GetVMEntityFilterResultsOK, error)

	GetVMEntityFilterResultsConnection(params *GetVMEntityFilterResultsConnectionParams, opts ...ClientOption) (*GetVMEntityFilterResultsConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetVMEntityFilterResults get Vm entity filter results API
*/
func (a *Client) GetVMEntityFilterResults(params *GetVMEntityFilterResultsParams, opts ...ClientOption) (*GetVMEntityFilterResultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMEntityFilterResultsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVmEntityFilterResults",
		Method:             "POST",
		PathPattern:        "/get-vm-entity-filter-results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVMEntityFilterResultsReader{formats: a.formats},
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
	success, ok := result.(*GetVMEntityFilterResultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVmEntityFilterResults: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVMEntityFilterResultsConnection get Vm entity filter results connection API
*/
func (a *Client) GetVMEntityFilterResultsConnection(params *GetVMEntityFilterResultsConnectionParams, opts ...ClientOption) (*GetVMEntityFilterResultsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMEntityFilterResultsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVmEntityFilterResultsConnection",
		Method:             "POST",
		PathPattern:        "/get-vm-entity-filter-results-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVMEntityFilterResultsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetVMEntityFilterResultsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVmEntityFilterResultsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
