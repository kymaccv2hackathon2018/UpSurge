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

// NewGetSopPaymentResponseUsingGETParams creates a new GetSopPaymentResponseUsingGETParams object
// with the default values initialized.
func NewGetSopPaymentResponseUsingGETParams() *GetSopPaymentResponseUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetSopPaymentResponseUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSopPaymentResponseUsingGETParamsWithTimeout creates a new GetSopPaymentResponseUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSopPaymentResponseUsingGETParamsWithTimeout(timeout time.Duration) *GetSopPaymentResponseUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetSopPaymentResponseUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetSopPaymentResponseUsingGETParamsWithContext creates a new GetSopPaymentResponseUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSopPaymentResponseUsingGETParamsWithContext(ctx context.Context) *GetSopPaymentResponseUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetSopPaymentResponseUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetSopPaymentResponseUsingGETParamsWithHTTPClient creates a new GetSopPaymentResponseUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSopPaymentResponseUsingGETParamsWithHTTPClient(client *http.Client) *GetSopPaymentResponseUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetSopPaymentResponseUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetSopPaymentResponseUsingGETParams contains all the parameters to send to the API endpoint
for the get sop payment response using g e t operation typically these are written to a http.Request
*/
type GetSopPaymentResponseUsingGETParams struct {

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

// WithTimeout adds the timeout to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) WithTimeout(timeout time.Duration) *GetSopPaymentResponseUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) WithContext(ctx context.Context) *GetSopPaymentResponseUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) WithHTTPClient(client *http.Client) *GetSopPaymentResponseUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) WithBaseSiteID(baseSiteID string) *GetSopPaymentResponseUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) WithCartID(cartID string) *GetSopPaymentResponseUsingGETParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithFields adds the fields to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) WithFields(fields *string) *GetSopPaymentResponseUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) WithUserID(userID string) *GetSopPaymentResponseUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get sop payment response using g e t params
func (o *GetSopPaymentResponseUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSopPaymentResponseUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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