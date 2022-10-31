/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Revision struct for Revision
type Revision struct {
	Version *int32 `json:"version,omitempty"`
	Real *int32 `json:"real,omitempty"`
	IsRepack *bool `json:"isRepack,omitempty"`
}

// NewRevision instantiates a new Revision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevision() *Revision {
	this := Revision{}
	return &this
}

// NewRevisionWithDefaults instantiates a new Revision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionWithDefaults() *Revision {
	this := Revision{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Revision) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Revision) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Revision) SetVersion(v int32) {
	o.Version = &v
}

// GetReal returns the Real field value if set, zero value otherwise.
func (o *Revision) GetReal() int32 {
	if o == nil || o.Real == nil {
		var ret int32
		return ret
	}
	return *o.Real
}

// GetRealOk returns a tuple with the Real field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetRealOk() (*int32, bool) {
	if o == nil || o.Real == nil {
		return nil, false
	}
	return o.Real, true
}

// HasReal returns a boolean if a field has been set.
func (o *Revision) HasReal() bool {
	if o != nil && o.Real != nil {
		return true
	}

	return false
}

// SetReal gets a reference to the given int32 and assigns it to the Real field.
func (o *Revision) SetReal(v int32) {
	o.Real = &v
}

// GetIsRepack returns the IsRepack field value if set, zero value otherwise.
func (o *Revision) GetIsRepack() bool {
	if o == nil || o.IsRepack == nil {
		var ret bool
		return ret
	}
	return *o.IsRepack
}

// GetIsRepackOk returns a tuple with the IsRepack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Revision) GetIsRepackOk() (*bool, bool) {
	if o == nil || o.IsRepack == nil {
		return nil, false
	}
	return o.IsRepack, true
}

// HasIsRepack returns a boolean if a field has been set.
func (o *Revision) HasIsRepack() bool {
	if o != nil && o.IsRepack != nil {
		return true
	}

	return false
}

// SetIsRepack gets a reference to the given bool and assigns it to the IsRepack field.
func (o *Revision) SetIsRepack(v bool) {
	o.IsRepack = &v
}

func (o Revision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Real != nil {
		toSerialize["real"] = o.Real
	}
	if o.IsRepack != nil {
		toSerialize["isRepack"] = o.IsRepack
	}
	return json.Marshal(toSerialize)
}

type NullableRevision struct {
	value *Revision
	isSet bool
}

func (v NullableRevision) Get() *Revision {
	return v.value
}

func (v *NullableRevision) Set(val *Revision) {
	v.value = val
	v.isSet = true
}

func (v NullableRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevision(val *Revision) *NullableRevision {
	return &NullableRevision{value: val, isSet: true}
}

func (v NullableRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


