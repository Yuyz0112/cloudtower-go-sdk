// Code generated by go-swagger; DO NOT EDIT.

package alert_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alert rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alert rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAlertRules(params *GetAlertRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertRulesOK, error)

	GetAlertRulesConnection(params *GetAlertRulesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertRulesConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAlertRules get alert rules API
*/
func (a *Client) GetAlertRules(params *GetAlertRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAlertRules",
		Method:             "POST",
		PathPattern:        "/get-alert-rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAlertRulesReader{formats: a.formats},
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
	success, ok := result.(*GetAlertRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAlertRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAlertRulesConnection get alert rules connection API
*/
func (a *Client) GetAlertRulesConnection(params *GetAlertRulesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAlertRulesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertRulesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAlertRulesConnection",
		Method:             "POST",
		PathPattern:        "/get-alert-rules-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAlertRulesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetAlertRulesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAlertRulesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}