/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EndpointTransformationSimulateIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointTransformationSimulateIn{}

// EndpointTransformationSimulateIn struct for EndpointTransformationSimulateIn
type EndpointTransformationSimulateIn struct {
	Channels []string `json:"channels,omitempty"`
	Code string `json:"code"`
	// The event type's name
	EventType string `json:"eventType"`
	Payload map[string]interface{} `json:"payload"`
}

type _EndpointTransformationSimulateIn EndpointTransformationSimulateIn

// NewEndpointTransformationSimulateIn instantiates a new EndpointTransformationSimulateIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTransformationSimulateIn(code string, eventType string, payload map[string]interface{}) *EndpointTransformationSimulateIn {
	this := EndpointTransformationSimulateIn{}
	this.Code = code
	this.EventType = eventType
	this.Payload = payload
	return &this
}

// NewEndpointTransformationSimulateInWithDefaults instantiates a new EndpointTransformationSimulateIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTransformationSimulateInWithDefaults() *EndpointTransformationSimulateIn {
	this := EndpointTransformationSimulateIn{}
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointTransformationSimulateIn) GetChannels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointTransformationSimulateIn) GetChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *EndpointTransformationSimulateIn) HasChannels() bool {
	if o != nil && IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *EndpointTransformationSimulateIn) SetChannels(v []string) {
	o.Channels = v
}

// GetCode returns the Code field value
func (o *EndpointTransformationSimulateIn) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EndpointTransformationSimulateIn) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EndpointTransformationSimulateIn) SetCode(v string) {
	o.Code = v
}

// GetEventType returns the EventType field value
func (o *EndpointTransformationSimulateIn) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EndpointTransformationSimulateIn) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EndpointTransformationSimulateIn) SetEventType(v string) {
	o.EventType = v
}

// GetPayload returns the Payload field value
func (o *EndpointTransformationSimulateIn) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *EndpointTransformationSimulateIn) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *EndpointTransformationSimulateIn) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

func (o EndpointTransformationSimulateIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTransformationSimulateIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	toSerialize["code"] = o.Code
	toSerialize["eventType"] = o.EventType
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

func (o *EndpointTransformationSimulateIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"eventType",
		"payload",
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

	varEndpointTransformationSimulateIn := _EndpointTransformationSimulateIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointTransformationSimulateIn)

	if err != nil {
		return err
	}

	*o = EndpointTransformationSimulateIn(varEndpointTransformationSimulateIn)

	return err
}

type NullableEndpointTransformationSimulateIn struct {
	value *EndpointTransformationSimulateIn
	isSet bool
}

func (v NullableEndpointTransformationSimulateIn) Get() *EndpointTransformationSimulateIn {
	return v.value
}

func (v *NullableEndpointTransformationSimulateIn) Set(val *EndpointTransformationSimulateIn) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTransformationSimulateIn) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTransformationSimulateIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTransformationSimulateIn(val *EndpointTransformationSimulateIn) *NullableEndpointTransformationSimulateIn {
	return &NullableEndpointTransformationSimulateIn{value: val, isSet: true}
}

func (v NullableEndpointTransformationSimulateIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTransformationSimulateIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


