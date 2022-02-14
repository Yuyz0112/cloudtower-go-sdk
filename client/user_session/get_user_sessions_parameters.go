// Code generated by go-swagger; DO NOT EDIT.

package user_session

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

// NewGetUserSessionsParams creates a new GetUserSessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserSessionsParams() *GetUserSessionsParams {
	return &GetUserSessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserSessionsParamsWithTimeout creates a new GetUserSessionsParams object
// with the ability to set a timeout on a request.
func NewGetUserSessionsParamsWithTimeout(timeout time.Duration) *GetUserSessionsParams {
	return &GetUserSessionsParams{
		timeout: timeout,
	}
}

// NewGetUserSessionsParamsWithContext creates a new GetUserSessionsParams object
// with the ability to set a context for a request.
func NewGetUserSessionsParamsWithContext(ctx context.Context) *GetUserSessionsParams {
	return &GetUserSessionsParams{
		Context: ctx,
	}
}

// NewGetUserSessionsParamsWithHTTPClient creates a new GetUserSessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserSessionsParamsWithHTTPClient(client *http.Client) *GetUserSessionsParams {
	return &GetUserSessionsParams{
		HTTPClient: client,
	}
}

/* GetUserSessionsParams contains all the parameters to send to the API endpoint
   for the get user sessions operation.

   Typically these are written to a http.Request.
*/
type GetUserSessionsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUserSessionsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserSessionsParams) WithDefaults() *GetUserSessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserSessionsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUserSessionsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get user sessions params
func (o *GetUserSessionsParams) WithTimeout(timeout time.Duration) *GetUserSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user sessions params
func (o *GetUserSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user sessions params
func (o *GetUserSessionsParams) WithContext(ctx context.Context) *GetUserSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user sessions params
func (o *GetUserSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user sessions params
func (o *GetUserSessionsParams) WithHTTPClient(client *http.Client) *GetUserSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user sessions params
func (o *GetUserSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get user sessions params
func (o *GetUserSessionsParams) WithContentLanguage(contentLanguage *string) *GetUserSessionsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get user sessions params
func (o *GetUserSessionsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get user sessions params
func (o *GetUserSessionsParams) WithRequestBody(requestBody *models.GetUserSessionsRequestBody) *GetUserSessionsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get user sessions params
func (o *GetUserSessionsParams) SetRequestBody(requestBody *models.GetUserSessionsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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