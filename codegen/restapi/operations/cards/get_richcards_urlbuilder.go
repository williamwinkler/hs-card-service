// Code generated by go-swagger; DO NOT EDIT.

package cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetRichcardsURL generates an URL for the get richcards operation
type GetRichcardsURL struct {
	Attack   *int64
	Class    *int64
	Health   *int64
	Limit    *int64
	ManaCost *int64
	Name     *string
	Page     *int64
	Rarity   *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRichcardsURL) WithBasePath(bp string) *GetRichcardsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRichcardsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetRichcardsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/richcards"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var attackQ string
	if o.Attack != nil {
		attackQ = swag.FormatInt64(*o.Attack)
	}
	if attackQ != "" {
		qs.Set("attack", attackQ)
	}

	var classQ string
	if o.Class != nil {
		classQ = swag.FormatInt64(*o.Class)
	}
	if classQ != "" {
		qs.Set("class", classQ)
	}

	var healthQ string
	if o.Health != nil {
		healthQ = swag.FormatInt64(*o.Health)
	}
	if healthQ != "" {
		qs.Set("health", healthQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var manaCostQ string
	if o.ManaCost != nil {
		manaCostQ = swag.FormatInt64(*o.ManaCost)
	}
	if manaCostQ != "" {
		qs.Set("manaCost", manaCostQ)
	}

	var nameQ string
	if o.Name != nil {
		nameQ = *o.Name
	}
	if nameQ != "" {
		qs.Set("name", nameQ)
	}

	var pageQ string
	if o.Page != nil {
		pageQ = swag.FormatInt64(*o.Page)
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	var rarityQ string
	if o.Rarity != nil {
		rarityQ = swag.FormatInt64(*o.Rarity)
	}
	if rarityQ != "" {
		qs.Set("rarity", rarityQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetRichcardsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetRichcardsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetRichcardsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetRichcardsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetRichcardsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetRichcardsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}