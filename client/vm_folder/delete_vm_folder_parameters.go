// Code generated by go-swagger; DO NOT EDIT.

package vm_folder

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

// NewDeleteVMFolderParams creates a new DeleteVMFolderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVMFolderParams() *DeleteVMFolderParams {
	return &DeleteVMFolderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVMFolderParamsWithTimeout creates a new DeleteVMFolderParams object
// with the ability to set a timeout on a request.
func NewDeleteVMFolderParamsWithTimeout(timeout time.Duration) *DeleteVMFolderParams {
	return &DeleteVMFolderParams{
		timeout: timeout,
	}
}

// NewDeleteVMFolderParamsWithContext creates a new DeleteVMFolderParams object
// with the ability to set a context for a request.
func NewDeleteVMFolderParamsWithContext(ctx context.Context) *DeleteVMFolderParams {
	return &DeleteVMFolderParams{
		Context: ctx,
	}
}

// NewDeleteVMFolderParamsWithHTTPClient creates a new DeleteVMFolderParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVMFolderParamsWithHTTPClient(client *http.Client) *DeleteVMFolderParams {
	return &DeleteVMFolderParams{
		HTTPClient: client,
	}
}

/* DeleteVMFolderParams contains all the parameters to send to the API endpoint
   for the delete Vm folder operation.

   Typically these are written to a http.Request.
*/
type DeleteVMFolderParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMFolderDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Vm folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMFolderParams) WithDefaults() *DeleteVMFolderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Vm folder params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMFolderParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteVMFolderParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete Vm folder params
func (o *DeleteVMFolderParams) WithTimeout(timeout time.Duration) *DeleteVMFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Vm folder params
func (o *DeleteVMFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Vm folder params
func (o *DeleteVMFolderParams) WithContext(ctx context.Context) *DeleteVMFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Vm folder params
func (o *DeleteVMFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Vm folder params
func (o *DeleteVMFolderParams) WithHTTPClient(client *http.Client) *DeleteVMFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Vm folder params
func (o *DeleteVMFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete Vm folder params
func (o *DeleteVMFolderParams) WithContentLanguage(contentLanguage *string) *DeleteVMFolderParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete Vm folder params
func (o *DeleteVMFolderParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete Vm folder params
func (o *DeleteVMFolderParams) WithRequestBody(requestBody *models.VMFolderDeletionParams) *DeleteVMFolderParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete Vm folder params
func (o *DeleteVMFolderParams) SetRequestBody(requestBody *models.VMFolderDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVMFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
