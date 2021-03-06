// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SuggestionWsDTO suggestion ws d t o
// swagger:model SuggestionWsDTO
type SuggestionWsDTO struct {

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this suggestion ws d t o
func (m *SuggestionWsDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SuggestionWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuggestionWsDTO) UnmarshalBinary(b []byte) error {
	var res SuggestionWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
