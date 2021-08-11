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

// NewUpdateVMTemplateFromVMParams creates a new UpdateVMTemplateFromVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMTemplateFromVMParams() *UpdateVMTemplateFromVMParams {
	return &UpdateVMTemplateFromVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMTemplateFromVMParamsWithTimeout creates a new UpdateVMTemplateFromVMParams object
// with the ability to set a timeout on a request.
func NewUpdateVMTemplateFromVMParamsWithTimeout(timeout time.Duration) *UpdateVMTemplateFromVMParams {
	return &UpdateVMTemplateFromVMParams{
		timeout: timeout,
	}
}

// NewUpdateVMTemplateFromVMParamsWithContext creates a new UpdateVMTemplateFromVMParams object
// with the ability to set a context for a request.
func NewUpdateVMTemplateFromVMParamsWithContext(ctx context.Context) *UpdateVMTemplateFromVMParams {
	return &UpdateVMTemplateFromVMParams{
		Context: ctx,
	}
}

// NewUpdateVMTemplateFromVMParamsWithHTTPClient creates a new UpdateVMTemplateFromVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMTemplateFromVMParamsWithHTTPClient(client *http.Client) *UpdateVMTemplateFromVMParams {
	return &UpdateVMTemplateFromVMParams{
		HTTPClient: client,
	}
}

/* UpdateVMTemplateFromVMParams contains all the parameters to send to the API endpoint
   for the update Vm template from Vm operation.

   Typically these are written to a http.Request.
*/
type UpdateVMTemplateFromVMParams struct {

	// RequestBody.
	RequestBody []*models.VMTemplateUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm template from Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMTemplateFromVMParams) WithDefaults() *UpdateVMTemplateFromVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm template from Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMTemplateFromVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Vm template from Vm params
func (o *UpdateVMTemplateFromVMParams) WithTimeout(timeout time.Duration) *UpdateVMTemplateFromVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm template from Vm params
func (o *UpdateVMTemplateFromVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm template from Vm params
func (o *UpdateVMTemplateFromVMParams) WithContext(ctx context.Context) *UpdateVMTemplateFromVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm template from Vm params
func (o *UpdateVMTemplateFromVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm template from Vm params
func (o *UpdateVMTemplateFromVMParams) WithHTTPClient(client *http.Client) *UpdateVMTemplateFromVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm template from Vm params
func (o *UpdateVMTemplateFromVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update Vm template from Vm params
func (o *UpdateVMTemplateFromVMParams) WithRequestBody(requestBody []*models.VMTemplateUpdationParams) *UpdateVMTemplateFromVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm template from Vm params
func (o *UpdateVMTemplateFromVMParams) SetRequestBody(requestBody []*models.VMTemplateUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMTemplateFromVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
