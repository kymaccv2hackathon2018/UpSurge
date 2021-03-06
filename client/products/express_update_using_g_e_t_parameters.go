// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewExpressUpdateUsingGETParams creates a new ExpressUpdateUsingGETParams object
// with the default values initialized.
func NewExpressUpdateUsingGETParams() *ExpressUpdateUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &ExpressUpdateUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewExpressUpdateUsingGETParamsWithTimeout creates a new ExpressUpdateUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExpressUpdateUsingGETParamsWithTimeout(timeout time.Duration) *ExpressUpdateUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &ExpressUpdateUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewExpressUpdateUsingGETParamsWithContext creates a new ExpressUpdateUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewExpressUpdateUsingGETParamsWithContext(ctx context.Context) *ExpressUpdateUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &ExpressUpdateUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewExpressUpdateUsingGETParamsWithHTTPClient creates a new ExpressUpdateUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExpressUpdateUsingGETParamsWithHTTPClient(client *http.Client) *ExpressUpdateUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &ExpressUpdateUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*ExpressUpdateUsingGETParams contains all the parameters to send to the API endpoint
for the express update using g e t operation typically these are written to a http.Request
*/
type ExpressUpdateUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*Catalog
	  Only products from this catalog are returned. Format: catalogId:catalogVersion

	*/
	Catalog *string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*Timestamp
	  Only items newer than the given parameter are retrieved from the queue. This parameter should be in ISO-8601 format.

	*/
	Timestamp string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) WithTimeout(timeout time.Duration) *ExpressUpdateUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) WithContext(ctx context.Context) *ExpressUpdateUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) WithHTTPClient(client *http.Client) *ExpressUpdateUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) WithBaseSiteID(baseSiteID string) *ExpressUpdateUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCatalog adds the catalog to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) WithCatalog(catalog *string) *ExpressUpdateUsingGETParams {
	o.SetCatalog(catalog)
	return o
}

// SetCatalog adds the catalog to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) SetCatalog(catalog *string) {
	o.Catalog = catalog
}

// WithFields adds the fields to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) WithFields(fields *string) *ExpressUpdateUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithTimestamp adds the timestamp to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) WithTimestamp(timestamp string) *ExpressUpdateUsingGETParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the express update using g e t params
func (o *ExpressUpdateUsingGETParams) SetTimestamp(timestamp string) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *ExpressUpdateUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	if o.Catalog != nil {

		// query param catalog
		var qrCatalog string
		if o.Catalog != nil {
			qrCatalog = *o.Catalog
		}
		qCatalog := qrCatalog
		if qCatalog != "" {
			if err := r.SetQueryParam("catalog", qCatalog); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// query param timestamp
	qrTimestamp := o.Timestamp
	qTimestamp := qrTimestamp
	if qTimestamp != "" {
		if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
