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

// NewSuspendVMParams creates a new SuspendVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSuspendVMParams() *SuspendVMParams {
	return &SuspendVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSuspendVMParamsWithTimeout creates a new SuspendVMParams object
// with the ability to set a timeout on a request.
func NewSuspendVMParamsWithTimeout(timeout time.Duration) *SuspendVMParams {
	return &SuspendVMParams{
		timeout: timeout,
	}
}

// NewSuspendVMParamsWithContext creates a new SuspendVMParams object
// with the ability to set a context for a request.
func NewSuspendVMParamsWithContext(ctx context.Context) *SuspendVMParams {
	return &SuspendVMParams{
		Context: ctx,
	}
}

// NewSuspendVMParamsWithHTTPClient creates a new SuspendVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewSuspendVMParamsWithHTTPClient(client *http.Client) *SuspendVMParams {
	return &SuspendVMParams{
		HTTPClient: client,
	}
}

/* SuspendVMParams contains all the parameters to send to the API endpoint
   for the suspend Vm operation.

   Typically these are written to a http.Request.
*/
type SuspendVMParams struct {

	// RequestBody.
	RequestBody *models.VMOperateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the suspend Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendVMParams) WithDefaults() *SuspendVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the suspend Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the suspend Vm params
func (o *SuspendVMParams) WithTimeout(timeout time.Duration) *SuspendVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suspend Vm params
func (o *SuspendVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suspend Vm params
func (o *SuspendVMParams) WithContext(ctx context.Context) *SuspendVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suspend Vm params
func (o *SuspendVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suspend Vm params
func (o *SuspendVMParams) WithHTTPClient(client *http.Client) *SuspendVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suspend Vm params
func (o *SuspendVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the suspend Vm params
func (o *SuspendVMParams) WithRequestBody(requestBody *models.VMOperateParams) *SuspendVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the suspend Vm params
func (o *SuspendVMParams) SetRequestBody(requestBody *models.VMOperateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *SuspendVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
