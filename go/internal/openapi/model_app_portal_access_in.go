/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AppPortalAccessIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPortalAccessIn{}

// AppPortalAccessIn struct for AppPortalAccessIn
type AppPortalAccessIn struct {
	// The set of feature flags the created token will have access to.
	FeatureFlags []string `json:"featureFlags,omitempty"`
}

// NewAppPortalAccessIn instantiates a new AppPortalAccessIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPortalAccessIn() *AppPortalAccessIn {
	this := AppPortalAccessIn{}
	return &this
}

// NewAppPortalAccessInWithDefaults instantiates a new AppPortalAccessIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPortalAccessInWithDefaults() *AppPortalAccessIn {
	this := AppPortalAccessIn{}
	return &this
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *AppPortalAccessIn) GetFeatureFlags() []string {
	if o == nil || IsNil(o.FeatureFlags) {
		var ret []string
		return ret
	}
	return o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPortalAccessIn) GetFeatureFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.FeatureFlags) {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *AppPortalAccessIn) HasFeatureFlags() bool {
	if o != nil && !IsNil(o.FeatureFlags) {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given []string and assigns it to the FeatureFlags field.
func (o *AppPortalAccessIn) SetFeatureFlags(v []string) {
	o.FeatureFlags = v
}

func (o AppPortalAccessIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPortalAccessIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureFlags) {
		toSerialize["featureFlags"] = o.FeatureFlags
	}
	return toSerialize, nil
}

type NullableAppPortalAccessIn struct {
	value *AppPortalAccessIn
	isSet bool
}

func (v NullableAppPortalAccessIn) Get() *AppPortalAccessIn {
	return v.value
}

func (v *NullableAppPortalAccessIn) Set(val *AppPortalAccessIn) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPortalAccessIn) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPortalAccessIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPortalAccessIn(val *AppPortalAccessIn) *NullableAppPortalAccessIn {
	return &NullableAppPortalAccessIn{value: val, isSet: true}
}

func (v NullableAppPortalAccessIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPortalAccessIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


