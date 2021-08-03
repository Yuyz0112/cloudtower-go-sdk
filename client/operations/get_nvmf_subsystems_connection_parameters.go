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

// NewGetNvmfSubsystemsConnectionParams creates a new GetNvmfSubsystemsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNvmfSubsystemsConnectionParams() *GetNvmfSubsystemsConnectionParams {
	return &GetNvmfSubsystemsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNvmfSubsystemsConnectionParamsWithTimeout creates a new GetNvmfSubsystemsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetNvmfSubsystemsConnectionParamsWithTimeout(timeout time.Duration) *GetNvmfSubsystemsConnectionParams {
	return &GetNvmfSubsystemsConnectionParams{
		timeout: timeout,
	}
}

// NewGetNvmfSubsystemsConnectionParamsWithContext creates a new GetNvmfSubsystemsConnectionParams object
// with the ability to set a context for a request.
func NewGetNvmfSubsystemsConnectionParamsWithContext(ctx context.Context) *GetNvmfSubsystemsConnectionParams {
	return &GetNvmfSubsystemsConnectionParams{
		Context: ctx,
	}
}

// NewGetNvmfSubsystemsConnectionParamsWithHTTPClient creates a new GetNvmfSubsystemsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNvmfSubsystemsConnectionParamsWithHTTPClient(client *http.Client) *GetNvmfSubsystemsConnectionParams {
	return &GetNvmfSubsystemsConnectionParams{
		HTTPClient: client,
	}
}

/* GetNvmfSubsystemsConnectionParams contains all the parameters to send to the API endpoint
   for the get nvmf subsystems connection operation.

   Typically these are written to a http.Request.
*/
type GetNvmfSubsystemsConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetNvmfSubsystemsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nvmf subsystems connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfSubsystemsConnectionParams) WithDefaults() *GetNvmfSubsystemsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nvmf subsystems connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfSubsystemsConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nvmf subsystems connection params
func (o *GetNvmfSubsystemsConnectionParams) WithTimeout(timeout time.Duration) *GetNvmfSubsystemsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nvmf subsystems connection params
func (o *GetNvmfSubsystemsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nvmf subsystems connection params
func (o *GetNvmfSubsystemsConnectionParams) WithContext(ctx context.Context) *GetNvmfSubsystemsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nvmf subsystems connection params
func (o *GetNvmfSubsystemsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nvmf subsystems connection params
func (o *GetNvmfSubsystemsConnectionParams) WithHTTPClient(client *http.Client) *GetNvmfSubsystemsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nvmf subsystems connection params
func (o *GetNvmfSubsystemsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get nvmf subsystems connection params
func (o *GetNvmfSubsystemsConnectionParams) WithRequestBody(requestBody *models.GetNvmfSubsystemsConnectionRequestBody) *GetNvmfSubsystemsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nvmf subsystems connection params
func (o *GetNvmfSubsystemsConnectionParams) SetRequestBody(requestBody *models.GetNvmfSubsystemsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNvmfSubsystemsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
