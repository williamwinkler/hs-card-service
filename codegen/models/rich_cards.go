// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RichCards rich cards
//
// swagger:model richCards
type RichCards struct {

	// card count
	CardCount int64 `json:"cardCount,omitempty"`

	// if there a no cards, the array is null
	// Example: []
	Cards []*RichCard `json:"cards"`

	// page
	Page int64 `json:"page,omitempty"`

	// page count
	PageCount int64 `json:"pageCount,omitempty"`
}

// Validate validates this rich cards
func (m *RichCards) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RichCards) validateCards(formats strfmt.Registry) error {
	if swag.IsZero(m.Cards) { // not required
		return nil
	}

	for i := 0; i < len(m.Cards); i++ {
		if swag.IsZero(m.Cards[i]) { // not required
			continue
		}

		if m.Cards[i] != nil {
			if err := m.Cards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rich cards based on the context it is used
func (m *RichCards) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RichCards) contextValidateCards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cards); i++ {

		if m.Cards[i] != nil {
			if err := m.Cards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RichCards) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RichCards) UnmarshalBinary(b []byte) error {
	var res RichCards
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
