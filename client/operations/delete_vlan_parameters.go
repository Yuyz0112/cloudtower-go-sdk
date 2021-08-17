// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Yuyz0112/cloudtower-go-sdk/models"
)

// NewDeleteVlanParams creates a new DeleteVlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVlanParams() *DeleteVlanParams {
	return &DeleteVlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVlanParamsWithTimeout creates a new DeleteVlanParams object
// with the ability to set a timeout on a request.
func NewDeleteVlanParamsWithTimeout(timeout time.Duration) *DeleteVlanParams {
	return &DeleteVlanParams{
		timeout: timeout,
	}
}

// NewDeleteVlanParamsWithContext creates a new DeleteVlanParams object
// with the ability to set a context for a request.
func NewDeleteVlanParamsWithContext(ctx context.Context) *DeleteVlanParams {
	return &DeleteVlanParams{
		Context: ctx,
	}
}

// NewDeleteVlanParamsWithHTTPClient creates a new DeleteVlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVlanParamsWithHTTPClient(client *http.Client) *DeleteVlanParams {
	return &DeleteVlanParams{
		HTTPClient: client,
	}
}

/* DeleteVlanParams contains all the parameters to send to the API endpoint
   for the delete vlan operation.

   Typically these are written to a http.Request.
*/
type DeleteVlanParams struct {

	// RequestBody.
	RequestBody *models.VlanDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVlanParams) WithDefaults() *DeleteVlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete vlan params
func (o *DeleteVlanParams) WithTimeout(timeout time.Duration) *DeleteVlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vlan params
func (o *DeleteVlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vlan params
func (o *DeleteVlanParams) WithContext(ctx context.Context) *DeleteVlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vlan params
func (o *DeleteVlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vlan params
func (o *DeleteVlanParams) WithHTTPClient(client *http.Client) *DeleteVlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vlan params
func (o *DeleteVlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete vlan params
func (o *DeleteVlanParams) WithRequestBody(requestBody *models.VlanDeletionParams) *DeleteVlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete vlan params
func (o *DeleteVlanParams) SetRequestBody(requestBody *models.VlanDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
