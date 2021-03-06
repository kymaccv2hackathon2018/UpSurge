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
	"github.com/go-openapi/swag"
	"golang.org/x/net/context"
)

// NewGetSopPaymentRequestDetailsUsingGETParams creates a new GetSopPaymentRequestDetailsUsingGETParams object
// with the default values initialized.
func NewGetSopPaymentRequestDetailsUsingGETParams() *GetSopPaymentRequestDetailsUsingGETParams {
	var (
		extendedMerchantCallbackDefault = bool(false)
		fieldsDefault                   = string("DEFAULT")
	)
	return &GetSopPaymentRequestDetailsUsingGETParams{
		ExtendedMerchantCallback: &extendedMerchantCallbackDefault,
		Fields:                   &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSopPaymentRequestDetailsUsingGETParamsWithTimeout creates a new GetSopPaymentRequestDetailsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSopPaymentRequestDetailsUsingGETParamsWithTimeout(timeout time.Duration) *GetSopPaymentRequestDetailsUsingGETParams {
	var (
		extendedMerchantCallbackDefault = bool(false)
		fieldsDefault                   = string("DEFAULT")
	)
	return &GetSopPaymentRequestDetailsUsingGETParams{
		ExtendedMerchantCallback: &extendedMerchantCallbackDefault,
		Fields:                   &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetSopPaymentRequestDetailsUsingGETParamsWithContext creates a new GetSopPaymentRequestDetailsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSopPaymentRequestDetailsUsingGETParamsWithContext(ctx context.Context) *GetSopPaymentRequestDetailsUsingGETParams {
	var (
		extendedMerchantCallbackDefault = bool(false)
		fieldsDefault                   = string("DEFAULT")
	)
	return &GetSopPaymentRequestDetailsUsingGETParams{
		ExtendedMerchantCallback: &extendedMerchantCallbackDefault,
		Fields:                   &fieldsDefault,

		Context: ctx,
	}
}

// NewGetSopPaymentRequestDetailsUsingGETParamsWithHTTPClient creates a new GetSopPaymentRequestDetailsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSopPaymentRequestDetailsUsingGETParamsWithHTTPClient(client *http.Client) *GetSopPaymentRequestDetailsUsingGETParams {
	var (
		extendedMerchantCallbackDefault = bool(false)
		fieldsDefault                   = string("DEFAULT")
	)
	return &GetSopPaymentRequestDetailsUsingGETParams{
		ExtendedMerchantCallback: &extendedMerchantCallbackDefault,
		Fields:                   &fieldsDefault,
		HTTPClient:               client,
	}
}

/*GetSopPaymentRequestDetailsUsingGETParams contains all the parameters to send to the API endpoint
for the get sop payment request details using g e t operation typically these are written to a http.Request
*/
type GetSopPaymentRequestDetailsUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CartID
	  Cart identifier: cart code for logged in user, cart guid for anonymous user, 'current' for the last modified cart

	*/
	CartID string
	/*ExtendedMerchantCallback
	  Define which url should be returned

	*/
	ExtendedMerchantCallback *bool
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*ResponseURL
	  The URL that the payment provider uses to return payment information. Possible values for responseUrl include the following: “orderPage_cancelResponseURL”, “orderPage_declineResponseURL”, and “orderPage_receiptResponseURL”.

	*/
	ResponseURL string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithTimeout(timeout time.Duration) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithContext(ctx context.Context) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithHTTPClient(client *http.Client) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithBaseSiteID(baseSiteID string) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithCartID(cartID string) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithExtendedMerchantCallback adds the extendedMerchantCallback to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithExtendedMerchantCallback(extendedMerchantCallback *bool) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetExtendedMerchantCallback(extendedMerchantCallback)
	return o
}

// SetExtendedMerchantCallback adds the extendedMerchantCallback to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetExtendedMerchantCallback(extendedMerchantCallback *bool) {
	o.ExtendedMerchantCallback = extendedMerchantCallback
}

// WithFields adds the fields to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithFields(fields *string) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithResponseURL adds the responseURL to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithResponseURL(responseURL string) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetResponseURL(responseURL)
	return o
}

// SetResponseURL adds the responseUrl to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetResponseURL(responseURL string) {
	o.ResponseURL = responseURL
}

// WithUserID adds the userID to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) WithUserID(userID string) *GetSopPaymentRequestDetailsUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get sop payment request details using g e t params
func (o *GetSopPaymentRequestDetailsUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSopPaymentRequestDetailsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExtendedMerchantCallback != nil {

		// query param extendedMerchantCallback
		var qrExtendedMerchantCallback bool
		if o.ExtendedMerchantCallback != nil {
			qrExtendedMerchantCallback = *o.ExtendedMerchantCallback
		}
		qExtendedMerchantCallback := swag.FormatBool(qrExtendedMerchantCallback)
		if qExtendedMerchantCallback != "" {
			if err := r.SetQueryParam("extendedMerchantCallback", qExtendedMerchantCallback); err != nil {
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

	// query param responseUrl
	qrResponseURL := o.ResponseURL
	qResponseURL := qrResponseURL
	if qResponseURL != "" {
		if err := r.SetQueryParam("responseUrl", qResponseURL); err != nil {
			return err
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
