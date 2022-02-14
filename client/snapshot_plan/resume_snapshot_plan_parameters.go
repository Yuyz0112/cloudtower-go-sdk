// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan

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

// NewResumeSnapshotPlanParams creates a new ResumeSnapshotPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResumeSnapshotPlanParams() *ResumeSnapshotPlanParams {
	return &ResumeSnapshotPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResumeSnapshotPlanParamsWithTimeout creates a new ResumeSnapshotPlanParams object
// with the ability to set a timeout on a request.
func NewResumeSnapshotPlanParamsWithTimeout(timeout time.Duration) *ResumeSnapshotPlanParams {
	return &ResumeSnapshotPlanParams{
		timeout: timeout,
	}
}

// NewResumeSnapshotPlanParamsWithContext creates a new ResumeSnapshotPlanParams object
// with the ability to set a context for a request.
func NewResumeSnapshotPlanParamsWithContext(ctx context.Context) *ResumeSnapshotPlanParams {
	return &ResumeSnapshotPlanParams{
		Context: ctx,
	}
}

// NewResumeSnapshotPlanParamsWithHTTPClient creates a new ResumeSnapshotPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewResumeSnapshotPlanParamsWithHTTPClient(client *http.Client) *ResumeSnapshotPlanParams {
	return &ResumeSnapshotPlanParams{
		HTTPClient: client,
	}
}

/* ResumeSnapshotPlanParams contains all the parameters to send to the API endpoint
   for the resume snapshot plan operation.

   Typically these are written to a http.Request.
*/
type ResumeSnapshotPlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.SnapshotPlanResumeParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resume snapshot plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeSnapshotPlanParams) WithDefaults() *ResumeSnapshotPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resume snapshot plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeSnapshotPlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := ResumeSnapshotPlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) WithTimeout(timeout time.Duration) *ResumeSnapshotPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) WithContext(ctx context.Context) *ResumeSnapshotPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) WithHTTPClient(client *http.Client) *ResumeSnapshotPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) WithContentLanguage(contentLanguage *string) *ResumeSnapshotPlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) WithRequestBody(requestBody *models.SnapshotPlanResumeParams) *ResumeSnapshotPlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the resume snapshot plan params
func (o *ResumeSnapshotPlanParams) SetRequestBody(requestBody *models.SnapshotPlanResumeParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *ResumeSnapshotPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
