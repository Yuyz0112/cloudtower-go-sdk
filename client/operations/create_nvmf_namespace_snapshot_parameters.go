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

// NewCreateNvmfNamespaceSnapshotParams creates a new CreateNvmfNamespaceSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNvmfNamespaceSnapshotParams() *CreateNvmfNamespaceSnapshotParams {
	return &CreateNvmfNamespaceSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNvmfNamespaceSnapshotParamsWithTimeout creates a new CreateNvmfNamespaceSnapshotParams object
// with the ability to set a timeout on a request.
func NewCreateNvmfNamespaceSnapshotParamsWithTimeout(timeout time.Duration) *CreateNvmfNamespaceSnapshotParams {
	return &CreateNvmfNamespaceSnapshotParams{
		timeout: timeout,
	}
}

// NewCreateNvmfNamespaceSnapshotParamsWithContext creates a new CreateNvmfNamespaceSnapshotParams object
// with the ability to set a context for a request.
func NewCreateNvmfNamespaceSnapshotParamsWithContext(ctx context.Context) *CreateNvmfNamespaceSnapshotParams {
	return &CreateNvmfNamespaceSnapshotParams{
		Context: ctx,
	}
}

// NewCreateNvmfNamespaceSnapshotParamsWithHTTPClient creates a new CreateNvmfNamespaceSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNvmfNamespaceSnapshotParamsWithHTTPClient(client *http.Client) *CreateNvmfNamespaceSnapshotParams {
	return &CreateNvmfNamespaceSnapshotParams{
		HTTPClient: client,
	}
}

/* CreateNvmfNamespaceSnapshotParams contains all the parameters to send to the API endpoint
   for the create nvmf namespace snapshot operation.

   Typically these are written to a http.Request.
*/
type CreateNvmfNamespaceSnapshotParams struct {

	// RequestBody.
	RequestBody []*models.NvmfNamespaceSnapshotCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create nvmf namespace snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNvmfNamespaceSnapshotParams) WithDefaults() *CreateNvmfNamespaceSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create nvmf namespace snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNvmfNamespaceSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create nvmf namespace snapshot params
func (o *CreateNvmfNamespaceSnapshotParams) WithTimeout(timeout time.Duration) *CreateNvmfNamespaceSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create nvmf namespace snapshot params
func (o *CreateNvmfNamespaceSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create nvmf namespace snapshot params
func (o *CreateNvmfNamespaceSnapshotParams) WithContext(ctx context.Context) *CreateNvmfNamespaceSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create nvmf namespace snapshot params
func (o *CreateNvmfNamespaceSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create nvmf namespace snapshot params
func (o *CreateNvmfNamespaceSnapshotParams) WithHTTPClient(client *http.Client) *CreateNvmfNamespaceSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create nvmf namespace snapshot params
func (o *CreateNvmfNamespaceSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the create nvmf namespace snapshot params
func (o *CreateNvmfNamespaceSnapshotParams) WithRequestBody(requestBody []*models.NvmfNamespaceSnapshotCreationParams) *CreateNvmfNamespaceSnapshotParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create nvmf namespace snapshot params
func (o *CreateNvmfNamespaceSnapshotParams) SetRequestBody(requestBody []*models.NvmfNamespaceSnapshotCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNvmfNamespaceSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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