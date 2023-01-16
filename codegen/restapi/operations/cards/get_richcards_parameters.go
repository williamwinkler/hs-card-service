// Code generated by go-swagger; DO NOT EDIT.

package cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetRichcardsParams creates a new GetRichcardsParams object
// with the default values initialized.
func NewGetRichcardsParams() GetRichcardsParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(20)

		pageDefault = int64(1)
	)

	return GetRichcardsParams{
		Limit: &limitDefault,

		Page: &pageDefault,
	}
}

// GetRichcardsParams contains all the bound params for the get richcards operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetRichcards
type GetRichcardsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Minimum: 0
	  In: query
	*/
	Attack *int64
	/*
	  Maximum: 14
	  Minimum: 1
	  In: query
	*/
	Class *int64
	/*
	  Minimum: 0
	  In: query
	*/
	Health *int64
	/*
	  Maximum: 100
	  Minimum: 1
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*
	  Minimum: 0
	  In: query
	*/
	ManaCost *int64
	/*
	  Min Length: 1
	  In: query
	*/
	Name *string
	/*
	  Minimum: 1
	  In: query
	  Default: 1
	*/
	Page *int64
	/*
	  Maximum: 5
	  Minimum: 1
	  In: query
	*/
	Rarity *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRichcardsParams() beforehand.
func (o *GetRichcardsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAttack, qhkAttack, _ := qs.GetOK("attack")
	if err := o.bindAttack(qAttack, qhkAttack, route.Formats); err != nil {
		res = append(res, err)
	}

	qClass, qhkClass, _ := qs.GetOK("class")
	if err := o.bindClass(qClass, qhkClass, route.Formats); err != nil {
		res = append(res, err)
	}

	qHealth, qhkHealth, _ := qs.GetOK("health")
	if err := o.bindHealth(qHealth, qhkHealth, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qManaCost, qhkManaCost, _ := qs.GetOK("manaCost")
	if err := o.bindManaCost(qManaCost, qhkManaCost, route.Formats); err != nil {
		res = append(res, err)
	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qRarity, qhkRarity, _ := qs.GetOK("rarity")
	if err := o.bindRarity(qRarity, qhkRarity, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAttack binds and validates parameter Attack from query.
func (o *GetRichcardsParams) bindAttack(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("attack", "query", "int64", raw)
	}
	o.Attack = &value

	if err := o.validateAttack(formats); err != nil {
		return err
	}

	return nil
}

// validateAttack carries on validations for parameter Attack
func (o *GetRichcardsParams) validateAttack(formats strfmt.Registry) error {

	if err := validate.MinimumInt("attack", "query", *o.Attack, 0, false); err != nil {
		return err
	}

	return nil
}

// bindClass binds and validates parameter Class from query.
func (o *GetRichcardsParams) bindClass(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("class", "query", "int64", raw)
	}
	o.Class = &value

	if err := o.validateClass(formats); err != nil {
		return err
	}

	return nil
}

// validateClass carries on validations for parameter Class
func (o *GetRichcardsParams) validateClass(formats strfmt.Registry) error {

	if err := validate.MinimumInt("class", "query", *o.Class, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("class", "query", *o.Class, 14, false); err != nil {
		return err
	}

	return nil
}

// bindHealth binds and validates parameter Health from query.
func (o *GetRichcardsParams) bindHealth(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("health", "query", "int64", raw)
	}
	o.Health = &value

	if err := o.validateHealth(formats); err != nil {
		return err
	}

	return nil
}

// validateHealth carries on validations for parameter Health
func (o *GetRichcardsParams) validateHealth(formats strfmt.Registry) error {

	if err := validate.MinimumInt("health", "query", *o.Health, 0, false); err != nil {
		return err
	}

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetRichcardsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRichcardsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetRichcardsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", *o.Limit, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "query", *o.Limit, 100, false); err != nil {
		return err
	}

	return nil
}

// bindManaCost binds and validates parameter ManaCost from query.
func (o *GetRichcardsParams) bindManaCost(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("manaCost", "query", "int64", raw)
	}
	o.ManaCost = &value

	if err := o.validateManaCost(formats); err != nil {
		return err
	}

	return nil
}

// validateManaCost carries on validations for parameter ManaCost
func (o *GetRichcardsParams) validateManaCost(formats strfmt.Registry) error {

	if err := validate.MinimumInt("manaCost", "query", *o.ManaCost, 0, false); err != nil {
		return err
	}

	return nil
}

// bindName binds and validates parameter Name from query.
func (o *GetRichcardsParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Name = &raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

// validateName carries on validations for parameter Name
func (o *GetRichcardsParams) validateName(formats strfmt.Registry) error {

	if err := validate.MinLength("name", "query", *o.Name, 1); err != nil {
		return err
	}

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetRichcardsParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRichcardsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	if err := o.validatePage(formats); err != nil {
		return err
	}

	return nil
}

// validatePage carries on validations for parameter Page
func (o *GetRichcardsParams) validatePage(formats strfmt.Registry) error {

	if err := validate.MinimumInt("page", "query", *o.Page, 1, false); err != nil {
		return err
	}

	return nil
}

// bindRarity binds and validates parameter Rarity from query.
func (o *GetRichcardsParams) bindRarity(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("rarity", "query", "int64", raw)
	}
	o.Rarity = &value

	if err := o.validateRarity(formats); err != nil {
		return err
	}

	return nil
}

// validateRarity carries on validations for parameter Rarity
func (o *GetRichcardsParams) validateRarity(formats strfmt.Registry) error {

	if err := validate.MinimumInt("rarity", "query", *o.Rarity, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rarity", "query", *o.Rarity, 5, false); err != nil {
		return err
	}

	return nil
}
