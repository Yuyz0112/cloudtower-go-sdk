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

// NewDeleteReportTemplateParams creates a new DeleteReportTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteReportTemplateParams() *DeleteReportTemplateParams {
	return &DeleteReportTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReportTemplateParamsWithTimeout creates a new DeleteReportTemplateParams object
// with the ability to set a timeout on a request.
func NewDeleteReportTemplateParamsWithTimeout(timeout time.Duration) *DeleteReportTemplateParams {
	return &DeleteReportTemplateParams{
		timeout: timeout,
	}
}

// NewDeleteReportTemplateParamsWithContext creates a new DeleteReportTemplateParams object
// with the ability to set a context for a request.
func NewDeleteReportTemplateParamsWithContext(ctx context.Context) *DeleteReportTemplateParams {
	return &DeleteReportTemplateParams{
		Context: ctx,
	}
}

// NewDeleteReportTemplateParamsWithHTTPClient creates a new DeleteReportTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteReportTemplateParamsWithHTTPClient(client *http.Client) *DeleteReportTemplateParams {
	return &DeleteReportTemplateParams{
		HTTPClient: client,
	}
}

/* DeleteReportTemplateParams contains all the parameters to send to the API endpoint
   for the delete report template operation.

   Typically these are written to a http.Request.
*/
type DeleteReportTemplateParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ReportTemplateDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete report template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteReportTemplateParams) WithDefaults() *DeleteReportTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete report template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteReportTemplateParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteReportTemplateParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete report template params
func (o *DeleteReportTemplateParams) WithTimeout(timeout time.Duration) *DeleteReportTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete report template params
func (o *DeleteReportTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete report template params
func (o *DeleteReportTemplateParams) WithContext(ctx context.Context) *DeleteReportTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete report template params
func (o *DeleteReportTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete report template params
func (o *DeleteReportTemplateParams) WithHTTPClient(client *http.Client) *DeleteReportTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete report template params
func (o *DeleteReportTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete report template params
func (o *DeleteReportTemplateParams) WithContentLanguage(contentLanguage *string) *DeleteReportTemplateParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete report template params
func (o *DeleteReportTemplateParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete report template params
func (o *DeleteReportTemplateParams) WithRequestBody(requestBody *models.ReportTemplateDeletionParams) *DeleteReportTemplateParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete report template params
func (o *DeleteReportTemplateParams) SetRequestBody(requestBody *models.ReportTemplateDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReportTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
