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

// NewGetVdsesParams creates a new GetVdsesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVdsesParams() *GetVdsesParams {
	return &GetVdsesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVdsesParamsWithTimeout creates a new GetVdsesParams object
// with the ability to set a timeout on a request.
func NewGetVdsesParamsWithTimeout(timeout time.Duration) *GetVdsesParams {
	return &GetVdsesParams{
		timeout: timeout,
	}
}

// NewGetVdsesParamsWithContext creates a new GetVdsesParams object
// with the ability to set a context for a request.
func NewGetVdsesParamsWithContext(ctx context.Context) *GetVdsesParams {
	return &GetVdsesParams{
		Context: ctx,
	}
}

// NewGetVdsesParamsWithHTTPClient creates a new GetVdsesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVdsesParamsWithHTTPClient(client *http.Client) *GetVdsesParams {
	return &GetVdsesParams{
		HTTPClient: client,
	}
}

/* GetVdsesParams contains all the parameters to send to the API endpoint
   for the get vdses operation.

   Typically these are written to a http.Request.
*/
type GetVdsesParams struct {

	// RequestBody.
	RequestBody *models.GetVdsesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vdses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVdsesParams) WithDefaults() *GetVdsesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vdses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVdsesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vdses params
func (o *GetVdsesParams) WithTimeout(timeout time.Duration) *GetVdsesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vdses params
func (o *GetVdsesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vdses params
func (o *GetVdsesParams) WithContext(ctx context.Context) *GetVdsesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vdses params
func (o *GetVdsesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vdses params
func (o *GetVdsesParams) WithHTTPClient(client *http.Client) *GetVdsesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vdses params
func (o *GetVdsesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get vdses params
func (o *GetVdsesParams) WithRequestBody(requestBody *models.GetVdsesRequestBody) *GetVdsesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get vdses params
func (o *GetVdsesParams) SetRequestBody(requestBody *models.GetVdsesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVdsesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
