// Code generated by go-swagger; DO NOT EDIT.

package vm_placement_group

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

// NewDeleteVMPlacementGroupParams creates a new DeleteVMPlacementGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVMPlacementGroupParams() *DeleteVMPlacementGroupParams {
	return &DeleteVMPlacementGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVMPlacementGroupParamsWithTimeout creates a new DeleteVMPlacementGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteVMPlacementGroupParamsWithTimeout(timeout time.Duration) *DeleteVMPlacementGroupParams {
	return &DeleteVMPlacementGroupParams{
		timeout: timeout,
	}
}

// NewDeleteVMPlacementGroupParamsWithContext creates a new DeleteVMPlacementGroupParams object
// with the ability to set a context for a request.
func NewDeleteVMPlacementGroupParamsWithContext(ctx context.Context) *DeleteVMPlacementGroupParams {
	return &DeleteVMPlacementGroupParams{
		Context: ctx,
	}
}

// NewDeleteVMPlacementGroupParamsWithHTTPClient creates a new DeleteVMPlacementGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVMPlacementGroupParamsWithHTTPClient(client *http.Client) *DeleteVMPlacementGroupParams {
	return &DeleteVMPlacementGroupParams{
		HTTPClient: client,
	}
}

/* DeleteVMPlacementGroupParams contains all the parameters to send to the API endpoint
   for the delete Vm placement group operation.

   Typically these are written to a http.Request.
*/
type DeleteVMPlacementGroupParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMPlacementGroupDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Vm placement group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMPlacementGroupParams) WithDefaults() *DeleteVMPlacementGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Vm placement group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMPlacementGroupParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteVMPlacementGroupParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) WithTimeout(timeout time.Duration) *DeleteVMPlacementGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) WithContext(ctx context.Context) *DeleteVMPlacementGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) WithHTTPClient(client *http.Client) *DeleteVMPlacementGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) WithContentLanguage(contentLanguage *string) *DeleteVMPlacementGroupParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) WithRequestBody(requestBody *models.VMPlacementGroupDeletionParams) *DeleteVMPlacementGroupParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete Vm placement group params
func (o *DeleteVMPlacementGroupParams) SetRequestBody(requestBody *models.VMPlacementGroupDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVMPlacementGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
