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

// DelayProfileResource struct for DelayProfileResource
type DelayProfileResource struct {
	Id *int32 `json:"id,omitempty"`
	EnableUsenet *bool `json:"enableUsenet,omitempty"`
	EnableTorrent *bool `json:"enableTorrent,omitempty"`
	PreferredProtocol *DownloadProtocol `json:"preferredProtocol,omitempty"`
	UsenetDelay *int32 `json:"usenetDelay,omitempty"`
	TorrentDelay *int32 `json:"torrentDelay,omitempty"`
	BypassIfHighestQuality *bool `json:"bypassIfHighestQuality,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
}

// NewDelayProfileResource instantiates a new DelayProfileResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelayProfileResource() *DelayProfileResource {
	this := DelayProfileResource{}
	return &this
}

// NewDelayProfileResourceWithDefaults instantiates a new DelayProfileResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelayProfileResourceWithDefaults() *DelayProfileResource {
	this := DelayProfileResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DelayProfileResource) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DelayProfileResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DelayProfileResource) SetId(v int32) {
	o.Id = &v
}

// GetEnableUsenet returns the EnableUsenet field value if set, zero value otherwise.
func (o *DelayProfileResource) GetEnableUsenet() bool {
	if o == nil || o.EnableUsenet == nil {
		var ret bool
		return ret
	}
	return *o.EnableUsenet
}

// GetEnableUsenetOk returns a tuple with the EnableUsenet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetEnableUsenetOk() (*bool, bool) {
	if o == nil || o.EnableUsenet == nil {
		return nil, false
	}
	return o.EnableUsenet, true
}

// HasEnableUsenet returns a boolean if a field has been set.
func (o *DelayProfileResource) HasEnableUsenet() bool {
	if o != nil && o.EnableUsenet != nil {
		return true
	}

	return false
}

// SetEnableUsenet gets a reference to the given bool and assigns it to the EnableUsenet field.
func (o *DelayProfileResource) SetEnableUsenet(v bool) {
	o.EnableUsenet = &v
}

// GetEnableTorrent returns the EnableTorrent field value if set, zero value otherwise.
func (o *DelayProfileResource) GetEnableTorrent() bool {
	if o == nil || o.EnableTorrent == nil {
		var ret bool
		return ret
	}
	return *o.EnableTorrent
}

// GetEnableTorrentOk returns a tuple with the EnableTorrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetEnableTorrentOk() (*bool, bool) {
	if o == nil || o.EnableTorrent == nil {
		return nil, false
	}
	return o.EnableTorrent, true
}

// HasEnableTorrent returns a boolean if a field has been set.
func (o *DelayProfileResource) HasEnableTorrent() bool {
	if o != nil && o.EnableTorrent != nil {
		return true
	}

	return false
}

// SetEnableTorrent gets a reference to the given bool and assigns it to the EnableTorrent field.
func (o *DelayProfileResource) SetEnableTorrent(v bool) {
	o.EnableTorrent = &v
}

// GetPreferredProtocol returns the PreferredProtocol field value if set, zero value otherwise.
func (o *DelayProfileResource) GetPreferredProtocol() DownloadProtocol {
	if o == nil || o.PreferredProtocol == nil {
		var ret DownloadProtocol
		return ret
	}
	return *o.PreferredProtocol
}

// GetPreferredProtocolOk returns a tuple with the PreferredProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetPreferredProtocolOk() (*DownloadProtocol, bool) {
	if o == nil || o.PreferredProtocol == nil {
		return nil, false
	}
	return o.PreferredProtocol, true
}

// HasPreferredProtocol returns a boolean if a field has been set.
func (o *DelayProfileResource) HasPreferredProtocol() bool {
	if o != nil && o.PreferredProtocol != nil {
		return true
	}

	return false
}

// SetPreferredProtocol gets a reference to the given DownloadProtocol and assigns it to the PreferredProtocol field.
func (o *DelayProfileResource) SetPreferredProtocol(v DownloadProtocol) {
	o.PreferredProtocol = &v
}

// GetUsenetDelay returns the UsenetDelay field value if set, zero value otherwise.
func (o *DelayProfileResource) GetUsenetDelay() int32 {
	if o == nil || o.UsenetDelay == nil {
		var ret int32
		return ret
	}
	return *o.UsenetDelay
}

// GetUsenetDelayOk returns a tuple with the UsenetDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetUsenetDelayOk() (*int32, bool) {
	if o == nil || o.UsenetDelay == nil {
		return nil, false
	}
	return o.UsenetDelay, true
}

// HasUsenetDelay returns a boolean if a field has been set.
func (o *DelayProfileResource) HasUsenetDelay() bool {
	if o != nil && o.UsenetDelay != nil {
		return true
	}

	return false
}

// SetUsenetDelay gets a reference to the given int32 and assigns it to the UsenetDelay field.
func (o *DelayProfileResource) SetUsenetDelay(v int32) {
	o.UsenetDelay = &v
}

// GetTorrentDelay returns the TorrentDelay field value if set, zero value otherwise.
func (o *DelayProfileResource) GetTorrentDelay() int32 {
	if o == nil || o.TorrentDelay == nil {
		var ret int32
		return ret
	}
	return *o.TorrentDelay
}

// GetTorrentDelayOk returns a tuple with the TorrentDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetTorrentDelayOk() (*int32, bool) {
	if o == nil || o.TorrentDelay == nil {
		return nil, false
	}
	return o.TorrentDelay, true
}

// HasTorrentDelay returns a boolean if a field has been set.
func (o *DelayProfileResource) HasTorrentDelay() bool {
	if o != nil && o.TorrentDelay != nil {
		return true
	}

	return false
}

// SetTorrentDelay gets a reference to the given int32 and assigns it to the TorrentDelay field.
func (o *DelayProfileResource) SetTorrentDelay(v int32) {
	o.TorrentDelay = &v
}

// GetBypassIfHighestQuality returns the BypassIfHighestQuality field value if set, zero value otherwise.
func (o *DelayProfileResource) GetBypassIfHighestQuality() bool {
	if o == nil || o.BypassIfHighestQuality == nil {
		var ret bool
		return ret
	}
	return *o.BypassIfHighestQuality
}

// GetBypassIfHighestQualityOk returns a tuple with the BypassIfHighestQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetBypassIfHighestQualityOk() (*bool, bool) {
	if o == nil || o.BypassIfHighestQuality == nil {
		return nil, false
	}
	return o.BypassIfHighestQuality, true
}

// HasBypassIfHighestQuality returns a boolean if a field has been set.
func (o *DelayProfileResource) HasBypassIfHighestQuality() bool {
	if o != nil && o.BypassIfHighestQuality != nil {
		return true
	}

	return false
}

// SetBypassIfHighestQuality gets a reference to the given bool and assigns it to the BypassIfHighestQuality field.
func (o *DelayProfileResource) SetBypassIfHighestQuality(v bool) {
	o.BypassIfHighestQuality = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *DelayProfileResource) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelayProfileResource) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *DelayProfileResource) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *DelayProfileResource) SetOrder(v int32) {
	o.Order = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DelayProfileResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DelayProfileResource) GetTagsOk() ([]int32, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DelayProfileResource) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *DelayProfileResource) SetTags(v []int32) {
	o.Tags = v
}

func (o DelayProfileResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EnableUsenet != nil {
		toSerialize["enableUsenet"] = o.EnableUsenet
	}
	if o.EnableTorrent != nil {
		toSerialize["enableTorrent"] = o.EnableTorrent
	}
	if o.PreferredProtocol != nil {
		toSerialize["preferredProtocol"] = o.PreferredProtocol
	}
	if o.UsenetDelay != nil {
		toSerialize["usenetDelay"] = o.UsenetDelay
	}
	if o.TorrentDelay != nil {
		toSerialize["torrentDelay"] = o.TorrentDelay
	}
	if o.BypassIfHighestQuality != nil {
		toSerialize["bypassIfHighestQuality"] = o.BypassIfHighestQuality
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableDelayProfileResource struct {
	value *DelayProfileResource
	isSet bool
}

func (v NullableDelayProfileResource) Get() *DelayProfileResource {
	return v.value
}

func (v *NullableDelayProfileResource) Set(val *DelayProfileResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDelayProfileResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDelayProfileResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelayProfileResource(val *DelayProfileResource) *NullableDelayProfileResource {
	return &NullableDelayProfileResource{value: val, isSet: true}
}

func (v NullableDelayProfileResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelayProfileResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


