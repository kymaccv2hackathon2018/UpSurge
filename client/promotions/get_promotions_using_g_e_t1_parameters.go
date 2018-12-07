// Code generated by go-swagger; DO NOT EDIT.

package promotions

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

// NewGetPromotionsUsingGET1Params creates a new GetPromotionsUsingGET1Params object
// with the default values initialized.
func NewGetPromotionsUsingGET1Params() *GetPromotionsUsingGET1Params {
	var (
		fieldsDefault = string("BASIC")
	)
	return &GetPromotionsUsingGET1Params{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPromotionsUsingGET1ParamsWithTimeout creates a new GetPromotionsUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPromotionsUsingGET1ParamsWithTimeout(timeout time.Duration) *GetPromotionsUsingGET1Params {
	var (
		fieldsDefault = string("BASIC")
	)
	return &GetPromotionsUsingGET1Params{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetPromotionsUsingGET1ParamsWithContext creates a new GetPromotionsUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetPromotionsUsingGET1ParamsWithContext(ctx context.Context) *GetPromotionsUsingGET1Params {
	var (
		fieldsDefault = string("BASIC")
	)
	return &GetPromotionsUsingGET1Params{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetPromotionsUsingGET1ParamsWithHTTPClient creates a new GetPromotionsUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPromotionsUsingGET1ParamsWithHTTPClient(client *http.Client) *GetPromotionsUsingGET1Params {
	var (
		fieldsDefault = string("BASIC")
	)
	return &GetPromotionsUsingGET1Params{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetPromotionsUsingGET1Params contains all the parameters to send to the API endpoint
for the get promotions using g e t 1 operation typically these are written to a http.Request
*/
type GetPromotionsUsingGET1Params struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*PromotionGroup
	  Only promotions from this group are returned

	*/
	PromotionGroup *string
	/*Type
	  Defines what type of promotions should be returned. Values supported for that parameter are: <ul><li>all: All available promotions are returned</li><li>product: Only product promotions are returned</li><li>order: Only order promotions are returned</li></ul>

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) WithTimeout(timeout time.Duration) *GetPromotionsUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) WithContext(ctx context.Context) *GetPromotionsUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) WithHTTPClient(client *http.Client) *GetPromotionsUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) WithBaseSiteID(baseSiteID string) *GetPromotionsUsingGET1Params {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithFields adds the fields to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) WithFields(fields *string) *GetPromotionsUsingGET1Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) SetFields(fields *string) {
	o.Fields = fields
}

// WithPromotionGroup adds the promotionGroup to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) WithPromotionGroup(promotionGroup *string) *GetPromotionsUsingGET1Params {
	o.SetPromotionGroup(promotionGroup)
	return o
}

// SetPromotionGroup adds the promotionGroup to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) SetPromotionGroup(promotionGroup *string) {
	o.PromotionGroup = promotionGroup
}

// WithType adds the typeVar to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) WithType(typeVar string) *GetPromotionsUsingGET1Params {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get promotions using g e t 1 params
func (o *GetPromotionsUsingGET1Params) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPromotionsUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
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

	if o.PromotionGroup != nil {

		// query param promotionGroup
		var qrPromotionGroup string
		if o.PromotionGroup != nil {
			qrPromotionGroup = *o.PromotionGroup
		}
		qPromotionGroup := qrPromotionGroup
		if qPromotionGroup != "" {
			if err := r.SetQueryParam("promotionGroup", qPromotionGroup); err != nil {
				return err
			}
		}

	}

	// query param type
	qrType := o.Type
	qType := qrType
	if qType != "" {
		if err := r.SetQueryParam("type", qType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
