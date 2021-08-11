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

// NewDeleteNfsExportParams creates a new DeleteNfsExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNfsExportParams() *DeleteNfsExportParams {
	return &DeleteNfsExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNfsExportParamsWithTimeout creates a new DeleteNfsExportParams object
// with the ability to set a timeout on a request.
func NewDeleteNfsExportParamsWithTimeout(timeout time.Duration) *DeleteNfsExportParams {
	return &DeleteNfsExportParams{
		timeout: timeout,
	}
}

// NewDeleteNfsExportParamsWithContext creates a new DeleteNfsExportParams object
// with the ability to set a context for a request.
func NewDeleteNfsExportParamsWithContext(ctx context.Context) *DeleteNfsExportParams {
	return &DeleteNfsExportParams{
		Context: ctx,
	}
}

// NewDeleteNfsExportParamsWithHTTPClient creates a new DeleteNfsExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNfsExportParamsWithHTTPClient(client *http.Client) *DeleteNfsExportParams {
	return &DeleteNfsExportParams{
		HTTPClient: client,
	}
}

/* DeleteNfsExportParams contains all the parameters to send to the API endpoint
   for the delete nfs export operation.

   Typically these are written to a http.Request.
*/
type DeleteNfsExportParams struct {

	// RequestBody.
	RequestBody []*models.NfsExportDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete nfs export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNfsExportParams) WithDefaults() *DeleteNfsExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete nfs export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNfsExportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete nfs export params
func (o *DeleteNfsExportParams) WithTimeout(timeout time.Duration) *DeleteNfsExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nfs export params
func (o *DeleteNfsExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nfs export params
func (o *DeleteNfsExportParams) WithContext(ctx context.Context) *DeleteNfsExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nfs export params
func (o *DeleteNfsExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nfs export params
func (o *DeleteNfsExportParams) WithHTTPClient(client *http.Client) *DeleteNfsExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nfs export params
func (o *DeleteNfsExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete nfs export params
func (o *DeleteNfsExportParams) WithRequestBody(requestBody []*models.NfsExportDeletionParams) *DeleteNfsExportParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete nfs export params
func (o *DeleteNfsExportParams) SetRequestBody(requestBody []*models.NfsExportDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNfsExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
