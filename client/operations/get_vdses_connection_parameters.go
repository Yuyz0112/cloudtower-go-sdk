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

// NewGetVdsesConnectionParams creates a new GetVdsesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVdsesConnectionParams() *GetVdsesConnectionParams {
	return &GetVdsesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVdsesConnectionParamsWithTimeout creates a new GetVdsesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVdsesConnectionParamsWithTimeout(timeout time.Duration) *GetVdsesConnectionParams {
	return &GetVdsesConnectionParams{
		timeout: timeout,
	}
}

// NewGetVdsesConnectionParamsWithContext creates a new GetVdsesConnectionParams object
// with the ability to set a context for a request.
func NewGetVdsesConnectionParamsWithContext(ctx context.Context) *GetVdsesConnectionParams {
	return &GetVdsesConnectionParams{
		Context: ctx,
	}
}

// NewGetVdsesConnectionParamsWithHTTPClient creates a new GetVdsesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVdsesConnectionParamsWithHTTPClient(client *http.Client) *GetVdsesConnectionParams {
	return &GetVdsesConnectionParams{
		HTTPClient: client,
	}
}

/* GetVdsesConnectionParams contains all the parameters to send to the API endpoint
   for the get vdses connection operation.

   Typically these are written to a http.Request.
*/
type GetVdsesConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetVdsesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vdses connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVdsesConnectionParams) WithDefaults() *GetVdsesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vdses connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVdsesConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vdses connection params
func (o *GetVdsesConnectionParams) WithTimeout(timeout time.Duration) *GetVdsesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vdses connection params
func (o *GetVdsesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vdses connection params
func (o *GetVdsesConnectionParams) WithContext(ctx context.Context) *GetVdsesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vdses connection params
func (o *GetVdsesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vdses connection params
func (o *GetVdsesConnectionParams) WithHTTPClient(client *http.Client) *GetVdsesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vdses connection params
func (o *GetVdsesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get vdses connection params
func (o *GetVdsesConnectionParams) WithRequestBody(requestBody *models.GetVdsesConnectionRequestBody) *GetVdsesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get vdses connection params
func (o *GetVdsesConnectionParams) SetRequestBody(requestBody *models.GetVdsesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVdsesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
