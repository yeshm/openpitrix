// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixSendTaskResponse openpitrix send task response
// swagger:model openpitrixSendTaskResponse
type OpenpitrixSendTaskResponse struct {

	// task id
	TaskID string `json:"task_id,omitempty"`
}

// Validate validates this openpitrix send task response
func (m *OpenpitrixSendTaskResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixSendTaskResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixSendTaskResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixSendTaskResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}