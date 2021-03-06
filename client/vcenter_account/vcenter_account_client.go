// Code generated by go-swagger; DO NOT EDIT.

package vcenter_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vcenter account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vcenter account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetVcenterAccounts(params *GetVcenterAccountsParams, opts ...ClientOption) (*GetVcenterAccountsOK, error)

	GetVcenterAccountsConnection(params *GetVcenterAccountsConnectionParams, opts ...ClientOption) (*GetVcenterAccountsConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetVcenterAccounts get vcenter accounts API
*/
func (a *Client) GetVcenterAccounts(params *GetVcenterAccountsParams, opts ...ClientOption) (*GetVcenterAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVcenterAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVcenterAccounts",
		Method:             "POST",
		PathPattern:        "/get-vcenter-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVcenterAccountsReader{formats: a.formats},
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
	success, ok := result.(*GetVcenterAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVcenterAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVcenterAccountsConnection get vcenter accounts connection API
*/
func (a *Client) GetVcenterAccountsConnection(params *GetVcenterAccountsConnectionParams, opts ...ClientOption) (*GetVcenterAccountsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVcenterAccountsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVcenterAccountsConnection",
		Method:             "POST",
		PathPattern:        "/get-vcenter-accounts-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVcenterAccountsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetVcenterAccountsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVcenterAccountsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
