// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Card card
//
// swagger:model card
type Card struct {

	// artist name
	ArtistName string `json:"artistName"`

	// attack
	Attack int64 `json:"attack"`

	// card set Id
	CardSetID int64 `json:"cardSetId"`

	// card type Id
	CardTypeID int64 `json:"cardTypeId"`

	// class Id
	ClassID int64 `json:"classId"`

	// collectible
	Collectible int64 `json:"collectible"`

	// duals
	Duals *Duals `json:"duals,omitempty"`

	// flavor text
	FlavorText string `json:"flavorText,omitempty"`

	// health
	Health int64 `json:"health"`

	// This is the ID from blizzards API
	ID int64 `json:"id,omitempty"`

	// Links to a png-image of the card
	Image string `json:"image,omitempty"`

	// Links to a png-image of the golden version of the card
	ImageGold string `json:"imageGold,omitempty"`

	// keyword ids
	KeywordIds []int64 `json:"keywordIds"`

	// mana cost
	ManaCost int64 `json:"manaCost"`

	// name
	Name string `json:"name,omitempty"`

	// parent Id
	ParentID int64 `json:"parentId"`

	// rarity Id
	RarityID int64 `json:"rarityId"`

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this card
func (m *Card) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Card) validateDuals(formats strfmt.Registry) error {
	if swag.IsZero(m.Duals) { // not required
		return nil
	}

	if m.Duals != nil {
		if err := m.Duals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("duals")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this card based on the context it is used
func (m *Card) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDuals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Card) contextValidateDuals(ctx context.Context, formats strfmt.Registry) error {

	if m.Duals != nil {
		if err := m.Duals.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duals")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("duals")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Card) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Card) UnmarshalBinary(b []byte) error {
	var res Card
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
