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

// NewGetNvmfNamespaceSnapshotsConnectionParams creates a new GetNvmfNamespaceSnapshotsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNvmfNamespaceSnapshotsConnectionParams() *GetNvmfNamespaceSnapshotsConnectionParams {
	return &GetNvmfNamespaceSnapshotsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNvmfNamespaceSnapshotsConnectionParamsWithTimeout creates a new GetNvmfNamespaceSnapshotsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetNvmfNamespaceSnapshotsConnectionParamsWithTimeout(timeout time.Duration) *GetNvmfNamespaceSnapshotsConnectionParams {
	return &GetNvmfNamespaceSnapshotsConnectionParams{
		timeout: timeout,
	}
}

// NewGetNvmfNamespaceSnapshotsConnectionParamsWithContext creates a new GetNvmfNamespaceSnapshotsConnectionParams object
// with the ability to set a context for a request.
func NewGetNvmfNamespaceSnapshotsConnectionParamsWithContext(ctx context.Context) *GetNvmfNamespaceSnapshotsConnectionParams {
	return &GetNvmfNamespaceSnapshotsConnectionParams{
		Context: ctx,
	}
}

// NewGetNvmfNamespaceSnapshotsConnectionParamsWithHTTPClient creates a new GetNvmfNamespaceSnapshotsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNvmfNamespaceSnapshotsConnectionParamsWithHTTPClient(client *http.Client) *GetNvmfNamespaceSnapshotsConnectionParams {
	return &GetNvmfNamespaceSnapshotsConnectionParams{
		HTTPClient: client,
	}
}

/* GetNvmfNamespaceSnapshotsConnectionParams contains all the parameters to send to the API endpoint
   for the get nvmf namespace snapshots connection operation.

   Typically these are written to a http.Request.
*/
type GetNvmfNamespaceSnapshotsConnectionParams struct {

	// RequestBody.
	RequestBody *models.GetNvmfNamespaceSnapshotsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nvmf namespace snapshots connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfNamespaceSnapshotsConnectionParams) WithDefaults() *GetNvmfNamespaceSnapshotsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nvmf namespace snapshots connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfNamespaceSnapshotsConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nvmf namespace snapshots connection params
func (o *GetNvmfNamespaceSnapshotsConnectionParams) WithTimeout(timeout time.Duration) *GetNvmfNamespaceSnapshotsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nvmf namespace snapshots connection params
func (o *GetNvmfNamespaceSnapshotsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nvmf namespace snapshots connection params
func (o *GetNvmfNamespaceSnapshotsConnectionParams) WithContext(ctx context.Context) *GetNvmfNamespaceSnapshotsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nvmf namespace snapshots connection params
func (o *GetNvmfNamespaceSnapshotsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nvmf namespace snapshots connection params
func (o *GetNvmfNamespaceSnapshotsConnectionParams) WithHTTPClient(client *http.Client) *GetNvmfNamespaceSnapshotsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nvmf namespace snapshots connection params
func (o *GetNvmfNamespaceSnapshotsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get nvmf namespace snapshots connection params
func (o *GetNvmfNamespaceSnapshotsConnectionParams) WithRequestBody(requestBody *models.GetNvmfNamespaceSnapshotsConnectionRequestBody) *GetNvmfNamespaceSnapshotsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nvmf namespace snapshots connection params
func (o *GetNvmfNamespaceSnapshotsConnectionParams) SetRequestBody(requestBody *models.GetNvmfNamespaceSnapshotsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNvmfNamespaceSnapshotsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
