// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EvaluationBatchRequest evaluation batch request
// swagger:model evaluationBatchRequest
type EvaluationBatchRequest struct {

	// entity context
	EntityContext interface{} `json:"entityContext,omitempty"`

	// entity ID
	EntityID string `json:"entityID,omitempty"`
}

// Validate validates this evaluation batch request
func (m *EvaluationBatchRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EvaluationBatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EvaluationBatchRequest) UnmarshalBinary(b []byte) error {
	var res EvaluationBatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
