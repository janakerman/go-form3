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

// AdviceSubmission advice submission
// swagger:model AdviceSubmission
type AdviceSubmission struct {

	// attributes
	Attributes *AdviceSubmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *AdviceSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func AdviceSubmissionWithDefaults(defaults client.Defaults) *AdviceSubmission {
	return &AdviceSubmission{

		Attributes: AdviceSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("AdviceSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("AdviceSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("AdviceSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("AdviceSubmission", "organisation_id"),

		Relationships: AdviceSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("AdviceSubmission", "type"),

		Version: defaults.GetInt64Ptr("AdviceSubmission", "version"),
	}
}

func (m *AdviceSubmission) WithAttributes(attributes AdviceSubmissionAttributes) *AdviceSubmission {

	m.Attributes = &attributes

	return m
}

func (m *AdviceSubmission) WithoutAttributes() *AdviceSubmission {
	m.Attributes = nil
	return m
}

func (m *AdviceSubmission) WithCreatedOn(createdOn strfmt.DateTime) *AdviceSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *AdviceSubmission) WithoutCreatedOn() *AdviceSubmission {
	m.CreatedOn = nil
	return m
}

func (m *AdviceSubmission) WithID(id strfmt.UUID) *AdviceSubmission {

	m.ID = &id

	return m
}

func (m *AdviceSubmission) WithoutID() *AdviceSubmission {
	m.ID = nil
	return m
}

func (m *AdviceSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *AdviceSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *AdviceSubmission) WithoutModifiedOn() *AdviceSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *AdviceSubmission) WithOrganisationID(organisationID strfmt.UUID) *AdviceSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *AdviceSubmission) WithoutOrganisationID() *AdviceSubmission {
	m.OrganisationID = nil
	return m
}

func (m *AdviceSubmission) WithRelationships(relationships AdviceSubmissionRelationships) *AdviceSubmission {

	m.Relationships = &relationships

	return m
}

func (m *AdviceSubmission) WithoutRelationships() *AdviceSubmission {
	m.Relationships = nil
	return m
}

func (m *AdviceSubmission) WithType(typeVar string) *AdviceSubmission {

	m.Type = typeVar

	return m
}

func (m *AdviceSubmission) WithVersion(version int64) *AdviceSubmission {

	m.Version = &version

	return m
}

func (m *AdviceSubmission) WithoutVersion() *AdviceSubmission {
	m.Version = nil
	return m
}

// Validate validates this advice submission
func (m *AdviceSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
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

func (m *AdviceSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *AdviceSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmission) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

func (m *AdviceSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdviceSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdviceSubmission) UnmarshalBinary(b []byte) error {
	var res AdviceSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AdviceSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AdviceSubmissionAttributes advice submission attributes
// swagger:model AdviceSubmissionAttributes
type AdviceSubmissionAttributes struct {

	// status
	Status AdviceSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`

	// transaction start datetime
	// Read Only: true
	// Format: date-time
	TransactionStartDatetime strfmt.DateTime `json:"transaction_start_datetime,omitempty"`
}

func AdviceSubmissionAttributesWithDefaults(defaults client.Defaults) *AdviceSubmissionAttributes {
	return &AdviceSubmissionAttributes{

		// TODO Status: AdviceSubmissionStatus,

		StatusReason: defaults.GetString("AdviceSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("AdviceSubmissionAttributes", "submission_datetime"),

		TransactionStartDatetime: defaults.GetStrfmtDateTime("AdviceSubmissionAttributes", "transaction_start_datetime"),
	}
}

func (m *AdviceSubmissionAttributes) WithStatus(status AdviceSubmissionStatus) *AdviceSubmissionAttributes {

	m.Status = status

	return m
}

func (m *AdviceSubmissionAttributes) WithStatusReason(statusReason string) *AdviceSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *AdviceSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *AdviceSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

func (m *AdviceSubmissionAttributes) WithTransactionStartDatetime(transactionStartDatetime strfmt.DateTime) *AdviceSubmissionAttributes {

	m.TransactionStartDatetime = transactionStartDatetime

	return m
}

// Validate validates this advice submission attributes
func (m *AdviceSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdviceSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *AdviceSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdviceSubmissionAttributes) validateTransactionStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"transaction_start_datetime", "body", "date-time", m.TransactionStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdviceSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdviceSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res AdviceSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AdviceSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
