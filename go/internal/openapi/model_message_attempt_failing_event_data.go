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

// checks if the MessageAttemptFailingEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAttemptFailingEventData{}

// MessageAttemptFailingEventData Sent when a message delivery has failed (all of the retry attempts have been exhausted) as a \"message.attempt.exhausted\" type or after it's failed four times as a \"message.attempt.failing\" event.
type MessageAttemptFailingEventData struct {
	// The app's ID
	AppId string `json:"appId"`
	// The app's UID
	AppUid NullableString `json:"appUid,omitempty"`
	// The ep's ID
	EndpointId string `json:"endpointId"`
	LastAttempt MessageAttemptFailedData `json:"lastAttempt"`
	// The msg's UID
	MsgEventId NullableString `json:"msgEventId,omitempty"`
	// The msg's ID
	MsgId string `json:"msgId"`
}

type _MessageAttemptFailingEventData MessageAttemptFailingEventData

// NewMessageAttemptFailingEventData instantiates a new MessageAttemptFailingEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttemptFailingEventData(appId string, endpointId string, lastAttempt MessageAttemptFailedData, msgId string) *MessageAttemptFailingEventData {
	this := MessageAttemptFailingEventData{}
	this.AppId = appId
	this.EndpointId = endpointId
	this.LastAttempt = lastAttempt
	this.MsgId = msgId
	return &this
}

// NewMessageAttemptFailingEventDataWithDefaults instantiates a new MessageAttemptFailingEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttemptFailingEventDataWithDefaults() *MessageAttemptFailingEventData {
	this := MessageAttemptFailingEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *MessageAttemptFailingEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptFailingEventData) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MessageAttemptFailingEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttemptFailingEventData) GetAppUid() string {
	if o == nil || IsNil(o.AppUid.Get()) {
		var ret string
		return ret
	}
	return *o.AppUid.Get()
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttemptFailingEventData) GetAppUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppUid.Get(), o.AppUid.IsSet()
}

// HasAppUid returns a boolean if a field has been set.
func (o *MessageAttemptFailingEventData) HasAppUid() bool {
	if o != nil && o.AppUid.IsSet() {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given NullableString and assigns it to the AppUid field.
func (o *MessageAttemptFailingEventData) SetAppUid(v string) {
	o.AppUid.Set(&v)
}
// SetAppUidNil sets the value for AppUid to be an explicit nil
func (o *MessageAttemptFailingEventData) SetAppUidNil() {
	o.AppUid.Set(nil)
}

// UnsetAppUid ensures that no value is present for AppUid, not even an explicit nil
func (o *MessageAttemptFailingEventData) UnsetAppUid() {
	o.AppUid.Unset()
}

// GetEndpointId returns the EndpointId field value
func (o *MessageAttemptFailingEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptFailingEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *MessageAttemptFailingEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetLastAttempt returns the LastAttempt field value
func (o *MessageAttemptFailingEventData) GetLastAttempt() MessageAttemptFailedData {
	if o == nil {
		var ret MessageAttemptFailedData
		return ret
	}

	return o.LastAttempt
}

// GetLastAttemptOk returns a tuple with the LastAttempt field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptFailingEventData) GetLastAttemptOk() (*MessageAttemptFailedData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastAttempt, true
}

// SetLastAttempt sets field value
func (o *MessageAttemptFailingEventData) SetLastAttempt(v MessageAttemptFailedData) {
	o.LastAttempt = v
}

// GetMsgEventId returns the MsgEventId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageAttemptFailingEventData) GetMsgEventId() string {
	if o == nil || IsNil(o.MsgEventId.Get()) {
		var ret string
		return ret
	}
	return *o.MsgEventId.Get()
}

// GetMsgEventIdOk returns a tuple with the MsgEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageAttemptFailingEventData) GetMsgEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MsgEventId.Get(), o.MsgEventId.IsSet()
}

// HasMsgEventId returns a boolean if a field has been set.
func (o *MessageAttemptFailingEventData) HasMsgEventId() bool {
	if o != nil && o.MsgEventId.IsSet() {
		return true
	}

	return false
}

// SetMsgEventId gets a reference to the given NullableString and assigns it to the MsgEventId field.
func (o *MessageAttemptFailingEventData) SetMsgEventId(v string) {
	o.MsgEventId.Set(&v)
}
// SetMsgEventIdNil sets the value for MsgEventId to be an explicit nil
func (o *MessageAttemptFailingEventData) SetMsgEventIdNil() {
	o.MsgEventId.Set(nil)
}

// UnsetMsgEventId ensures that no value is present for MsgEventId, not even an explicit nil
func (o *MessageAttemptFailingEventData) UnsetMsgEventId() {
	o.MsgEventId.Unset()
}

// GetMsgId returns the MsgId field value
func (o *MessageAttemptFailingEventData) GetMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptFailingEventData) GetMsgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsgId, true
}

// SetMsgId sets field value
func (o *MessageAttemptFailingEventData) SetMsgId(v string) {
	o.MsgId = v
}

func (o MessageAttemptFailingEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAttemptFailingEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	if o.AppUid.IsSet() {
		toSerialize["appUid"] = o.AppUid.Get()
	}
	toSerialize["endpointId"] = o.EndpointId
	toSerialize["lastAttempt"] = o.LastAttempt
	if o.MsgEventId.IsSet() {
		toSerialize["msgEventId"] = o.MsgEventId.Get()
	}
	toSerialize["msgId"] = o.MsgId
	return toSerialize, nil
}

func (o *MessageAttemptFailingEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appId",
		"endpointId",
		"lastAttempt",
		"msgId",
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

	varMessageAttemptFailingEventData := _MessageAttemptFailingEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageAttemptFailingEventData)

	if err != nil {
		return err
	}

	*o = MessageAttemptFailingEventData(varMessageAttemptFailingEventData)

	return err
}

type NullableMessageAttemptFailingEventData struct {
	value *MessageAttemptFailingEventData
	isSet bool
}

func (v NullableMessageAttemptFailingEventData) Get() *MessageAttemptFailingEventData {
	return v.value
}

func (v *NullableMessageAttemptFailingEventData) Set(val *MessageAttemptFailingEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttemptFailingEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttemptFailingEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttemptFailingEventData(val *MessageAttemptFailingEventData) *NullableMessageAttemptFailingEventData {
	return &NullableMessageAttemptFailingEventData{value: val, isSet: true}
}

func (v NullableMessageAttemptFailingEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttemptFailingEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


