// Code generated by go-swagger; DO NOT EDIT.

package content_library_vm_template

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

// NewGetContentLibraryVMTemplatesParams creates a new GetContentLibraryVMTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContentLibraryVMTemplatesParams() *GetContentLibraryVMTemplatesParams {
	return &GetContentLibraryVMTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentLibraryVMTemplatesParamsWithTimeout creates a new GetContentLibraryVMTemplatesParams object
// with the ability to set a timeout on a request.
func NewGetContentLibraryVMTemplatesParamsWithTimeout(timeout time.Duration) *GetContentLibraryVMTemplatesParams {
	return &GetContentLibraryVMTemplatesParams{
		timeout: timeout,
	}
}

// NewGetContentLibraryVMTemplatesParamsWithContext creates a new GetContentLibraryVMTemplatesParams object
// with the ability to set a context for a request.
func NewGetContentLibraryVMTemplatesParamsWithContext(ctx context.Context) *GetContentLibraryVMTemplatesParams {
	return &GetContentLibraryVMTemplatesParams{
		Context: ctx,
	}
}

// NewGetContentLibraryVMTemplatesParamsWithHTTPClient creates a new GetContentLibraryVMTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContentLibraryVMTemplatesParamsWithHTTPClient(client *http.Client) *GetContentLibraryVMTemplatesParams {
	return &GetContentLibraryVMTemplatesParams{
		HTTPClient: client,
	}
}

/* GetContentLibraryVMTemplatesParams contains all the parameters to send to the API endpoint
   for the get content library Vm templates operation.

   Typically these are written to a http.Request.
*/
type GetContentLibraryVMTemplatesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetContentLibraryVMTemplatesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get content library Vm templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentLibraryVMTemplatesParams) WithDefaults() *GetContentLibraryVMTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get content library Vm templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContentLibraryVMTemplatesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetContentLibraryVMTemplatesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) WithTimeout(timeout time.Duration) *GetContentLibraryVMTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) WithContext(ctx context.Context) *GetContentLibraryVMTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) WithHTTPClient(client *http.Client) *GetContentLibraryVMTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) WithContentLanguage(contentLanguage *string) *GetContentLibraryVMTemplatesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) WithRequestBody(requestBody *models.GetContentLibraryVMTemplatesRequestBody) *GetContentLibraryVMTemplatesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get content library Vm templates params
func (o *GetContentLibraryVMTemplatesParams) SetRequestBody(requestBody *models.GetContentLibraryVMTemplatesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentLibraryVMTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
