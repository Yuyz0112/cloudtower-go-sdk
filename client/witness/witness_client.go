// Code generated by go-swagger; DO NOT EDIT.

package witness

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new witness API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for witness API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetWitnesses(params *GetWitnessesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWitnessesOK, error)

	GetWitnessesConnection(params *GetWitnessesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWitnessesConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetWitnesses get witnesses API
*/
func (a *Client) GetWitnesses(params *GetWitnessesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWitnessesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWitnessesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWitnesses",
		Method:             "POST",
		PathPattern:        "/get-witnesses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWitnessesReader{formats: a.formats},
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
	success, ok := result.(*GetWitnessesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWitnesses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWitnessesConnection get witnesses connection API
*/
func (a *Client) GetWitnessesConnection(params *GetWitnessesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWitnessesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWitnessesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWitnessesConnection",
		Method:             "POST",
		PathPattern:        "/get-witnesses-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWitnessesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetWitnessesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWitnessesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
