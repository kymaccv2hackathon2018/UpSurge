// Code generated by go-swagger; DO NOT EDIT.

package catalogs

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

// NewGetCategoriesUsingGETParams creates a new GetCategoriesUsingGETParams object
// with the default values initialized.
func NewGetCategoriesUsingGETParams() *GetCategoriesUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCategoriesUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCategoriesUsingGETParamsWithTimeout creates a new GetCategoriesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCategoriesUsingGETParamsWithTimeout(timeout time.Duration) *GetCategoriesUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCategoriesUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetCategoriesUsingGETParamsWithContext creates a new GetCategoriesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCategoriesUsingGETParamsWithContext(ctx context.Context) *GetCategoriesUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCategoriesUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetCategoriesUsingGETParamsWithHTTPClient creates a new GetCategoriesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCategoriesUsingGETParamsWithHTTPClient(client *http.Client) *GetCategoriesUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCategoriesUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetCategoriesUsingGETParams contains all the parameters to send to the API endpoint
for the get categories using g e t operation typically these are written to a http.Request
*/
type GetCategoriesUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CatalogID
	  Catalog identifier

	*/
	CatalogID string
	/*CatalogVersionID
	  Catalog version identifier

	*/
	CatalogVersionID string
	/*CategoryID
	  Category identifier

	*/
	CategoryID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) WithTimeout(timeout time.Duration) *GetCategoriesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) WithContext(ctx context.Context) *GetCategoriesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) WithHTTPClient(client *http.Client) *GetCategoriesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) WithBaseSiteID(baseSiteID string) *GetCategoriesUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCatalogID adds the catalogID to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) WithCatalogID(catalogID string) *GetCategoriesUsingGETParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithCatalogVersionID adds the catalogVersionID to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) WithCatalogVersionID(catalogVersionID string) *GetCategoriesUsingGETParams {
	o.SetCatalogVersionID(catalogVersionID)
	return o
}

// SetCatalogVersionID adds the catalogVersionId to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) SetCatalogVersionID(catalogVersionID string) {
	o.CatalogVersionID = catalogVersionID
}

// WithCategoryID adds the categoryID to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) WithCategoryID(categoryID string) *GetCategoriesUsingGETParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithFields adds the fields to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) WithFields(fields *string) *GetCategoriesUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get categories using g e t params
func (o *GetCategoriesUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetCategoriesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// path param catalogId
	if err := r.SetPathParam("catalogId", o.CatalogID); err != nil {
		return err
	}

	// path param catalogVersionId
	if err := r.SetPathParam("catalogVersionId", o.CatalogVersionID); err != nil {
		return err
	}

	// path param categoryId
	if err := r.SetPathParam("categoryId", o.CategoryID); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
