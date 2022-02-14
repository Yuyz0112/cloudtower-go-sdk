// Code generated by go-swagger; DO NOT EDIT.

package vm_snapshot

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

// NewGetVMSnapshotsConnectionParams creates a new GetVMSnapshotsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMSnapshotsConnectionParams() *GetVMSnapshotsConnectionParams {
	return &GetVMSnapshotsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMSnapshotsConnectionParamsWithTimeout creates a new GetVMSnapshotsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVMSnapshotsConnectionParamsWithTimeout(timeout time.Duration) *GetVMSnapshotsConnectionParams {
	return &GetVMSnapshotsConnectionParams{
		timeout: timeout,
	}
}

// NewGetVMSnapshotsConnectionParamsWithContext creates a new GetVMSnapshotsConnectionParams object
// with the ability to set a context for a request.
func NewGetVMSnapshotsConnectionParamsWithContext(ctx context.Context) *GetVMSnapshotsConnectionParams {
	return &GetVMSnapshotsConnectionParams{
		Context: ctx,
	}
}

// NewGetVMSnapshotsConnectionParamsWithHTTPClient creates a new GetVMSnapshotsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMSnapshotsConnectionParamsWithHTTPClient(client *http.Client) *GetVMSnapshotsConnectionParams {
	return &GetVMSnapshotsConnectionParams{
		HTTPClient: client,
	}
}

/* GetVMSnapshotsConnectionParams contains all the parameters to send to the API endpoint
   for the get Vm snapshots connection operation.

   Typically these are written to a http.Request.
*/
type GetVMSnapshotsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMSnapshotsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm snapshots connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMSnapshotsConnectionParams) WithDefaults() *GetVMSnapshotsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm snapshots connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMSnapshotsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMSnapshotsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) WithTimeout(timeout time.Duration) *GetVMSnapshotsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) WithContext(ctx context.Context) *GetVMSnapshotsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) WithHTTPClient(client *http.Client) *GetVMSnapshotsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) WithContentLanguage(contentLanguage *string) *GetVMSnapshotsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) WithRequestBody(requestBody *models.GetVMSnapshotsConnectionRequestBody) *GetVMSnapshotsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm snapshots connection params
func (o *GetVMSnapshotsConnectionParams) SetRequestBody(requestBody *models.GetVMSnapshotsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMSnapshotsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
