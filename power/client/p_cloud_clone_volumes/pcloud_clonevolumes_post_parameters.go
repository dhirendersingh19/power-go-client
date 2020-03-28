// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_clone_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudClonevolumesPostParams creates a new PcloudClonevolumesPostParams object
// with the default values initialized.
func NewPcloudClonevolumesPostParams() *PcloudClonevolumesPostParams {
	var ()
	return &PcloudClonevolumesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudClonevolumesPostParamsWithTimeout creates a new PcloudClonevolumesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudClonevolumesPostParamsWithTimeout(timeout time.Duration) *PcloudClonevolumesPostParams {
	var ()
	return &PcloudClonevolumesPostParams{

		timeout: timeout,
	}
}

// NewPcloudClonevolumesPostParamsWithContext creates a new PcloudClonevolumesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudClonevolumesPostParamsWithContext(ctx context.Context) *PcloudClonevolumesPostParams {
	var ()
	return &PcloudClonevolumesPostParams{

		Context: ctx,
	}
}

// NewPcloudClonevolumesPostParamsWithHTTPClient creates a new PcloudClonevolumesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudClonevolumesPostParamsWithHTTPClient(client *http.Client) *PcloudClonevolumesPostParams {
	var ()
	return &PcloudClonevolumesPostParams{
		HTTPClient: client,
	}
}

/*PcloudClonevolumesPostParams contains all the parameters to send to the API endpoint
for the pcloud clonevolumes post operation typically these are written to a http.Request
*/
type PcloudClonevolumesPostParams struct {

	/*Body
	  Parameters for the cloning of volumes

	*/
	Body *models.CloneVolumesRequest
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) WithTimeout(timeout time.Duration) *PcloudClonevolumesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) WithContext(ctx context.Context) *PcloudClonevolumesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) WithHTTPClient(client *http.Client) *PcloudClonevolumesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) WithBody(body *models.CloneVolumesRequest) *PcloudClonevolumesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) SetBody(body *models.CloneVolumesRequest) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudClonevolumesPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud clonevolumes post params
func (o *PcloudClonevolumesPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudClonevolumesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
