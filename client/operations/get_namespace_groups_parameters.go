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

// NewGetNamespaceGroupsParams creates a new GetNamespaceGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNamespaceGroupsParams() *GetNamespaceGroupsParams {
	return &GetNamespaceGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamespaceGroupsParamsWithTimeout creates a new GetNamespaceGroupsParams object
// with the ability to set a timeout on a request.
func NewGetNamespaceGroupsParamsWithTimeout(timeout time.Duration) *GetNamespaceGroupsParams {
	return &GetNamespaceGroupsParams{
		timeout: timeout,
	}
}

// NewGetNamespaceGroupsParamsWithContext creates a new GetNamespaceGroupsParams object
// with the ability to set a context for a request.
func NewGetNamespaceGroupsParamsWithContext(ctx context.Context) *GetNamespaceGroupsParams {
	return &GetNamespaceGroupsParams{
		Context: ctx,
	}
}

// NewGetNamespaceGroupsParamsWithHTTPClient creates a new GetNamespaceGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNamespaceGroupsParamsWithHTTPClient(client *http.Client) *GetNamespaceGroupsParams {
	return &GetNamespaceGroupsParams{
		HTTPClient: client,
	}
}

/* GetNamespaceGroupsParams contains all the parameters to send to the API endpoint
   for the get namespace groups operation.

   Typically these are written to a http.Request.
*/
type GetNamespaceGroupsParams struct {

	// RequestBody.
	RequestBody *models.GetNamespaceGroupsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get namespace groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamespaceGroupsParams) WithDefaults() *GetNamespaceGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get namespace groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamespaceGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get namespace groups params
func (o *GetNamespaceGroupsParams) WithTimeout(timeout time.Duration) *GetNamespaceGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get namespace groups params
func (o *GetNamespaceGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get namespace groups params
func (o *GetNamespaceGroupsParams) WithContext(ctx context.Context) *GetNamespaceGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get namespace groups params
func (o *GetNamespaceGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get namespace groups params
func (o *GetNamespaceGroupsParams) WithHTTPClient(client *http.Client) *GetNamespaceGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get namespace groups params
func (o *GetNamespaceGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get namespace groups params
func (o *GetNamespaceGroupsParams) WithRequestBody(requestBody *models.GetNamespaceGroupsRequestBody) *GetNamespaceGroupsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get namespace groups params
func (o *GetNamespaceGroupsParams) SetRequestBody(requestBody *models.GetNamespaceGroupsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNamespaceGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
