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

// checks if the EventTypePatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventTypePatch{}

// EventTypePatch struct for EventTypePatch
type EventTypePatch struct {
	Archived *bool `json:"archived,omitempty"`
	Description *string `json:"description,omitempty"`
	FeatureFlag NullableString `json:"featureFlag,omitempty"`
	Schemas map[string]map[string]interface{} `json:"schemas,omitempty"`
}

// NewEventTypePatch instantiates a new EventTypePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypePatch() *EventTypePatch {
	this := EventTypePatch{}
	return &this
}

// NewEventTypePatchWithDefaults instantiates a new EventTypePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypePatchWithDefaults() *EventTypePatch {
	this := EventTypePatch{}
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *EventTypePatch) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypePatch) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *EventTypePatch) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *EventTypePatch) SetArchived(v bool) {
	o.Archived = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventTypePatch) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypePatch) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventTypePatch) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventTypePatch) SetDescription(v string) {
	o.Description = &v
}

// GetFeatureFlag returns the FeatureFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypePatch) GetFeatureFlag() string {
	if o == nil || IsNil(o.FeatureFlag.Get()) {
		var ret string
		return ret
	}
	return *o.FeatureFlag.Get()
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypePatch) GetFeatureFlagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlag.Get(), o.FeatureFlag.IsSet()
}

// HasFeatureFlag returns a boolean if a field has been set.
func (o *EventTypePatch) HasFeatureFlag() bool {
	if o != nil && o.FeatureFlag.IsSet() {
		return true
	}

	return false
}

// SetFeatureFlag gets a reference to the given NullableString and assigns it to the FeatureFlag field.
func (o *EventTypePatch) SetFeatureFlag(v string) {
	o.FeatureFlag.Set(&v)
}
// SetFeatureFlagNil sets the value for FeatureFlag to be an explicit nil
func (o *EventTypePatch) SetFeatureFlagNil() {
	o.FeatureFlag.Set(nil)
}

// UnsetFeatureFlag ensures that no value is present for FeatureFlag, not even an explicit nil
func (o *EventTypePatch) UnsetFeatureFlag() {
	o.FeatureFlag.Unset()
}

// GetSchemas returns the Schemas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypePatch) GetSchemas() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypePatch) GetSchemasOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Schemas) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *EventTypePatch) HasSchemas() bool {
	if o != nil && IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given map[string]map[string]interface{} and assigns it to the Schemas field.
func (o *EventTypePatch) SetSchemas(v map[string]map[string]interface{}) {
	o.Schemas = v
}

func (o EventTypePatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventTypePatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.FeatureFlag.IsSet() {
		toSerialize["featureFlag"] = o.FeatureFlag.Get()
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	return toSerialize, nil
}

type NullableEventTypePatch struct {
	value *EventTypePatch
	isSet bool
}

func (v NullableEventTypePatch) Get() *EventTypePatch {
	return v.value
}

func (v *NullableEventTypePatch) Set(val *EventTypePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypePatch(val *EventTypePatch) *NullableEventTypePatch {
	return &NullableEventTypePatch{value: val, isSet: true}
}

func (v NullableEventTypePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


