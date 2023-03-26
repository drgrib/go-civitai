// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Metadata metadata
//
// swagger:model Metadata
type Metadata struct {

	// The the current page you are at.
	// Example: 2
	CurrentPage int64 `json:"currentPage,omitempty"`

	// The url to get the next batch of items.
	// Example: https://civitai.com/api/v1/creators?limit=3\u0026page=3
	NextPage string `json:"nextPage,omitempty"`

	// The the size of the batch.
	// Example: 3
	PageSize int64 `json:"pageSize,omitempty"`

	// The url to get the previous batch of items.
	// Example: https://civitai.com/api/v1/creators?limit=3\u0026page=1
	PrevPage string `json:"prevPage,omitempty"`

	// The total number of items available.
	// Example: 46
	TotalItems int64 `json:"totalItems,omitempty"`

	// The total number of pages.
	// Example: 16
	TotalPages int64 `json:"totalPages,omitempty"`
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metadata based on context it is used
func (m *Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metadata) UnmarshalBinary(b []byte) error {
	var res Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}