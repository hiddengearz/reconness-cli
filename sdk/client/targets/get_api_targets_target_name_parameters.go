// Code generated by go-swagger; DO NOT EDIT.

package targets

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

// NewGetAPITargetsTargetNameParams creates a new GetAPITargetsTargetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPITargetsTargetNameParams() *GetAPITargetsTargetNameParams {
	return &GetAPITargetsTargetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPITargetsTargetNameParamsWithTimeout creates a new GetAPITargetsTargetNameParams object
// with the ability to set a timeout on a request.
func NewGetAPITargetsTargetNameParamsWithTimeout(timeout time.Duration) *GetAPITargetsTargetNameParams {
	return &GetAPITargetsTargetNameParams{
		timeout: timeout,
	}
}

// NewGetAPITargetsTargetNameParamsWithContext creates a new GetAPITargetsTargetNameParams object
// with the ability to set a context for a request.
func NewGetAPITargetsTargetNameParamsWithContext(ctx context.Context) *GetAPITargetsTargetNameParams {
	return &GetAPITargetsTargetNameParams{
		Context: ctx,
	}
}

// NewGetAPITargetsTargetNameParamsWithHTTPClient creates a new GetAPITargetsTargetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPITargetsTargetNameParamsWithHTTPClient(client *http.Client) *GetAPITargetsTargetNameParams {
	return &GetAPITargetsTargetNameParams{
		HTTPClient: client,
	}
}

/* GetAPITargetsTargetNameParams contains all the parameters to send to the API endpoint
   for the get API targets target name operation.

   Typically these are written to a http.Request.
*/
type GetAPITargetsTargetNameParams struct {

	// TargetName.
	TargetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API targets target name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITargetsTargetNameParams) WithDefaults() *GetAPITargetsTargetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API targets target name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITargetsTargetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API targets target name params
func (o *GetAPITargetsTargetNameParams) WithTimeout(timeout time.Duration) *GetAPITargetsTargetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API targets target name params
func (o *GetAPITargetsTargetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API targets target name params
func (o *GetAPITargetsTargetNameParams) WithContext(ctx context.Context) *GetAPITargetsTargetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API targets target name params
func (o *GetAPITargetsTargetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API targets target name params
func (o *GetAPITargetsTargetNameParams) WithHTTPClient(client *http.Client) *GetAPITargetsTargetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API targets target name params
func (o *GetAPITargetsTargetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTargetName adds the targetName to the get API targets target name params
func (o *GetAPITargetsTargetNameParams) WithTargetName(targetName string) *GetAPITargetsTargetNameParams {
	o.SetTargetName(targetName)
	return o
}

// SetTargetName adds the targetName to the get API targets target name params
func (o *GetAPITargetsTargetNameParams) SetTargetName(targetName string) {
	o.TargetName = targetName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPITargetsTargetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param targetName
	if err := r.SetPathParam("targetName", o.TargetName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
