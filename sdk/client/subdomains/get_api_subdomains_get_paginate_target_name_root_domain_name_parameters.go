// Code generated by go-swagger; DO NOT EDIT.

package subdomains

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
	"github.com/go-openapi/swag"
)

// NewGetAPISubdomainsGetPaginateTargetNameRootDomainNameParams creates a new GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPISubdomainsGetPaginateTargetNameRootDomainNameParams() *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	return &GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPISubdomainsGetPaginateTargetNameRootDomainNameParamsWithTimeout creates a new GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams object
// with the ability to set a timeout on a request.
func NewGetAPISubdomainsGetPaginateTargetNameRootDomainNameParamsWithTimeout(timeout time.Duration) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	return &GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams{
		timeout: timeout,
	}
}

// NewGetAPISubdomainsGetPaginateTargetNameRootDomainNameParamsWithContext creates a new GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams object
// with the ability to set a context for a request.
func NewGetAPISubdomainsGetPaginateTargetNameRootDomainNameParamsWithContext(ctx context.Context) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	return &GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams{
		Context: ctx,
	}
}

// NewGetAPISubdomainsGetPaginateTargetNameRootDomainNameParamsWithHTTPClient creates a new GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPISubdomainsGetPaginateTargetNameRootDomainNameParamsWithHTTPClient(client *http.Client) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	return &GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams{
		HTTPClient: client,
	}
}

/* GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams contains all the parameters to send to the API endpoint
   for the get API subdomains get paginate target name root domain name operation.

   Typically these are written to a http.Request.
*/
type GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams struct {

	// Ascending.
	//
	// Format: int32
	Ascending *int32

	// ByColumn.
	//
	// Format: int32
	ByColumn *int32

	// Limit.
	//
	// Format: int32
	Limit *int32

	// Page.
	//
	// Format: int32
	Page *int32

	// Query.
	Query *string

	// RootDomainName.
	RootDomainName string

	// TargetName.
	TargetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API subdomains get paginate target name root domain name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithDefaults() *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API subdomains get paginate target name root domain name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithTimeout(timeout time.Duration) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithContext(ctx context.Context) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithHTTPClient(client *http.Client) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAscending adds the ascending to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithAscending(ascending *int32) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetAscending(ascending)
	return o
}

// SetAscending adds the ascending to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetAscending(ascending *int32) {
	o.Ascending = ascending
}

// WithByColumn adds the byColumn to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithByColumn(byColumn *int32) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetByColumn(byColumn)
	return o
}

// SetByColumn adds the byColumn to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetByColumn(byColumn *int32) {
	o.ByColumn = byColumn
}

// WithLimit adds the limit to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithLimit(limit *int32) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithPage adds the page to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithPage(page *int32) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetPage(page *int32) {
	o.Page = page
}

// WithQuery adds the query to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithQuery(query *string) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetQuery(query *string) {
	o.Query = query
}

// WithRootDomainName adds the rootDomainName to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithRootDomainName(rootDomainName string) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetRootDomainName(rootDomainName)
	return o
}

// SetRootDomainName adds the rootDomainName to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetRootDomainName(rootDomainName string) {
	o.RootDomainName = rootDomainName
}

// WithTargetName adds the targetName to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WithTargetName(targetName string) *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams {
	o.SetTargetName(targetName)
	return o
}

// SetTargetName adds the targetName to the get API subdomains get paginate target name root domain name params
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) SetTargetName(targetName string) {
	o.TargetName = targetName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPISubdomainsGetPaginateTargetNameRootDomainNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ascending != nil {

		// query param Ascending
		var qrAscending int32

		if o.Ascending != nil {
			qrAscending = *o.Ascending
		}
		qAscending := swag.FormatInt32(qrAscending)
		if qAscending != "" {

			if err := r.SetQueryParam("Ascending", qAscending); err != nil {
				return err
			}
		}
	}

	if o.ByColumn != nil {

		// query param ByColumn
		var qrByColumn int32

		if o.ByColumn != nil {
			qrByColumn = *o.ByColumn
		}
		qByColumn := swag.FormatInt32(qrByColumn)
		if qByColumn != "" {

			if err := r.SetQueryParam("ByColumn", qByColumn); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param Limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("Limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param Page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("Page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param Query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("Query", qQuery); err != nil {
				return err
			}
		}
	}

	// path param rootDomainName
	if err := r.SetPathParam("rootDomainName", o.RootDomainName); err != nil {
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
