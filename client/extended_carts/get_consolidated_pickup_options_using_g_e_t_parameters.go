// Code generated by go-swagger; DO NOT EDIT.

package extended_carts

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

// NewGetConsolidatedPickupOptionsUsingGETParams creates a new GetConsolidatedPickupOptionsUsingGETParams object
// with the default values initialized.
func NewGetConsolidatedPickupOptionsUsingGETParams() *GetConsolidatedPickupOptionsUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetConsolidatedPickupOptionsUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConsolidatedPickupOptionsUsingGETParamsWithTimeout creates a new GetConsolidatedPickupOptionsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConsolidatedPickupOptionsUsingGETParamsWithTimeout(timeout time.Duration) *GetConsolidatedPickupOptionsUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetConsolidatedPickupOptionsUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetConsolidatedPickupOptionsUsingGETParamsWithContext creates a new GetConsolidatedPickupOptionsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConsolidatedPickupOptionsUsingGETParamsWithContext(ctx context.Context) *GetConsolidatedPickupOptionsUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetConsolidatedPickupOptionsUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetConsolidatedPickupOptionsUsingGETParamsWithHTTPClient creates a new GetConsolidatedPickupOptionsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConsolidatedPickupOptionsUsingGETParamsWithHTTPClient(client *http.Client) *GetConsolidatedPickupOptionsUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetConsolidatedPickupOptionsUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetConsolidatedPickupOptionsUsingGETParams contains all the parameters to send to the API endpoint
for the get consolidated pickup options using g e t operation typically these are written to a http.Request
*/
type GetConsolidatedPickupOptionsUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CartID
	  Cart identifier: cart code for logged in user, cart guid for anonymous user, 'current' for the last modified cart

	*/
	CartID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) WithTimeout(timeout time.Duration) *GetConsolidatedPickupOptionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) WithContext(ctx context.Context) *GetConsolidatedPickupOptionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) WithHTTPClient(client *http.Client) *GetConsolidatedPickupOptionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) WithBaseSiteID(baseSiteID string) *GetConsolidatedPickupOptionsUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) WithCartID(cartID string) *GetConsolidatedPickupOptionsUsingGETParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithFields adds the fields to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) WithFields(fields *string) *GetConsolidatedPickupOptionsUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) WithUserID(userID string) *GetConsolidatedPickupOptionsUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get consolidated pickup options using g e t params
func (o *GetConsolidatedPickupOptionsUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConsolidatedPickupOptionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
