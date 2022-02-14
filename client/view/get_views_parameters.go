// Code generated by go-swagger; DO NOT EDIT.

package view

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

// NewGetViewsParams creates a new GetViewsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetViewsParams() *GetViewsParams {
	return &GetViewsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetViewsParamsWithTimeout creates a new GetViewsParams object
// with the ability to set a timeout on a request.
func NewGetViewsParamsWithTimeout(timeout time.Duration) *GetViewsParams {
	return &GetViewsParams{
		timeout: timeout,
	}
}

// NewGetViewsParamsWithContext creates a new GetViewsParams object
// with the ability to set a context for a request.
func NewGetViewsParamsWithContext(ctx context.Context) *GetViewsParams {
	return &GetViewsParams{
		Context: ctx,
	}
}

// NewGetViewsParamsWithHTTPClient creates a new GetViewsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetViewsParamsWithHTTPClient(client *http.Client) *GetViewsParams {
	return &GetViewsParams{
		HTTPClient: client,
	}
}

/* GetViewsParams contains all the parameters to send to the API endpoint
   for the get views operation.

   Typically these are written to a http.Request.
*/
type GetViewsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetViewsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetViewsParams) WithDefaults() *GetViewsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get views params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetViewsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetViewsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get views params
func (o *GetViewsParams) WithTimeout(timeout time.Duration) *GetViewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get views params
func (o *GetViewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get views params
func (o *GetViewsParams) WithContext(ctx context.Context) *GetViewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get views params
func (o *GetViewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get views params
func (o *GetViewsParams) WithHTTPClient(client *http.Client) *GetViewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get views params
func (o *GetViewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get views params
func (o *GetViewsParams) WithContentLanguage(contentLanguage *string) *GetViewsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get views params
func (o *GetViewsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get views params
func (o *GetViewsParams) WithRequestBody(requestBody *models.GetViewsRequestBody) *GetViewsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get views params
func (o *GetViewsParams) SetRequestBody(requestBody *models.GetViewsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetViewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentLanguage != nil {

		// header param content-language
		if err := r.SetHeaderParam("content-language", *o.ContentLanguage); err != nil {
			return err
		}
	}
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