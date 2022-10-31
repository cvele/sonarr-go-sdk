/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ProperDownloadTypes the model 'ProperDownloadTypes'
type ProperDownloadTypes string

// List of ProperDownloadTypes
const (
	PREFER_AND_UPGRADE ProperDownloadTypes = "preferAndUpgrade"
	DO_NOT_UPGRADE ProperDownloadTypes = "doNotUpgrade"
	DO_NOT_PREFER ProperDownloadTypes = "doNotPrefer"
)

// All allowed values of ProperDownloadTypes enum
var AllowedProperDownloadTypesEnumValues = []ProperDownloadTypes{
	"preferAndUpgrade",
	"doNotUpgrade",
	"doNotPrefer",
}

func (v *ProperDownloadTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProperDownloadTypes(value)
	for _, existing := range AllowedProperDownloadTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProperDownloadTypes", value)
}

// NewProperDownloadTypesFromValue returns a pointer to a valid ProperDownloadTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProperDownloadTypesFromValue(v string) (*ProperDownloadTypes, error) {
	ev := ProperDownloadTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProperDownloadTypes: valid values are %v", v, AllowedProperDownloadTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProperDownloadTypes) IsValid() bool {
	for _, existing := range AllowedProperDownloadTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProperDownloadTypes value
func (v ProperDownloadTypes) Ptr() *ProperDownloadTypes {
	return &v
}

type NullableProperDownloadTypes struct {
	value *ProperDownloadTypes
	isSet bool
}

func (v NullableProperDownloadTypes) Get() *ProperDownloadTypes {
	return v.value
}

func (v *NullableProperDownloadTypes) Set(val *ProperDownloadTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableProperDownloadTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableProperDownloadTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProperDownloadTypes(val *ProperDownloadTypes) *NullableProperDownloadTypes {
	return &NullableProperDownloadTypes{value: val, isSet: true}
}

func (v NullableProperDownloadTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProperDownloadTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
