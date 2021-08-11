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

// NewUpdateMigrationVlanParams creates a new UpdateMigrationVlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMigrationVlanParams() *UpdateMigrationVlanParams {
	return &UpdateMigrationVlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMigrationVlanParamsWithTimeout creates a new UpdateMigrationVlanParams object
// with the ability to set a timeout on a request.
func NewUpdateMigrationVlanParamsWithTimeout(timeout time.Duration) *UpdateMigrationVlanParams {
	return &UpdateMigrationVlanParams{
		timeout: timeout,
	}
}

// NewUpdateMigrationVlanParamsWithContext creates a new UpdateMigrationVlanParams object
// with the ability to set a context for a request.
func NewUpdateMigrationVlanParamsWithContext(ctx context.Context) *UpdateMigrationVlanParams {
	return &UpdateMigrationVlanParams{
		Context: ctx,
	}
}

// NewUpdateMigrationVlanParamsWithHTTPClient creates a new UpdateMigrationVlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMigrationVlanParamsWithHTTPClient(client *http.Client) *UpdateMigrationVlanParams {
	return &UpdateMigrationVlanParams{
		HTTPClient: client,
	}
}

/* UpdateMigrationVlanParams contains all the parameters to send to the API endpoint
   for the update migration vlan operation.

   Typically these are written to a http.Request.
*/
type UpdateMigrationVlanParams struct {

	// RequestBody.
	RequestBody []*models.MigrationVlanUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update migration vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMigrationVlanParams) WithDefaults() *UpdateMigrationVlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update migration vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMigrationVlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update migration vlan params
func (o *UpdateMigrationVlanParams) WithTimeout(timeout time.Duration) *UpdateMigrationVlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update migration vlan params
func (o *UpdateMigrationVlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update migration vlan params
func (o *UpdateMigrationVlanParams) WithContext(ctx context.Context) *UpdateMigrationVlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update migration vlan params
func (o *UpdateMigrationVlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update migration vlan params
func (o *UpdateMigrationVlanParams) WithHTTPClient(client *http.Client) *UpdateMigrationVlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update migration vlan params
func (o *UpdateMigrationVlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update migration vlan params
func (o *UpdateMigrationVlanParams) WithRequestBody(requestBody []*models.MigrationVlanUpdationParams) *UpdateMigrationVlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update migration vlan params
func (o *UpdateMigrationVlanParams) SetRequestBody(requestBody []*models.MigrationVlanUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMigrationVlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
