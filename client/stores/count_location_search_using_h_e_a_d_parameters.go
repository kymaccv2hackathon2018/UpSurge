// Code generated by go-swagger; DO NOT EDIT.

package stores

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"golang.org/x/net/context"
)

// NewCountLocationSearchUsingHEADParams creates a new CountLocationSearchUsingHEADParams object
// with the default values initialized.
func NewCountLocationSearchUsingHEADParams() *CountLocationSearchUsingHEADParams {
	var (
		accuracyDefault = float64(0)
		radiusDefault   = float64(100000)
	)
	return &CountLocationSearchUsingHEADParams{
		Accuracy: &accuracyDefault,
		Radius:   &radiusDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCountLocationSearchUsingHEADParamsWithTimeout creates a new CountLocationSearchUsingHEADParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCountLocationSearchUsingHEADParamsWithTimeout(timeout time.Duration) *CountLocationSearchUsingHEADParams {
	var (
		accuracyDefault = float64(0)
		radiusDefault   = float64(100000)
	)
	return &CountLocationSearchUsingHEADParams{
		Accuracy: &accuracyDefault,
		Radius:   &radiusDefault,

		timeout: timeout,
	}
}

// NewCountLocationSearchUsingHEADParamsWithContext creates a new CountLocationSearchUsingHEADParams object
// with the default values initialized, and the ability to set a context for a request
func NewCountLocationSearchUsingHEADParamsWithContext(ctx context.Context) *CountLocationSearchUsingHEADParams {
	var (
		accuracyDefault = float64(0)
		radiusDefault   = float64(100000)
	)
	return &CountLocationSearchUsingHEADParams{
		Accuracy: &accuracyDefault,
		Radius:   &radiusDefault,

		Context: ctx,
	}
}

// NewCountLocationSearchUsingHEADParamsWithHTTPClient creates a new CountLocationSearchUsingHEADParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCountLocationSearchUsingHEADParamsWithHTTPClient(client *http.Client) *CountLocationSearchUsingHEADParams {
	var (
		accuracyDefault = float64(0)
		radiusDefault   = float64(100000)
	)
	return &CountLocationSearchUsingHEADParams{
		Accuracy:   &accuracyDefault,
		Radius:     &radiusDefault,
		HTTPClient: client,
	}
}

/*CountLocationSearchUsingHEADParams contains all the parameters to send to the API endpoint
for the count location search using h e a d operation typically these are written to a http.Request
*/
type CountLocationSearchUsingHEADParams struct {

	/*Accuracy
	  Accuracy in meters.

	*/
	Accuracy *float64
	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*Latitude
	  Coordinate that specifies the north-south position of a point on the Earth's surface.

	*/
	Latitude *float64
	/*Longitude
	  Coordinate that specifies the east-west position of a point on the Earth's surface.

	*/
	Longitude *float64
	/*Query
	  Location in natural language i.e. city or country.

	*/
	Query *string
	/*Radius
	  Radius in meters. Max value: 40075000.0 (Earth's perimeter).

	*/
	Radius *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithTimeout(timeout time.Duration) *CountLocationSearchUsingHEADParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithContext(ctx context.Context) *CountLocationSearchUsingHEADParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithHTTPClient(client *http.Client) *CountLocationSearchUsingHEADParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccuracy adds the accuracy to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithAccuracy(accuracy *float64) *CountLocationSearchUsingHEADParams {
	o.SetAccuracy(accuracy)
	return o
}

// SetAccuracy adds the accuracy to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetAccuracy(accuracy *float64) {
	o.Accuracy = accuracy
}

// WithBaseSiteID adds the baseSiteID to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithBaseSiteID(baseSiteID string) *CountLocationSearchUsingHEADParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithLatitude adds the latitude to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithLatitude(latitude *float64) *CountLocationSearchUsingHEADParams {
	o.SetLatitude(latitude)
	return o
}

// SetLatitude adds the latitude to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetLatitude(latitude *float64) {
	o.Latitude = latitude
}

// WithLongitude adds the longitude to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithLongitude(longitude *float64) *CountLocationSearchUsingHEADParams {
	o.SetLongitude(longitude)
	return o
}

// SetLongitude adds the longitude to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetLongitude(longitude *float64) {
	o.Longitude = longitude
}

// WithQuery adds the query to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithQuery(query *string) *CountLocationSearchUsingHEADParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetQuery(query *string) {
	o.Query = query
}

// WithRadius adds the radius to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) WithRadius(radius *float64) *CountLocationSearchUsingHEADParams {
	o.SetRadius(radius)
	return o
}

// SetRadius adds the radius to the count location search using h e a d params
func (o *CountLocationSearchUsingHEADParams) SetRadius(radius *float64) {
	o.Radius = radius
}

// WriteToRequest writes these params to a swagger request
func (o *CountLocationSearchUsingHEADParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Accuracy != nil {

		// query param accuracy
		var qrAccuracy float64
		if o.Accuracy != nil {
			qrAccuracy = *o.Accuracy
		}
		qAccuracy := swag.FormatFloat64(qrAccuracy)
		if qAccuracy != "" {
			if err := r.SetQueryParam("accuracy", qAccuracy); err != nil {
				return err
			}
		}

	}

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	if o.Latitude != nil {

		// query param latitude
		var qrLatitude float64
		if o.Latitude != nil {
			qrLatitude = *o.Latitude
		}
		qLatitude := swag.FormatFloat64(qrLatitude)
		if qLatitude != "" {
			if err := r.SetQueryParam("latitude", qLatitude); err != nil {
				return err
			}
		}

	}

	if o.Longitude != nil {

		// query param longitude
		var qrLongitude float64
		if o.Longitude != nil {
			qrLongitude = *o.Longitude
		}
		qLongitude := swag.FormatFloat64(qrLongitude)
		if qLongitude != "" {
			if err := r.SetQueryParam("longitude", qLongitude); err != nil {
				return err
			}
		}

	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.Radius != nil {

		// query param radius
		var qrRadius float64
		if o.Radius != nil {
			qrRadius = *o.Radius
		}
		qRadius := swag.FormatFloat64(qrRadius)
		if qRadius != "" {
			if err := r.SetQueryParam("radius", qRadius); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
