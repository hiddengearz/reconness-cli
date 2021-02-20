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
	"github.com/go-openapi/validate"
)

// RootDomainDto root domain dto
//
// swagger:model RootDomainDto
type RootDomainDto struct {

	// agents ran before
	AgentsRanBefore *string `json:"agentsRanBefore,omitempty"`

	// id
	// Format: uuid
	ID *strfmt.UUID `json:"id,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// notes
	Notes *string `json:"notes,omitempty"`

	// subdomains
	Subdomains []*SubdomainDto `json:"subdomains"`
}

// Validate validates this root domain dto
func (m *RootDomainDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubdomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RootDomainDto) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RootDomainDto) validateSubdomains(formats strfmt.Registry) error {
	if swag.IsZero(m.Subdomains) { // not required
		return nil
	}

	for i := 0; i < len(m.Subdomains); i++ {
		if swag.IsZero(m.Subdomains[i]) { // not required
			continue
		}

		if m.Subdomains[i] != nil {
			if err := m.Subdomains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subdomains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this root domain dto based on the context it is used
func (m *RootDomainDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubdomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RootDomainDto) contextValidateSubdomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Subdomains); i++ {

		if m.Subdomains[i] != nil {
			if err := m.Subdomains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subdomains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RootDomainDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RootDomainDto) UnmarshalBinary(b []byte) error {
	var res RootDomainDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}