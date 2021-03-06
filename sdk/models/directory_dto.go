// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectoryDto directory dto
//
// swagger:model DirectoryDto
type DirectoryDto struct {

	// id
	// Format: uuid
	ID *strfmt.UUID `json:"id,omitempty"`

	// method
	Method *string `json:"method,omitempty"`

	// size
	Size *string `json:"size,omitempty"`

	// status code
	StatusCode *string `json:"statusCode,omitempty"`

	// uri
	URI *string `json:"uri,omitempty"`
}

// Validate validates this directory dto
func (m *DirectoryDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectoryDto) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this directory dto based on context it is used
func (m *DirectoryDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DirectoryDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectoryDto) UnmarshalBinary(b []byte) error {
	var res DirectoryDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
