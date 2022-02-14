// Code generated by go-swagger; DO NOT EDIT.

package backup_plan_execution

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

// NewGetBackupPlanExecutionsConnectionParams creates a new GetBackupPlanExecutionsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBackupPlanExecutionsConnectionParams() *GetBackupPlanExecutionsConnectionParams {
	return &GetBackupPlanExecutionsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupPlanExecutionsConnectionParamsWithTimeout creates a new GetBackupPlanExecutionsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetBackupPlanExecutionsConnectionParamsWithTimeout(timeout time.Duration) *GetBackupPlanExecutionsConnectionParams {
	return &GetBackupPlanExecutionsConnectionParams{
		timeout: timeout,
	}
}

// NewGetBackupPlanExecutionsConnectionParamsWithContext creates a new GetBackupPlanExecutionsConnectionParams object
// with the ability to set a context for a request.
func NewGetBackupPlanExecutionsConnectionParamsWithContext(ctx context.Context) *GetBackupPlanExecutionsConnectionParams {
	return &GetBackupPlanExecutionsConnectionParams{
		Context: ctx,
	}
}

// NewGetBackupPlanExecutionsConnectionParamsWithHTTPClient creates a new GetBackupPlanExecutionsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBackupPlanExecutionsConnectionParamsWithHTTPClient(client *http.Client) *GetBackupPlanExecutionsConnectionParams {
	return &GetBackupPlanExecutionsConnectionParams{
		HTTPClient: client,
	}
}

/* GetBackupPlanExecutionsConnectionParams contains all the parameters to send to the API endpoint
   for the get backup plan executions connection operation.

   Typically these are written to a http.Request.
*/
type GetBackupPlanExecutionsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetBackupPlanExecutionsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get backup plan executions connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupPlanExecutionsConnectionParams) WithDefaults() *GetBackupPlanExecutionsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get backup plan executions connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupPlanExecutionsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetBackupPlanExecutionsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) WithTimeout(timeout time.Duration) *GetBackupPlanExecutionsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) WithContext(ctx context.Context) *GetBackupPlanExecutionsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) WithHTTPClient(client *http.Client) *GetBackupPlanExecutionsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) WithContentLanguage(contentLanguage *string) *GetBackupPlanExecutionsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) WithRequestBody(requestBody *models.GetBackupPlanExecutionsConnectionRequestBody) *GetBackupPlanExecutionsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get backup plan executions connection params
func (o *GetBackupPlanExecutionsConnectionParams) SetRequestBody(requestBody *models.GetBackupPlanExecutionsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupPlanExecutionsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
