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

// NewDeleteAPITargetsTargetNameParams creates a new DeleteAPITargetsTargetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPITargetsTargetNameParams() *DeleteAPITargetsTargetNameParams {
	return &DeleteAPITargetsTargetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPITargetsTargetNameParamsWithTimeout creates a new DeleteAPITargetsTargetNameParams object
// with the ability to set a timeout on a request.
func NewDeleteAPITargetsTargetNameParamsWithTimeout(timeout time.Duration) *DeleteAPITargetsTargetNameParams {
	return &DeleteAPITargetsTargetNameParams{
		timeout: timeout,
	}
}

// NewDeleteAPITargetsTargetNameParamsWithContext creates a new DeleteAPITargetsTargetNameParams object
// with the ability to set a context for a request.
func NewDeleteAPITargetsTargetNameParamsWithContext(ctx context.Context) *DeleteAPITargetsTargetNameParams {
	return &DeleteAPITargetsTargetNameParams{
		Context: ctx,
	}
}

// NewDeleteAPITargetsTargetNameParamsWithHTTPClient creates a new DeleteAPITargetsTargetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPITargetsTargetNameParamsWithHTTPClient(client *http.Client) *DeleteAPITargetsTargetNameParams {
	return &DeleteAPITargetsTargetNameParams{
		HTTPClient: client,
	}
}

/* DeleteAPITargetsTargetNameParams contains all the parameters to send to the API endpoint
   for the delete API targets target name operation.

   Typically these are written to a http.Request.
*/
type DeleteAPITargetsTargetNameParams struct {

	// TargetName.
	TargetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API targets target name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITargetsTargetNameParams) WithDefaults() *DeleteAPITargetsTargetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API targets target name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITargetsTargetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API targets target name params
func (o *DeleteAPITargetsTargetNameParams) WithTimeout(timeout time.Duration) *DeleteAPITargetsTargetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API targets target name params
func (o *DeleteAPITargetsTargetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API targets target name params
func (o *DeleteAPITargetsTargetNameParams) WithContext(ctx context.Context) *DeleteAPITargetsTargetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API targets target name params
func (o *DeleteAPITargetsTargetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API targets target name params
func (o *DeleteAPITargetsTargetNameParams) WithHTTPClient(client *http.Client) *DeleteAPITargetsTargetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API targets target name params
func (o *DeleteAPITargetsTargetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTargetName adds the targetName to the delete API targets target name params
func (o *DeleteAPITargetsTargetNameParams) WithTargetName(targetName string) *DeleteAPITargetsTargetNameParams {
	o.SetTargetName(targetName)
	return o
}

// SetTargetName adds the targetName to the delete API targets target name params
func (o *DeleteAPITargetsTargetNameParams) SetTargetName(targetName string) {
	o.TargetName = targetName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPITargetsTargetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
