// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BreadcrumbWsDTO breadcrumb ws d t o
// swagger:model BreadcrumbWsDTO
type BreadcrumbWsDTO struct {

	// facet code
	FacetCode string `json:"facetCode,omitempty"`

	// facet name
	FacetName string `json:"facetName,omitempty"`

	// facet value code
	FacetValueCode string `json:"facetValueCode,omitempty"`

	// facet value name
	FacetValueName string `json:"facetValueName,omitempty"`

	// remove query
	RemoveQuery *SearchStateWsDTO `json:"removeQuery,omitempty"`

	// truncate query
	TruncateQuery *SearchStateWsDTO `json:"truncateQuery,omitempty"`
}

// Validate validates this breadcrumb ws d t o
func (m *BreadcrumbWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoveQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTruncateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BreadcrumbWsDTO) validateRemoveQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.RemoveQuery) { // not required
		return nil
	}

	if m.RemoveQuery != nil {
		if err := m.RemoveQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("removeQuery")
			}
			return err
		}
	}

	return nil
}

func (m *BreadcrumbWsDTO) validateTruncateQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.TruncateQuery) { // not required
		return nil
	}

	if m.TruncateQuery != nil {
		if err := m.TruncateQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("truncateQuery")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BreadcrumbWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BreadcrumbWsDTO) UnmarshalBinary(b []byte) error {
	var res BreadcrumbWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}