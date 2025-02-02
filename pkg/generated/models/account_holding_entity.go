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

// AccountHoldingEntity account holding entity
// swagger:model AccountHoldingEntity
type AccountHoldingEntity struct {

	// Financial institution address
	BankAddress []string `json:"bank_address,omitempty"`

	// Financial institution identification
	BankID string `json:"bank_id,omitempty"`

	// bank id code
	BankIDCode BankIDCode `json:"bank_id_code,omitempty"`

	// Financial institution name
	BankName string `json:"bank_name,omitempty"`

	// Identifier of the financial institution which services the account
	BankPartyID string `json:"bank_party_id,omitempty"`
}

func AccountHoldingEntityWithDefaults(defaults client.Defaults) *AccountHoldingEntity {
	return &AccountHoldingEntity{

		BankAddress: make([]string, 0),

		BankID: defaults.GetString("AccountHoldingEntity", "bank_id"),

		// TODO BankIDCode: BankIDCode,

		BankName: defaults.GetString("AccountHoldingEntity", "bank_name"),

		BankPartyID: defaults.GetString("AccountHoldingEntity", "bank_party_id"),
	}
}

func (m *AccountHoldingEntity) WithBankAddress(bankAddress []string) *AccountHoldingEntity {

	m.BankAddress = bankAddress

	return m
}

func (m *AccountHoldingEntity) WithBankID(bankID string) *AccountHoldingEntity {

	m.BankID = bankID

	return m
}

func (m *AccountHoldingEntity) WithBankIDCode(bankIDCode BankIDCode) *AccountHoldingEntity {

	m.BankIDCode = bankIDCode

	return m
}

func (m *AccountHoldingEntity) WithBankName(bankName string) *AccountHoldingEntity {

	m.BankName = bankName

	return m
}

func (m *AccountHoldingEntity) WithBankPartyID(bankPartyID string) *AccountHoldingEntity {

	m.BankPartyID = bankPartyID

	return m
}

// Validate validates this account holding entity
func (m *AccountHoldingEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountHoldingEntity) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := m.BankIDCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bank_id_code")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountHoldingEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountHoldingEntity) UnmarshalBinary(b []byte) error {
	var res AccountHoldingEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountHoldingEntity) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
