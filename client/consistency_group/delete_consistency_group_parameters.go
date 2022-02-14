// Code generated by go-swagger; DO NOT EDIT.

package consistency_group

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

// NewDeleteConsistencyGroupParams creates a new DeleteConsistencyGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteConsistencyGroupParams() *DeleteConsistencyGroupParams {
	return &DeleteConsistencyGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConsistencyGroupParamsWithTimeout creates a new DeleteConsistencyGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteConsistencyGroupParamsWithTimeout(timeout time.Duration) *DeleteConsistencyGroupParams {
	return &DeleteConsistencyGroupParams{
		timeout: timeout,
	}
}

// NewDeleteConsistencyGroupParamsWithContext creates a new DeleteConsistencyGroupParams object
// with the ability to set a context for a request.
func NewDeleteConsistencyGroupParamsWithContext(ctx context.Context) *DeleteConsistencyGroupParams {
	return &DeleteConsistencyGroupParams{
		Context: ctx,
	}
}

// NewDeleteConsistencyGroupParamsWithHTTPClient creates a new DeleteConsistencyGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteConsistencyGroupParamsWithHTTPClient(client *http.Client) *DeleteConsistencyGroupParams {
	return &DeleteConsistencyGroupParams{
		HTTPClient: client,
	}
}

/* DeleteConsistencyGroupParams contains all the parameters to send to the API endpoint
   for the delete consistency group operation.

   Typically these are written to a http.Request.
*/
type DeleteConsistencyGroupParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ConsistencyGroupDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete consistency group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConsistencyGroupParams) WithDefaults() *DeleteConsistencyGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete consistency group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConsistencyGroupParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteConsistencyGroupParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete consistency group params
func (o *DeleteConsistencyGroupParams) WithTimeout(timeout time.Duration) *DeleteConsistencyGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete consistency group params
func (o *DeleteConsistencyGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete consistency group params
func (o *DeleteConsistencyGroupParams) WithContext(ctx context.Context) *DeleteConsistencyGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete consistency group params
func (o *DeleteConsistencyGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete consistency group params
func (o *DeleteConsistencyGroupParams) WithHTTPClient(client *http.Client) *DeleteConsistencyGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete consistency group params
func (o *DeleteConsistencyGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete consistency group params
func (o *DeleteConsistencyGroupParams) WithContentLanguage(contentLanguage *string) *DeleteConsistencyGroupParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete consistency group params
func (o *DeleteConsistencyGroupParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete consistency group params
func (o *DeleteConsistencyGroupParams) WithRequestBody(requestBody *models.ConsistencyGroupDeletionParams) *DeleteConsistencyGroupParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete consistency group params
func (o *DeleteConsistencyGroupParams) SetRequestBody(requestBody *models.ConsistencyGroupDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConsistencyGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
