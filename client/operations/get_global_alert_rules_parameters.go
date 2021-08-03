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

// NewGetGlobalAlertRulesParams creates a new GetGlobalAlertRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGlobalAlertRulesParams() *GetGlobalAlertRulesParams {
	return &GetGlobalAlertRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalAlertRulesParamsWithTimeout creates a new GetGlobalAlertRulesParams object
// with the ability to set a timeout on a request.
func NewGetGlobalAlertRulesParamsWithTimeout(timeout time.Duration) *GetGlobalAlertRulesParams {
	return &GetGlobalAlertRulesParams{
		timeout: timeout,
	}
}

// NewGetGlobalAlertRulesParamsWithContext creates a new GetGlobalAlertRulesParams object
// with the ability to set a context for a request.
func NewGetGlobalAlertRulesParamsWithContext(ctx context.Context) *GetGlobalAlertRulesParams {
	return &GetGlobalAlertRulesParams{
		Context: ctx,
	}
}

// NewGetGlobalAlertRulesParamsWithHTTPClient creates a new GetGlobalAlertRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGlobalAlertRulesParamsWithHTTPClient(client *http.Client) *GetGlobalAlertRulesParams {
	return &GetGlobalAlertRulesParams{
		HTTPClient: client,
	}
}

/* GetGlobalAlertRulesParams contains all the parameters to send to the API endpoint
   for the get global alert rules operation.

   Typically these are written to a http.Request.
*/
type GetGlobalAlertRulesParams struct {

	// RequestBody.
	RequestBody *models.GetGlobalAlertRulesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get global alert rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalAlertRulesParams) WithDefaults() *GetGlobalAlertRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get global alert rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalAlertRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get global alert rules params
func (o *GetGlobalAlertRulesParams) WithTimeout(timeout time.Duration) *GetGlobalAlertRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global alert rules params
func (o *GetGlobalAlertRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global alert rules params
func (o *GetGlobalAlertRulesParams) WithContext(ctx context.Context) *GetGlobalAlertRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global alert rules params
func (o *GetGlobalAlertRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global alert rules params
func (o *GetGlobalAlertRulesParams) WithHTTPClient(client *http.Client) *GetGlobalAlertRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global alert rules params
func (o *GetGlobalAlertRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get global alert rules params
func (o *GetGlobalAlertRulesParams) WithRequestBody(requestBody *models.GetGlobalAlertRulesRequestBody) *GetGlobalAlertRulesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get global alert rules params
func (o *GetGlobalAlertRulesParams) SetRequestBody(requestBody *models.GetGlobalAlertRulesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalAlertRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
