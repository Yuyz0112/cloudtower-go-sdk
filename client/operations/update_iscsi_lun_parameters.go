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

// NewUpdateIscsiLunParams creates a new UpdateIscsiLunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIscsiLunParams() *UpdateIscsiLunParams {
	return &UpdateIscsiLunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIscsiLunParamsWithTimeout creates a new UpdateIscsiLunParams object
// with the ability to set a timeout on a request.
func NewUpdateIscsiLunParamsWithTimeout(timeout time.Duration) *UpdateIscsiLunParams {
	return &UpdateIscsiLunParams{
		timeout: timeout,
	}
}

// NewUpdateIscsiLunParamsWithContext creates a new UpdateIscsiLunParams object
// with the ability to set a context for a request.
func NewUpdateIscsiLunParamsWithContext(ctx context.Context) *UpdateIscsiLunParams {
	return &UpdateIscsiLunParams{
		Context: ctx,
	}
}

// NewUpdateIscsiLunParamsWithHTTPClient creates a new UpdateIscsiLunParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIscsiLunParamsWithHTTPClient(client *http.Client) *UpdateIscsiLunParams {
	return &UpdateIscsiLunParams{
		HTTPClient: client,
	}
}

/* UpdateIscsiLunParams contains all the parameters to send to the API endpoint
   for the update iscsi lun operation.

   Typically these are written to a http.Request.
*/
type UpdateIscsiLunParams struct {

	// RequestBody.
	RequestBody *models.IscsiLunUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update iscsi lun params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIscsiLunParams) WithDefaults() *UpdateIscsiLunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update iscsi lun params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIscsiLunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update iscsi lun params
func (o *UpdateIscsiLunParams) WithTimeout(timeout time.Duration) *UpdateIscsiLunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update iscsi lun params
func (o *UpdateIscsiLunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update iscsi lun params
func (o *UpdateIscsiLunParams) WithContext(ctx context.Context) *UpdateIscsiLunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update iscsi lun params
func (o *UpdateIscsiLunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update iscsi lun params
func (o *UpdateIscsiLunParams) WithHTTPClient(client *http.Client) *UpdateIscsiLunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update iscsi lun params
func (o *UpdateIscsiLunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update iscsi lun params
func (o *UpdateIscsiLunParams) WithRequestBody(requestBody *models.IscsiLunUpdationParams) *UpdateIscsiLunParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update iscsi lun params
func (o *UpdateIscsiLunParams) SetRequestBody(requestBody *models.IscsiLunUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIscsiLunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
