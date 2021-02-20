// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewPostAPIAgentsRunParams creates a new PostAPIAgentsRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIAgentsRunParams() *PostAPIAgentsRunParams {
	return &PostAPIAgentsRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIAgentsRunParamsWithTimeout creates a new PostAPIAgentsRunParams object
// with the ability to set a timeout on a request.
func NewPostAPIAgentsRunParamsWithTimeout(timeout time.Duration) *PostAPIAgentsRunParams {
	return &PostAPIAgentsRunParams{
		timeout: timeout,
	}
}

// NewPostAPIAgentsRunParamsWithContext creates a new PostAPIAgentsRunParams object
// with the ability to set a context for a request.
func NewPostAPIAgentsRunParamsWithContext(ctx context.Context) *PostAPIAgentsRunParams {
	return &PostAPIAgentsRunParams{
		Context: ctx,
	}
}

// NewPostAPIAgentsRunParamsWithHTTPClient creates a new PostAPIAgentsRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIAgentsRunParamsWithHTTPClient(client *http.Client) *PostAPIAgentsRunParams {
	return &PostAPIAgentsRunParams{
		HTTPClient: client,
	}
}

/* PostAPIAgentsRunParams contains all the parameters to send to the API endpoint
   for the post API agents run operation.

   Typically these are written to a http.Request.
*/
type PostAPIAgentsRunParams struct {

	// Body.
	Body *models.AgentRunnerDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API agents run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIAgentsRunParams) WithDefaults() *PostAPIAgentsRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API agents run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIAgentsRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API agents run params
func (o *PostAPIAgentsRunParams) WithTimeout(timeout time.Duration) *PostAPIAgentsRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API agents run params
func (o *PostAPIAgentsRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API agents run params
func (o *PostAPIAgentsRunParams) WithContext(ctx context.Context) *PostAPIAgentsRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API agents run params
func (o *PostAPIAgentsRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API agents run params
func (o *PostAPIAgentsRunParams) WithHTTPClient(client *http.Client) *PostAPIAgentsRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API agents run params
func (o *PostAPIAgentsRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API agents run params
func (o *PostAPIAgentsRunParams) WithBody(body *models.AgentRunnerDto) *PostAPIAgentsRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API agents run params
func (o *PostAPIAgentsRunParams) SetBody(body *models.AgentRunnerDto) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIAgentsRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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