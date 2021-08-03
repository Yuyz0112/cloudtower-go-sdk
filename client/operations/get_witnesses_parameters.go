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

// NewGetWitnessesParams creates a new GetWitnessesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWitnessesParams() *GetWitnessesParams {
	return &GetWitnessesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWitnessesParamsWithTimeout creates a new GetWitnessesParams object
// with the ability to set a timeout on a request.
func NewGetWitnessesParamsWithTimeout(timeout time.Duration) *GetWitnessesParams {
	return &GetWitnessesParams{
		timeout: timeout,
	}
}

// NewGetWitnessesParamsWithContext creates a new GetWitnessesParams object
// with the ability to set a context for a request.
func NewGetWitnessesParamsWithContext(ctx context.Context) *GetWitnessesParams {
	return &GetWitnessesParams{
		Context: ctx,
	}
}

// NewGetWitnessesParamsWithHTTPClient creates a new GetWitnessesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWitnessesParamsWithHTTPClient(client *http.Client) *GetWitnessesParams {
	return &GetWitnessesParams{
		HTTPClient: client,
	}
}

/* GetWitnessesParams contains all the parameters to send to the API endpoint
   for the get witnesses operation.

   Typically these are written to a http.Request.
*/
type GetWitnessesParams struct {

	// RequestBody.
	RequestBody *models.GetWitnessesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get witnesses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWitnessesParams) WithDefaults() *GetWitnessesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get witnesses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWitnessesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get witnesses params
func (o *GetWitnessesParams) WithTimeout(timeout time.Duration) *GetWitnessesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get witnesses params
func (o *GetWitnessesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get witnesses params
func (o *GetWitnessesParams) WithContext(ctx context.Context) *GetWitnessesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get witnesses params
func (o *GetWitnessesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get witnesses params
func (o *GetWitnessesParams) WithHTTPClient(client *http.Client) *GetWitnessesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get witnesses params
func (o *GetWitnessesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get witnesses params
func (o *GetWitnessesParams) WithRequestBody(requestBody *models.GetWitnessesRequestBody) *GetWitnessesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get witnesses params
func (o *GetWitnessesParams) SetRequestBody(requestBody *models.GetWitnessesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetWitnessesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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