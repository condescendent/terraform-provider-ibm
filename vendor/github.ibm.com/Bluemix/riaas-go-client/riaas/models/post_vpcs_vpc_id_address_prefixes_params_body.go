// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostVpcsVpcIDAddressPrefixesParamsBody post vpcs vpc Id address prefixes params body
// swagger:model postVpcsVpcIdAddressPrefixesParamsBody
type PostVpcsVpcIDAddressPrefixesParamsBody struct {

	// The CIDR block for this prefix.
	Cidr string `json:"cidr,omitempty"`

	// The user-defined name for this prefix. By default, the base IP address will be the name. For example, for 10.0.0.0/24 the name will be 10.0.0.0.
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// zone
	Zone *PostVpcsVpcIDAddressPrefixesParamsBodyZone `json:"zone,omitempty"`
}

// Validate validates this post vpcs vpc Id address prefixes params body
func (m *PostVpcsVpcIDAddressPrefixesParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostVpcsVpcIDAddressPrefixesParamsBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostVpcsVpcIDAddressPrefixesParamsBody) validateZone(formats strfmt.Registry) error {

	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {
		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostVpcsVpcIDAddressPrefixesParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostVpcsVpcIDAddressPrefixesParamsBody) UnmarshalBinary(b []byte) error {
	var res PostVpcsVpcIDAddressPrefixesParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}