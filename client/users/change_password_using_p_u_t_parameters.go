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

// NewChangePasswordUsingPUTParams creates a new ChangePasswordUsingPUTParams object
// with the default values initialized.
func NewChangePasswordUsingPUTParams() *ChangePasswordUsingPUTParams {
	var ()
	return &ChangePasswordUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangePasswordUsingPUTParamsWithTimeout creates a new ChangePasswordUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangePasswordUsingPUTParamsWithTimeout(timeout time.Duration) *ChangePasswordUsingPUTParams {
	var ()
	return &ChangePasswordUsingPUTParams{

		timeout: timeout,
	}
}

// NewChangePasswordUsingPUTParamsWithContext creates a new ChangePasswordUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangePasswordUsingPUTParamsWithContext(ctx context.Context) *ChangePasswordUsingPUTParams {
	var ()
	return &ChangePasswordUsingPUTParams{

		Context: ctx,
	}
}

// NewChangePasswordUsingPUTParamsWithHTTPClient creates a new ChangePasswordUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangePasswordUsingPUTParamsWithHTTPClient(client *http.Client) *ChangePasswordUsingPUTParams {
	var ()
	return &ChangePasswordUsingPUTParams{
		HTTPClient: client,
	}
}

/*ChangePasswordUsingPUTParams contains all the parameters to send to the API endpoint
for the change password using p u t operation typically these are written to a http.Request
*/
type ChangePasswordUsingPUTParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*New
	  New password.

	*/
	New string
	/*Old
	  Old password. Required only for ROLE_CUSTOMERGROUP

	*/
	Old *string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) WithTimeout(timeout time.Duration) *ChangePasswordUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) WithContext(ctx context.Context) *ChangePasswordUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) WithHTTPClient(client *http.Client) *ChangePasswordUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) WithBaseSiteID(baseSiteID string) *ChangePasswordUsingPUTParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithNew adds the new to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) WithNew(new string) *ChangePasswordUsingPUTParams {
	o.SetNew(new)
	return o
}

// SetNew adds the new to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) SetNew(new string) {
	o.New = new
}

// WithOld adds the old to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) WithOld(old *string) *ChangePasswordUsingPUTParams {
	o.SetOld(old)
	return o
}

// SetOld adds the old to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) SetOld(old *string) {
	o.Old = old
}

// WithUserID adds the userID to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) WithUserID(userID string) *ChangePasswordUsingPUTParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the change password using p u t params
func (o *ChangePasswordUsingPUTParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangePasswordUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// query param new
	qrNew := o.New
	qNew := qrNew
	if qNew != "" {
		if err := r.SetQueryParam("new", qNew); err != nil {
			return err
		}
	}

	if o.Old != nil {

		// query param old
		var qrOld string
		if o.Old != nil {
			qrOld = *o.Old
		}
		qOld := qrOld
		if qOld != "" {
			if err := r.SetQueryParam("old", qOld); err != nil {
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