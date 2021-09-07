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

// NewCreateVMPlacementGroupParams creates a new CreateVMPlacementGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVMPlacementGroupParams() *CreateVMPlacementGroupParams {
	return &CreateVMPlacementGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVMPlacementGroupParamsWithTimeout creates a new CreateVMPlacementGroupParams object
// with the ability to set a timeout on a request.
func NewCreateVMPlacementGroupParamsWithTimeout(timeout time.Duration) *CreateVMPlacementGroupParams {
	return &CreateVMPlacementGroupParams{
		timeout: timeout,
	}
}

// NewCreateVMPlacementGroupParamsWithContext creates a new CreateVMPlacementGroupParams object
// with the ability to set a context for a request.
func NewCreateVMPlacementGroupParamsWithContext(ctx context.Context) *CreateVMPlacementGroupParams {
	return &CreateVMPlacementGroupParams{
		Context: ctx,
	}
}

// NewCreateVMPlacementGroupParamsWithHTTPClient creates a new CreateVMPlacementGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVMPlacementGroupParamsWithHTTPClient(client *http.Client) *CreateVMPlacementGroupParams {
	return &CreateVMPlacementGroupParams{
		HTTPClient: client,
	}
}

/* CreateVMPlacementGroupParams contains all the parameters to send to the API endpoint
   for the create Vm placement group operation.

   Typically these are written to a http.Request.
*/
type CreateVMPlacementGroupParams struct {

	// RequestBody.
	RequestBody []*models.VMPlacementGroupCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create Vm placement group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMPlacementGroupParams) WithDefaults() *CreateVMPlacementGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create Vm placement group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMPlacementGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create Vm placement group params
func (o *CreateVMPlacementGroupParams) WithTimeout(timeout time.Duration) *CreateVMPlacementGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Vm placement group params
func (o *CreateVMPlacementGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Vm placement group params
func (o *CreateVMPlacementGroupParams) WithContext(ctx context.Context) *CreateVMPlacementGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Vm placement group params
func (o *CreateVMPlacementGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Vm placement group params
func (o *CreateVMPlacementGroupParams) WithHTTPClient(client *http.Client) *CreateVMPlacementGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Vm placement group params
func (o *CreateVMPlacementGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the create Vm placement group params
func (o *CreateVMPlacementGroupParams) WithRequestBody(requestBody []*models.VMPlacementGroupCreationParams) *CreateVMPlacementGroupParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create Vm placement group params
func (o *CreateVMPlacementGroupParams) SetRequestBody(requestBody []*models.VMPlacementGroupCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVMPlacementGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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