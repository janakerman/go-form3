// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MandateRelationships mandate relationships
// swagger:model MandateRelationships
type MandateRelationships struct {

	// mandate admission
	MandateAdmission *MandateRelationshipsMandateAdmission `json:"mandate_admission,omitempty"`

	// mandate return
	MandateReturn *MandateRelationshipsMandateReturn `json:"mandate_return,omitempty"`

	// mandate submission
	MandateSubmission *MandateRelationshipsMandateSubmission `json:"mandate_submission,omitempty"`
}

func MandateRelationshipsWithDefaults(defaults client.Defaults) *MandateRelationships {
	return &MandateRelationships{

		MandateAdmission: MandateRelationshipsMandateAdmissionWithDefaults(defaults),

		MandateReturn: MandateRelationshipsMandateReturnWithDefaults(defaults),

		MandateSubmission: MandateRelationshipsMandateSubmissionWithDefaults(defaults),
	}
}

func (m *MandateRelationships) WithMandateAdmission(mandateAdmission MandateRelationshipsMandateAdmission) *MandateRelationships {

	m.MandateAdmission = &mandateAdmission

	return m
}

func (m *MandateRelationships) WithoutMandateAdmission() *MandateRelationships {
	m.MandateAdmission = nil
	return m
}

func (m *MandateRelationships) WithMandateReturn(mandateReturn MandateRelationshipsMandateReturn) *MandateRelationships {

	m.MandateReturn = &mandateReturn

	return m
}

func (m *MandateRelationships) WithoutMandateReturn() *MandateRelationships {
	m.MandateReturn = nil
	return m
}

func (m *MandateRelationships) WithMandateSubmission(mandateSubmission MandateRelationshipsMandateSubmission) *MandateRelationships {

	m.MandateSubmission = &mandateSubmission

	return m
}

func (m *MandateRelationships) WithoutMandateSubmission() *MandateRelationships {
	m.MandateSubmission = nil
	return m
}

// Validate validates this mandate relationships
func (m *MandateRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMandateAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMandateReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMandateSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateRelationships) validateMandateAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.MandateAdmission) { // not required
		return nil
	}

	if m.MandateAdmission != nil {
		if err := m.MandateAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandate_admission")
			}
			return err
		}
	}

	return nil
}

func (m *MandateRelationships) validateMandateReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.MandateReturn) { // not required
		return nil
	}

	if m.MandateReturn != nil {
		if err := m.MandateReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandate_return")
			}
			return err
		}
	}

	return nil
}

func (m *MandateRelationships) validateMandateSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.MandateSubmission) { // not required
		return nil
	}

	if m.MandateSubmission != nil {
		if err := m.MandateSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandate_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateRelationships) UnmarshalBinary(b []byte) error {
	var res MandateRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateRelationshipsMandateAdmission mandate relationships mandate admission
// swagger:model MandateRelationshipsMandateAdmission
type MandateRelationshipsMandateAdmission struct {

	// data
	Data []*MandateAdmission `json:"data"`
}

func MandateRelationshipsMandateAdmissionWithDefaults(defaults client.Defaults) *MandateRelationshipsMandateAdmission {
	return &MandateRelationshipsMandateAdmission{

		Data: make([]*MandateAdmission, 0),
	}
}

func (m *MandateRelationshipsMandateAdmission) WithData(data []*MandateAdmission) *MandateRelationshipsMandateAdmission {

	m.Data = data

	return m
}

// Validate validates this mandate relationships mandate admission
func (m *MandateRelationshipsMandateAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateRelationshipsMandateAdmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mandate_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateRelationshipsMandateAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateRelationshipsMandateAdmission) UnmarshalBinary(b []byte) error {
	var res MandateRelationshipsMandateAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateRelationshipsMandateAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateRelationshipsMandateReturn mandate relationships mandate return
// swagger:model MandateRelationshipsMandateReturn
type MandateRelationshipsMandateReturn struct {

	// data
	Data []*MandateReturn `json:"data"`
}

func MandateRelationshipsMandateReturnWithDefaults(defaults client.Defaults) *MandateRelationshipsMandateReturn {
	return &MandateRelationshipsMandateReturn{

		Data: make([]*MandateReturn, 0),
	}
}

func (m *MandateRelationshipsMandateReturn) WithData(data []*MandateReturn) *MandateRelationshipsMandateReturn {

	m.Data = data

	return m
}

// Validate validates this mandate relationships mandate return
func (m *MandateRelationshipsMandateReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateRelationshipsMandateReturn) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mandate_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateRelationshipsMandateReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateRelationshipsMandateReturn) UnmarshalBinary(b []byte) error {
	var res MandateRelationshipsMandateReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateRelationshipsMandateReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateRelationshipsMandateSubmission mandate relationships mandate submission
// swagger:model MandateRelationshipsMandateSubmission
type MandateRelationshipsMandateSubmission struct {

	// data
	Data []*MandateSubmission `json:"data"`
}

func MandateRelationshipsMandateSubmissionWithDefaults(defaults client.Defaults) *MandateRelationshipsMandateSubmission {
	return &MandateRelationshipsMandateSubmission{

		Data: make([]*MandateSubmission, 0),
	}
}

func (m *MandateRelationshipsMandateSubmission) WithData(data []*MandateSubmission) *MandateRelationshipsMandateSubmission {

	m.Data = data

	return m
}

// Validate validates this mandate relationships mandate submission
func (m *MandateRelationshipsMandateSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateRelationshipsMandateSubmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mandate_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateRelationshipsMandateSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateRelationshipsMandateSubmission) UnmarshalBinary(b []byte) error {
	var res MandateRelationshipsMandateSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateRelationshipsMandateSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
