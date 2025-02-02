// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RecallDecisionAdmissionRelationships recall decision admission relationships
// swagger:model RecallDecisionAdmissionRelationships
type RecallDecisionAdmissionRelationships struct {

	// ID of the payment resource related to the recall decicsion admission
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// ID of the recall resource related to the recall decicsion admission
	Recall *RelationshipLinks `json:"recall,omitempty"`

	// ID of the recall decision resource related to the recall decicsion admission
	RecallDecision *RelationshipLinks `json:"recall_decision,omitempty"`
}

func RecallDecisionAdmissionRelationshipsWithDefaults(defaults client.Defaults) *RecallDecisionAdmissionRelationships {
	return &RecallDecisionAdmissionRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		Recall: RelationshipLinksWithDefaults(defaults),

		RecallDecision: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *RecallDecisionAdmissionRelationships) WithPayment(payment RelationshipLinks) *RecallDecisionAdmissionRelationships {

	m.Payment = &payment

	return m
}

func (m *RecallDecisionAdmissionRelationships) WithoutPayment() *RecallDecisionAdmissionRelationships {
	m.Payment = nil
	return m
}

func (m *RecallDecisionAdmissionRelationships) WithRecall(recall RelationshipLinks) *RecallDecisionAdmissionRelationships {

	m.Recall = &recall

	return m
}

func (m *RecallDecisionAdmissionRelationships) WithoutRecall() *RecallDecisionAdmissionRelationships {
	m.Recall = nil
	return m
}

func (m *RecallDecisionAdmissionRelationships) WithRecallDecision(recallDecision RelationshipLinks) *RecallDecisionAdmissionRelationships {

	m.RecallDecision = &recallDecision

	return m
}

func (m *RecallDecisionAdmissionRelationships) WithoutRecallDecision() *RecallDecisionAdmissionRelationships {
	m.RecallDecision = nil
	return m
}

// Validate validates this recall decision admission relationships
func (m *RecallDecisionAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallDecision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionAdmissionRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionAdmissionRelationships) validateRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.Recall) { // not required
		return nil
	}

	if m.Recall != nil {
		if err := m.Recall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall")
			}
			return err
		}
	}

	return nil
}

func (m *RecallDecisionAdmissionRelationships) validateRecallDecision(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallDecision) { // not required
		return nil
	}

	if m.RecallDecision != nil {
		if err := m.RecallDecision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall_decision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res RecallDecisionAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
