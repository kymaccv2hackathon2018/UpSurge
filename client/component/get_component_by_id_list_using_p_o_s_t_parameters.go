// Code generated by go-swagger; DO NOT EDIT.

package component

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
	models "github.com/gopdf/models"
	"golang.org/x/net/context"
)

// NewGetComponentByIDListUsingPOSTParams creates a new GetComponentByIDListUsingPOSTParams object
// with the default values initialized.
func NewGetComponentByIDListUsingPOSTParams() *GetComponentByIDListUsingPOSTParams {
	var (
		currentPageDefault = int32(0)
		fieldsDefault      = string("DEFAULT")
		pageSizeDefault    = int32(10)
	)
	return &GetComponentByIDListUsingPOSTParams{
		CurrentPage: &currentPageDefault,
		Fields:      &fieldsDefault,
		PageSize:    &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetComponentByIDListUsingPOSTParamsWithTimeout creates a new GetComponentByIDListUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetComponentByIDListUsingPOSTParamsWithTimeout(timeout time.Duration) *GetComponentByIDListUsingPOSTParams {
	var (
		currentPageDefault = int32(0)
		fieldsDefault      = string("DEFAULT")
		pageSizeDefault    = int32(10)
	)
	return &GetComponentByIDListUsingPOSTParams{
		CurrentPage: &currentPageDefault,
		Fields:      &fieldsDefault,
		PageSize:    &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetComponentByIDListUsingPOSTParamsWithContext creates a new GetComponentByIDListUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetComponentByIDListUsingPOSTParamsWithContext(ctx context.Context) *GetComponentByIDListUsingPOSTParams {
	var (
		currentPageDefault = int32(0)
		fieldsDefault      = string("DEFAULT")
		pageSizeDefault    = int32(10)
	)
	return &GetComponentByIDListUsingPOSTParams{
		CurrentPage: &currentPageDefault,
		Fields:      &fieldsDefault,
		PageSize:    &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetComponentByIDListUsingPOSTParamsWithHTTPClient creates a new GetComponentByIDListUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetComponentByIDListUsingPOSTParamsWithHTTPClient(client *http.Client) *GetComponentByIDListUsingPOSTParams {
	var (
		currentPageDefault = int32(0)
		fieldsDefault      = string("DEFAULT")
		pageSizeDefault    = int32(10)
	)
	return &GetComponentByIDListUsingPOSTParams{
		CurrentPage: &currentPageDefault,
		Fields:      &fieldsDefault,
		PageSize:    &pageSizeDefault,
		HTTPClient:  client,
	}
}

/*GetComponentByIDListUsingPOSTParams contains all the parameters to send to the API endpoint
for the get component by Id list using p o s t operation typically these are written to a http.Request
*/
type GetComponentByIDListUsingPOSTParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CatalogCode
	  Catalog code

	*/
	CatalogCode *string
	/*CategoryCode
	  Category code

	*/
	CategoryCode *string
	/*ComponentIDList
	  List of Component identifiers

	*/
	ComponentIDList *models.ComponentIDListWsDTO
	/*CurrentPage
	  Optional pagination parameter. Default value 0.

	*/
	CurrentPage *int32
	/*Fields
	  Response configuration (list of fields, which should be returned in response)

	*/
	Fields *string
	/*PageSize
	  Optional pagination parameter. Default value 10.

	*/
	PageSize *int32
	/*ProductCode
	  Product code

	*/
	ProductCode *string
	/*Sort
	  Optional sort criterion. No default value.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithTimeout(timeout time.Duration) *GetComponentByIDListUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithContext(ctx context.Context) *GetComponentByIDListUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithHTTPClient(client *http.Client) *GetComponentByIDListUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithBaseSiteID(baseSiteID string) *GetComponentByIDListUsingPOSTParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCatalogCode adds the catalogCode to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithCatalogCode(catalogCode *string) *GetComponentByIDListUsingPOSTParams {
	o.SetCatalogCode(catalogCode)
	return o
}

// SetCatalogCode adds the catalogCode to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetCatalogCode(catalogCode *string) {
	o.CatalogCode = catalogCode
}

// WithCategoryCode adds the categoryCode to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithCategoryCode(categoryCode *string) *GetComponentByIDListUsingPOSTParams {
	o.SetCategoryCode(categoryCode)
	return o
}

// SetCategoryCode adds the categoryCode to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetCategoryCode(categoryCode *string) {
	o.CategoryCode = categoryCode
}

// WithComponentIDList adds the componentIDList to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithComponentIDList(componentIDList *models.ComponentIDListWsDTO) *GetComponentByIDListUsingPOSTParams {
	o.SetComponentIDList(componentIDList)
	return o
}

// SetComponentIDList adds the componentIdList to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetComponentIDList(componentIDList *models.ComponentIDListWsDTO) {
	o.ComponentIDList = componentIDList
}

// WithCurrentPage adds the currentPage to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithCurrentPage(currentPage *int32) *GetComponentByIDListUsingPOSTParams {
	o.SetCurrentPage(currentPage)
	return o
}

// SetCurrentPage adds the currentPage to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetCurrentPage(currentPage *int32) {
	o.CurrentPage = currentPage
}

// WithFields adds the fields to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithFields(fields *string) *GetComponentByIDListUsingPOSTParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithPageSize adds the pageSize to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithPageSize(pageSize *int32) *GetComponentByIDListUsingPOSTParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithProductCode adds the productCode to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithProductCode(productCode *string) *GetComponentByIDListUsingPOSTParams {
	o.SetProductCode(productCode)
	return o
}

// SetProductCode adds the productCode to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetProductCode(productCode *string) {
	o.ProductCode = productCode
}

// WithSort adds the sort to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) WithSort(sort *string) *GetComponentByIDListUsingPOSTParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get component by Id list using p o s t params
func (o *GetComponentByIDListUsingPOSTParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetComponentByIDListUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	if o.CatalogCode != nil {

		// query param catalogCode
		var qrCatalogCode string
		if o.CatalogCode != nil {
			qrCatalogCode = *o.CatalogCode
		}
		qCatalogCode := qrCatalogCode
		if qCatalogCode != "" {
			if err := r.SetQueryParam("catalogCode", qCatalogCode); err != nil {
				return err
			}
		}

	}

	if o.CategoryCode != nil {

		// query param categoryCode
		var qrCategoryCode string
		if o.CategoryCode != nil {
			qrCategoryCode = *o.CategoryCode
		}
		qCategoryCode := qrCategoryCode
		if qCategoryCode != "" {
			if err := r.SetQueryParam("categoryCode", qCategoryCode); err != nil {
				return err
			}
		}

	}

	if o.ComponentIDList != nil {
		if err := r.SetBodyParam(o.ComponentIDList); err != nil {
			return err
		}
	}

	if o.CurrentPage != nil {

		// query param currentPage
		var qrCurrentPage int32
		if o.CurrentPage != nil {
			qrCurrentPage = *o.CurrentPage
		}
		qCurrentPage := swag.FormatInt32(qrCurrentPage)
		if qCurrentPage != "" {
			if err := r.SetQueryParam("currentPage", qCurrentPage); err != nil {
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

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.ProductCode != nil {

		// query param productCode
		var qrProductCode string
		if o.ProductCode != nil {
			qrProductCode = *o.ProductCode
		}
		qProductCode := qrProductCode
		if qProductCode != "" {
			if err := r.SetQueryParam("productCode", qProductCode); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}