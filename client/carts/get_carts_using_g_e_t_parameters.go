// Code generated by go-swagger; DO NOT EDIT.

package carts

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

// NewGetCartsUsingGETParams creates a new GetCartsUsingGETParams object
// with the default values initialized.
func NewGetCartsUsingGETParams() *GetCartsUsingGETParams {
	var (
		currentPageDefault    = int32(0)
		fieldsDefault         = string("DEFAULT")
		pageSizeDefault       = int32(20)
		savedCartsOnlyDefault = bool(false)
	)
	return &GetCartsUsingGETParams{
		CurrentPage:    &currentPageDefault,
		Fields:         &fieldsDefault,
		PageSize:       &pageSizeDefault,
		SavedCartsOnly: &savedCartsOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCartsUsingGETParamsWithTimeout creates a new GetCartsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCartsUsingGETParamsWithTimeout(timeout time.Duration) *GetCartsUsingGETParams {
	var (
		currentPageDefault    = int32(0)
		fieldsDefault         = string("DEFAULT")
		pageSizeDefault       = int32(20)
		savedCartsOnlyDefault = bool(false)
	)
	return &GetCartsUsingGETParams{
		CurrentPage:    &currentPageDefault,
		Fields:         &fieldsDefault,
		PageSize:       &pageSizeDefault,
		SavedCartsOnly: &savedCartsOnlyDefault,

		timeout: timeout,
	}
}

// NewGetCartsUsingGETParamsWithContext creates a new GetCartsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCartsUsingGETParamsWithContext(ctx context.Context) *GetCartsUsingGETParams {
	var (
		currentPageDefault    = int32(0)
		fieldsDefault         = string("DEFAULT")
		pageSizeDefault       = int32(20)
		savedCartsOnlyDefault = bool(false)
	)
	return &GetCartsUsingGETParams{
		CurrentPage:    &currentPageDefault,
		Fields:         &fieldsDefault,
		PageSize:       &pageSizeDefault,
		SavedCartsOnly: &savedCartsOnlyDefault,

		Context: ctx,
	}
}

// NewGetCartsUsingGETParamsWithHTTPClient creates a new GetCartsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCartsUsingGETParamsWithHTTPClient(client *http.Client) *GetCartsUsingGETParams {
	var (
		currentPageDefault    = int32(0)
		fieldsDefault         = string("DEFAULT")
		pageSizeDefault       = int32(20)
		savedCartsOnlyDefault = bool(false)
	)
	return &GetCartsUsingGETParams{
		CurrentPage:    &currentPageDefault,
		Fields:         &fieldsDefault,
		PageSize:       &pageSizeDefault,
		SavedCartsOnly: &savedCartsOnlyDefault,
		HTTPClient:     client,
	}
}

/*GetCartsUsingGETParams contains all the parameters to send to the API endpoint
for the get carts using g e t operation typically these are written to a http.Request
*/
type GetCartsUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CurrentPage
	  Optional pagination parameter in case of savedCartsOnly == true. Default value 0.

	*/
	CurrentPage *int32
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*PageSize
	  Optional {@link PaginationData} parameter in case of savedCartsOnly == true. Default value 20.

	*/
	PageSize *int32
	/*SavedCartsOnly
	  Optional parameter. If the parameter is provided and its value is true, only saved carts are returned.

	*/
	SavedCartsOnly *bool
	/*Sort
	  Optional sort criterion in case of savedCartsOnly == true. No default value.

	*/
	Sort *string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithTimeout(timeout time.Duration) *GetCartsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithContext(ctx context.Context) *GetCartsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithHTTPClient(client *http.Client) *GetCartsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithBaseSiteID(baseSiteID string) *GetCartsUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCurrentPage adds the currentPage to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithCurrentPage(currentPage *int32) *GetCartsUsingGETParams {
	o.SetCurrentPage(currentPage)
	return o
}

// SetCurrentPage adds the currentPage to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetCurrentPage(currentPage *int32) {
	o.CurrentPage = currentPage
}

// WithFields adds the fields to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithFields(fields *string) *GetCartsUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithPageSize adds the pageSize to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithPageSize(pageSize *int32) *GetCartsUsingGETParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSavedCartsOnly adds the savedCartsOnly to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithSavedCartsOnly(savedCartsOnly *bool) *GetCartsUsingGETParams {
	o.SetSavedCartsOnly(savedCartsOnly)
	return o
}

// SetSavedCartsOnly adds the savedCartsOnly to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetSavedCartsOnly(savedCartsOnly *bool) {
	o.SavedCartsOnly = savedCartsOnly
}

// WithSort adds the sort to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithSort(sort *string) *GetCartsUsingGETParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithUserID adds the userID to the get carts using g e t params
func (o *GetCartsUsingGETParams) WithUserID(userID string) *GetCartsUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get carts using g e t params
func (o *GetCartsUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCartsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
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

	if o.SavedCartsOnly != nil {

		// query param savedCartsOnly
		var qrSavedCartsOnly bool
		if o.SavedCartsOnly != nil {
			qrSavedCartsOnly = *o.SavedCartsOnly
		}
		qSavedCartsOnly := swag.FormatBool(qrSavedCartsOnly)
		if qSavedCartsOnly != "" {
			if err := r.SetQueryParam("savedCartsOnly", qSavedCartsOnly); err != nil {
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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
