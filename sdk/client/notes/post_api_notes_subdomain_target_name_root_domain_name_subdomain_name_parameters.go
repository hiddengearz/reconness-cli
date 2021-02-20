// Code generated by go-swagger; DO NOT EDIT.

package notes

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

// NewPostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams creates a new PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams() *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	return &PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParamsWithTimeout creates a new PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams object
// with the ability to set a timeout on a request.
func NewPostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParamsWithTimeout(timeout time.Duration) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	return &PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams{
		timeout: timeout,
	}
}

// NewPostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParamsWithContext creates a new PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams object
// with the ability to set a context for a request.
func NewPostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParamsWithContext(ctx context.Context) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	return &PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams{
		Context: ctx,
	}
}

// NewPostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParamsWithHTTPClient creates a new PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParamsWithHTTPClient(client *http.Client) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	return &PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams{
		HTTPClient: client,
	}
}

/* PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams contains all the parameters to send to the API endpoint
   for the post API notes subdomain target name root domain name subdomain name operation.

   Typically these are written to a http.Request.
*/
type PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams struct {

	// Body.
	Body *models.NoteDto

	// RootDomainName.
	RootDomainName string

	// SubdomainName.
	SubdomainName string

	// TargetName.
	TargetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API notes subdomain target name root domain name subdomain name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WithDefaults() *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API notes subdomain target name root domain name subdomain name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WithTimeout(timeout time.Duration) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WithContext(ctx context.Context) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WithHTTPClient(client *http.Client) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WithBody(body *models.NoteDto) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) SetBody(body *models.NoteDto) {
	o.Body = body
}

// WithRootDomainName adds the rootDomainName to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WithRootDomainName(rootDomainName string) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	o.SetRootDomainName(rootDomainName)
	return o
}

// SetRootDomainName adds the rootDomainName to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) SetRootDomainName(rootDomainName string) {
	o.RootDomainName = rootDomainName
}

// WithSubdomainName adds the subdomainName to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WithSubdomainName(subdomainName string) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	o.SetSubdomainName(subdomainName)
	return o
}

// SetSubdomainName adds the subdomainName to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) SetSubdomainName(subdomainName string) {
	o.SubdomainName = subdomainName
}

// WithTargetName adds the targetName to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WithTargetName(targetName string) *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams {
	o.SetTargetName(targetName)
	return o
}

// SetTargetName adds the targetName to the post API notes subdomain target name root domain name subdomain name params
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) SetTargetName(targetName string) {
	o.TargetName = targetName
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPINotesSubdomainTargetNameRootDomainNameSubdomainNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param rootDomainName
	if err := r.SetPathParam("rootDomainName", o.RootDomainName); err != nil {
		return err
	}

	// path param subdomainName
	if err := r.SetPathParam("subdomainName", o.SubdomainName); err != nil {
		return err
	}

	// path param targetName
	if err := r.SetPathParam("targetName", o.TargetName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
