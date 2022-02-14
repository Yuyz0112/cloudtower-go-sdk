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

// NewCreateSnapshotPlanParams creates a new CreateSnapshotPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSnapshotPlanParams() *CreateSnapshotPlanParams {
	return &CreateSnapshotPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSnapshotPlanParamsWithTimeout creates a new CreateSnapshotPlanParams object
// with the ability to set a timeout on a request.
func NewCreateSnapshotPlanParamsWithTimeout(timeout time.Duration) *CreateSnapshotPlanParams {
	return &CreateSnapshotPlanParams{
		timeout: timeout,
	}
}

// NewCreateSnapshotPlanParamsWithContext creates a new CreateSnapshotPlanParams object
// with the ability to set a context for a request.
func NewCreateSnapshotPlanParamsWithContext(ctx context.Context) *CreateSnapshotPlanParams {
	return &CreateSnapshotPlanParams{
		Context: ctx,
	}
}

// NewCreateSnapshotPlanParamsWithHTTPClient creates a new CreateSnapshotPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSnapshotPlanParamsWithHTTPClient(client *http.Client) *CreateSnapshotPlanParams {
	return &CreateSnapshotPlanParams{
		HTTPClient: client,
	}
}

/* CreateSnapshotPlanParams contains all the parameters to send to the API endpoint
   for the create snapshot plan operation.

   Typically these are written to a http.Request.
*/
type CreateSnapshotPlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.SnapshotPlanCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create snapshot plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSnapshotPlanParams) WithDefaults() *CreateSnapshotPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create snapshot plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSnapshotPlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateSnapshotPlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create snapshot plan params
func (o *CreateSnapshotPlanParams) WithTimeout(timeout time.Duration) *CreateSnapshotPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create snapshot plan params
func (o *CreateSnapshotPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create snapshot plan params
func (o *CreateSnapshotPlanParams) WithContext(ctx context.Context) *CreateSnapshotPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create snapshot plan params
func (o *CreateSnapshotPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create snapshot plan params
func (o *CreateSnapshotPlanParams) WithHTTPClient(client *http.Client) *CreateSnapshotPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create snapshot plan params
func (o *CreateSnapshotPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create snapshot plan params
func (o *CreateSnapshotPlanParams) WithContentLanguage(contentLanguage *string) *CreateSnapshotPlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create snapshot plan params
func (o *CreateSnapshotPlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create snapshot plan params
func (o *CreateSnapshotPlanParams) WithRequestBody(requestBody []*models.SnapshotPlanCreationParams) *CreateSnapshotPlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create snapshot plan params
func (o *CreateSnapshotPlanParams) SetRequestBody(requestBody []*models.SnapshotPlanCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSnapshotPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
