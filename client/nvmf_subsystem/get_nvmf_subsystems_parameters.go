// Code generated by go-swagger; DO NOT EDIT.

package nvmf_subsystem

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

// NewGetNvmfSubsystemsParams creates a new GetNvmfSubsystemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNvmfSubsystemsParams() *GetNvmfSubsystemsParams {
	return &GetNvmfSubsystemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNvmfSubsystemsParamsWithTimeout creates a new GetNvmfSubsystemsParams object
// with the ability to set a timeout on a request.
func NewGetNvmfSubsystemsParamsWithTimeout(timeout time.Duration) *GetNvmfSubsystemsParams {
	return &GetNvmfSubsystemsParams{
		timeout: timeout,
	}
}

// NewGetNvmfSubsystemsParamsWithContext creates a new GetNvmfSubsystemsParams object
// with the ability to set a context for a request.
func NewGetNvmfSubsystemsParamsWithContext(ctx context.Context) *GetNvmfSubsystemsParams {
	return &GetNvmfSubsystemsParams{
		Context: ctx,
	}
}

// NewGetNvmfSubsystemsParamsWithHTTPClient creates a new GetNvmfSubsystemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNvmfSubsystemsParamsWithHTTPClient(client *http.Client) *GetNvmfSubsystemsParams {
	return &GetNvmfSubsystemsParams{
		HTTPClient: client,
	}
}

/* GetNvmfSubsystemsParams contains all the parameters to send to the API endpoint
   for the get nvmf subsystems operation.

   Typically these are written to a http.Request.
*/
type GetNvmfSubsystemsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetNvmfSubsystemsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nvmf subsystems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfSubsystemsParams) WithDefaults() *GetNvmfSubsystemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nvmf subsystems params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfSubsystemsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetNvmfSubsystemsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) WithTimeout(timeout time.Duration) *GetNvmfSubsystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) WithContext(ctx context.Context) *GetNvmfSubsystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) WithHTTPClient(client *http.Client) *GetNvmfSubsystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) WithContentLanguage(contentLanguage *string) *GetNvmfSubsystemsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) WithRequestBody(requestBody *models.GetNvmfSubsystemsRequestBody) *GetNvmfSubsystemsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nvmf subsystems params
func (o *GetNvmfSubsystemsParams) SetRequestBody(requestBody *models.GetNvmfSubsystemsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNvmfSubsystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
