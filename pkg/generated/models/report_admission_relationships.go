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

// ReportAdmissionRelationships report admission relationships
// swagger:model ReportAdmissionRelationships
type ReportAdmissionRelationships struct {

	// report
	Report *ReportAdmissionRelationshipsReport `json:"report,omitempty"`
}

func ReportAdmissionRelationshipsWithDefaults(defaults client.Defaults) *ReportAdmissionRelationships {
	return &ReportAdmissionRelationships{

		Report: ReportAdmissionRelationshipsReportWithDefaults(defaults),
	}
}

func (m *ReportAdmissionRelationships) WithReport(report ReportAdmissionRelationshipsReport) *ReportAdmissionRelationships {

	m.Report = &report

	return m
}

func (m *ReportAdmissionRelationships) WithoutReport() *ReportAdmissionRelationships {
	m.Report = nil
	return m
}

// Validate validates this report admission relationships
func (m *ReportAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportAdmissionRelationships) validateReport(formats strfmt.Registry) error {

	if swag.IsZero(m.Report) { // not required
		return nil
	}

	if m.Report != nil {
		if err := m.Report.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("report")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ReportAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReportAdmissionRelationshipsReport report admission relationships report
// swagger:model ReportAdmissionRelationshipsReport
type ReportAdmissionRelationshipsReport struct {

	// data
	Data []*Report `json:"data"`
}

func ReportAdmissionRelationshipsReportWithDefaults(defaults client.Defaults) *ReportAdmissionRelationshipsReport {
	return &ReportAdmissionRelationshipsReport{

		Data: make([]*Report, 0),
	}
}

func (m *ReportAdmissionRelationshipsReport) WithData(data []*Report) *ReportAdmissionRelationshipsReport {

	m.Data = data

	return m
}

// Validate validates this report admission relationships report
func (m *ReportAdmissionRelationshipsReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportAdmissionRelationshipsReport) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("report" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportAdmissionRelationshipsReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportAdmissionRelationshipsReport) UnmarshalBinary(b []byte) error {
	var res ReportAdmissionRelationshipsReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReportAdmissionRelationshipsReport) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
