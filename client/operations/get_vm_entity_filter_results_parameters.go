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

// NewGetVMEntityFilterResultsParams creates a new GetVMEntityFilterResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMEntityFilterResultsParams() *GetVMEntityFilterResultsParams {
	return &GetVMEntityFilterResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMEntityFilterResultsParamsWithTimeout creates a new GetVMEntityFilterResultsParams object
// with the ability to set a timeout on a request.
func NewGetVMEntityFilterResultsParamsWithTimeout(timeout time.Duration) *GetVMEntityFilterResultsParams {
	return &GetVMEntityFilterResultsParams{
		timeout: timeout,
	}
}

// NewGetVMEntityFilterResultsParamsWithContext creates a new GetVMEntityFilterResultsParams object
// with the ability to set a context for a request.
func NewGetVMEntityFilterResultsParamsWithContext(ctx context.Context) *GetVMEntityFilterResultsParams {
	return &GetVMEntityFilterResultsParams{
		Context: ctx,
	}
}

// NewGetVMEntityFilterResultsParamsWithHTTPClient creates a new GetVMEntityFilterResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMEntityFilterResultsParamsWithHTTPClient(client *http.Client) *GetVMEntityFilterResultsParams {
	return &GetVMEntityFilterResultsParams{
		HTTPClient: client,
	}
}

/* GetVMEntityFilterResultsParams contains all the parameters to send to the API endpoint
   for the get Vm entity filter results operation.

   Typically these are written to a http.Request.
*/
type GetVMEntityFilterResultsParams struct {

	// RequestBody.
	RequestBody *models.GetVMEntityFilterResultsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm entity filter results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMEntityFilterResultsParams) WithDefaults() *GetVMEntityFilterResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm entity filter results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMEntityFilterResultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Vm entity filter results params
func (o *GetVMEntityFilterResultsParams) WithTimeout(timeout time.Duration) *GetVMEntityFilterResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm entity filter results params
func (o *GetVMEntityFilterResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm entity filter results params
func (o *GetVMEntityFilterResultsParams) WithContext(ctx context.Context) *GetVMEntityFilterResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm entity filter results params
func (o *GetVMEntityFilterResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm entity filter results params
func (o *GetVMEntityFilterResultsParams) WithHTTPClient(client *http.Client) *GetVMEntityFilterResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm entity filter results params
func (o *GetVMEntityFilterResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get Vm entity filter results params
func (o *GetVMEntityFilterResultsParams) WithRequestBody(requestBody *models.GetVMEntityFilterResultsRequestBody) *GetVMEntityFilterResultsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm entity filter results params
func (o *GetVMEntityFilterResultsParams) SetRequestBody(requestBody *models.GetVMEntityFilterResultsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMEntityFilterResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
