// Code generated by go-swagger; DO NOT EDIT.

package zone_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new zone topo API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for zone topo API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetZoneTopoes(params *GetZoneTopoesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZoneTopoesOK, error)

	GetZoneTopoesConnection(params *GetZoneTopoesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZoneTopoesConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetZoneTopoes get zone topoes API
*/
func (a *Client) GetZoneTopoes(params *GetZoneTopoesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZoneTopoesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZoneTopoesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetZoneTopoes",
		Method:             "POST",
		PathPattern:        "/get-zone-topoes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetZoneTopoesReader{formats: a.formats},
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
	success, ok := result.(*GetZoneTopoesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetZoneTopoes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetZoneTopoesConnection get zone topoes connection API
*/
func (a *Client) GetZoneTopoesConnection(params *GetZoneTopoesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZoneTopoesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZoneTopoesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetZoneTopoesConnection",
		Method:             "POST",
		PathPattern:        "/get-zone-topoes-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetZoneTopoesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetZoneTopoesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetZoneTopoesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}