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

// NewRemoveVMDiskParams creates a new RemoveVMDiskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveVMDiskParams() *RemoveVMDiskParams {
	return &RemoveVMDiskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveVMDiskParamsWithTimeout creates a new RemoveVMDiskParams object
// with the ability to set a timeout on a request.
func NewRemoveVMDiskParamsWithTimeout(timeout time.Duration) *RemoveVMDiskParams {
	return &RemoveVMDiskParams{
		timeout: timeout,
	}
}

// NewRemoveVMDiskParamsWithContext creates a new RemoveVMDiskParams object
// with the ability to set a context for a request.
func NewRemoveVMDiskParamsWithContext(ctx context.Context) *RemoveVMDiskParams {
	return &RemoveVMDiskParams{
		Context: ctx,
	}
}

// NewRemoveVMDiskParamsWithHTTPClient creates a new RemoveVMDiskParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveVMDiskParamsWithHTTPClient(client *http.Client) *RemoveVMDiskParams {
	return &RemoveVMDiskParams{
		HTTPClient: client,
	}
}

/* RemoveVMDiskParams contains all the parameters to send to the API endpoint
   for the remove Vm disk operation.

   Typically these are written to a http.Request.
*/
type RemoveVMDiskParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMRemoveDiskParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove Vm disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVMDiskParams) WithDefaults() *RemoveVMDiskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove Vm disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVMDiskParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := RemoveVMDiskParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the remove Vm disk params
func (o *RemoveVMDiskParams) WithTimeout(timeout time.Duration) *RemoveVMDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove Vm disk params
func (o *RemoveVMDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove Vm disk params
func (o *RemoveVMDiskParams) WithContext(ctx context.Context) *RemoveVMDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove Vm disk params
func (o *RemoveVMDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove Vm disk params
func (o *RemoveVMDiskParams) WithHTTPClient(client *http.Client) *RemoveVMDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove Vm disk params
func (o *RemoveVMDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the remove Vm disk params
func (o *RemoveVMDiskParams) WithContentLanguage(contentLanguage *string) *RemoveVMDiskParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the remove Vm disk params
func (o *RemoveVMDiskParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the remove Vm disk params
func (o *RemoveVMDiskParams) WithRequestBody(requestBody *models.VMRemoveDiskParams) *RemoveVMDiskParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the remove Vm disk params
func (o *RemoveVMDiskParams) SetRequestBody(requestBody *models.VMRemoveDiskParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveVMDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
