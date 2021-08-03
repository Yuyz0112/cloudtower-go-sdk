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

// NewGetClusterUpgradeHistoriesParams creates a new GetClusterUpgradeHistoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterUpgradeHistoriesParams() *GetClusterUpgradeHistoriesParams {
	return &GetClusterUpgradeHistoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterUpgradeHistoriesParamsWithTimeout creates a new GetClusterUpgradeHistoriesParams object
// with the ability to set a timeout on a request.
func NewGetClusterUpgradeHistoriesParamsWithTimeout(timeout time.Duration) *GetClusterUpgradeHistoriesParams {
	return &GetClusterUpgradeHistoriesParams{
		timeout: timeout,
	}
}

// NewGetClusterUpgradeHistoriesParamsWithContext creates a new GetClusterUpgradeHistoriesParams object
// with the ability to set a context for a request.
func NewGetClusterUpgradeHistoriesParamsWithContext(ctx context.Context) *GetClusterUpgradeHistoriesParams {
	return &GetClusterUpgradeHistoriesParams{
		Context: ctx,
	}
}

// NewGetClusterUpgradeHistoriesParamsWithHTTPClient creates a new GetClusterUpgradeHistoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterUpgradeHistoriesParamsWithHTTPClient(client *http.Client) *GetClusterUpgradeHistoriesParams {
	return &GetClusterUpgradeHistoriesParams{
		HTTPClient: client,
	}
}

/* GetClusterUpgradeHistoriesParams contains all the parameters to send to the API endpoint
   for the get cluster upgrade histories operation.

   Typically these are written to a http.Request.
*/
type GetClusterUpgradeHistoriesParams struct {

	// RequestBody.
	RequestBody *models.GetClusterUpgradeHistoriesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster upgrade histories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterUpgradeHistoriesParams) WithDefaults() *GetClusterUpgradeHistoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster upgrade histories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterUpgradeHistoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster upgrade histories params
func (o *GetClusterUpgradeHistoriesParams) WithTimeout(timeout time.Duration) *GetClusterUpgradeHistoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster upgrade histories params
func (o *GetClusterUpgradeHistoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster upgrade histories params
func (o *GetClusterUpgradeHistoriesParams) WithContext(ctx context.Context) *GetClusterUpgradeHistoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster upgrade histories params
func (o *GetClusterUpgradeHistoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster upgrade histories params
func (o *GetClusterUpgradeHistoriesParams) WithHTTPClient(client *http.Client) *GetClusterUpgradeHistoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster upgrade histories params
func (o *GetClusterUpgradeHistoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get cluster upgrade histories params
func (o *GetClusterUpgradeHistoriesParams) WithRequestBody(requestBody *models.GetClusterUpgradeHistoriesRequestBody) *GetClusterUpgradeHistoriesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get cluster upgrade histories params
func (o *GetClusterUpgradeHistoriesParams) SetRequestBody(requestBody *models.GetClusterUpgradeHistoriesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterUpgradeHistoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
