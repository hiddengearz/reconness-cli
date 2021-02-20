// Code generated by go-swagger; DO NOT EDIT.

package references

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
)

// NewGetAPIReferencesParams creates a new GetAPIReferencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIReferencesParams() *GetAPIReferencesParams {
	return &GetAPIReferencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIReferencesParamsWithTimeout creates a new GetAPIReferencesParams object
// with the ability to set a timeout on a request.
func NewGetAPIReferencesParamsWithTimeout(timeout time.Duration) *GetAPIReferencesParams {
	return &GetAPIReferencesParams{
		timeout: timeout,
	}
}

// NewGetAPIReferencesParamsWithContext creates a new GetAPIReferencesParams object
// with the ability to set a context for a request.
func NewGetAPIReferencesParamsWithContext(ctx context.Context) *GetAPIReferencesParams {
	return &GetAPIReferencesParams{
		Context: ctx,
	}
}

// NewGetAPIReferencesParamsWithHTTPClient creates a new GetAPIReferencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIReferencesParamsWithHTTPClient(client *http.Client) *GetAPIReferencesParams {
	return &GetAPIReferencesParams{
		HTTPClient: client,
	}
}

/* GetAPIReferencesParams contains all the parameters to send to the API endpoint
   for the get API references operation.

   Typically these are written to a http.Request.
*/
type GetAPIReferencesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API references params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIReferencesParams) WithDefaults() *GetAPIReferencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API references params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIReferencesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API references params
func (o *GetAPIReferencesParams) WithTimeout(timeout time.Duration) *GetAPIReferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API references params
func (o *GetAPIReferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API references params
func (o *GetAPIReferencesParams) WithContext(ctx context.Context) *GetAPIReferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API references params
func (o *GetAPIReferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API references params
func (o *GetAPIReferencesParams) WithHTTPClient(client *http.Client) *GetAPIReferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API references params
func (o *GetAPIReferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIReferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
