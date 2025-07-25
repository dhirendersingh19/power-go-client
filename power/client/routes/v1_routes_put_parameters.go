// Code generated by go-swagger; DO NOT EDIT.

package routes

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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewV1RoutesPutParams creates a new V1RoutesPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1RoutesPutParams() *V1RoutesPutParams {
	return &V1RoutesPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1RoutesPutParamsWithTimeout creates a new V1RoutesPutParams object
// with the ability to set a timeout on a request.
func NewV1RoutesPutParamsWithTimeout(timeout time.Duration) *V1RoutesPutParams {
	return &V1RoutesPutParams{
		timeout: timeout,
	}
}

// NewV1RoutesPutParamsWithContext creates a new V1RoutesPutParams object
// with the ability to set a context for a request.
func NewV1RoutesPutParamsWithContext(ctx context.Context) *V1RoutesPutParams {
	return &V1RoutesPutParams{
		Context: ctx,
	}
}

// NewV1RoutesPutParamsWithHTTPClient creates a new V1RoutesPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1RoutesPutParamsWithHTTPClient(client *http.Client) *V1RoutesPutParams {
	return &V1RoutesPutParams{
		HTTPClient: client,
	}
}

/*
V1RoutesPutParams contains all the parameters to send to the API endpoint

	for the v1 routes put operation.

	Typically these are written to a http.Request.
*/
type V1RoutesPutParams struct {

	/* Body.

	   Parameters for updating a route
	*/
	Body *models.RouteUpdate

	/* RouteID.

	   Route ID
	*/
	RouteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 routes put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1RoutesPutParams) WithDefaults() *V1RoutesPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 routes put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1RoutesPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 routes put params
func (o *V1RoutesPutParams) WithTimeout(timeout time.Duration) *V1RoutesPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 routes put params
func (o *V1RoutesPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 routes put params
func (o *V1RoutesPutParams) WithContext(ctx context.Context) *V1RoutesPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 routes put params
func (o *V1RoutesPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 routes put params
func (o *V1RoutesPutParams) WithHTTPClient(client *http.Client) *V1RoutesPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 routes put params
func (o *V1RoutesPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 routes put params
func (o *V1RoutesPutParams) WithBody(body *models.RouteUpdate) *V1RoutesPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 routes put params
func (o *V1RoutesPutParams) SetBody(body *models.RouteUpdate) {
	o.Body = body
}

// WithRouteID adds the routeID to the v1 routes put params
func (o *V1RoutesPutParams) WithRouteID(routeID string) *V1RoutesPutParams {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the v1 routes put params
func (o *V1RoutesPutParams) SetRouteID(routeID string) {
	o.RouteID = routeID
}

// WriteToRequest writes these params to a swagger request
func (o *V1RoutesPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param route_id
	if err := r.SetPathParam("route_id", o.RouteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
