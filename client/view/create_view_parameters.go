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

// NewCreateViewParams creates a new CreateViewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateViewParams() *CreateViewParams {
	return &CreateViewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateViewParamsWithTimeout creates a new CreateViewParams object
// with the ability to set a timeout on a request.
func NewCreateViewParamsWithTimeout(timeout time.Duration) *CreateViewParams {
	return &CreateViewParams{
		timeout: timeout,
	}
}

// NewCreateViewParamsWithContext creates a new CreateViewParams object
// with the ability to set a context for a request.
func NewCreateViewParamsWithContext(ctx context.Context) *CreateViewParams {
	return &CreateViewParams{
		Context: ctx,
	}
}

// NewCreateViewParamsWithHTTPClient creates a new CreateViewParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateViewParamsWithHTTPClient(client *http.Client) *CreateViewParams {
	return &CreateViewParams{
		HTTPClient: client,
	}
}

/* CreateViewParams contains all the parameters to send to the API endpoint
   for the create view operation.

   Typically these are written to a http.Request.
*/
type CreateViewParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.ViewCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateViewParams) WithDefaults() *CreateViewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateViewParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateViewParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create view params
func (o *CreateViewParams) WithTimeout(timeout time.Duration) *CreateViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create view params
func (o *CreateViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create view params
func (o *CreateViewParams) WithContext(ctx context.Context) *CreateViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create view params
func (o *CreateViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create view params
func (o *CreateViewParams) WithHTTPClient(client *http.Client) *CreateViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create view params
func (o *CreateViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create view params
func (o *CreateViewParams) WithContentLanguage(contentLanguage *string) *CreateViewParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create view params
func (o *CreateViewParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create view params
func (o *CreateViewParams) WithRequestBody(requestBody []*models.ViewCreationParams) *CreateViewParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create view params
func (o *CreateViewParams) SetRequestBody(requestBody []*models.ViewCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
