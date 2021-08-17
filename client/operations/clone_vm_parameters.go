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

// NewCloneVMParams creates a new CloneVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloneVMParams() *CloneVMParams {
	return &CloneVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloneVMParamsWithTimeout creates a new CloneVMParams object
// with the ability to set a timeout on a request.
func NewCloneVMParamsWithTimeout(timeout time.Duration) *CloneVMParams {
	return &CloneVMParams{
		timeout: timeout,
	}
}

// NewCloneVMParamsWithContext creates a new CloneVMParams object
// with the ability to set a context for a request.
func NewCloneVMParamsWithContext(ctx context.Context) *CloneVMParams {
	return &CloneVMParams{
		Context: ctx,
	}
}

// NewCloneVMParamsWithHTTPClient creates a new CloneVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloneVMParamsWithHTTPClient(client *http.Client) *CloneVMParams {
	return &CloneVMParams{
		HTTPClient: client,
	}
}

/* CloneVMParams contains all the parameters to send to the API endpoint
   for the clone Vm operation.

   Typically these are written to a http.Request.
*/
type CloneVMParams struct {

	// RequestBody.
	RequestBody []*models.VMCloneParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clone Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneVMParams) WithDefaults() *CloneVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clone Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the clone Vm params
func (o *CloneVMParams) WithTimeout(timeout time.Duration) *CloneVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone Vm params
func (o *CloneVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone Vm params
func (o *CloneVMParams) WithContext(ctx context.Context) *CloneVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone Vm params
func (o *CloneVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone Vm params
func (o *CloneVMParams) WithHTTPClient(client *http.Client) *CloneVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone Vm params
func (o *CloneVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the clone Vm params
func (o *CloneVMParams) WithRequestBody(requestBody []*models.VMCloneParams) *CloneVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the clone Vm params
func (o *CloneVMParams) SetRequestBody(requestBody []*models.VMCloneParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CloneVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
