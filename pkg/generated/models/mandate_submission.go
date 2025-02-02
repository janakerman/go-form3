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
	"github.com/go-openapi/validate"
)

// MandateSubmission mandate submission
// swagger:model MandateSubmission
type MandateSubmission struct {

	// attributes
	Attributes *MandateSubmissionAttributes `json:"attributes,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func MandateSubmissionWithDefaults(defaults client.Defaults) *MandateSubmission {
	return &MandateSubmission{

		Attributes: MandateSubmissionAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("MandateSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUID("MandateSubmission", "organisation_id"),

		Type: defaults.GetString("MandateSubmission", "type"),

		Version: defaults.GetInt64Ptr("MandateSubmission", "version"),
	}
}

func (m *MandateSubmission) WithAttributes(attributes MandateSubmissionAttributes) *MandateSubmission {

	m.Attributes = &attributes

	return m
}

func (m *MandateSubmission) WithoutAttributes() *MandateSubmission {
	m.Attributes = nil
	return m
}

func (m *MandateSubmission) WithID(id strfmt.UUID) *MandateSubmission {

	m.ID = id

	return m
}

func (m *MandateSubmission) WithOrganisationID(organisationID strfmt.UUID) *MandateSubmission {

	m.OrganisationID = organisationID

	return m
}

func (m *MandateSubmission) WithType(typeVar string) *MandateSubmission {

	m.Type = typeVar

	return m
}

func (m *MandateSubmission) WithVersion(version int64) *MandateSubmission {

	m.Version = &version

	return m
}

func (m *MandateSubmission) WithoutVersion() *MandateSubmission {
	m.Version = nil
	return m
}

// Validate validates this mandate submission
func (m *MandateSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateSubmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *MandateSubmission) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmission) UnmarshalBinary(b []byte) error {
	var res MandateSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateSubmissionAttributes mandate submission attributes
// swagger:model MandateSubmissionAttributes
type MandateSubmissionAttributes struct {

	// last payment date
	// Format: date
	LastPaymentDate strfmt.Date `json:"last_payment_date,omitempty"`

	// original mandate
	OriginalMandate *MandateAttributes `json:"original_mandate,omitempty"`

	// status
	Status MandateSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission date
	// Format: date-time
	SubmissionDate strfmt.DateTime `json:"submission_date,omitempty"`

	// submission reason
	SubmissionReason string `json:"submission_reason,omitempty"`

	// submitted mandate
	SubmittedMandate *MandateAttributes `json:"submitted_mandate,omitempty"`
}

func MandateSubmissionAttributesWithDefaults(defaults client.Defaults) *MandateSubmissionAttributes {
	return &MandateSubmissionAttributes{

		LastPaymentDate: defaults.GetStrfmtDate("MandateSubmissionAttributes", "last_payment_date"),

		OriginalMandate: MandateAttributesWithDefaults(defaults),

		// TODO Status: MandateSubmissionStatus,

		StatusReason: defaults.GetString("MandateSubmissionAttributes", "status_reason"),

		SubmissionDate: defaults.GetStrfmtDateTime("MandateSubmissionAttributes", "submission_date"),

		SubmissionReason: defaults.GetString("MandateSubmissionAttributes", "submission_reason"),

		SubmittedMandate: MandateAttributesWithDefaults(defaults),
	}
}

func (m *MandateSubmissionAttributes) WithLastPaymentDate(lastPaymentDate strfmt.Date) *MandateSubmissionAttributes {

	m.LastPaymentDate = lastPaymentDate

	return m
}

func (m *MandateSubmissionAttributes) WithOriginalMandate(originalMandate MandateAttributes) *MandateSubmissionAttributes {

	m.OriginalMandate = &originalMandate

	return m
}

func (m *MandateSubmissionAttributes) WithoutOriginalMandate() *MandateSubmissionAttributes {
	m.OriginalMandate = nil
	return m
}

func (m *MandateSubmissionAttributes) WithStatus(status MandateSubmissionStatus) *MandateSubmissionAttributes {

	m.Status = status

	return m
}

func (m *MandateSubmissionAttributes) WithStatusReason(statusReason string) *MandateSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *MandateSubmissionAttributes) WithSubmissionDate(submissionDate strfmt.DateTime) *MandateSubmissionAttributes {

	m.SubmissionDate = submissionDate

	return m
}

func (m *MandateSubmissionAttributes) WithSubmissionReason(submissionReason string) *MandateSubmissionAttributes {

	m.SubmissionReason = submissionReason

	return m
}

func (m *MandateSubmissionAttributes) WithSubmittedMandate(submittedMandate MandateAttributes) *MandateSubmissionAttributes {

	m.SubmittedMandate = &submittedMandate

	return m
}

func (m *MandateSubmissionAttributes) WithoutSubmittedMandate() *MandateSubmissionAttributes {
	m.SubmittedMandate = nil
	return m
}

// Validate validates this mandate submission attributes
func (m *MandateSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastPaymentDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalMandate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedMandate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateSubmissionAttributes) validateLastPaymentDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastPaymentDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"last_payment_date", "body", "date", m.LastPaymentDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateOriginalMandate(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginalMandate) { // not required
		return nil
	}

	if m.OriginalMandate != nil {
		if err := m.OriginalMandate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "original_mandate")
			}
			return err
		}
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateSubmissionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_date", "body", "date-time", m.SubmissionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MandateSubmissionAttributes) validateSubmittedMandate(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmittedMandate) { // not required
		return nil
	}

	if m.SubmittedMandate != nil {
		if err := m.SubmittedMandate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "submitted_mandate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res MandateSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
