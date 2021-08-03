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

// NewGetNvmfNamespaceSnapshotsParams creates a new GetNvmfNamespaceSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNvmfNamespaceSnapshotsParams() *GetNvmfNamespaceSnapshotsParams {
	return &GetNvmfNamespaceSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNvmfNamespaceSnapshotsParamsWithTimeout creates a new GetNvmfNamespaceSnapshotsParams object
// with the ability to set a timeout on a request.
func NewGetNvmfNamespaceSnapshotsParamsWithTimeout(timeout time.Duration) *GetNvmfNamespaceSnapshotsParams {
	return &GetNvmfNamespaceSnapshotsParams{
		timeout: timeout,
	}
}

// NewGetNvmfNamespaceSnapshotsParamsWithContext creates a new GetNvmfNamespaceSnapshotsParams object
// with the ability to set a context for a request.
func NewGetNvmfNamespaceSnapshotsParamsWithContext(ctx context.Context) *GetNvmfNamespaceSnapshotsParams {
	return &GetNvmfNamespaceSnapshotsParams{
		Context: ctx,
	}
}

// NewGetNvmfNamespaceSnapshotsParamsWithHTTPClient creates a new GetNvmfNamespaceSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNvmfNamespaceSnapshotsParamsWithHTTPClient(client *http.Client) *GetNvmfNamespaceSnapshotsParams {
	return &GetNvmfNamespaceSnapshotsParams{
		HTTPClient: client,
	}
}

/* GetNvmfNamespaceSnapshotsParams contains all the parameters to send to the API endpoint
   for the get nvmf namespace snapshots operation.

   Typically these are written to a http.Request.
*/
type GetNvmfNamespaceSnapshotsParams struct {

	// RequestBody.
	RequestBody *models.GetNvmfNamespaceSnapshotsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nvmf namespace snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfNamespaceSnapshotsParams) WithDefaults() *GetNvmfNamespaceSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nvmf namespace snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfNamespaceSnapshotsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nvmf namespace snapshots params
func (o *GetNvmfNamespaceSnapshotsParams) WithTimeout(timeout time.Duration) *GetNvmfNamespaceSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nvmf namespace snapshots params
func (o *GetNvmfNamespaceSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nvmf namespace snapshots params
func (o *GetNvmfNamespaceSnapshotsParams) WithContext(ctx context.Context) *GetNvmfNamespaceSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nvmf namespace snapshots params
func (o *GetNvmfNamespaceSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nvmf namespace snapshots params
func (o *GetNvmfNamespaceSnapshotsParams) WithHTTPClient(client *http.Client) *GetNvmfNamespaceSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nvmf namespace snapshots params
func (o *GetNvmfNamespaceSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get nvmf namespace snapshots params
func (o *GetNvmfNamespaceSnapshotsParams) WithRequestBody(requestBody *models.GetNvmfNamespaceSnapshotsRequestBody) *GetNvmfNamespaceSnapshotsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nvmf namespace snapshots params
func (o *GetNvmfNamespaceSnapshotsParams) SetRequestBody(requestBody *models.GetNvmfNamespaceSnapshotsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNvmfNamespaceSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
