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
)

// NewDeleteAPIAgentsAgentNameParams creates a new DeleteAPIAgentsAgentNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIAgentsAgentNameParams() *DeleteAPIAgentsAgentNameParams {
	return &DeleteAPIAgentsAgentNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIAgentsAgentNameParamsWithTimeout creates a new DeleteAPIAgentsAgentNameParams object
// with the ability to set a timeout on a request.
func NewDeleteAPIAgentsAgentNameParamsWithTimeout(timeout time.Duration) *DeleteAPIAgentsAgentNameParams {
	return &DeleteAPIAgentsAgentNameParams{
		timeout: timeout,
	}
}

// NewDeleteAPIAgentsAgentNameParamsWithContext creates a new DeleteAPIAgentsAgentNameParams object
// with the ability to set a context for a request.
func NewDeleteAPIAgentsAgentNameParamsWithContext(ctx context.Context) *DeleteAPIAgentsAgentNameParams {
	return &DeleteAPIAgentsAgentNameParams{
		Context: ctx,
	}
}

// NewDeleteAPIAgentsAgentNameParamsWithHTTPClient creates a new DeleteAPIAgentsAgentNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIAgentsAgentNameParamsWithHTTPClient(client *http.Client) *DeleteAPIAgentsAgentNameParams {
	return &DeleteAPIAgentsAgentNameParams{
		HTTPClient: client,
	}
}

/* DeleteAPIAgentsAgentNameParams contains all the parameters to send to the API endpoint
   for the delete API agents agent name operation.

   Typically these are written to a http.Request.
*/
type DeleteAPIAgentsAgentNameParams struct {

	// AgentName.
	AgentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API agents agent name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIAgentsAgentNameParams) WithDefaults() *DeleteAPIAgentsAgentNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API agents agent name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIAgentsAgentNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API agents agent name params
func (o *DeleteAPIAgentsAgentNameParams) WithTimeout(timeout time.Duration) *DeleteAPIAgentsAgentNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API agents agent name params
func (o *DeleteAPIAgentsAgentNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API agents agent name params
func (o *DeleteAPIAgentsAgentNameParams) WithContext(ctx context.Context) *DeleteAPIAgentsAgentNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API agents agent name params
func (o *DeleteAPIAgentsAgentNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API agents agent name params
func (o *DeleteAPIAgentsAgentNameParams) WithHTTPClient(client *http.Client) *DeleteAPIAgentsAgentNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API agents agent name params
func (o *DeleteAPIAgentsAgentNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentName adds the agentName to the delete API agents agent name params
func (o *DeleteAPIAgentsAgentNameParams) WithAgentName(agentName string) *DeleteAPIAgentsAgentNameParams {
	o.SetAgentName(agentName)
	return o
}

// SetAgentName adds the agentName to the delete API agents agent name params
func (o *DeleteAPIAgentsAgentNameParams) SetAgentName(agentName string) {
	o.AgentName = agentName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIAgentsAgentNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentName
	if err := r.SetPathParam("agentName", o.AgentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}