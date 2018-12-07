// Code generated by go-swagger; DO NOT EDIT.

package address

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

// NewPutAddressUsingPUT1Params creates a new PutAddressUsingPUT1Params object
// with the default values initialized.
func NewPutAddressUsingPUT1Params() *PutAddressUsingPUT1Params {
	var ()
	return &PutAddressUsingPUT1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAddressUsingPUT1ParamsWithTimeout creates a new PutAddressUsingPUT1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAddressUsingPUT1ParamsWithTimeout(timeout time.Duration) *PutAddressUsingPUT1Params {
	var ()
	return &PutAddressUsingPUT1Params{

		timeout: timeout,
	}
}

// NewPutAddressUsingPUT1ParamsWithContext creates a new PutAddressUsingPUT1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPutAddressUsingPUT1ParamsWithContext(ctx context.Context) *PutAddressUsingPUT1Params {
	var ()
	return &PutAddressUsingPUT1Params{

		Context: ctx,
	}
}

// NewPutAddressUsingPUT1ParamsWithHTTPClient creates a new PutAddressUsingPUT1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAddressUsingPUT1ParamsWithHTTPClient(client *http.Client) *PutAddressUsingPUT1Params {
	var ()
	return &PutAddressUsingPUT1Params{
		HTTPClient: client,
	}
}

/*PutAddressUsingPUT1Params contains all the parameters to send to the API endpoint
for the put address using p u t 1 operation typically these are written to a http.Request
*/
type PutAddressUsingPUT1Params struct {

	/*Address
	  Address object.

	*/
	Address *models.AddressWsDTO
	/*AddressID
	  Address identifier.

	*/
	AddressID string
	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) WithTimeout(timeout time.Duration) *PutAddressUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) WithContext(ctx context.Context) *PutAddressUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) WithHTTPClient(client *http.Client) *PutAddressUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) WithAddress(address *models.AddressWsDTO) *PutAddressUsingPUT1Params {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) SetAddress(address *models.AddressWsDTO) {
	o.Address = address
}

// WithAddressID adds the addressID to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) WithAddressID(addressID string) *PutAddressUsingPUT1Params {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) SetAddressID(addressID string) {
	o.AddressID = addressID
}

// WithBaseSiteID adds the baseSiteID to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) WithBaseSiteID(baseSiteID string) *PutAddressUsingPUT1Params {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithUserID adds the userID to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) WithUserID(userID string) *PutAddressUsingPUT1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the put address using p u t 1 params
func (o *PutAddressUsingPUT1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAddressUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Address != nil {
		if err := r.SetBodyParam(o.Address); err != nil {
			return err
		}
	}

	// path param addressId
	if err := r.SetPathParam("addressId", o.AddressID); err != nil {
		return err
	}

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
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
