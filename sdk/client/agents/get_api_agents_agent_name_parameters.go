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

// NewGetAPIAgentsAgentNameParams creates a new GetAPIAgentsAgentNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIAgentsAgentNameParams() *GetAPIAgentsAgentNameParams {
	return &GetAPIAgentsAgentNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIAgentsAgentNameParamsWithTimeout creates a new GetAPIAgentsAgentNameParams object
// with the ability to set a timeout on a request.
func NewGetAPIAgentsAgentNameParamsWithTimeout(timeout time.Duration) *GetAPIAgentsAgentNameParams {
	return &GetAPIAgentsAgentNameParams{
		timeout: timeout,
	}
}

// NewGetAPIAgentsAgentNameParamsWithContext creates a new GetAPIAgentsAgentNameParams object
// with the ability to set a context for a request.
func NewGetAPIAgentsAgentNameParamsWithContext(ctx context.Context) *GetAPIAgentsAgentNameParams {
	return &GetAPIAgentsAgentNameParams{
		Context: ctx,
	}
}

// NewGetAPIAgentsAgentNameParamsWithHTTPClient creates a new GetAPIAgentsAgentNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIAgentsAgentNameParamsWithHTTPClient(client *http.Client) *GetAPIAgentsAgentNameParams {
	return &GetAPIAgentsAgentNameParams{
		HTTPClient: client,
	}
}

/* GetAPIAgentsAgentNameParams contains all the parameters to send to the API endpoint
   for the get API agents agent name operation.

   Typically these are written to a http.Request.
*/
type GetAPIAgentsAgentNameParams struct {

	// AgentName.
	AgentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API agents agent name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIAgentsAgentNameParams) WithDefaults() *GetAPIAgentsAgentNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API agents agent name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIAgentsAgentNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API agents agent name params
func (o *GetAPIAgentsAgentNameParams) WithTimeout(timeout time.Duration) *GetAPIAgentsAgentNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API agents agent name params
func (o *GetAPIAgentsAgentNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API agents agent name params
func (o *GetAPIAgentsAgentNameParams) WithContext(ctx context.Context) *GetAPIAgentsAgentNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API agents agent name params
func (o *GetAPIAgentsAgentNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API agents agent name params
func (o *GetAPIAgentsAgentNameParams) WithHTTPClient(client *http.Client) *GetAPIAgentsAgentNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API agents agent name params
func (o *GetAPIAgentsAgentNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentName adds the agentName to the get API agents agent name params
func (o *GetAPIAgentsAgentNameParams) WithAgentName(agentName string) *GetAPIAgentsAgentNameParams {
	o.SetAgentName(agentName)
	return o
}

// SetAgentName adds the agentName to the get API agents agent name params
func (o *GetAPIAgentsAgentNameParams) SetAgentName(agentName string) {
	o.AgentName = agentName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIAgentsAgentNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
