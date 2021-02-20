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

// AgentDto agent dto
//
// swagger:model AgentDto
type AgentDto struct {

	// agent type
	AgentType *string `json:"agentType,omitempty"`

	// categories
	Categories []string `json:"categories"`

	// command
	Command *string `json:"command,omitempty"`

	// id
	// Format: uuid
	ID *strfmt.UUID `json:"id,omitempty"`

	// last run
	// Format: date-time
	LastRun strfmt.DateTime `json:"lastRun,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// repository
	Repository *string `json:"repository,omitempty"`

	// script
	Script *string `json:"script,omitempty"`

	// trigger rootdomain has bounty
	TriggerRootdomainHasBounty bool `json:"triggerRootdomainHasBounty,omitempty"`

	// trigger rootdomain inc exc name
	TriggerRootdomainIncExcName *string `json:"triggerRootdomainIncExcName,omitempty"`

	// trigger rootdomain name
	TriggerRootdomainName *string `json:"triggerRootdomainName,omitempty"`

	// trigger skip if run before
	TriggerSkipIfRunBefore bool `json:"triggerSkipIfRunBefore,omitempty"`

	// trigger subdomain has bounty
	TriggerSubdomainHasBounty bool `json:"triggerSubdomainHasBounty,omitempty"`

	// trigger subdomain has Http or Https open
	TriggerSubdomainHasHTTPOrHTTPSOpen bool `json:"triggerSubdomainHasHttpOrHttpsOpen,omitempty"`

	// trigger subdomain IP
	TriggerSubdomainIP *string `json:"triggerSubdomainIP,omitempty"`

	// trigger subdomain inc exc IP
	TriggerSubdomainIncExcIP *string `json:"triggerSubdomainIncExcIP,omitempty"`

	// trigger subdomain inc exc label
	TriggerSubdomainIncExcLabel *string `json:"triggerSubdomainIncExcLabel,omitempty"`

	// trigger subdomain inc exc name
	TriggerSubdomainIncExcName *string `json:"triggerSubdomainIncExcName,omitempty"`

	// trigger subdomain inc exc service port
	TriggerSubdomainIncExcServicePort *string `json:"triggerSubdomainIncExcServicePort,omitempty"`

	// trigger subdomain inc exc technology
	TriggerSubdomainIncExcTechnology *string `json:"triggerSubdomainIncExcTechnology,omitempty"`

	// trigger subdomain is alive
	TriggerSubdomainIsAlive bool `json:"triggerSubdomainIsAlive,omitempty"`

	// trigger subdomain is main portal
	TriggerSubdomainIsMainPortal bool `json:"triggerSubdomainIsMainPortal,omitempty"`

	// trigger subdomain label
	TriggerSubdomainLabel *string `json:"triggerSubdomainLabel,omitempty"`

	// trigger subdomain name
	TriggerSubdomainName *string `json:"triggerSubdomainName,omitempty"`

	// trigger subdomain service port
	TriggerSubdomainServicePort *string `json:"triggerSubdomainServicePort,omitempty"`

	// trigger subdomain technology
	TriggerSubdomainTechnology *string `json:"triggerSubdomainTechnology,omitempty"`

	// trigger target has bounty
	TriggerTargetHasBounty bool `json:"triggerTargetHasBounty,omitempty"`

	// trigger target inc exc name
	TriggerTargetIncExcName *string `json:"triggerTargetIncExcName,omitempty"`

	// trigger target name
	TriggerTargetName *string `json:"triggerTargetName,omitempty"`
}

// Validate validates this agent dto
func (m *AgentDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentDto) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AgentDto) validateLastRun(formats strfmt.Registry) error {
	if swag.IsZero(m.LastRun) { // not required
		return nil
	}

	if err := validate.FormatOf("lastRun", "body", "date-time", m.LastRun.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this agent dto based on context it is used
func (m *AgentDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AgentDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentDto) UnmarshalBinary(b []byte) error {
	var res AgentDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
