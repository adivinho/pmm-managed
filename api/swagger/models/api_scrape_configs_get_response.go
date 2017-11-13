// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIScrapeConfigsGetResponse api scrape configs get response
// swagger:model apiScrapeConfigsGetResponse

type APIScrapeConfigsGetResponse struct {

	// scrape config
	ScrapeConfig *APIScrapeConfig `json:"scrape_config,omitempty"`
}

/* polymorph apiScrapeConfigsGetResponse scrape_config false */

// Validate validates this api scrape configs get response
func (m *APIScrapeConfigsGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScrapeConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIScrapeConfigsGetResponse) validateScrapeConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ScrapeConfig) { // not required
		return nil
	}

	if m.ScrapeConfig != nil {

		if err := m.ScrapeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scrape_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIScrapeConfigsGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIScrapeConfigsGetResponse) UnmarshalBinary(b []byte) error {
	var res APIScrapeConfigsGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}