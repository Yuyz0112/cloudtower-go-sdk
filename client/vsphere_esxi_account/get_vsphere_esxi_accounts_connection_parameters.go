// Code generated by go-swagger; DO NOT EDIT.

package vsphere_esxi_account

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

// NewGetVsphereEsxiAccountsConnectionParams creates a new GetVsphereEsxiAccountsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVsphereEsxiAccountsConnectionParams() *GetVsphereEsxiAccountsConnectionParams {
	return &GetVsphereEsxiAccountsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVsphereEsxiAccountsConnectionParamsWithTimeout creates a new GetVsphereEsxiAccountsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVsphereEsxiAccountsConnectionParamsWithTimeout(timeout time.Duration) *GetVsphereEsxiAccountsConnectionParams {
	return &GetVsphereEsxiAccountsConnectionParams{
		timeout: timeout,
	}
}

// NewGetVsphereEsxiAccountsConnectionParamsWithContext creates a new GetVsphereEsxiAccountsConnectionParams object
// with the ability to set a context for a request.
func NewGetVsphereEsxiAccountsConnectionParamsWithContext(ctx context.Context) *GetVsphereEsxiAccountsConnectionParams {
	return &GetVsphereEsxiAccountsConnectionParams{
		Context: ctx,
	}
}

// NewGetVsphereEsxiAccountsConnectionParamsWithHTTPClient creates a new GetVsphereEsxiAccountsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVsphereEsxiAccountsConnectionParamsWithHTTPClient(client *http.Client) *GetVsphereEsxiAccountsConnectionParams {
	return &GetVsphereEsxiAccountsConnectionParams{
		HTTPClient: client,
	}
}

/* GetVsphereEsxiAccountsConnectionParams contains all the parameters to send to the API endpoint
   for the get vsphere esxi accounts connection operation.

   Typically these are written to a http.Request.
*/
type GetVsphereEsxiAccountsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVsphereEsxiAccountsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vsphere esxi accounts connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVsphereEsxiAccountsConnectionParams) WithDefaults() *GetVsphereEsxiAccountsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vsphere esxi accounts connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVsphereEsxiAccountsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVsphereEsxiAccountsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) WithTimeout(timeout time.Duration) *GetVsphereEsxiAccountsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) WithContext(ctx context.Context) *GetVsphereEsxiAccountsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) WithHTTPClient(client *http.Client) *GetVsphereEsxiAccountsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) WithContentLanguage(contentLanguage *string) *GetVsphereEsxiAccountsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) WithRequestBody(requestBody *models.GetVsphereEsxiAccountsConnectionRequestBody) *GetVsphereEsxiAccountsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get vsphere esxi accounts connection params
func (o *GetVsphereEsxiAccountsConnectionParams) SetRequestBody(requestBody *models.GetVsphereEsxiAccountsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVsphereEsxiAccountsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
