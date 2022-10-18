// Code generated by go-swagger; DO NOT EDIT.

package beers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAllBeersParams creates a new GetAllBeersParams object
// with the default values initialized.
func NewGetAllBeersParams() GetAllBeersParams {

	var (
		// initialize parameters with default values

		limitDefault = int32(10)
	)

	return GetAllBeersParams{
		Limit: &limitDefault,
	}
}

// GetAllBeersParams contains all the bound params for the get all beers operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAllBeers
type GetAllBeersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: 10
	*/
	Limit *int32
	/*
	  In: query
	*/
	Origin *string
	/*
	  In: query
	*/
	Title *string
	/*
	  In: query
	*/
	VolumePercentage *float32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAllBeersParams() beforehand.
func (o *GetAllBeersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrigin, qhkOrigin, _ := qs.GetOK("origin")
	if err := o.bindOrigin(qOrigin, qhkOrigin, route.Formats); err != nil {
		res = append(res, err)
	}

	qTitle, qhkTitle, _ := qs.GetOK("title")
	if err := o.bindTitle(qTitle, qhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	qVolumePercentage, qhkVolumePercentage, _ := qs.GetOK("volume-percentage")
	if err := o.bindVolumePercentage(qVolumePercentage, qhkVolumePercentage, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetAllBeersParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAllBeersParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int32", raw)
	}
	o.Limit = &value

	return nil
}

// bindOrigin binds and validates parameter Origin from query.
func (o *GetAllBeersParams) bindOrigin(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Origin = &raw

	return nil
}

// bindTitle binds and validates parameter Title from query.
func (o *GetAllBeersParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Title = &raw

	return nil
}

// bindVolumePercentage binds and validates parameter VolumePercentage from query.
func (o *GetAllBeersParams) bindVolumePercentage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat32(raw)
	if err != nil {
		return errors.InvalidType("volume-percentage", "query", "float32", raw)
	}
	o.VolumePercentage = &value

	return nil
}
