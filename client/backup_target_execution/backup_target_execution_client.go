// Code generated by go-swagger; DO NOT EDIT.

package backup_target_execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backup target execution API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backup target execution API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetBackupTargetExecutions(params *GetBackupTargetExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupTargetExecutionsOK, error)

	GetBackupTargetExecutionsConnection(params *GetBackupTargetExecutionsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupTargetExecutionsConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetBackupTargetExecutions get backup target executions API
*/
func (a *Client) GetBackupTargetExecutions(params *GetBackupTargetExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupTargetExecutionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupTargetExecutionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupTargetExecutions",
		Method:             "POST",
		PathPattern:        "/get-backup-target-executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackupTargetExecutionsReader{formats: a.formats},
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
	success, ok := result.(*GetBackupTargetExecutionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBackupTargetExecutions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBackupTargetExecutionsConnection get backup target executions connection API
*/
func (a *Client) GetBackupTargetExecutionsConnection(params *GetBackupTargetExecutionsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBackupTargetExecutionsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupTargetExecutionsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupTargetExecutionsConnection",
		Method:             "POST",
		PathPattern:        "/get-backup-target-executions-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackupTargetExecutionsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetBackupTargetExecutionsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBackupTargetExecutionsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
