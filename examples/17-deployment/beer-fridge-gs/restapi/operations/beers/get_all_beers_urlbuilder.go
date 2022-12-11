// Code generated by go-swagger; DO NOT EDIT.

package beers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetAllBeersURL generates an URL for the get all beers operation
type GetAllBeersURL struct {
	Limit            *int32
	Origin           *string
	Title            *string
	VolumePercentage *float32

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAllBeersURL) WithBasePath(bp string) *GetAllBeersURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAllBeersURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetAllBeersURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/beers"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt32(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var originQ string
	if o.Origin != nil {
		originQ = *o.Origin
	}
	if originQ != "" {
		qs.Set("origin", originQ)
	}

	var titleQ string
	if o.Title != nil {
		titleQ = *o.Title
	}
	if titleQ != "" {
		qs.Set("title", titleQ)
	}

	var volumePercentageQ string
	if o.VolumePercentage != nil {
		volumePercentageQ = swag.FormatFloat32(*o.VolumePercentage)
	}
	if volumePercentageQ != "" {
		qs.Set("volume-percentage", volumePercentageQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetAllBeersURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetAllBeersURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetAllBeersURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAllBeersURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAllBeersURL")
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
func (o *GetAllBeersURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
