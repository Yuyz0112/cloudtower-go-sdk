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

// NewGetDiscoverHostsParams creates a new GetDiscoverHostsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDiscoverHostsParams() *GetDiscoverHostsParams {
	return &GetDiscoverHostsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiscoverHostsParamsWithTimeout creates a new GetDiscoverHostsParams object
// with the ability to set a timeout on a request.
func NewGetDiscoverHostsParamsWithTimeout(timeout time.Duration) *GetDiscoverHostsParams {
	return &GetDiscoverHostsParams{
		timeout: timeout,
	}
}

// NewGetDiscoverHostsParamsWithContext creates a new GetDiscoverHostsParams object
// with the ability to set a context for a request.
func NewGetDiscoverHostsParamsWithContext(ctx context.Context) *GetDiscoverHostsParams {
	return &GetDiscoverHostsParams{
		Context: ctx,
	}
}

// NewGetDiscoverHostsParamsWithHTTPClient creates a new GetDiscoverHostsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDiscoverHostsParamsWithHTTPClient(client *http.Client) *GetDiscoverHostsParams {
	return &GetDiscoverHostsParams{
		HTTPClient: client,
	}
}

/* GetDiscoverHostsParams contains all the parameters to send to the API endpoint
   for the get discover hosts operation.

   Typically these are written to a http.Request.
*/
type GetDiscoverHostsParams struct {

	// RequestBody.
	RequestBody *models.GetDiscoverHostsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get discover hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiscoverHostsParams) WithDefaults() *GetDiscoverHostsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get discover hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiscoverHostsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get discover hosts params
func (o *GetDiscoverHostsParams) WithTimeout(timeout time.Duration) *GetDiscoverHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get discover hosts params
func (o *GetDiscoverHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get discover hosts params
func (o *GetDiscoverHostsParams) WithContext(ctx context.Context) *GetDiscoverHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get discover hosts params
func (o *GetDiscoverHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get discover hosts params
func (o *GetDiscoverHostsParams) WithHTTPClient(client *http.Client) *GetDiscoverHostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get discover hosts params
func (o *GetDiscoverHostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get discover hosts params
func (o *GetDiscoverHostsParams) WithRequestBody(requestBody *models.GetDiscoverHostsRequestBody) *GetDiscoverHostsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get discover hosts params
func (o *GetDiscoverHostsParams) SetRequestBody(requestBody *models.GetDiscoverHostsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiscoverHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
