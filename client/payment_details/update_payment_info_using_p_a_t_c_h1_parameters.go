// Code generated by go-swagger; DO NOT EDIT.

package payment_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
	"golang.org/x/net/context"
)

// NewUpdatePaymentInfoUsingPATCH1Params creates a new UpdatePaymentInfoUsingPATCH1Params object
// with the default values initialized.
func NewUpdatePaymentInfoUsingPATCH1Params() *UpdatePaymentInfoUsingPATCH1Params {
	var ()
	return &UpdatePaymentInfoUsingPATCH1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePaymentInfoUsingPATCH1ParamsWithTimeout creates a new UpdatePaymentInfoUsingPATCH1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePaymentInfoUsingPATCH1ParamsWithTimeout(timeout time.Duration) *UpdatePaymentInfoUsingPATCH1Params {
	var ()
	return &UpdatePaymentInfoUsingPATCH1Params{

		timeout: timeout,
	}
}

// NewUpdatePaymentInfoUsingPATCH1ParamsWithContext creates a new UpdatePaymentInfoUsingPATCH1Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePaymentInfoUsingPATCH1ParamsWithContext(ctx context.Context) *UpdatePaymentInfoUsingPATCH1Params {
	var ()
	return &UpdatePaymentInfoUsingPATCH1Params{

		Context: ctx,
	}
}

// NewUpdatePaymentInfoUsingPATCH1ParamsWithHTTPClient creates a new UpdatePaymentInfoUsingPATCH1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePaymentInfoUsingPATCH1ParamsWithHTTPClient(client *http.Client) *UpdatePaymentInfoUsingPATCH1Params {
	var ()
	return &UpdatePaymentInfoUsingPATCH1Params{
		HTTPClient: client,
	}
}

/*UpdatePaymentInfoUsingPATCH1Params contains all the parameters to send to the API endpoint
for the update payment info using p a t c h 1 operation typically these are written to a http.Request
*/
type UpdatePaymentInfoUsingPATCH1Params struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*PaymentDetails
	  Payment details object

	*/
	PaymentDetails *models.PaymentDetailsWsDTO
	/*PaymentDetailsID
	  Payment details identifier.

	*/
	PaymentDetailsID string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) WithTimeout(timeout time.Duration) *UpdatePaymentInfoUsingPATCH1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) WithContext(ctx context.Context) *UpdatePaymentInfoUsingPATCH1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) WithHTTPClient(client *http.Client) *UpdatePaymentInfoUsingPATCH1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) WithBaseSiteID(baseSiteID string) *UpdatePaymentInfoUsingPATCH1Params {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithPaymentDetails adds the paymentDetails to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) WithPaymentDetails(paymentDetails *models.PaymentDetailsWsDTO) *UpdatePaymentInfoUsingPATCH1Params {
	o.SetPaymentDetails(paymentDetails)
	return o
}

// SetPaymentDetails adds the paymentDetails to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) SetPaymentDetails(paymentDetails *models.PaymentDetailsWsDTO) {
	o.PaymentDetails = paymentDetails
}

// WithPaymentDetailsID adds the paymentDetailsID to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) WithPaymentDetailsID(paymentDetailsID string) *UpdatePaymentInfoUsingPATCH1Params {
	o.SetPaymentDetailsID(paymentDetailsID)
	return o
}

// SetPaymentDetailsID adds the paymentDetailsId to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) SetPaymentDetailsID(paymentDetailsID string) {
	o.PaymentDetailsID = paymentDetailsID
}

// WithUserID adds the userID to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) WithUserID(userID string) *UpdatePaymentInfoUsingPATCH1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update payment info using p a t c h 1 params
func (o *UpdatePaymentInfoUsingPATCH1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePaymentInfoUsingPATCH1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	if o.PaymentDetails != nil {
		if err := r.SetBodyParam(o.PaymentDetails); err != nil {
			return err
		}
	}

	// path param paymentDetailsId
	if err := r.SetPathParam("paymentDetailsId", o.PaymentDetailsID); err != nil {
		return err
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
