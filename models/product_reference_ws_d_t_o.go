// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProductReferenceWsDTO product reference ws d t o
// swagger:model ProductReferenceWsDTO
type ProductReferenceWsDTO struct {

	// description
	Description string `json:"description,omitempty"`

	// preselected
	Preselected bool `json:"preselected,omitempty"`

	// quantity
	Quantity int32 `json:"quantity,omitempty"`

	// reference type
	ReferenceType string `json:"referenceType,omitempty"`

	// target
	Target *ProductWsDTO `json:"target,omitempty"`
}

// Validate validates this product reference ws d t o
func (m *ProductReferenceWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductReferenceWsDTO) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductReferenceWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductReferenceWsDTO) UnmarshalBinary(b []byte) error {
	var res ProductReferenceWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
