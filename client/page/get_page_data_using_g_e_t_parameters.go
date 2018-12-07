// Code generated by go-swagger; DO NOT EDIT.

package page

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

// NewGetPageDataUsingGETParams creates a new GetPageDataUsingGETParams object
// with the default values initialized.
func NewGetPageDataUsingGETParams() *GetPageDataUsingGETParams {
	var (
		fieldsDefault   = string("DEFAULT")
		pageTypeDefault = string("ContentPage")
	)
	return &GetPageDataUsingGETParams{
		Fields:   &fieldsDefault,
		PageType: &pageTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPageDataUsingGETParamsWithTimeout creates a new GetPageDataUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPageDataUsingGETParamsWithTimeout(timeout time.Duration) *GetPageDataUsingGETParams {
	var (
		fieldsDefault   = string("DEFAULT")
		pageTypeDefault = string("ContentPage")
	)
	return &GetPageDataUsingGETParams{
		Fields:   &fieldsDefault,
		PageType: &pageTypeDefault,

		timeout: timeout,
	}
}

// NewGetPageDataUsingGETParamsWithContext creates a new GetPageDataUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPageDataUsingGETParamsWithContext(ctx context.Context) *GetPageDataUsingGETParams {
	var (
		fieldsDefault   = string("DEFAULT")
		pageTypeDefault = string("ContentPage")
	)
	return &GetPageDataUsingGETParams{
		Fields:   &fieldsDefault,
		PageType: &pageTypeDefault,

		Context: ctx,
	}
}

// NewGetPageDataUsingGETParamsWithHTTPClient creates a new GetPageDataUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPageDataUsingGETParamsWithHTTPClient(client *http.Client) *GetPageDataUsingGETParams {
	var (
		fieldsDefault   = string("DEFAULT")
		pageTypeDefault = string("ContentPage")
	)
	return &GetPageDataUsingGETParams{
		Fields:     &fieldsDefault,
		PageType:   &pageTypeDefault,
		HTTPClient: client,
	}
}

/*GetPageDataUsingGETParams contains all the parameters to send to the API endpoint
for the get page data using g e t operation typically these are written to a http.Request
*/
type GetPageDataUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*Code
	  If pageType is ProductPage, code should be product code; if pageType is CategoryPage, code should be category code; if pageType is CatalogPage, code should be catalog code

	*/
	Code *string
	/*Fields
	  Response configuration (list of fields, which should be returned in response)

	*/
	Fields *string
	/*PageLabelOrID
	  Page Label or Id

	*/
	PageLabelOrID *string
	/*PageType
	  page type

	*/
	PageType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get page data using g e t params
func (o *GetPageDataUsingGETParams) WithTimeout(timeout time.Duration) *GetPageDataUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get page data using g e t params
func (o *GetPageDataUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get page data using g e t params
func (o *GetPageDataUsingGETParams) WithContext(ctx context.Context) *GetPageDataUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get page data using g e t params
func (o *GetPageDataUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get page data using g e t params
func (o *GetPageDataUsingGETParams) WithHTTPClient(client *http.Client) *GetPageDataUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get page data using g e t params
func (o *GetPageDataUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get page data using g e t params
func (o *GetPageDataUsingGETParams) WithBaseSiteID(baseSiteID string) *GetPageDataUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get page data using g e t params
func (o *GetPageDataUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCode adds the code to the get page data using g e t params
func (o *GetPageDataUsingGETParams) WithCode(code *string) *GetPageDataUsingGETParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the get page data using g e t params
func (o *GetPageDataUsingGETParams) SetCode(code *string) {
	o.Code = code
}

// WithFields adds the fields to the get page data using g e t params
func (o *GetPageDataUsingGETParams) WithFields(fields *string) *GetPageDataUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get page data using g e t params
func (o *GetPageDataUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithPageLabelOrID adds the pageLabelOrID to the get page data using g e t params
func (o *GetPageDataUsingGETParams) WithPageLabelOrID(pageLabelOrID *string) *GetPageDataUsingGETParams {
	o.SetPageLabelOrID(pageLabelOrID)
	return o
}

// SetPageLabelOrID adds the pageLabelOrId to the get page data using g e t params
func (o *GetPageDataUsingGETParams) SetPageLabelOrID(pageLabelOrID *string) {
	o.PageLabelOrID = pageLabelOrID
}

// WithPageType adds the pageType to the get page data using g e t params
func (o *GetPageDataUsingGETParams) WithPageType(pageType *string) *GetPageDataUsingGETParams {
	o.SetPageType(pageType)
	return o
}

// SetPageType adds the pageType to the get page data using g e t params
func (o *GetPageDataUsingGETParams) SetPageType(pageType *string) {
	o.PageType = pageType
}

// WriteToRequest writes these params to a swagger request
func (o *GetPageDataUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	if o.Code != nil {

		// query param code
		var qrCode string
		if o.Code != nil {
			qrCode = *o.Code
		}
		qCode := qrCode
		if qCode != "" {
			if err := r.SetQueryParam("code", qCode); err != nil {
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

	if o.PageLabelOrID != nil {

		// query param pageLabelOrId
		var qrPageLabelOrID string
		if o.PageLabelOrID != nil {
			qrPageLabelOrID = *o.PageLabelOrID
		}
		qPageLabelOrID := qrPageLabelOrID
		if qPageLabelOrID != "" {
			if err := r.SetQueryParam("pageLabelOrId", qPageLabelOrID); err != nil {
				return err
			}
		}

	}

	if o.PageType != nil {

		// query param pageType
		var qrPageType string
		if o.PageType != nil {
			qrPageType = *o.PageType
		}
		qPageType := qrPageType
		if qPageType != "" {
			if err := r.SetQueryParam("pageType", qPageType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}