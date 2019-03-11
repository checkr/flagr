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

// EvalResult eval result
// swagger:model evalResult
type EvalResult struct {

	// flag key
	// Required: true
	FlagKey *string `json:"flagKey"`

	// flag value
	// Required: true
	// Min Length: 1
	FlagValue *string `json:"flagValue"`

	// payload
	// Required: true
	Payload interface{} `json:"payload"`
}

// Validate validates this eval result
func (m *EvalResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlagKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlagValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EvalResult) validateFlagKey(formats strfmt.Registry) error {

	if err := validate.Required("flagKey", "body", m.FlagKey); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validateFlagValue(formats strfmt.Registry) error {

	if err := validate.Required("flagValue", "body", m.FlagValue); err != nil {
		return err
	}

	if err := validate.MinLength("flagValue", "body", string(*m.FlagValue), 1); err != nil {
		return err
	}

	return nil
}

func (m *EvalResult) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EvalResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvalResult) UnmarshalBinary(b []byte) error {
	var res EvalResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
