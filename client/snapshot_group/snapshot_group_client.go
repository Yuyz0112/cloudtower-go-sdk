// Code generated by go-swagger; DO NOT EDIT.

package snapshot_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new snapshot group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for snapshot group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CloneSnapshotGroup(params *CloneSnapshotGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloneSnapshotGroupOK, error)

	DeleteSnapshotGroup(params *DeleteSnapshotGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSnapshotGroupOK, error)

	GetSnapshotGroups(params *GetSnapshotGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSnapshotGroupsOK, error)

	GetSnapshotGroupsConnection(params *GetSnapshotGroupsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSnapshotGroupsConnectionOK, error)

	KeepSnapshotGroup(params *KeepSnapshotGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeepSnapshotGroupOK, error)

	RollbackSnapshotGroup(params *RollbackSnapshotGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RollbackSnapshotGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CloneSnapshotGroup clone snapshot group API
*/
func (a *Client) CloneSnapshotGroup(params *CloneSnapshotGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CloneSnapshotGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCloneSnapshotGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CloneSnapshotGroup",
		Method:             "POST",
		PathPattern:        "/clone-snapshot-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CloneSnapshotGroupReader{formats: a.formats},
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
	success, ok := result.(*CloneSnapshotGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CloneSnapshotGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSnapshotGroup delete snapshot group API
*/
func (a *Client) DeleteSnapshotGroup(params *DeleteSnapshotGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSnapshotGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSnapshotGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSnapshotGroup",
		Method:             "POST",
		PathPattern:        "/delete-snapshot-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSnapshotGroupReader{formats: a.formats},
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
	success, ok := result.(*DeleteSnapshotGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteSnapshotGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSnapshotGroups get snapshot groups API
*/
func (a *Client) GetSnapshotGroups(params *GetSnapshotGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSnapshotGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSnapshotGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSnapshotGroups",
		Method:             "POST",
		PathPattern:        "/get-snapshot-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSnapshotGroupsReader{formats: a.formats},
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
	success, ok := result.(*GetSnapshotGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSnapshotGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSnapshotGroupsConnection get snapshot groups connection API
*/
func (a *Client) GetSnapshotGroupsConnection(params *GetSnapshotGroupsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSnapshotGroupsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSnapshotGroupsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSnapshotGroupsConnection",
		Method:             "POST",
		PathPattern:        "/get-snapshot-groups-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSnapshotGroupsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetSnapshotGroupsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSnapshotGroupsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KeepSnapshotGroup keep snapshot group API
*/
func (a *Client) KeepSnapshotGroup(params *KeepSnapshotGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeepSnapshotGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeepSnapshotGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KeepSnapshotGroup",
		Method:             "POST",
		PathPattern:        "/keep-snapshot-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KeepSnapshotGroupReader{formats: a.formats},
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
	success, ok := result.(*KeepSnapshotGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for KeepSnapshotGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RollbackSnapshotGroup rollback snapshot group API
*/
func (a *Client) RollbackSnapshotGroup(params *RollbackSnapshotGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RollbackSnapshotGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRollbackSnapshotGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RollbackSnapshotGroup",
		Method:             "POST",
		PathPattern:        "/rollback-snapshot-group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RollbackSnapshotGroupReader{formats: a.formats},
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
	success, ok := result.(*RollbackSnapshotGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RollbackSnapshotGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}