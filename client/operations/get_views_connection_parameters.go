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

// NewGetViewsConnectionParams creates a new GetViewsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetViewsConnectionParams() *GetViewsConnectionParams {
	return &GetViewsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetViewsConnectionParamsWithTimeout creates a new GetViewsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetViewsConnectionParamsWithTimeout(timeout time.Duration) *GetViewsConnectionParams {
	return &GetViewsConnectionParams{
		timeout: timeout,
	}
}

// NewGetViewsConnectionParamsWithContext creates a new GetViewsConnectionParams object
// with the ability to set a context for a request.
func NewGetViewsConnectionParamsWithContext(ctx context.Context) *GetViewsConnectionParams {
	return &GetViewsConnectionParams{
		Context: ctx,
	}
}

// NewGetViewsConnectionParamsWithHTTPClient creates a new GetViewsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetViewsConnectionParamsWithHTTPClient(client *http.Client) *GetViewsConnectionParams {
	return &GetViewsConnectionParams{
		HTTPClient: client,
	}
}

/* GetViewsConnectionParams contains all the parameters to send to the API endpoint
   for the get views connection operation.

   Typically these are written to a http.Request.
*/
type GetViewsConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetViewsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get views connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetViewsConnectionParams) WithDefaults() *GetViewsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get views connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetViewsConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get views connection params
func (o *GetViewsConnectionParams) WithTimeout(timeout time.Duration) *GetViewsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get views connection params
func (o *GetViewsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get views connection params
func (o *GetViewsConnectionParams) WithContext(ctx context.Context) *GetViewsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get views connection params
func (o *GetViewsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get views connection params
func (o *GetViewsConnectionParams) WithHTTPClient(client *http.Client) *GetViewsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get views connection params
func (o *GetViewsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get views connection params
func (o *GetViewsConnectionParams) WithRequestBody(requestBody *models.GetViewsConnectionRequestBody) *GetViewsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get views connection params
func (o *GetViewsConnectionParams) SetRequestBody(requestBody *models.GetViewsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetViewsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
