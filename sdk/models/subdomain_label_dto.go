// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubdomainLabelDto subdomain label dto
//
// swagger:model SubdomainLabelDto
type SubdomainLabelDto struct {

	// label
	Label *string `json:"label,omitempty"`
}

// Validate validates this subdomain label dto
func (m *SubdomainLabelDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subdomain label dto based on context it is used
func (m *SubdomainLabelDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubdomainLabelDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubdomainLabelDto) UnmarshalBinary(b []byte) error {
	var res SubdomainLabelDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}