// Code generated by go-swagger; DO NOT EDIT.

package cluster_upgrade_history

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

// NewGetClusterUpgradeHistoriesConnectionParams creates a new GetClusterUpgradeHistoriesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterUpgradeHistoriesConnectionParams() *GetClusterUpgradeHistoriesConnectionParams {
	return &GetClusterUpgradeHistoriesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterUpgradeHistoriesConnectionParamsWithTimeout creates a new GetClusterUpgradeHistoriesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetClusterUpgradeHistoriesConnectionParamsWithTimeout(timeout time.Duration) *GetClusterUpgradeHistoriesConnectionParams {
	return &GetClusterUpgradeHistoriesConnectionParams{
		timeout: timeout,
	}
}

// NewGetClusterUpgradeHistoriesConnectionParamsWithContext creates a new GetClusterUpgradeHistoriesConnectionParams object
// with the ability to set a context for a request.
func NewGetClusterUpgradeHistoriesConnectionParamsWithContext(ctx context.Context) *GetClusterUpgradeHistoriesConnectionParams {
	return &GetClusterUpgradeHistoriesConnectionParams{
		Context: ctx,
	}
}

// NewGetClusterUpgradeHistoriesConnectionParamsWithHTTPClient creates a new GetClusterUpgradeHistoriesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterUpgradeHistoriesConnectionParamsWithHTTPClient(client *http.Client) *GetClusterUpgradeHistoriesConnectionParams {
	return &GetClusterUpgradeHistoriesConnectionParams{
		HTTPClient: client,
	}
}

/* GetClusterUpgradeHistoriesConnectionParams contains all the parameters to send to the API endpoint
   for the get cluster upgrade histories connection operation.

   Typically these are written to a http.Request.
*/
type GetClusterUpgradeHistoriesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetClusterUpgradeHistoriesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster upgrade histories connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterUpgradeHistoriesConnectionParams) WithDefaults() *GetClusterUpgradeHistoriesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster upgrade histories connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterUpgradeHistoriesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetClusterUpgradeHistoriesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) WithTimeout(timeout time.Duration) *GetClusterUpgradeHistoriesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) WithContext(ctx context.Context) *GetClusterUpgradeHistoriesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) WithHTTPClient(client *http.Client) *GetClusterUpgradeHistoriesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) WithContentLanguage(contentLanguage *string) *GetClusterUpgradeHistoriesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) WithRequestBody(requestBody *models.GetClusterUpgradeHistoriesConnectionRequestBody) *GetClusterUpgradeHistoriesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get cluster upgrade histories connection params
func (o *GetClusterUpgradeHistoriesConnectionParams) SetRequestBody(requestBody *models.GetClusterUpgradeHistoriesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterUpgradeHistoriesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
