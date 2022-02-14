// Code generated by go-swagger; DO NOT EDIT.

package witness_service

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

// NewGetWitnessServicesParams creates a new GetWitnessServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWitnessServicesParams() *GetWitnessServicesParams {
	return &GetWitnessServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWitnessServicesParamsWithTimeout creates a new GetWitnessServicesParams object
// with the ability to set a timeout on a request.
func NewGetWitnessServicesParamsWithTimeout(timeout time.Duration) *GetWitnessServicesParams {
	return &GetWitnessServicesParams{
		timeout: timeout,
	}
}

// NewGetWitnessServicesParamsWithContext creates a new GetWitnessServicesParams object
// with the ability to set a context for a request.
func NewGetWitnessServicesParamsWithContext(ctx context.Context) *GetWitnessServicesParams {
	return &GetWitnessServicesParams{
		Context: ctx,
	}
}

// NewGetWitnessServicesParamsWithHTTPClient creates a new GetWitnessServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWitnessServicesParamsWithHTTPClient(client *http.Client) *GetWitnessServicesParams {
	return &GetWitnessServicesParams{
		HTTPClient: client,
	}
}

/* GetWitnessServicesParams contains all the parameters to send to the API endpoint
   for the get witness services operation.

   Typically these are written to a http.Request.
*/
type GetWitnessServicesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetWitnessServicesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get witness services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWitnessServicesParams) WithDefaults() *GetWitnessServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get witness services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWitnessServicesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetWitnessServicesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get witness services params
func (o *GetWitnessServicesParams) WithTimeout(timeout time.Duration) *GetWitnessServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get witness services params
func (o *GetWitnessServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get witness services params
func (o *GetWitnessServicesParams) WithContext(ctx context.Context) *GetWitnessServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get witness services params
func (o *GetWitnessServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get witness services params
func (o *GetWitnessServicesParams) WithHTTPClient(client *http.Client) *GetWitnessServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get witness services params
func (o *GetWitnessServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get witness services params
func (o *GetWitnessServicesParams) WithContentLanguage(contentLanguage *string) *GetWitnessServicesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get witness services params
func (o *GetWitnessServicesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get witness services params
func (o *GetWitnessServicesParams) WithRequestBody(requestBody *models.GetWitnessServicesRequestBody) *GetWitnessServicesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get witness services params
func (o *GetWitnessServicesParams) SetRequestBody(requestBody *models.GetWitnessServicesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetWitnessServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
