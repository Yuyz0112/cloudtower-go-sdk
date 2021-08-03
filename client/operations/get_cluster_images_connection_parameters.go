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

	"cloudtower-go-sdk/models"
)

// NewGetClusterImagesConnectionParams creates a new GetClusterImagesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterImagesConnectionParams() *GetClusterImagesConnectionParams {
	return &GetClusterImagesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterImagesConnectionParamsWithTimeout creates a new GetClusterImagesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetClusterImagesConnectionParamsWithTimeout(timeout time.Duration) *GetClusterImagesConnectionParams {
	return &GetClusterImagesConnectionParams{
		timeout: timeout,
	}
}

// NewGetClusterImagesConnectionParamsWithContext creates a new GetClusterImagesConnectionParams object
// with the ability to set a context for a request.
func NewGetClusterImagesConnectionParamsWithContext(ctx context.Context) *GetClusterImagesConnectionParams {
	return &GetClusterImagesConnectionParams{
		Context: ctx,
	}
}

// NewGetClusterImagesConnectionParamsWithHTTPClient creates a new GetClusterImagesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterImagesConnectionParamsWithHTTPClient(client *http.Client) *GetClusterImagesConnectionParams {
	return &GetClusterImagesConnectionParams{
		HTTPClient: client,
	}
}

/* GetClusterImagesConnectionParams contains all the parameters to send to the API endpoint
   for the get cluster images connection operation.

   Typically these are written to a http.Request.
*/
type GetClusterImagesConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetClusterImagesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster images connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterImagesConnectionParams) WithDefaults() *GetClusterImagesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster images connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterImagesConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster images connection params
func (o *GetClusterImagesConnectionParams) WithTimeout(timeout time.Duration) *GetClusterImagesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster images connection params
func (o *GetClusterImagesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster images connection params
func (o *GetClusterImagesConnectionParams) WithContext(ctx context.Context) *GetClusterImagesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster images connection params
func (o *GetClusterImagesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster images connection params
func (o *GetClusterImagesConnectionParams) WithHTTPClient(client *http.Client) *GetClusterImagesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster images connection params
func (o *GetClusterImagesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get cluster images connection params
func (o *GetClusterImagesConnectionParams) WithRequestBody(requestBody *models.GetClusterImagesConnectionRequestBody) *GetClusterImagesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get cluster images connection params
func (o *GetClusterImagesConnectionParams) SetRequestBody(requestBody *models.GetClusterImagesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterImagesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
