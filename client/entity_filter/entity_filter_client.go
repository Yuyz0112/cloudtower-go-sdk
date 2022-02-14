// Code generated by go-swagger; DO NOT EDIT.

package entity_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new entity filter API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for entity filter API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEntityFilter(params *CreateEntityFilterParams, opts ...ClientOption) (*CreateEntityFilterOK, error)

	DeleteEntityFilter(params *DeleteEntityFilterParams, opts ...ClientOption) (*DeleteEntityFilterOK, error)

	GetEntityFilters(params *GetEntityFiltersParams, opts ...ClientOption) (*GetEntityFiltersOK, error)

	GetEntityFiltersConnection(params *GetEntityFiltersConnectionParams, opts ...ClientOption) (*GetEntityFiltersConnectionOK, error)

	UpdateEntityFilter(params *UpdateEntityFilterParams, opts ...ClientOption) (*UpdateEntityFilterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateEntityFilter create entity filter API
*/
func (a *Client) CreateEntityFilter(params *CreateEntityFilterParams, opts ...ClientOption) (*CreateEntityFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEntityFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateEntityFilter",
		Method:             "POST",
		PathPattern:        "/create-entity-filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateEntityFilterReader{formats: a.formats},
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
	success, ok := result.(*CreateEntityFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateEntityFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteEntityFilter delete entity filter API
*/
func (a *Client) DeleteEntityFilter(params *DeleteEntityFilterParams, opts ...ClientOption) (*DeleteEntityFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEntityFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteEntityFilter",
		Method:             "POST",
		PathPattern:        "/delete-entity-filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEntityFilterReader{formats: a.formats},
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
	success, ok := result.(*DeleteEntityFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteEntityFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEntityFilters get entity filters API
*/
func (a *Client) GetEntityFilters(params *GetEntityFiltersParams, opts ...ClientOption) (*GetEntityFiltersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntityFiltersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEntityFilters",
		Method:             "POST",
		PathPattern:        "/get-entity-filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEntityFiltersReader{formats: a.formats},
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
	success, ok := result.(*GetEntityFiltersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEntityFilters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEntityFiltersConnection get entity filters connection API
*/
func (a *Client) GetEntityFiltersConnection(params *GetEntityFiltersConnectionParams, opts ...ClientOption) (*GetEntityFiltersConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntityFiltersConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetEntityFiltersConnection",
		Method:             "POST",
		PathPattern:        "/get-entity-filters-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEntityFiltersConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetEntityFiltersConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetEntityFiltersConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateEntityFilter update entity filter API
*/
func (a *Client) UpdateEntityFilter(params *UpdateEntityFilterParams, opts ...ClientOption) (*UpdateEntityFilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEntityFilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateEntityFilter",
		Method:             "POST",
		PathPattern:        "/update-entity-filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateEntityFilterReader{formats: a.formats},
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
	success, ok := result.(*UpdateEntityFilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateEntityFilter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
