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

// NewTriggerDiskBlinkParams creates a new TriggerDiskBlinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriggerDiskBlinkParams() *TriggerDiskBlinkParams {
	return &TriggerDiskBlinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerDiskBlinkParamsWithTimeout creates a new TriggerDiskBlinkParams object
// with the ability to set a timeout on a request.
func NewTriggerDiskBlinkParamsWithTimeout(timeout time.Duration) *TriggerDiskBlinkParams {
	return &TriggerDiskBlinkParams{
		timeout: timeout,
	}
}

// NewTriggerDiskBlinkParamsWithContext creates a new TriggerDiskBlinkParams object
// with the ability to set a context for a request.
func NewTriggerDiskBlinkParamsWithContext(ctx context.Context) *TriggerDiskBlinkParams {
	return &TriggerDiskBlinkParams{
		Context: ctx,
	}
}

// NewTriggerDiskBlinkParamsWithHTTPClient creates a new TriggerDiskBlinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriggerDiskBlinkParamsWithHTTPClient(client *http.Client) *TriggerDiskBlinkParams {
	return &TriggerDiskBlinkParams{
		HTTPClient: client,
	}
}

/* TriggerDiskBlinkParams contains all the parameters to send to the API endpoint
   for the trigger disk blink operation.

   Typically these are written to a http.Request.
*/
type TriggerDiskBlinkParams struct {

	// RequestBody.
	RequestBody []*models.TriggerDiskBlinkParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the trigger disk blink params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerDiskBlinkParams) WithDefaults() *TriggerDiskBlinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trigger disk blink params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerDiskBlinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the trigger disk blink params
func (o *TriggerDiskBlinkParams) WithTimeout(timeout time.Duration) *TriggerDiskBlinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger disk blink params
func (o *TriggerDiskBlinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger disk blink params
func (o *TriggerDiskBlinkParams) WithContext(ctx context.Context) *TriggerDiskBlinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger disk blink params
func (o *TriggerDiskBlinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger disk blink params
func (o *TriggerDiskBlinkParams) WithHTTPClient(client *http.Client) *TriggerDiskBlinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger disk blink params
func (o *TriggerDiskBlinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the trigger disk blink params
func (o *TriggerDiskBlinkParams) WithRequestBody(requestBody []*models.TriggerDiskBlinkParams) *TriggerDiskBlinkParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the trigger disk blink params
func (o *TriggerDiskBlinkParams) SetRequestBody(requestBody []*models.TriggerDiskBlinkParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerDiskBlinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
