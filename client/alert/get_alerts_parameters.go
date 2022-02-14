// Code generated by go-swagger; DO NOT EDIT.

package alert

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

// NewGetAlertsParams creates a new GetAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertsParams() *GetAlertsParams {
	return &GetAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertsParamsWithTimeout creates a new GetAlertsParams object
// with the ability to set a timeout on a request.
func NewGetAlertsParamsWithTimeout(timeout time.Duration) *GetAlertsParams {
	return &GetAlertsParams{
		timeout: timeout,
	}
}

// NewGetAlertsParamsWithContext creates a new GetAlertsParams object
// with the ability to set a context for a request.
func NewGetAlertsParamsWithContext(ctx context.Context) *GetAlertsParams {
	return &GetAlertsParams{
		Context: ctx,
	}
}

// NewGetAlertsParamsWithHTTPClient creates a new GetAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertsParamsWithHTTPClient(client *http.Client) *GetAlertsParams {
	return &GetAlertsParams{
		HTTPClient: client,
	}
}

/* GetAlertsParams contains all the parameters to send to the API endpoint
   for the get alerts operation.

   Typically these are written to a http.Request.
*/
type GetAlertsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetAlertsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertsParams) WithDefaults() *GetAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetAlertsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get alerts params
func (o *GetAlertsParams) WithTimeout(timeout time.Duration) *GetAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alerts params
func (o *GetAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alerts params
func (o *GetAlertsParams) WithContext(ctx context.Context) *GetAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alerts params
func (o *GetAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alerts params
func (o *GetAlertsParams) WithHTTPClient(client *http.Client) *GetAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alerts params
func (o *GetAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get alerts params
func (o *GetAlertsParams) WithContentLanguage(contentLanguage *string) *GetAlertsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get alerts params
func (o *GetAlertsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get alerts params
func (o *GetAlertsParams) WithRequestBody(requestBody *models.GetAlertsRequestBody) *GetAlertsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get alerts params
func (o *GetAlertsParams) SetRequestBody(requestBody *models.GetAlertsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
