// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewPostAPIAuthLoginParams creates a new PostAPIAuthLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIAuthLoginParams() *PostAPIAuthLoginParams {
	return &PostAPIAuthLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIAuthLoginParamsWithTimeout creates a new PostAPIAuthLoginParams object
// with the ability to set a timeout on a request.
func NewPostAPIAuthLoginParamsWithTimeout(timeout time.Duration) *PostAPIAuthLoginParams {
	return &PostAPIAuthLoginParams{
		timeout: timeout,
	}
}

// NewPostAPIAuthLoginParamsWithContext creates a new PostAPIAuthLoginParams object
// with the ability to set a context for a request.
func NewPostAPIAuthLoginParamsWithContext(ctx context.Context) *PostAPIAuthLoginParams {
	return &PostAPIAuthLoginParams{
		Context: ctx,
	}
}

// NewPostAPIAuthLoginParamsWithHTTPClient creates a new PostAPIAuthLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIAuthLoginParamsWithHTTPClient(client *http.Client) *PostAPIAuthLoginParams {
	return &PostAPIAuthLoginParams{
		HTTPClient: client,
	}
}

/* PostAPIAuthLoginParams contains all the parameters to send to the API endpoint
   for the post API auth login operation.

   Typically these are written to a http.Request.
*/
type PostAPIAuthLoginParams struct {

	// Body.
	Body *models.CredentialsViewModel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API auth login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIAuthLoginParams) WithDefaults() *PostAPIAuthLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API auth login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIAuthLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API auth login params
func (o *PostAPIAuthLoginParams) WithTimeout(timeout time.Duration) *PostAPIAuthLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API auth login params
func (o *PostAPIAuthLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API auth login params
func (o *PostAPIAuthLoginParams) WithContext(ctx context.Context) *PostAPIAuthLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API auth login params
func (o *PostAPIAuthLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API auth login params
func (o *PostAPIAuthLoginParams) WithHTTPClient(client *http.Client) *PostAPIAuthLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API auth login params
func (o *PostAPIAuthLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API auth login params
func (o *PostAPIAuthLoginParams) WithBody(body *models.CredentialsViewModel) *PostAPIAuthLoginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API auth login params
func (o *PostAPIAuthLoginParams) SetBody(body *models.CredentialsViewModel) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIAuthLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
