// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewUpdateGlobalRecycleBinSettingParams creates a new UpdateGlobalRecycleBinSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGlobalRecycleBinSettingParams() *UpdateGlobalRecycleBinSettingParams {
	return &UpdateGlobalRecycleBinSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGlobalRecycleBinSettingParamsWithTimeout creates a new UpdateGlobalRecycleBinSettingParams object
// with the ability to set a timeout on a request.
func NewUpdateGlobalRecycleBinSettingParamsWithTimeout(timeout time.Duration) *UpdateGlobalRecycleBinSettingParams {
	return &UpdateGlobalRecycleBinSettingParams{
		timeout: timeout,
	}
}

// NewUpdateGlobalRecycleBinSettingParamsWithContext creates a new UpdateGlobalRecycleBinSettingParams object
// with the ability to set a context for a request.
func NewUpdateGlobalRecycleBinSettingParamsWithContext(ctx context.Context) *UpdateGlobalRecycleBinSettingParams {
	return &UpdateGlobalRecycleBinSettingParams{
		Context: ctx,
	}
}

// NewUpdateGlobalRecycleBinSettingParamsWithHTTPClient creates a new UpdateGlobalRecycleBinSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGlobalRecycleBinSettingParamsWithHTTPClient(client *http.Client) *UpdateGlobalRecycleBinSettingParams {
	return &UpdateGlobalRecycleBinSettingParams{
		HTTPClient: client,
	}
}

/* UpdateGlobalRecycleBinSettingParams contains all the parameters to send to the API endpoint
   for the update global recycle bin setting operation.

   Typically these are written to a http.Request.
*/
type UpdateGlobalRecycleBinSettingParams struct {

	// RequestBody.
	RequestBody *models.GlobalRecycleBinUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update global recycle bin setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGlobalRecycleBinSettingParams) WithDefaults() *UpdateGlobalRecycleBinSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update global recycle bin setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGlobalRecycleBinSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update global recycle bin setting params
func (o *UpdateGlobalRecycleBinSettingParams) WithTimeout(timeout time.Duration) *UpdateGlobalRecycleBinSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update global recycle bin setting params
func (o *UpdateGlobalRecycleBinSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update global recycle bin setting params
func (o *UpdateGlobalRecycleBinSettingParams) WithContext(ctx context.Context) *UpdateGlobalRecycleBinSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update global recycle bin setting params
func (o *UpdateGlobalRecycleBinSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update global recycle bin setting params
func (o *UpdateGlobalRecycleBinSettingParams) WithHTTPClient(client *http.Client) *UpdateGlobalRecycleBinSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update global recycle bin setting params
func (o *UpdateGlobalRecycleBinSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update global recycle bin setting params
func (o *UpdateGlobalRecycleBinSettingParams) WithRequestBody(requestBody *models.GlobalRecycleBinUpdationParams) *UpdateGlobalRecycleBinSettingParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update global recycle bin setting params
func (o *UpdateGlobalRecycleBinSettingParams) SetRequestBody(requestBody *models.GlobalRecycleBinUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGlobalRecycleBinSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
