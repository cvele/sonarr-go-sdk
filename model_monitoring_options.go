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

// MonitoringOptions struct for MonitoringOptions
type MonitoringOptions struct {
	IgnoreEpisodesWithFiles *bool `json:"ignoreEpisodesWithFiles,omitempty"`
	IgnoreEpisodesWithoutFiles *bool `json:"ignoreEpisodesWithoutFiles,omitempty"`
	Monitor *MonitorTypes `json:"monitor,omitempty"`
}

// NewMonitoringOptions instantiates a new MonitoringOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringOptions() *MonitoringOptions {
	this := MonitoringOptions{}
	return &this
}

// NewMonitoringOptionsWithDefaults instantiates a new MonitoringOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringOptionsWithDefaults() *MonitoringOptions {
	this := MonitoringOptions{}
	return &this
}

// GetIgnoreEpisodesWithFiles returns the IgnoreEpisodesWithFiles field value if set, zero value otherwise.
func (o *MonitoringOptions) GetIgnoreEpisodesWithFiles() bool {
	if o == nil || o.IgnoreEpisodesWithFiles == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreEpisodesWithFiles
}

// GetIgnoreEpisodesWithFilesOk returns a tuple with the IgnoreEpisodesWithFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringOptions) GetIgnoreEpisodesWithFilesOk() (*bool, bool) {
	if o == nil || o.IgnoreEpisodesWithFiles == nil {
		return nil, false
	}
	return o.IgnoreEpisodesWithFiles, true
}

// HasIgnoreEpisodesWithFiles returns a boolean if a field has been set.
func (o *MonitoringOptions) HasIgnoreEpisodesWithFiles() bool {
	if o != nil && o.IgnoreEpisodesWithFiles != nil {
		return true
	}

	return false
}

// SetIgnoreEpisodesWithFiles gets a reference to the given bool and assigns it to the IgnoreEpisodesWithFiles field.
func (o *MonitoringOptions) SetIgnoreEpisodesWithFiles(v bool) {
	o.IgnoreEpisodesWithFiles = &v
}

// GetIgnoreEpisodesWithoutFiles returns the IgnoreEpisodesWithoutFiles field value if set, zero value otherwise.
func (o *MonitoringOptions) GetIgnoreEpisodesWithoutFiles() bool {
	if o == nil || o.IgnoreEpisodesWithoutFiles == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreEpisodesWithoutFiles
}

// GetIgnoreEpisodesWithoutFilesOk returns a tuple with the IgnoreEpisodesWithoutFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringOptions) GetIgnoreEpisodesWithoutFilesOk() (*bool, bool) {
	if o == nil || o.IgnoreEpisodesWithoutFiles == nil {
		return nil, false
	}
	return o.IgnoreEpisodesWithoutFiles, true
}

// HasIgnoreEpisodesWithoutFiles returns a boolean if a field has been set.
func (o *MonitoringOptions) HasIgnoreEpisodesWithoutFiles() bool {
	if o != nil && o.IgnoreEpisodesWithoutFiles != nil {
		return true
	}

	return false
}

// SetIgnoreEpisodesWithoutFiles gets a reference to the given bool and assigns it to the IgnoreEpisodesWithoutFiles field.
func (o *MonitoringOptions) SetIgnoreEpisodesWithoutFiles(v bool) {
	o.IgnoreEpisodesWithoutFiles = &v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *MonitoringOptions) GetMonitor() MonitorTypes {
	if o == nil || o.Monitor == nil {
		var ret MonitorTypes
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringOptions) GetMonitorOk() (*MonitorTypes, bool) {
	if o == nil || o.Monitor == nil {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *MonitoringOptions) HasMonitor() bool {
	if o != nil && o.Monitor != nil {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given MonitorTypes and assigns it to the Monitor field.
func (o *MonitoringOptions) SetMonitor(v MonitorTypes) {
	o.Monitor = &v
}

func (o MonitoringOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IgnoreEpisodesWithFiles != nil {
		toSerialize["ignoreEpisodesWithFiles"] = o.IgnoreEpisodesWithFiles
	}
	if o.IgnoreEpisodesWithoutFiles != nil {
		toSerialize["ignoreEpisodesWithoutFiles"] = o.IgnoreEpisodesWithoutFiles
	}
	if o.Monitor != nil {
		toSerialize["monitor"] = o.Monitor
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringOptions struct {
	value *MonitoringOptions
	isSet bool
}

func (v NullableMonitoringOptions) Get() *MonitoringOptions {
	return v.value
}

func (v *NullableMonitoringOptions) Set(val *MonitoringOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringOptions(val *MonitoringOptions) *NullableMonitoringOptions {
	return &NullableMonitoringOptions{value: val, isSet: true}
}

func (v NullableMonitoringOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


