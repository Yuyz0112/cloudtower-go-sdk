// Code generated by go-swagger; DO NOT EDIT.

package backup_restore_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backup restore point API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backup restore point API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetBackupRestorePoints(params *GetBackupRestorePointsParams, opts ...ClientOption) (*GetBackupRestorePointsOK, error)

	GetBackupRestorePointsConnection(params *GetBackupRestorePointsConnectionParams, opts ...ClientOption) (*GetBackupRestorePointsConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetBackupRestorePoints get backup restore points API
*/
func (a *Client) GetBackupRestorePoints(params *GetBackupRestorePointsParams, opts ...ClientOption) (*GetBackupRestorePointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupRestorePointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupRestorePoints",
		Method:             "POST",
		PathPattern:        "/get-backup-restore-points",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackupRestorePointsReader{formats: a.formats},
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
	success, ok := result.(*GetBackupRestorePointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBackupRestorePoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBackupRestorePointsConnection get backup restore points connection API
*/
func (a *Client) GetBackupRestorePointsConnection(params *GetBackupRestorePointsConnectionParams, opts ...ClientOption) (*GetBackupRestorePointsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupRestorePointsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBackupRestorePointsConnection",
		Method:             "POST",
		PathPattern:        "/get-backup-restore-points-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBackupRestorePointsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetBackupRestorePointsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBackupRestorePointsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
