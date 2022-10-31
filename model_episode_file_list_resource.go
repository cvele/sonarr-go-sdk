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

// EpisodeFileListResource struct for EpisodeFileListResource
type EpisodeFileListResource struct {
	EpisodeFileIds []int32 `json:"episodeFileIds,omitempty"`
	Language *Language `json:"language,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	SceneName NullableString `json:"sceneName,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
}

// NewEpisodeFileListResource instantiates a new EpisodeFileListResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpisodeFileListResource() *EpisodeFileListResource {
	this := EpisodeFileListResource{}
	return &this
}

// NewEpisodeFileListResourceWithDefaults instantiates a new EpisodeFileListResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpisodeFileListResourceWithDefaults() *EpisodeFileListResource {
	this := EpisodeFileListResource{}
	return &this
}

// GetEpisodeFileIds returns the EpisodeFileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileListResource) GetEpisodeFileIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.EpisodeFileIds
}

// GetEpisodeFileIdsOk returns a tuple with the EpisodeFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileListResource) GetEpisodeFileIdsOk() ([]int32, bool) {
	if o == nil || o.EpisodeFileIds == nil {
		return nil, false
	}
	return o.EpisodeFileIds, true
}

// HasEpisodeFileIds returns a boolean if a field has been set.
func (o *EpisodeFileListResource) HasEpisodeFileIds() bool {
	if o != nil && o.EpisodeFileIds != nil {
		return true
	}

	return false
}

// SetEpisodeFileIds gets a reference to the given []int32 and assigns it to the EpisodeFileIds field.
func (o *EpisodeFileListResource) SetEpisodeFileIds(v []int32) {
	o.EpisodeFileIds = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *EpisodeFileListResource) GetLanguage() Language {
	if o == nil || o.Language == nil {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileListResource) GetLanguageOk() (*Language, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *EpisodeFileListResource) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *EpisodeFileListResource) SetLanguage(v Language) {
	o.Language = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *EpisodeFileListResource) GetQuality() QualityModel {
	if o == nil || o.Quality == nil {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeFileListResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || o.Quality == nil {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *EpisodeFileListResource) HasQuality() bool {
	if o != nil && o.Quality != nil {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *EpisodeFileListResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileListResource) GetSceneName() string {
	if o == nil || o.SceneName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SceneName.Get()
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileListResource) GetSceneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SceneName.Get(), o.SceneName.IsSet()
}

// HasSceneName returns a boolean if a field has been set.
func (o *EpisodeFileListResource) HasSceneName() bool {
	if o != nil && o.SceneName.IsSet() {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given NullableString and assigns it to the SceneName field.
func (o *EpisodeFileListResource) SetSceneName(v string) {
	o.SceneName.Set(&v)
}
// SetSceneNameNil sets the value for SceneName to be an explicit nil
func (o *EpisodeFileListResource) SetSceneNameNil() {
	o.SceneName.Set(nil)
}

// UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
func (o *EpisodeFileListResource) UnsetSceneName() {
	o.SceneName.Unset()
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EpisodeFileListResource) GetReleaseGroup() string {
	if o == nil || o.ReleaseGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EpisodeFileListResource) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *EpisodeFileListResource) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *EpisodeFileListResource) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *EpisodeFileListResource) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *EpisodeFileListResource) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

func (o EpisodeFileListResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EpisodeFileIds != nil {
		toSerialize["episodeFileIds"] = o.EpisodeFileIds
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Quality != nil {
		toSerialize["quality"] = o.Quality
	}
	if o.SceneName.IsSet() {
		toSerialize["sceneName"] = o.SceneName.Get()
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEpisodeFileListResource struct {
	value *EpisodeFileListResource
	isSet bool
}

func (v NullableEpisodeFileListResource) Get() *EpisodeFileListResource {
	return v.value
}

func (v *NullableEpisodeFileListResource) Set(val *EpisodeFileListResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEpisodeFileListResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEpisodeFileListResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpisodeFileListResource(val *EpisodeFileListResource) *NullableEpisodeFileListResource {
	return &NullableEpisodeFileListResource{value: val, isSet: true}
}

func (v NullableEpisodeFileListResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpisodeFileListResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


