// Code generated by go-swagger; DO NOT EDIT.

package backup_package

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

// NewGetBackupPackagesParams creates a new GetBackupPackagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBackupPackagesParams() *GetBackupPackagesParams {
	return &GetBackupPackagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupPackagesParamsWithTimeout creates a new GetBackupPackagesParams object
// with the ability to set a timeout on a request.
func NewGetBackupPackagesParamsWithTimeout(timeout time.Duration) *GetBackupPackagesParams {
	return &GetBackupPackagesParams{
		timeout: timeout,
	}
}

// NewGetBackupPackagesParamsWithContext creates a new GetBackupPackagesParams object
// with the ability to set a context for a request.
func NewGetBackupPackagesParamsWithContext(ctx context.Context) *GetBackupPackagesParams {
	return &GetBackupPackagesParams{
		Context: ctx,
	}
}

// NewGetBackupPackagesParamsWithHTTPClient creates a new GetBackupPackagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBackupPackagesParamsWithHTTPClient(client *http.Client) *GetBackupPackagesParams {
	return &GetBackupPackagesParams{
		HTTPClient: client,
	}
}

/* GetBackupPackagesParams contains all the parameters to send to the API endpoint
   for the get backup packages operation.

   Typically these are written to a http.Request.
*/
type GetBackupPackagesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetBackupPackagesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get backup packages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupPackagesParams) WithDefaults() *GetBackupPackagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get backup packages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupPackagesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetBackupPackagesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get backup packages params
func (o *GetBackupPackagesParams) WithTimeout(timeout time.Duration) *GetBackupPackagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup packages params
func (o *GetBackupPackagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup packages params
func (o *GetBackupPackagesParams) WithContext(ctx context.Context) *GetBackupPackagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup packages params
func (o *GetBackupPackagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup packages params
func (o *GetBackupPackagesParams) WithHTTPClient(client *http.Client) *GetBackupPackagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup packages params
func (o *GetBackupPackagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get backup packages params
func (o *GetBackupPackagesParams) WithContentLanguage(contentLanguage *string) *GetBackupPackagesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get backup packages params
func (o *GetBackupPackagesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get backup packages params
func (o *GetBackupPackagesParams) WithRequestBody(requestBody *models.GetBackupPackagesRequestBody) *GetBackupPackagesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get backup packages params
func (o *GetBackupPackagesParams) SetRequestBody(requestBody *models.GetBackupPackagesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupPackagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
