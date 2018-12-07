// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProductExpressUpdateElementWsDTO product express update element ws d t o
// swagger:model ProductExpressUpdateElementWsDTO
type ProductExpressUpdateElementWsDTO struct {

	// catalog Id
	CatalogID string `json:"catalogId,omitempty"`

	// catalog version
	CatalogVersion string `json:"catalogVersion,omitempty"`

	// code
	Code string `json:"code,omitempty"`
}

// Validate validates this product express update element ws d t o
func (m *ProductExpressUpdateElementWsDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProductExpressUpdateElementWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductExpressUpdateElementWsDTO) UnmarshalBinary(b []byte) error {
	var res ProductExpressUpdateElementWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
