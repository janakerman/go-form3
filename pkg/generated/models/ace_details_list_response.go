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

// AceDetailsListResponse ace details list response
// swagger:model AceDetailsListResponse
type AceDetailsListResponse struct {

	// data
	Data []*Ace `json:"data"`

	// links
	Links *Links `json:"links,omitempty"`
}

func AceDetailsListResponseWithDefaults(defaults client.Defaults) *AceDetailsListResponse {
	return &AceDetailsListResponse{

		Data: make([]*Ace, 0),

		Links: LinksWithDefaults(defaults),
	}
}

func (m *AceDetailsListResponse) WithData(data []*Ace) *AceDetailsListResponse {

	m.Data = data

	return m
}

func (m *AceDetailsListResponse) WithLinks(links Links) *AceDetailsListResponse {

	m.Links = &links

	return m
}

func (m *AceDetailsListResponse) WithoutLinks() *AceDetailsListResponse {
	m.Links = nil
	return m
}

// Validate validates this ace details list response
func (m *AceDetailsListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AceDetailsListResponse) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AceDetailsListResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AceDetailsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AceDetailsListResponse) UnmarshalBinary(b []byte) error {
	var res AceDetailsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AceDetailsListResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
