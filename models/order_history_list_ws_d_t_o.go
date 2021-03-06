// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrderHistoryListWsDTO order history list ws d t o
// swagger:model OrderHistoryListWsDTO
type OrderHistoryListWsDTO struct {

	// orders
	Orders []*OrderHistoryWsDTO `json:"orders"`

	// pagination
	Pagination *PaginationWsDTO `json:"pagination,omitempty"`

	// sorts
	Sorts []*SortWsDTO `json:"sorts"`
}

// Validate validates this order history list ws d t o
func (m *OrderHistoryListWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderHistoryListWsDTO) validateOrders(formats strfmt.Registry) error {

	if swag.IsZero(m.Orders) { // not required
		return nil
	}

	for i := 0; i < len(m.Orders); i++ {
		if swag.IsZero(m.Orders[i]) { // not required
			continue
		}

		if m.Orders[i] != nil {
			if err := m.Orders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderHistoryListWsDTO) validatePagination(formats strfmt.Registry) error {

	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *OrderHistoryListWsDTO) validateSorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Sorts) { // not required
		return nil
	}

	for i := 0; i < len(m.Sorts); i++ {
		if swag.IsZero(m.Sorts[i]) { // not required
			continue
		}

		if m.Sorts[i] != nil {
			if err := m.Sorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderHistoryListWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderHistoryListWsDTO) UnmarshalBinary(b []byte) error {
	var res OrderHistoryListWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
