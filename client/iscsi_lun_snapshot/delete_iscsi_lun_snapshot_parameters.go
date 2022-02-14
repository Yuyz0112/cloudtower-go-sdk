// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun_snapshot

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

// NewDeleteIscsiLunSnapshotParams creates a new DeleteIscsiLunSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIscsiLunSnapshotParams() *DeleteIscsiLunSnapshotParams {
	return &DeleteIscsiLunSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIscsiLunSnapshotParamsWithTimeout creates a new DeleteIscsiLunSnapshotParams object
// with the ability to set a timeout on a request.
func NewDeleteIscsiLunSnapshotParamsWithTimeout(timeout time.Duration) *DeleteIscsiLunSnapshotParams {
	return &DeleteIscsiLunSnapshotParams{
		timeout: timeout,
	}
}

// NewDeleteIscsiLunSnapshotParamsWithContext creates a new DeleteIscsiLunSnapshotParams object
// with the ability to set a context for a request.
func NewDeleteIscsiLunSnapshotParamsWithContext(ctx context.Context) *DeleteIscsiLunSnapshotParams {
	return &DeleteIscsiLunSnapshotParams{
		Context: ctx,
	}
}

// NewDeleteIscsiLunSnapshotParamsWithHTTPClient creates a new DeleteIscsiLunSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIscsiLunSnapshotParamsWithHTTPClient(client *http.Client) *DeleteIscsiLunSnapshotParams {
	return &DeleteIscsiLunSnapshotParams{
		HTTPClient: client,
	}
}

/* DeleteIscsiLunSnapshotParams contains all the parameters to send to the API endpoint
   for the delete iscsi lun snapshot operation.

   Typically these are written to a http.Request.
*/
type DeleteIscsiLunSnapshotParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.IscsiLunSnapshotDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete iscsi lun snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIscsiLunSnapshotParams) WithDefaults() *DeleteIscsiLunSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete iscsi lun snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIscsiLunSnapshotParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteIscsiLunSnapshotParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) WithTimeout(timeout time.Duration) *DeleteIscsiLunSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) WithContext(ctx context.Context) *DeleteIscsiLunSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) WithHTTPClient(client *http.Client) *DeleteIscsiLunSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) WithContentLanguage(contentLanguage *string) *DeleteIscsiLunSnapshotParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) WithRequestBody(requestBody *models.IscsiLunSnapshotDeletionParams) *DeleteIscsiLunSnapshotParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete iscsi lun snapshot params
func (o *DeleteIscsiLunSnapshotParams) SetRequestBody(requestBody *models.IscsiLunSnapshotDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIscsiLunSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
