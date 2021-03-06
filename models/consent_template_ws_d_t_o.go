// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConsentTemplateWsDTO consent template ws d t o
// swagger:model ConsentTemplateWsDTO
type ConsentTemplateWsDTO struct {

	// current consent
	CurrentConsent *ConsentWsDTO `json:"currentConsent,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this consent template ws d t o
func (m *ConsentTemplateWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentConsent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentTemplateWsDTO) validateCurrentConsent(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentConsent) { // not required
		return nil
	}

	if m.CurrentConsent != nil {
		if err := m.CurrentConsent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentConsent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentTemplateWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentTemplateWsDTO) UnmarshalBinary(b []byte) error {
	var res ConsentTemplateWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
