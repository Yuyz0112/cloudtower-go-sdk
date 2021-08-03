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

// NewGetElfImagesConnectionParams creates a new GetElfImagesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetElfImagesConnectionParams() *GetElfImagesConnectionParams {
	return &GetElfImagesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetElfImagesConnectionParamsWithTimeout creates a new GetElfImagesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetElfImagesConnectionParamsWithTimeout(timeout time.Duration) *GetElfImagesConnectionParams {
	return &GetElfImagesConnectionParams{
		timeout: timeout,
	}
}

// NewGetElfImagesConnectionParamsWithContext creates a new GetElfImagesConnectionParams object
// with the ability to set a context for a request.
func NewGetElfImagesConnectionParamsWithContext(ctx context.Context) *GetElfImagesConnectionParams {
	return &GetElfImagesConnectionParams{
		Context: ctx,
	}
}

// NewGetElfImagesConnectionParamsWithHTTPClient creates a new GetElfImagesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetElfImagesConnectionParamsWithHTTPClient(client *http.Client) *GetElfImagesConnectionParams {
	return &GetElfImagesConnectionParams{
		HTTPClient: client,
	}
}

/* GetElfImagesConnectionParams contains all the parameters to send to the API endpoint
   for the get elf images connection operation.

   Typically these are written to a http.Request.
*/
type GetElfImagesConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetElfImagesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get elf images connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetElfImagesConnectionParams) WithDefaults() *GetElfImagesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get elf images connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetElfImagesConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get elf images connection params
func (o *GetElfImagesConnectionParams) WithTimeout(timeout time.Duration) *GetElfImagesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get elf images connection params
func (o *GetElfImagesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get elf images connection params
func (o *GetElfImagesConnectionParams) WithContext(ctx context.Context) *GetElfImagesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get elf images connection params
func (o *GetElfImagesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get elf images connection params
func (o *GetElfImagesConnectionParams) WithHTTPClient(client *http.Client) *GetElfImagesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get elf images connection params
func (o *GetElfImagesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get elf images connection params
func (o *GetElfImagesConnectionParams) WithRequestBody(requestBody *models.GetElfImagesConnectionRequestBody) *GetElfImagesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get elf images connection params
func (o *GetElfImagesConnectionParams) SetRequestBody(requestBody *models.GetElfImagesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetElfImagesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
