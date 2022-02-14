// Code generated by go-swagger; DO NOT EDIT.

package cluster_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cluster settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetClusterSettingses(params *GetClusterSettingsesParams, opts ...ClientOption) (*GetClusterSettingsesOK, error)

	GetClusterSettingsesConnection(params *GetClusterSettingsesConnectionParams, opts ...ClientOption) (*GetClusterSettingsesConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetClusterSettingses get cluster settingses API
*/
func (a *Client) GetClusterSettingses(params *GetClusterSettingsesParams, opts ...ClientOption) (*GetClusterSettingsesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterSettingsesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetClusterSettingses",
		Method:             "POST",
		PathPattern:        "/get-cluster-settingses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterSettingsesReader{formats: a.formats},
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
	success, ok := result.(*GetClusterSettingsesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetClusterSettingses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClusterSettingsesConnection get cluster settingses connection API
*/
func (a *Client) GetClusterSettingsesConnection(params *GetClusterSettingsesConnectionParams, opts ...ClientOption) (*GetClusterSettingsesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterSettingsesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetClusterSettingsesConnection",
		Method:             "POST",
		PathPattern:        "/get-cluster-settingses-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterSettingsesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetClusterSettingsesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetClusterSettingsesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
