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

// NewAddVMNicParams creates a new AddVMNicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddVMNicParams() *AddVMNicParams {
	return &AddVMNicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddVMNicParamsWithTimeout creates a new AddVMNicParams object
// with the ability to set a timeout on a request.
func NewAddVMNicParamsWithTimeout(timeout time.Duration) *AddVMNicParams {
	return &AddVMNicParams{
		timeout: timeout,
	}
}

// NewAddVMNicParamsWithContext creates a new AddVMNicParams object
// with the ability to set a context for a request.
func NewAddVMNicParamsWithContext(ctx context.Context) *AddVMNicParams {
	return &AddVMNicParams{
		Context: ctx,
	}
}

// NewAddVMNicParamsWithHTTPClient creates a new AddVMNicParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddVMNicParamsWithHTTPClient(client *http.Client) *AddVMNicParams {
	return &AddVMNicParams{
		HTTPClient: client,
	}
}

/* AddVMNicParams contains all the parameters to send to the API endpoint
   for the add Vm nic operation.

   Typically these are written to a http.Request.
*/
type AddVMNicParams struct {

	// RequestBody.
	RequestBody []*models.VMAddNicParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add Vm nic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVMNicParams) WithDefaults() *AddVMNicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add Vm nic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddVMNicParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add Vm nic params
func (o *AddVMNicParams) WithTimeout(timeout time.Duration) *AddVMNicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Vm nic params
func (o *AddVMNicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Vm nic params
func (o *AddVMNicParams) WithContext(ctx context.Context) *AddVMNicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Vm nic params
func (o *AddVMNicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Vm nic params
func (o *AddVMNicParams) WithHTTPClient(client *http.Client) *AddVMNicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Vm nic params
func (o *AddVMNicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the add Vm nic params
func (o *AddVMNicParams) WithRequestBody(requestBody []*models.VMAddNicParams) *AddVMNicParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the add Vm nic params
func (o *AddVMNicParams) SetRequestBody(requestBody []*models.VMAddNicParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *AddVMNicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
