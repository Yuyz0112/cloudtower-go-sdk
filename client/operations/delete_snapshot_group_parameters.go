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

// NewDeleteSnapshotGroupParams creates a new DeleteSnapshotGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSnapshotGroupParams() *DeleteSnapshotGroupParams {
	return &DeleteSnapshotGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnapshotGroupParamsWithTimeout creates a new DeleteSnapshotGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteSnapshotGroupParamsWithTimeout(timeout time.Duration) *DeleteSnapshotGroupParams {
	return &DeleteSnapshotGroupParams{
		timeout: timeout,
	}
}

// NewDeleteSnapshotGroupParamsWithContext creates a new DeleteSnapshotGroupParams object
// with the ability to set a context for a request.
func NewDeleteSnapshotGroupParamsWithContext(ctx context.Context) *DeleteSnapshotGroupParams {
	return &DeleteSnapshotGroupParams{
		Context: ctx,
	}
}

// NewDeleteSnapshotGroupParamsWithHTTPClient creates a new DeleteSnapshotGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSnapshotGroupParamsWithHTTPClient(client *http.Client) *DeleteSnapshotGroupParams {
	return &DeleteSnapshotGroupParams{
		HTTPClient: client,
	}
}

/* DeleteSnapshotGroupParams contains all the parameters to send to the API endpoint
   for the delete snapshot group operation.

   Typically these are written to a http.Request.
*/
type DeleteSnapshotGroupParams struct {

	// RequestBody.
	RequestBody *models.SnapshotGroupDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete snapshot group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSnapshotGroupParams) WithDefaults() *DeleteSnapshotGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete snapshot group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSnapshotGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete snapshot group params
func (o *DeleteSnapshotGroupParams) WithTimeout(timeout time.Duration) *DeleteSnapshotGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snapshot group params
func (o *DeleteSnapshotGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snapshot group params
func (o *DeleteSnapshotGroupParams) WithContext(ctx context.Context) *DeleteSnapshotGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snapshot group params
func (o *DeleteSnapshotGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snapshot group params
func (o *DeleteSnapshotGroupParams) WithHTTPClient(client *http.Client) *DeleteSnapshotGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snapshot group params
func (o *DeleteSnapshotGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete snapshot group params
func (o *DeleteSnapshotGroupParams) WithRequestBody(requestBody *models.SnapshotGroupDeletionParams) *DeleteSnapshotGroupParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete snapshot group params
func (o *DeleteSnapshotGroupParams) SetRequestBody(requestBody *models.SnapshotGroupDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnapshotGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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