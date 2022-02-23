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
)

// TopologySelectorTerm A topology selector term represents the result of label queries.
//
// A null or empty topology selector term matches no objects.
// The requirements of them are ANDed.
// It provides a subset of functionality as NodeSelectorTerm.
// This is an alpha feature and may change in the future.
// +structType=atomic
//
// swagger:model TopologySelectorTerm
type TopologySelectorTerm struct {

	// A list of topology selector requirements by labels.
	// +optional
	MatchLabelExpressions []*TopologySelectorLabelRequirement `json:"matchLabelExpressions"`
}

// Validate validates this topology selector term
func (m *TopologySelectorTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchLabelExpressions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologySelectorTerm) validateMatchLabelExpressions(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchLabelExpressions) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchLabelExpressions); i++ {
		if swag.IsZero(m.MatchLabelExpressions[i]) { // not required
			continue
		}

		if m.MatchLabelExpressions[i] != nil {
			if err := m.MatchLabelExpressions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchLabelExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this topology selector term based on the context it is used
func (m *TopologySelectorTerm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatchLabelExpressions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopologySelectorTerm) contextValidateMatchLabelExpressions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MatchLabelExpressions); i++ {

		if m.MatchLabelExpressions[i] != nil {
			if err := m.MatchLabelExpressions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchLabelExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopologySelectorTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopologySelectorTerm) UnmarshalBinary(b []byte) error {
	var res TopologySelectorTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}