// Code generated by go-swagger; DO NOT EDIT.

package report_template

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

// NewCreateReportTemplateParams creates a new CreateReportTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateReportTemplateParams() *CreateReportTemplateParams {
	return &CreateReportTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateReportTemplateParamsWithTimeout creates a new CreateReportTemplateParams object
// with the ability to set a timeout on a request.
func NewCreateReportTemplateParamsWithTimeout(timeout time.Duration) *CreateReportTemplateParams {
	return &CreateReportTemplateParams{
		timeout: timeout,
	}
}

// NewCreateReportTemplateParamsWithContext creates a new CreateReportTemplateParams object
// with the ability to set a context for a request.
func NewCreateReportTemplateParamsWithContext(ctx context.Context) *CreateReportTemplateParams {
	return &CreateReportTemplateParams{
		Context: ctx,
	}
}

// NewCreateReportTemplateParamsWithHTTPClient creates a new CreateReportTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateReportTemplateParamsWithHTTPClient(client *http.Client) *CreateReportTemplateParams {
	return &CreateReportTemplateParams{
		HTTPClient: client,
	}
}

/* CreateReportTemplateParams contains all the parameters to send to the API endpoint
   for the create report template operation.

   Typically these are written to a http.Request.
*/
type CreateReportTemplateParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.ReportTemplateCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create report template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateReportTemplateParams) WithDefaults() *CreateReportTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create report template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateReportTemplateParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateReportTemplateParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create report template params
func (o *CreateReportTemplateParams) WithTimeout(timeout time.Duration) *CreateReportTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create report template params
func (o *CreateReportTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create report template params
func (o *CreateReportTemplateParams) WithContext(ctx context.Context) *CreateReportTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create report template params
func (o *CreateReportTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create report template params
func (o *CreateReportTemplateParams) WithHTTPClient(client *http.Client) *CreateReportTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create report template params
func (o *CreateReportTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create report template params
func (o *CreateReportTemplateParams) WithContentLanguage(contentLanguage *string) *CreateReportTemplateParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create report template params
func (o *CreateReportTemplateParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create report template params
func (o *CreateReportTemplateParams) WithRequestBody(requestBody []*models.ReportTemplateCreationParams) *CreateReportTemplateParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create report template params
func (o *CreateReportTemplateParams) SetRequestBody(requestBody []*models.ReportTemplateCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateReportTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
