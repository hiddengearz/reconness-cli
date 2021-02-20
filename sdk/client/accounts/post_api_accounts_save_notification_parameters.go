// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewPostAPIAccountsSaveNotificationParams creates a new PostAPIAccountsSaveNotificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIAccountsSaveNotificationParams() *PostAPIAccountsSaveNotificationParams {
	return &PostAPIAccountsSaveNotificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIAccountsSaveNotificationParamsWithTimeout creates a new PostAPIAccountsSaveNotificationParams object
// with the ability to set a timeout on a request.
func NewPostAPIAccountsSaveNotificationParamsWithTimeout(timeout time.Duration) *PostAPIAccountsSaveNotificationParams {
	return &PostAPIAccountsSaveNotificationParams{
		timeout: timeout,
	}
}

// NewPostAPIAccountsSaveNotificationParamsWithContext creates a new PostAPIAccountsSaveNotificationParams object
// with the ability to set a context for a request.
func NewPostAPIAccountsSaveNotificationParamsWithContext(ctx context.Context) *PostAPIAccountsSaveNotificationParams {
	return &PostAPIAccountsSaveNotificationParams{
		Context: ctx,
	}
}

// NewPostAPIAccountsSaveNotificationParamsWithHTTPClient creates a new PostAPIAccountsSaveNotificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIAccountsSaveNotificationParamsWithHTTPClient(client *http.Client) *PostAPIAccountsSaveNotificationParams {
	return &PostAPIAccountsSaveNotificationParams{
		HTTPClient: client,
	}
}

/* PostAPIAccountsSaveNotificationParams contains all the parameters to send to the API endpoint
   for the post API accounts save notification operation.

   Typically these are written to a http.Request.
*/
type PostAPIAccountsSaveNotificationParams struct {

	// Body.
	Body *models.NotificationDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API accounts save notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIAccountsSaveNotificationParams) WithDefaults() *PostAPIAccountsSaveNotificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API accounts save notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIAccountsSaveNotificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API accounts save notification params
func (o *PostAPIAccountsSaveNotificationParams) WithTimeout(timeout time.Duration) *PostAPIAccountsSaveNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API accounts save notification params
func (o *PostAPIAccountsSaveNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API accounts save notification params
func (o *PostAPIAccountsSaveNotificationParams) WithContext(ctx context.Context) *PostAPIAccountsSaveNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API accounts save notification params
func (o *PostAPIAccountsSaveNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API accounts save notification params
func (o *PostAPIAccountsSaveNotificationParams) WithHTTPClient(client *http.Client) *PostAPIAccountsSaveNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API accounts save notification params
func (o *PostAPIAccountsSaveNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API accounts save notification params
func (o *PostAPIAccountsSaveNotificationParams) WithBody(body *models.NotificationDto) *PostAPIAccountsSaveNotificationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API accounts save notification params
func (o *PostAPIAccountsSaveNotificationParams) SetBody(body *models.NotificationDto) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIAccountsSaveNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
