// Code generated by go-swagger; DO NOT EDIT.

package scrape_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/percona/pmm-managed/api/swagger/models"
)

// NewAddStaticTargetsParams creates a new AddStaticTargetsParams object
// with the default values initialized.
func NewAddStaticTargetsParams() *AddStaticTargetsParams {
	var ()
	return &AddStaticTargetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddStaticTargetsParamsWithTimeout creates a new AddStaticTargetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddStaticTargetsParamsWithTimeout(timeout time.Duration) *AddStaticTargetsParams {
	var ()
	return &AddStaticTargetsParams{

		timeout: timeout,
	}
}

// NewAddStaticTargetsParamsWithContext creates a new AddStaticTargetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddStaticTargetsParamsWithContext(ctx context.Context) *AddStaticTargetsParams {
	var ()
	return &AddStaticTargetsParams{

		Context: ctx,
	}
}

// NewAddStaticTargetsParamsWithHTTPClient creates a new AddStaticTargetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddStaticTargetsParamsWithHTTPClient(client *http.Client) *AddStaticTargetsParams {
	var ()
	return &AddStaticTargetsParams{
		HTTPClient: client,
	}
}

/*AddStaticTargetsParams contains all the parameters to send to the API endpoint
for the add static targets operation typically these are written to a http.Request
*/
type AddStaticTargetsParams struct {

	/*Body*/
	Body *models.APIScrapeConfigsAddStaticTargetsRequest
	/*JobName*/
	JobName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add static targets params
func (o *AddStaticTargetsParams) WithTimeout(timeout time.Duration) *AddStaticTargetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add static targets params
func (o *AddStaticTargetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add static targets params
func (o *AddStaticTargetsParams) WithContext(ctx context.Context) *AddStaticTargetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add static targets params
func (o *AddStaticTargetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add static targets params
func (o *AddStaticTargetsParams) WithHTTPClient(client *http.Client) *AddStaticTargetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add static targets params
func (o *AddStaticTargetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add static targets params
func (o *AddStaticTargetsParams) WithBody(body *models.APIScrapeConfigsAddStaticTargetsRequest) *AddStaticTargetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add static targets params
func (o *AddStaticTargetsParams) SetBody(body *models.APIScrapeConfigsAddStaticTargetsRequest) {
	o.Body = body
}

// WithJobName adds the jobName to the add static targets params
func (o *AddStaticTargetsParams) WithJobName(jobName string) *AddStaticTargetsParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the add static targets params
func (o *AddStaticTargetsParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WriteToRequest writes these params to a swagger request
func (o *AddStaticTargetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.APIScrapeConfigsAddStaticTargetsRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param job_name
	if err := r.SetPathParam("job_name", o.JobName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}