// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace

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

// NewUpdateNvmfNamespaceParams creates a new UpdateNvmfNamespaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNvmfNamespaceParams() *UpdateNvmfNamespaceParams {
	return &UpdateNvmfNamespaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNvmfNamespaceParamsWithTimeout creates a new UpdateNvmfNamespaceParams object
// with the ability to set a timeout on a request.
func NewUpdateNvmfNamespaceParamsWithTimeout(timeout time.Duration) *UpdateNvmfNamespaceParams {
	return &UpdateNvmfNamespaceParams{
		timeout: timeout,
	}
}

// NewUpdateNvmfNamespaceParamsWithContext creates a new UpdateNvmfNamespaceParams object
// with the ability to set a context for a request.
func NewUpdateNvmfNamespaceParamsWithContext(ctx context.Context) *UpdateNvmfNamespaceParams {
	return &UpdateNvmfNamespaceParams{
		Context: ctx,
	}
}

// NewUpdateNvmfNamespaceParamsWithHTTPClient creates a new UpdateNvmfNamespaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNvmfNamespaceParamsWithHTTPClient(client *http.Client) *UpdateNvmfNamespaceParams {
	return &UpdateNvmfNamespaceParams{
		HTTPClient: client,
	}
}

/* UpdateNvmfNamespaceParams contains all the parameters to send to the API endpoint
   for the update nvmf namespace operation.

   Typically these are written to a http.Request.
*/
type UpdateNvmfNamespaceParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.NvmfNamespaceUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update nvmf namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNvmfNamespaceParams) WithDefaults() *UpdateNvmfNamespaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update nvmf namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNvmfNamespaceParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateNvmfNamespaceParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) WithTimeout(timeout time.Duration) *UpdateNvmfNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) WithContext(ctx context.Context) *UpdateNvmfNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) WithHTTPClient(client *http.Client) *UpdateNvmfNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) WithContentLanguage(contentLanguage *string) *UpdateNvmfNamespaceParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) WithRequestBody(requestBody *models.NvmfNamespaceUpdationParams) *UpdateNvmfNamespaceParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update nvmf namespace params
func (o *UpdateNvmfNamespaceParams) SetRequestBody(requestBody *models.NvmfNamespaceUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNvmfNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
