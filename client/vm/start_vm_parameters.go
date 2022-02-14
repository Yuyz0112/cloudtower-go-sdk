// Code generated by go-swagger; DO NOT EDIT.

package vm

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

// NewStartVMParams creates a new StartVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartVMParams() *StartVMParams {
	return &StartVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartVMParamsWithTimeout creates a new StartVMParams object
// with the ability to set a timeout on a request.
func NewStartVMParamsWithTimeout(timeout time.Duration) *StartVMParams {
	return &StartVMParams{
		timeout: timeout,
	}
}

// NewStartVMParamsWithContext creates a new StartVMParams object
// with the ability to set a context for a request.
func NewStartVMParamsWithContext(ctx context.Context) *StartVMParams {
	return &StartVMParams{
		Context: ctx,
	}
}

// NewStartVMParamsWithHTTPClient creates a new StartVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartVMParamsWithHTTPClient(client *http.Client) *StartVMParams {
	return &StartVMParams{
		HTTPClient: client,
	}
}

/* StartVMParams contains all the parameters to send to the API endpoint
   for the start Vm operation.

   Typically these are written to a http.Request.
*/
type StartVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMStartParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartVMParams) WithDefaults() *StartVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := StartVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the start Vm params
func (o *StartVMParams) WithTimeout(timeout time.Duration) *StartVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start Vm params
func (o *StartVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start Vm params
func (o *StartVMParams) WithContext(ctx context.Context) *StartVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start Vm params
func (o *StartVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start Vm params
func (o *StartVMParams) WithHTTPClient(client *http.Client) *StartVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start Vm params
func (o *StartVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the start Vm params
func (o *StartVMParams) WithContentLanguage(contentLanguage *string) *StartVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the start Vm params
func (o *StartVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the start Vm params
func (o *StartVMParams) WithRequestBody(requestBody *models.VMStartParams) *StartVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the start Vm params
func (o *StartVMParams) SetRequestBody(requestBody *models.VMStartParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *StartVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentLanguage != nil {

		// header param content-language
		if err := r.SetHeaderParam("content-language", *o.ContentLanguage); err != nil {
			return err
		}
	}
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
