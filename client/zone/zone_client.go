// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new zone API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for zone API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetZones(params *GetZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZonesOK, error)

	GetZonesConnection(params *GetZonesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZonesConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetZones get zones API
*/
func (a *Client) GetZones(params *GetZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetZones",
		Method:             "POST",
		PathPattern:        "/get-zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetZonesReader{formats: a.formats},
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
	success, ok := result.(*GetZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetZones: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetZonesConnection get zones connection API
*/
func (a *Client) GetZonesConnection(params *GetZonesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZonesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetZonesConnection",
		Method:             "POST",
		PathPattern:        "/get-zones-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetZonesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetZonesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetZonesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}