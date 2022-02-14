// Code generated by go-swagger; DO NOT EDIT.

package vds

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

// NewCreateVdsParams creates a new CreateVdsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVdsParams() *CreateVdsParams {
	return &CreateVdsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVdsParamsWithTimeout creates a new CreateVdsParams object
// with the ability to set a timeout on a request.
func NewCreateVdsParamsWithTimeout(timeout time.Duration) *CreateVdsParams {
	return &CreateVdsParams{
		timeout: timeout,
	}
}

// NewCreateVdsParamsWithContext creates a new CreateVdsParams object
// with the ability to set a context for a request.
func NewCreateVdsParamsWithContext(ctx context.Context) *CreateVdsParams {
	return &CreateVdsParams{
		Context: ctx,
	}
}

// NewCreateVdsParamsWithHTTPClient creates a new CreateVdsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVdsParamsWithHTTPClient(client *http.Client) *CreateVdsParams {
	return &CreateVdsParams{
		HTTPClient: client,
	}
}

/* CreateVdsParams contains all the parameters to send to the API endpoint
   for the create vds operation.

   Typically these are written to a http.Request.
*/
type CreateVdsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VdsCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create vds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVdsParams) WithDefaults() *CreateVdsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create vds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVdsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVdsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create vds params
func (o *CreateVdsParams) WithTimeout(timeout time.Duration) *CreateVdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create vds params
func (o *CreateVdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create vds params
func (o *CreateVdsParams) WithContext(ctx context.Context) *CreateVdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create vds params
func (o *CreateVdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create vds params
func (o *CreateVdsParams) WithHTTPClient(client *http.Client) *CreateVdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create vds params
func (o *CreateVdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create vds params
func (o *CreateVdsParams) WithContentLanguage(contentLanguage *string) *CreateVdsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create vds params
func (o *CreateVdsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create vds params
func (o *CreateVdsParams) WithRequestBody(requestBody []*models.VdsCreationParams) *CreateVdsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create vds params
func (o *CreateVdsParams) SetRequestBody(requestBody []*models.VdsCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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