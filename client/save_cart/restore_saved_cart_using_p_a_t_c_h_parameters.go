// Code generated by go-swagger; DO NOT EDIT.

package save_cart

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

// NewRestoreSavedCartUsingPATCHParams creates a new RestoreSavedCartUsingPATCHParams object
// with the default values initialized.
func NewRestoreSavedCartUsingPATCHParams() *RestoreSavedCartUsingPATCHParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &RestoreSavedCartUsingPATCHParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreSavedCartUsingPATCHParamsWithTimeout creates a new RestoreSavedCartUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestoreSavedCartUsingPATCHParamsWithTimeout(timeout time.Duration) *RestoreSavedCartUsingPATCHParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &RestoreSavedCartUsingPATCHParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewRestoreSavedCartUsingPATCHParamsWithContext creates a new RestoreSavedCartUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestoreSavedCartUsingPATCHParamsWithContext(ctx context.Context) *RestoreSavedCartUsingPATCHParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &RestoreSavedCartUsingPATCHParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewRestoreSavedCartUsingPATCHParamsWithHTTPClient creates a new RestoreSavedCartUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestoreSavedCartUsingPATCHParamsWithHTTPClient(client *http.Client) *RestoreSavedCartUsingPATCHParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &RestoreSavedCartUsingPATCHParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*RestoreSavedCartUsingPATCHParams contains all the parameters to send to the API endpoint
for the restore saved cart using p a t c h operation typically these are written to a http.Request
*/
type RestoreSavedCartUsingPATCHParams struct {

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

// WithTimeout adds the timeout to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) WithTimeout(timeout time.Duration) *RestoreSavedCartUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) WithContext(ctx context.Context) *RestoreSavedCartUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) WithHTTPClient(client *http.Client) *RestoreSavedCartUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) WithBaseSiteID(baseSiteID string) *RestoreSavedCartUsingPATCHParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) WithCartID(cartID string) *RestoreSavedCartUsingPATCHParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithFields adds the fields to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) WithFields(fields *string) *RestoreSavedCartUsingPATCHParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) WithUserID(userID string) *RestoreSavedCartUsingPATCHParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the restore saved cart using p a t c h params
func (o *RestoreSavedCartUsingPATCHParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreSavedCartUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
