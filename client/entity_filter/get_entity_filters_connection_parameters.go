// Code generated by go-swagger; DO NOT EDIT.

package entity_filter

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

// NewGetEntityFiltersConnectionParams creates a new GetEntityFiltersConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEntityFiltersConnectionParams() *GetEntityFiltersConnectionParams {
	return &GetEntityFiltersConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEntityFiltersConnectionParamsWithTimeout creates a new GetEntityFiltersConnectionParams object
// with the ability to set a timeout on a request.
func NewGetEntityFiltersConnectionParamsWithTimeout(timeout time.Duration) *GetEntityFiltersConnectionParams {
	return &GetEntityFiltersConnectionParams{
		timeout: timeout,
	}
}

// NewGetEntityFiltersConnectionParamsWithContext creates a new GetEntityFiltersConnectionParams object
// with the ability to set a context for a request.
func NewGetEntityFiltersConnectionParamsWithContext(ctx context.Context) *GetEntityFiltersConnectionParams {
	return &GetEntityFiltersConnectionParams{
		Context: ctx,
	}
}

// NewGetEntityFiltersConnectionParamsWithHTTPClient creates a new GetEntityFiltersConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEntityFiltersConnectionParamsWithHTTPClient(client *http.Client) *GetEntityFiltersConnectionParams {
	return &GetEntityFiltersConnectionParams{
		HTTPClient: client,
	}
}

/* GetEntityFiltersConnectionParams contains all the parameters to send to the API endpoint
   for the get entity filters connection operation.

   Typically these are written to a http.Request.
*/
type GetEntityFiltersConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetEntityFiltersConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get entity filters connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEntityFiltersConnectionParams) WithDefaults() *GetEntityFiltersConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get entity filters connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEntityFiltersConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetEntityFiltersConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) WithTimeout(timeout time.Duration) *GetEntityFiltersConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) WithContext(ctx context.Context) *GetEntityFiltersConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) WithHTTPClient(client *http.Client) *GetEntityFiltersConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) WithContentLanguage(contentLanguage *string) *GetEntityFiltersConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) WithRequestBody(requestBody *models.GetEntityFiltersConnectionRequestBody) *GetEntityFiltersConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get entity filters connection params
func (o *GetEntityFiltersConnectionParams) SetRequestBody(requestBody *models.GetEntityFiltersConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetEntityFiltersConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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