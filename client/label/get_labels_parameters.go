// Code generated by go-swagger; DO NOT EDIT.

package label

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

// NewGetLabelsParams creates a new GetLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLabelsParams() *GetLabelsParams {
	return &GetLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelsParamsWithTimeout creates a new GetLabelsParams object
// with the ability to set a timeout on a request.
func NewGetLabelsParamsWithTimeout(timeout time.Duration) *GetLabelsParams {
	return &GetLabelsParams{
		timeout: timeout,
	}
}

// NewGetLabelsParamsWithContext creates a new GetLabelsParams object
// with the ability to set a context for a request.
func NewGetLabelsParamsWithContext(ctx context.Context) *GetLabelsParams {
	return &GetLabelsParams{
		Context: ctx,
	}
}

// NewGetLabelsParamsWithHTTPClient creates a new GetLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLabelsParamsWithHTTPClient(client *http.Client) *GetLabelsParams {
	return &GetLabelsParams{
		HTTPClient: client,
	}
}

/* GetLabelsParams contains all the parameters to send to the API endpoint
   for the get labels operation.

   Typically these are written to a http.Request.
*/
type GetLabelsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetLabelsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsParams) WithDefaults() *GetLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetLabelsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get labels params
func (o *GetLabelsParams) WithTimeout(timeout time.Duration) *GetLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get labels params
func (o *GetLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get labels params
func (o *GetLabelsParams) WithContext(ctx context.Context) *GetLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get labels params
func (o *GetLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get labels params
func (o *GetLabelsParams) WithHTTPClient(client *http.Client) *GetLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get labels params
func (o *GetLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get labels params
func (o *GetLabelsParams) WithContentLanguage(contentLanguage *string) *GetLabelsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get labels params
func (o *GetLabelsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get labels params
func (o *GetLabelsParams) WithRequestBody(requestBody *models.GetLabelsRequestBody) *GetLabelsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get labels params
func (o *GetLabelsParams) SetRequestBody(requestBody *models.GetLabelsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
