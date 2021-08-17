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

// NewRecoverVMFromRecycleBinParams creates a new RecoverVMFromRecycleBinParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecoverVMFromRecycleBinParams() *RecoverVMFromRecycleBinParams {
	return &RecoverVMFromRecycleBinParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecoverVMFromRecycleBinParamsWithTimeout creates a new RecoverVMFromRecycleBinParams object
// with the ability to set a timeout on a request.
func NewRecoverVMFromRecycleBinParamsWithTimeout(timeout time.Duration) *RecoverVMFromRecycleBinParams {
	return &RecoverVMFromRecycleBinParams{
		timeout: timeout,
	}
}

// NewRecoverVMFromRecycleBinParamsWithContext creates a new RecoverVMFromRecycleBinParams object
// with the ability to set a context for a request.
func NewRecoverVMFromRecycleBinParamsWithContext(ctx context.Context) *RecoverVMFromRecycleBinParams {
	return &RecoverVMFromRecycleBinParams{
		Context: ctx,
	}
}

// NewRecoverVMFromRecycleBinParamsWithHTTPClient creates a new RecoverVMFromRecycleBinParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecoverVMFromRecycleBinParamsWithHTTPClient(client *http.Client) *RecoverVMFromRecycleBinParams {
	return &RecoverVMFromRecycleBinParams{
		HTTPClient: client,
	}
}

/* RecoverVMFromRecycleBinParams contains all the parameters to send to the API endpoint
   for the recover Vm from recycle bin operation.

   Typically these are written to a http.Request.
*/
type RecoverVMFromRecycleBinParams struct {

	// RequestBody.
	RequestBody *models.VMOperateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the recover Vm from recycle bin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecoverVMFromRecycleBinParams) WithDefaults() *RecoverVMFromRecycleBinParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the recover Vm from recycle bin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecoverVMFromRecycleBinParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the recover Vm from recycle bin params
func (o *RecoverVMFromRecycleBinParams) WithTimeout(timeout time.Duration) *RecoverVMFromRecycleBinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recover Vm from recycle bin params
func (o *RecoverVMFromRecycleBinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recover Vm from recycle bin params
func (o *RecoverVMFromRecycleBinParams) WithContext(ctx context.Context) *RecoverVMFromRecycleBinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recover Vm from recycle bin params
func (o *RecoverVMFromRecycleBinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recover Vm from recycle bin params
func (o *RecoverVMFromRecycleBinParams) WithHTTPClient(client *http.Client) *RecoverVMFromRecycleBinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recover Vm from recycle bin params
func (o *RecoverVMFromRecycleBinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the recover Vm from recycle bin params
func (o *RecoverVMFromRecycleBinParams) WithRequestBody(requestBody *models.VMOperateParams) *RecoverVMFromRecycleBinParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the recover Vm from recycle bin params
func (o *RecoverVMFromRecycleBinParams) SetRequestBody(requestBody *models.VMOperateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RecoverVMFromRecycleBinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
