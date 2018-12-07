// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CurrencyWsDTO currency ws d t o
// swagger:model CurrencyWsDTO
type CurrencyWsDTO struct {

	// active
	Active bool `json:"active,omitempty"`

	// isocode
	Isocode string `json:"isocode,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// symbol
	Symbol string `json:"symbol,omitempty"`
}

// Validate validates this currency ws d t o
func (m *CurrencyWsDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CurrencyWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrencyWsDTO) UnmarshalBinary(b []byte) error {
	var res CurrencyWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
