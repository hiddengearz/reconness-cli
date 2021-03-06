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

// TargetDto target dto
//
// swagger:model TargetDto
type TargetDto struct {

	// agents ran before
	AgentsRanBefore *string `json:"agentsRanBefore,omitempty"`

	// bug bounty program Url
	BugBountyProgramURL *string `json:"bugBountyProgramUrl,omitempty"`

	// id
	// Format: uuid
	ID *strfmt.UUID `json:"id,omitempty"`

	// in scope
	InScope *string `json:"inScope,omitempty"`

	// is private
	IsPrivate bool `json:"isPrivate,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// out of scope
	OutOfScope *string `json:"outOfScope,omitempty"`

	// root domains
	RootDomains []*RootDomainDto `json:"rootDomains"`
}

// Validate validates this target dto
func (m *TargetDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetDto) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TargetDto) validateRootDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.RootDomains) { // not required
		return nil
	}

	for i := 0; i < len(m.RootDomains); i++ {
		if swag.IsZero(m.RootDomains[i]) { // not required
			continue
		}

		if m.RootDomains[i] != nil {
			if err := m.RootDomains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rootDomains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this target dto based on the context it is used
func (m *TargetDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRootDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetDto) contextValidateRootDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RootDomains); i++ {

		if m.RootDomains[i] != nil {
			if err := m.RootDomains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rootDomains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TargetDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TargetDto) UnmarshalBinary(b []byte) error {
	var res TargetDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
