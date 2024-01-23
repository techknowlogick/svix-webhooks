/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the MessageEndpointOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageEndpointOut{}

// MessageEndpointOut struct for MessageEndpointOut
type MessageEndpointOut struct {
	// List of message channels this endpoint listens to (omit for all)
	Channels []string `json:"channels,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	// An example endpoint name
	Description string `json:"description"`
	Disabled *bool `json:"disabled,omitempty"`
	FilterTypes []string `json:"filterTypes,omitempty"`
	// The ep's ID
	Id string `json:"id"`
	NextAttempt NullableTime `json:"nextAttempt,omitempty"`
	RateLimit NullableInt32 `json:"rateLimit,omitempty"`
	Status MessageStatus `json:"status"`
	// Optional unique identifier for the endpoint
	Uid NullableString `json:"uid,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	Url string `json:"url"`
	// Deprecated
	Version int32 `json:"version"`
}

type _MessageEndpointOut MessageEndpointOut

// NewMessageEndpointOut instantiates a new MessageEndpointOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageEndpointOut(createdAt time.Time, description string, id string, status MessageStatus, updatedAt time.Time, url string, version int32) *MessageEndpointOut {
	this := MessageEndpointOut{}
	this.CreatedAt = createdAt
	this.Description = description
	var disabled bool = false
	this.Disabled = &disabled
	this.Id = id
	this.Status = status
	this.UpdatedAt = updatedAt
	this.Url = url
	this.Version = version
	return &this
}

// NewMessageEndpointOutWithDefaults instantiates a new MessageEndpointOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageEndpointOutWithDefaults() *MessageEndpointOut {
	this := MessageEndpointOut{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEndpointOut) GetChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEndpointOut) GetChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *MessageEndpointOut) HasChannels() bool {
	if o != nil && IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *MessageEndpointOut) SetChannels(v []string) {
	o.Channels = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MessageEndpointOut) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MessageEndpointOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *MessageEndpointOut) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *MessageEndpointOut) SetDescription(v string) {
	o.Description = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *MessageEndpointOut) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *MessageEndpointOut) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *MessageEndpointOut) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEndpointOut) GetFilterTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEndpointOut) GetFilterTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterTypes) {
		return nil, false
	}
	return o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *MessageEndpointOut) HasFilterTypes() bool {
	if o != nil && IsNil(o.FilterTypes) {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *MessageEndpointOut) SetFilterTypes(v []string) {
	o.FilterTypes = v
}

// GetId returns the Id field value
func (o *MessageEndpointOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageEndpointOut) SetId(v string) {
	o.Id = v
}

// GetNextAttempt returns the NextAttempt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEndpointOut) GetNextAttempt() time.Time {
	if o == nil || IsNil(o.NextAttempt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.NextAttempt.Get()
}

// GetNextAttemptOk returns a tuple with the NextAttempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEndpointOut) GetNextAttemptOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextAttempt.Get(), o.NextAttempt.IsSet()
}

// HasNextAttempt returns a boolean if a field has been set.
func (o *MessageEndpointOut) HasNextAttempt() bool {
	if o != nil && o.NextAttempt.IsSet() {
		return true
	}

	return false
}

// SetNextAttempt gets a reference to the given NullableTime and assigns it to the NextAttempt field.
func (o *MessageEndpointOut) SetNextAttempt(v time.Time) {
	o.NextAttempt.Set(&v)
}
// SetNextAttemptNil sets the value for NextAttempt to be an explicit nil
func (o *MessageEndpointOut) SetNextAttemptNil() {
	o.NextAttempt.Set(nil)
}

// UnsetNextAttempt ensures that no value is present for NextAttempt, not even an explicit nil
func (o *MessageEndpointOut) UnsetNextAttempt() {
	o.NextAttempt.Unset()
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEndpointOut) GetRateLimit() int32 {
	if o == nil || IsNil(o.RateLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimit.Get()
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEndpointOut) GetRateLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimit.Get(), o.RateLimit.IsSet()
}

// HasRateLimit returns a boolean if a field has been set.
func (o *MessageEndpointOut) HasRateLimit() bool {
	if o != nil && o.RateLimit.IsSet() {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given NullableInt32 and assigns it to the RateLimit field.
func (o *MessageEndpointOut) SetRateLimit(v int32) {
	o.RateLimit.Set(&v)
}
// SetRateLimitNil sets the value for RateLimit to be an explicit nil
func (o *MessageEndpointOut) SetRateLimitNil() {
	o.RateLimit.Set(nil)
}

// UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
func (o *MessageEndpointOut) UnsetRateLimit() {
	o.RateLimit.Unset()
}

// GetStatus returns the Status field value
func (o *MessageEndpointOut) GetStatus() MessageStatus {
	if o == nil {
		var ret MessageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetStatusOk() (*MessageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MessageEndpointOut) SetStatus(v MessageStatus) {
	o.Status = v
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageEndpointOut) GetUid() string {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret string
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageEndpointOut) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *MessageEndpointOut) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableString and assigns it to the Uid field.
func (o *MessageEndpointOut) SetUid(v string) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *MessageEndpointOut) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *MessageEndpointOut) UnsetUid() {
	o.Uid.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *MessageEndpointOut) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *MessageEndpointOut) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUrl returns the Url field value
func (o *MessageEndpointOut) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MessageEndpointOut) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MessageEndpointOut) SetUrl(v string) {
	o.Url = v
}

// GetVersion returns the Version field value
// Deprecated
func (o *MessageEndpointOut) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *MessageEndpointOut) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
// Deprecated
func (o *MessageEndpointOut) SetVersion(v int32) {
	o.Version = v
}

func (o MessageEndpointOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageEndpointOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["description"] = o.Description
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if o.FilterTypes != nil {
		toSerialize["filterTypes"] = o.FilterTypes
	}
	toSerialize["id"] = o.Id
	if o.NextAttempt.IsSet() {
		toSerialize["nextAttempt"] = o.NextAttempt.Get()
	}
	if o.RateLimit.IsSet() {
		toSerialize["rateLimit"] = o.RateLimit.Get()
	}
	toSerialize["status"] = o.Status
	if o.Uid.IsSet() {
		toSerialize["uid"] = o.Uid.Get()
	}
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["url"] = o.Url
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *MessageEndpointOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"description",
		"id",
		"status",
		"updatedAt",
		"url",
		"version",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMessageEndpointOut := _MessageEndpointOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageEndpointOut)

	if err != nil {
		return err
	}

	*o = MessageEndpointOut(varMessageEndpointOut)

	return err
}

type NullableMessageEndpointOut struct {
	value *MessageEndpointOut
	isSet bool
}

func (v NullableMessageEndpointOut) Get() *MessageEndpointOut {
	return v.value
}

func (v *NullableMessageEndpointOut) Set(val *MessageEndpointOut) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageEndpointOut) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageEndpointOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageEndpointOut(val *MessageEndpointOut) *NullableMessageEndpointOut {
	return &NullableMessageEndpointOut{value: val, isSet: true}
}

func (v NullableMessageEndpointOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageEndpointOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


