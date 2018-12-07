// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewChangeLoginUsingPUTParams creates a new ChangeLoginUsingPUTParams object
// with the default values initialized.
func NewChangeLoginUsingPUTParams() *ChangeLoginUsingPUTParams {
	var ()
	return &ChangeLoginUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeLoginUsingPUTParamsWithTimeout creates a new ChangeLoginUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeLoginUsingPUTParamsWithTimeout(timeout time.Duration) *ChangeLoginUsingPUTParams {
	var ()
	return &ChangeLoginUsingPUTParams{

		timeout: timeout,
	}
}

// NewChangeLoginUsingPUTParamsWithContext creates a new ChangeLoginUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeLoginUsingPUTParamsWithContext(ctx context.Context) *ChangeLoginUsingPUTParams {
	var ()
	return &ChangeLoginUsingPUTParams{

		Context: ctx,
	}
}

// NewChangeLoginUsingPUTParamsWithHTTPClient creates a new ChangeLoginUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeLoginUsingPUTParamsWithHTTPClient(client *http.Client) *ChangeLoginUsingPUTParams {
	var ()
	return &ChangeLoginUsingPUTParams{
		HTTPClient: client,
	}
}

/*ChangeLoginUsingPUTParams contains all the parameters to send to the API endpoint
for the change login using p u t operation typically these are written to a http.Request
*/
type ChangeLoginUsingPUTParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*NewLogin
	  Customer's new login name. Customer login is case insensitive.

	*/
	NewLogin string
	/*Password
	  Customer's current password.

	*/
	Password string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) WithTimeout(timeout time.Duration) *ChangeLoginUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) WithContext(ctx context.Context) *ChangeLoginUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) WithHTTPClient(client *http.Client) *ChangeLoginUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) WithBaseSiteID(baseSiteID string) *ChangeLoginUsingPUTParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithNewLogin adds the newLogin to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) WithNewLogin(newLogin string) *ChangeLoginUsingPUTParams {
	o.SetNewLogin(newLogin)
	return o
}

// SetNewLogin adds the newLogin to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) SetNewLogin(newLogin string) {
	o.NewLogin = newLogin
}

// WithPassword adds the password to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) WithPassword(password string) *ChangeLoginUsingPUTParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) SetPassword(password string) {
	o.Password = password
}

// WithUserID adds the userID to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) WithUserID(userID string) *ChangeLoginUsingPUTParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the change login using p u t params
func (o *ChangeLoginUsingPUTParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeLoginUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// query param newLogin
	qrNewLogin := o.NewLogin
	qNewLogin := qrNewLogin
	if qNewLogin != "" {
		if err := r.SetQueryParam("newLogin", qNewLogin); err != nil {
			return err
		}
	}

	// query param password
	qrPassword := o.Password
	qPassword := qrPassword
	if qPassword != "" {
		if err := r.SetQueryParam("password", qPassword); err != nil {
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
