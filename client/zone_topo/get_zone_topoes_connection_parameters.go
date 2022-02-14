// Code generated by go-swagger; DO NOT EDIT.

package zone_topo

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

// NewGetZoneTopoesConnectionParams creates a new GetZoneTopoesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetZoneTopoesConnectionParams() *GetZoneTopoesConnectionParams {
	return &GetZoneTopoesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetZoneTopoesConnectionParamsWithTimeout creates a new GetZoneTopoesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetZoneTopoesConnectionParamsWithTimeout(timeout time.Duration) *GetZoneTopoesConnectionParams {
	return &GetZoneTopoesConnectionParams{
		timeout: timeout,
	}
}

// NewGetZoneTopoesConnectionParamsWithContext creates a new GetZoneTopoesConnectionParams object
// with the ability to set a context for a request.
func NewGetZoneTopoesConnectionParamsWithContext(ctx context.Context) *GetZoneTopoesConnectionParams {
	return &GetZoneTopoesConnectionParams{
		Context: ctx,
	}
}

// NewGetZoneTopoesConnectionParamsWithHTTPClient creates a new GetZoneTopoesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetZoneTopoesConnectionParamsWithHTTPClient(client *http.Client) *GetZoneTopoesConnectionParams {
	return &GetZoneTopoesConnectionParams{
		HTTPClient: client,
	}
}

/* GetZoneTopoesConnectionParams contains all the parameters to send to the API endpoint
   for the get zone topoes connection operation.

   Typically these are written to a http.Request.
*/
type GetZoneTopoesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetZoneTopoesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get zone topoes connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZoneTopoesConnectionParams) WithDefaults() *GetZoneTopoesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get zone topoes connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZoneTopoesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetZoneTopoesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) WithTimeout(timeout time.Duration) *GetZoneTopoesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) WithContext(ctx context.Context) *GetZoneTopoesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) WithHTTPClient(client *http.Client) *GetZoneTopoesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) WithContentLanguage(contentLanguage *string) *GetZoneTopoesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) WithRequestBody(requestBody *models.GetZoneTopoesConnectionRequestBody) *GetZoneTopoesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get zone topoes connection params
func (o *GetZoneTopoesConnectionParams) SetRequestBody(requestBody *models.GetZoneTopoesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetZoneTopoesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
