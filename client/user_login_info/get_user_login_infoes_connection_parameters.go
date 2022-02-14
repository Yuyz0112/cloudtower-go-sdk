// Code generated by go-swagger; DO NOT EDIT.

package user_login_info

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

// NewGetUserLoginInfoesConnectionParams creates a new GetUserLoginInfoesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserLoginInfoesConnectionParams() *GetUserLoginInfoesConnectionParams {
	return &GetUserLoginInfoesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserLoginInfoesConnectionParamsWithTimeout creates a new GetUserLoginInfoesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetUserLoginInfoesConnectionParamsWithTimeout(timeout time.Duration) *GetUserLoginInfoesConnectionParams {
	return &GetUserLoginInfoesConnectionParams{
		timeout: timeout,
	}
}

// NewGetUserLoginInfoesConnectionParamsWithContext creates a new GetUserLoginInfoesConnectionParams object
// with the ability to set a context for a request.
func NewGetUserLoginInfoesConnectionParamsWithContext(ctx context.Context) *GetUserLoginInfoesConnectionParams {
	return &GetUserLoginInfoesConnectionParams{
		Context: ctx,
	}
}

// NewGetUserLoginInfoesConnectionParamsWithHTTPClient creates a new GetUserLoginInfoesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserLoginInfoesConnectionParamsWithHTTPClient(client *http.Client) *GetUserLoginInfoesConnectionParams {
	return &GetUserLoginInfoesConnectionParams{
		HTTPClient: client,
	}
}

/* GetUserLoginInfoesConnectionParams contains all the parameters to send to the API endpoint
   for the get user login infoes connection operation.

   Typically these are written to a http.Request.
*/
type GetUserLoginInfoesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUserLoginInfoesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user login infoes connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserLoginInfoesConnectionParams) WithDefaults() *GetUserLoginInfoesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user login infoes connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserLoginInfoesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUserLoginInfoesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) WithTimeout(timeout time.Duration) *GetUserLoginInfoesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) WithContext(ctx context.Context) *GetUserLoginInfoesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) WithHTTPClient(client *http.Client) *GetUserLoginInfoesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) WithContentLanguage(contentLanguage *string) *GetUserLoginInfoesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) WithRequestBody(requestBody *models.GetUserLoginInfoesConnectionRequestBody) *GetUserLoginInfoesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get user login infoes connection params
func (o *GetUserLoginInfoesConnectionParams) SetRequestBody(requestBody *models.GetUserLoginInfoesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserLoginInfoesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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