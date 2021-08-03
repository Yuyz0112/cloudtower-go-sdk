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

// NewMoveVmsToRecycleBinParams creates a new MoveVmsToRecycleBinParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMoveVmsToRecycleBinParams() *MoveVmsToRecycleBinParams {
	return &MoveVmsToRecycleBinParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMoveVmsToRecycleBinParamsWithTimeout creates a new MoveVmsToRecycleBinParams object
// with the ability to set a timeout on a request.
func NewMoveVmsToRecycleBinParamsWithTimeout(timeout time.Duration) *MoveVmsToRecycleBinParams {
	return &MoveVmsToRecycleBinParams{
		timeout: timeout,
	}
}

// NewMoveVmsToRecycleBinParamsWithContext creates a new MoveVmsToRecycleBinParams object
// with the ability to set a context for a request.
func NewMoveVmsToRecycleBinParamsWithContext(ctx context.Context) *MoveVmsToRecycleBinParams {
	return &MoveVmsToRecycleBinParams{
		Context: ctx,
	}
}

// NewMoveVmsToRecycleBinParamsWithHTTPClient creates a new MoveVmsToRecycleBinParams object
// with the ability to set a custom HTTPClient for a request.
func NewMoveVmsToRecycleBinParamsWithHTTPClient(client *http.Client) *MoveVmsToRecycleBinParams {
	return &MoveVmsToRecycleBinParams{
		HTTPClient: client,
	}
}

/* MoveVmsToRecycleBinParams contains all the parameters to send to the API endpoint
   for the move vms to recycle bin operation.

   Typically these are written to a http.Request.
*/
type MoveVmsToRecycleBinParams struct {

	// RequestBody.
	RequestBody []*models.VMOperateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the move vms to recycle bin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MoveVmsToRecycleBinParams) WithDefaults() *MoveVmsToRecycleBinParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the move vms to recycle bin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MoveVmsToRecycleBinParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the move vms to recycle bin params
func (o *MoveVmsToRecycleBinParams) WithTimeout(timeout time.Duration) *MoveVmsToRecycleBinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move vms to recycle bin params
func (o *MoveVmsToRecycleBinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move vms to recycle bin params
func (o *MoveVmsToRecycleBinParams) WithContext(ctx context.Context) *MoveVmsToRecycleBinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move vms to recycle bin params
func (o *MoveVmsToRecycleBinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move vms to recycle bin params
func (o *MoveVmsToRecycleBinParams) WithHTTPClient(client *http.Client) *MoveVmsToRecycleBinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move vms to recycle bin params
func (o *MoveVmsToRecycleBinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the move vms to recycle bin params
func (o *MoveVmsToRecycleBinParams) WithRequestBody(requestBody []*models.VMOperateParams) *MoveVmsToRecycleBinParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the move vms to recycle bin params
func (o *MoveVmsToRecycleBinParams) SetRequestBody(requestBody []*models.VMOperateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *MoveVmsToRecycleBinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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