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

// DeliveryModeListWsDTO delivery mode list ws d t o
// swagger:model DeliveryModeListWsDTO
type DeliveryModeListWsDTO struct {

	// delivery modes
	DeliveryModes []*DeliveryModeWsDTO `json:"deliveryModes"`
}

// Validate validates this delivery mode list ws d t o
func (m *DeliveryModeListWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeliveryModes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryModeListWsDTO) validateDeliveryModes(formats strfmt.Registry) error {

	if swag.IsZero(m.DeliveryModes) { // not required
		return nil
	}

	for i := 0; i < len(m.DeliveryModes); i++ {
		if swag.IsZero(m.DeliveryModes[i]) { // not required
			continue
		}

		if m.DeliveryModes[i] != nil {
			if err := m.DeliveryModes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deliveryModes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeliveryModeListWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeliveryModeListWsDTO) UnmarshalBinary(b []byte) error {
	var res DeliveryModeListWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}