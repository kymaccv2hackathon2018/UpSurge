// Code generated by go-swagger; DO NOT EDIT.

package consents

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

// NewGetConsentTemplateUsingGETParams creates a new GetConsentTemplateUsingGETParams object
// with the default values initialized.
func NewGetConsentTemplateUsingGETParams() *GetConsentTemplateUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetConsentTemplateUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConsentTemplateUsingGETParamsWithTimeout creates a new GetConsentTemplateUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConsentTemplateUsingGETParamsWithTimeout(timeout time.Duration) *GetConsentTemplateUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetConsentTemplateUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetConsentTemplateUsingGETParamsWithContext creates a new GetConsentTemplateUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConsentTemplateUsingGETParamsWithContext(ctx context.Context) *GetConsentTemplateUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetConsentTemplateUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetConsentTemplateUsingGETParamsWithHTTPClient creates a new GetConsentTemplateUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConsentTemplateUsingGETParamsWithHTTPClient(client *http.Client) *GetConsentTemplateUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetConsentTemplateUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetConsentTemplateUsingGETParams contains all the parameters to send to the API endpoint
for the get consent template using g e t operation typically these are written to a http.Request
*/
type GetConsentTemplateUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*ConsentTemplateID
	  Consent template id.

	*/
	ConsentTemplateID string
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

// WithTimeout adds the timeout to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) WithTimeout(timeout time.Duration) *GetConsentTemplateUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) WithContext(ctx context.Context) *GetConsentTemplateUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) WithHTTPClient(client *http.Client) *GetConsentTemplateUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) WithBaseSiteID(baseSiteID string) *GetConsentTemplateUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithConsentTemplateID adds the consentTemplateID to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) WithConsentTemplateID(consentTemplateID string) *GetConsentTemplateUsingGETParams {
	o.SetConsentTemplateID(consentTemplateID)
	return o
}

// SetConsentTemplateID adds the consentTemplateId to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) SetConsentTemplateID(consentTemplateID string) {
	o.ConsentTemplateID = consentTemplateID
}

// WithFields adds the fields to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) WithFields(fields *string) *GetConsentTemplateUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) WithUserID(userID string) *GetConsentTemplateUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get consent template using g e t params
func (o *GetConsentTemplateUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConsentTemplateUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// path param consentTemplateId
	if err := r.SetPathParam("consentTemplateId", o.ConsentTemplateID); err != nil {
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