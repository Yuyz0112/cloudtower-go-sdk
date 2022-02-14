// Code generated by go-swagger; DO NOT EDIT.

package brick_topo

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

// NewCreateBrickTopoParams creates a new CreateBrickTopoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBrickTopoParams() *CreateBrickTopoParams {
	return &CreateBrickTopoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBrickTopoParamsWithTimeout creates a new CreateBrickTopoParams object
// with the ability to set a timeout on a request.
func NewCreateBrickTopoParamsWithTimeout(timeout time.Duration) *CreateBrickTopoParams {
	return &CreateBrickTopoParams{
		timeout: timeout,
	}
}

// NewCreateBrickTopoParamsWithContext creates a new CreateBrickTopoParams object
// with the ability to set a context for a request.
func NewCreateBrickTopoParamsWithContext(ctx context.Context) *CreateBrickTopoParams {
	return &CreateBrickTopoParams{
		Context: ctx,
	}
}

// NewCreateBrickTopoParamsWithHTTPClient creates a new CreateBrickTopoParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBrickTopoParamsWithHTTPClient(client *http.Client) *CreateBrickTopoParams {
	return &CreateBrickTopoParams{
		HTTPClient: client,
	}
}

/* CreateBrickTopoParams contains all the parameters to send to the API endpoint
   for the create brick topo operation.

   Typically these are written to a http.Request.
*/
type CreateBrickTopoParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.BrickTopoCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create brick topo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBrickTopoParams) WithDefaults() *CreateBrickTopoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create brick topo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBrickTopoParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateBrickTopoParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create brick topo params
func (o *CreateBrickTopoParams) WithTimeout(timeout time.Duration) *CreateBrickTopoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create brick topo params
func (o *CreateBrickTopoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create brick topo params
func (o *CreateBrickTopoParams) WithContext(ctx context.Context) *CreateBrickTopoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create brick topo params
func (o *CreateBrickTopoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create brick topo params
func (o *CreateBrickTopoParams) WithHTTPClient(client *http.Client) *CreateBrickTopoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create brick topo params
func (o *CreateBrickTopoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create brick topo params
func (o *CreateBrickTopoParams) WithContentLanguage(contentLanguage *string) *CreateBrickTopoParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create brick topo params
func (o *CreateBrickTopoParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create brick topo params
func (o *CreateBrickTopoParams) WithRequestBody(requestBody []*models.BrickTopoCreationParams) *CreateBrickTopoParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create brick topo params
func (o *CreateBrickTopoParams) SetRequestBody(requestBody []*models.BrickTopoCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBrickTopoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
