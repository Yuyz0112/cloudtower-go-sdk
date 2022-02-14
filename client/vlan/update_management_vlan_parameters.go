// Code generated by go-swagger; DO NOT EDIT.

package vlan

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

// NewUpdateManagementVlanParams creates a new UpdateManagementVlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateManagementVlanParams() *UpdateManagementVlanParams {
	return &UpdateManagementVlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateManagementVlanParamsWithTimeout creates a new UpdateManagementVlanParams object
// with the ability to set a timeout on a request.
func NewUpdateManagementVlanParamsWithTimeout(timeout time.Duration) *UpdateManagementVlanParams {
	return &UpdateManagementVlanParams{
		timeout: timeout,
	}
}

// NewUpdateManagementVlanParamsWithContext creates a new UpdateManagementVlanParams object
// with the ability to set a context for a request.
func NewUpdateManagementVlanParamsWithContext(ctx context.Context) *UpdateManagementVlanParams {
	return &UpdateManagementVlanParams{
		Context: ctx,
	}
}

// NewUpdateManagementVlanParamsWithHTTPClient creates a new UpdateManagementVlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateManagementVlanParamsWithHTTPClient(client *http.Client) *UpdateManagementVlanParams {
	return &UpdateManagementVlanParams{
		HTTPClient: client,
	}
}

/* UpdateManagementVlanParams contains all the parameters to send to the API endpoint
   for the update management vlan operation.

   Typically these are written to a http.Request.
*/
type UpdateManagementVlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ManagementVlanUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update management vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateManagementVlanParams) WithDefaults() *UpdateManagementVlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update management vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateManagementVlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateManagementVlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update management vlan params
func (o *UpdateManagementVlanParams) WithTimeout(timeout time.Duration) *UpdateManagementVlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update management vlan params
func (o *UpdateManagementVlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update management vlan params
func (o *UpdateManagementVlanParams) WithContext(ctx context.Context) *UpdateManagementVlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update management vlan params
func (o *UpdateManagementVlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update management vlan params
func (o *UpdateManagementVlanParams) WithHTTPClient(client *http.Client) *UpdateManagementVlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update management vlan params
func (o *UpdateManagementVlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update management vlan params
func (o *UpdateManagementVlanParams) WithContentLanguage(contentLanguage *string) *UpdateManagementVlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update management vlan params
func (o *UpdateManagementVlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update management vlan params
func (o *UpdateManagementVlanParams) WithRequestBody(requestBody *models.ManagementVlanUpdationParams) *UpdateManagementVlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update management vlan params
func (o *UpdateManagementVlanParams) SetRequestBody(requestBody *models.ManagementVlanUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateManagementVlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
