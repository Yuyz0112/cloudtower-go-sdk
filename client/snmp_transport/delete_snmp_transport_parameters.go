// Code generated by go-swagger; DO NOT EDIT.

package snmp_transport

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

// NewDeleteSnmpTransportParams creates a new DeleteSnmpTransportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSnmpTransportParams() *DeleteSnmpTransportParams {
	return &DeleteSnmpTransportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnmpTransportParamsWithTimeout creates a new DeleteSnmpTransportParams object
// with the ability to set a timeout on a request.
func NewDeleteSnmpTransportParamsWithTimeout(timeout time.Duration) *DeleteSnmpTransportParams {
	return &DeleteSnmpTransportParams{
		timeout: timeout,
	}
}

// NewDeleteSnmpTransportParamsWithContext creates a new DeleteSnmpTransportParams object
// with the ability to set a context for a request.
func NewDeleteSnmpTransportParamsWithContext(ctx context.Context) *DeleteSnmpTransportParams {
	return &DeleteSnmpTransportParams{
		Context: ctx,
	}
}

// NewDeleteSnmpTransportParamsWithHTTPClient creates a new DeleteSnmpTransportParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSnmpTransportParamsWithHTTPClient(client *http.Client) *DeleteSnmpTransportParams {
	return &DeleteSnmpTransportParams{
		HTTPClient: client,
	}
}

/* DeleteSnmpTransportParams contains all the parameters to send to the API endpoint
   for the delete snmp transport operation.

   Typically these are written to a http.Request.
*/
type DeleteSnmpTransportParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.SnmpTransportDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete snmp transport params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSnmpTransportParams) WithDefaults() *DeleteSnmpTransportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete snmp transport params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSnmpTransportParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteSnmpTransportParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete snmp transport params
func (o *DeleteSnmpTransportParams) WithTimeout(timeout time.Duration) *DeleteSnmpTransportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snmp transport params
func (o *DeleteSnmpTransportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snmp transport params
func (o *DeleteSnmpTransportParams) WithContext(ctx context.Context) *DeleteSnmpTransportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snmp transport params
func (o *DeleteSnmpTransportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snmp transport params
func (o *DeleteSnmpTransportParams) WithHTTPClient(client *http.Client) *DeleteSnmpTransportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snmp transport params
func (o *DeleteSnmpTransportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete snmp transport params
func (o *DeleteSnmpTransportParams) WithContentLanguage(contentLanguage *string) *DeleteSnmpTransportParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete snmp transport params
func (o *DeleteSnmpTransportParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete snmp transport params
func (o *DeleteSnmpTransportParams) WithRequestBody(requestBody *models.SnmpTransportDeletionParams) *DeleteSnmpTransportParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete snmp transport params
func (o *DeleteSnmpTransportParams) SetRequestBody(requestBody *models.SnmpTransportDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnmpTransportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
