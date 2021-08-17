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

// NewDeleteNvmfNamespaceParams creates a new DeleteNvmfNamespaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNvmfNamespaceParams() *DeleteNvmfNamespaceParams {
	return &DeleteNvmfNamespaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNvmfNamespaceParamsWithTimeout creates a new DeleteNvmfNamespaceParams object
// with the ability to set a timeout on a request.
func NewDeleteNvmfNamespaceParamsWithTimeout(timeout time.Duration) *DeleteNvmfNamespaceParams {
	return &DeleteNvmfNamespaceParams{
		timeout: timeout,
	}
}

// NewDeleteNvmfNamespaceParamsWithContext creates a new DeleteNvmfNamespaceParams object
// with the ability to set a context for a request.
func NewDeleteNvmfNamespaceParamsWithContext(ctx context.Context) *DeleteNvmfNamespaceParams {
	return &DeleteNvmfNamespaceParams{
		Context: ctx,
	}
}

// NewDeleteNvmfNamespaceParamsWithHTTPClient creates a new DeleteNvmfNamespaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNvmfNamespaceParamsWithHTTPClient(client *http.Client) *DeleteNvmfNamespaceParams {
	return &DeleteNvmfNamespaceParams{
		HTTPClient: client,
	}
}

/* DeleteNvmfNamespaceParams contains all the parameters to send to the API endpoint
   for the delete nvmf namespace operation.

   Typically these are written to a http.Request.
*/
type DeleteNvmfNamespaceParams struct {

	// RequestBody.
	RequestBody *models.NvmfNamespaceDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete nvmf namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNvmfNamespaceParams) WithDefaults() *DeleteNvmfNamespaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete nvmf namespace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNvmfNamespaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete nvmf namespace params
func (o *DeleteNvmfNamespaceParams) WithTimeout(timeout time.Duration) *DeleteNvmfNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nvmf namespace params
func (o *DeleteNvmfNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nvmf namespace params
func (o *DeleteNvmfNamespaceParams) WithContext(ctx context.Context) *DeleteNvmfNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nvmf namespace params
func (o *DeleteNvmfNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nvmf namespace params
func (o *DeleteNvmfNamespaceParams) WithHTTPClient(client *http.Client) *DeleteNvmfNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nvmf namespace params
func (o *DeleteNvmfNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete nvmf namespace params
func (o *DeleteNvmfNamespaceParams) WithRequestBody(requestBody *models.NvmfNamespaceDeletionParams) *DeleteNvmfNamespaceParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete nvmf namespace params
func (o *DeleteNvmfNamespaceParams) SetRequestBody(requestBody *models.NvmfNamespaceDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNvmfNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
