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

// NewUpdateVlanParams creates a new UpdateVlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVlanParams() *UpdateVlanParams {
	return &UpdateVlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVlanParamsWithTimeout creates a new UpdateVlanParams object
// with the ability to set a timeout on a request.
func NewUpdateVlanParamsWithTimeout(timeout time.Duration) *UpdateVlanParams {
	return &UpdateVlanParams{
		timeout: timeout,
	}
}

// NewUpdateVlanParamsWithContext creates a new UpdateVlanParams object
// with the ability to set a context for a request.
func NewUpdateVlanParamsWithContext(ctx context.Context) *UpdateVlanParams {
	return &UpdateVlanParams{
		Context: ctx,
	}
}

// NewUpdateVlanParamsWithHTTPClient creates a new UpdateVlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVlanParamsWithHTTPClient(client *http.Client) *UpdateVlanParams {
	return &UpdateVlanParams{
		HTTPClient: client,
	}
}

/* UpdateVlanParams contains all the parameters to send to the API endpoint
   for the update vlan operation.

   Typically these are written to a http.Request.
*/
type UpdateVlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMVlanUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVlanParams) WithDefaults() *UpdateVlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update vlan params
func (o *UpdateVlanParams) WithTimeout(timeout time.Duration) *UpdateVlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update vlan params
func (o *UpdateVlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update vlan params
func (o *UpdateVlanParams) WithContext(ctx context.Context) *UpdateVlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update vlan params
func (o *UpdateVlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update vlan params
func (o *UpdateVlanParams) WithHTTPClient(client *http.Client) *UpdateVlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update vlan params
func (o *UpdateVlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update vlan params
func (o *UpdateVlanParams) WithContentLanguage(contentLanguage *string) *UpdateVlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update vlan params
func (o *UpdateVlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update vlan params
func (o *UpdateVlanParams) WithRequestBody(requestBody *models.VMVlanUpdationParams) *UpdateVlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update vlan params
func (o *UpdateVlanParams) SetRequestBody(requestBody *models.VMVlanUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
