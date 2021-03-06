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

	"github.com/hiddengearz/reconness-cli/sdk/models"
)

// NewPutAPITargetsIDParams creates a new PutAPITargetsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAPITargetsIDParams() *PutAPITargetsIDParams {
	return &PutAPITargetsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPITargetsIDParamsWithTimeout creates a new PutAPITargetsIDParams object
// with the ability to set a timeout on a request.
func NewPutAPITargetsIDParamsWithTimeout(timeout time.Duration) *PutAPITargetsIDParams {
	return &PutAPITargetsIDParams{
		timeout: timeout,
	}
}

// NewPutAPITargetsIDParamsWithContext creates a new PutAPITargetsIDParams object
// with the ability to set a context for a request.
func NewPutAPITargetsIDParamsWithContext(ctx context.Context) *PutAPITargetsIDParams {
	return &PutAPITargetsIDParams{
		Context: ctx,
	}
}

// NewPutAPITargetsIDParamsWithHTTPClient creates a new PutAPITargetsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAPITargetsIDParamsWithHTTPClient(client *http.Client) *PutAPITargetsIDParams {
	return &PutAPITargetsIDParams{
		HTTPClient: client,
	}
}

/* PutAPITargetsIDParams contains all the parameters to send to the API endpoint
   for the put API targets ID operation.

   Typically these are written to a http.Request.
*/
type PutAPITargetsIDParams struct {

	// Body.
	Body *models.TargetDto

	// ID.
	//
	// Format: uuid
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put API targets ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPITargetsIDParams) WithDefaults() *PutAPITargetsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put API targets ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAPITargetsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put API targets ID params
func (o *PutAPITargetsIDParams) WithTimeout(timeout time.Duration) *PutAPITargetsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API targets ID params
func (o *PutAPITargetsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API targets ID params
func (o *PutAPITargetsIDParams) WithContext(ctx context.Context) *PutAPITargetsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API targets ID params
func (o *PutAPITargetsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API targets ID params
func (o *PutAPITargetsIDParams) WithHTTPClient(client *http.Client) *PutAPITargetsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API targets ID params
func (o *PutAPITargetsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put API targets ID params
func (o *PutAPITargetsIDParams) WithBody(body *models.TargetDto) *PutAPITargetsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put API targets ID params
func (o *PutAPITargetsIDParams) SetBody(body *models.TargetDto) {
	o.Body = body
}

// WithID adds the id to the put API targets ID params
func (o *PutAPITargetsIDParams) WithID(id strfmt.UUID) *PutAPITargetsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put API targets ID params
func (o *PutAPITargetsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPITargetsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
