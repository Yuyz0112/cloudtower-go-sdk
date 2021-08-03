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

	"cloudtower-go-sdk/models"
)

// NewUpdateClusterParams creates a new UpdateClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateClusterParams() *UpdateClusterParams {
	return &UpdateClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterParamsWithTimeout creates a new UpdateClusterParams object
// with the ability to set a timeout on a request.
func NewUpdateClusterParamsWithTimeout(timeout time.Duration) *UpdateClusterParams {
	return &UpdateClusterParams{
		timeout: timeout,
	}
}

// NewUpdateClusterParamsWithContext creates a new UpdateClusterParams object
// with the ability to set a context for a request.
func NewUpdateClusterParamsWithContext(ctx context.Context) *UpdateClusterParams {
	return &UpdateClusterParams{
		Context: ctx,
	}
}

// NewUpdateClusterParamsWithHTTPClient creates a new UpdateClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateClusterParamsWithHTTPClient(client *http.Client) *UpdateClusterParams {
	return &UpdateClusterParams{
		HTTPClient: client,
	}
}

/* UpdateClusterParams contains all the parameters to send to the API endpoint
   for the update cluster operation.

   Typically these are written to a http.Request.
*/
type UpdateClusterParams struct {

	// RequestBody.
	RequestBody []*models.ClusterUpdateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClusterParams) WithDefaults() *UpdateClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) WithTimeout(timeout time.Duration) *UpdateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster params
func (o *UpdateClusterParams) WithContext(ctx context.Context) *UpdateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster params
func (o *UpdateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) WithHTTPClient(client *http.Client) *UpdateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update cluster params
func (o *UpdateClusterParams) WithRequestBody(requestBody []*models.ClusterUpdateParams) *UpdateClusterParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update cluster params
func (o *UpdateClusterParams) SetRequestBody(requestBody []*models.ClusterUpdateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
