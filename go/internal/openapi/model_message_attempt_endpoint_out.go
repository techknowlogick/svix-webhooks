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

// checks if the MessageAttemptEndpointOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAttemptEndpointOut{}

// MessageAttemptEndpointOut struct for MessageAttemptEndpointOut
type MessageAttemptEndpointOut struct {
	// The ep's ID
	EndpointId string `json:"endpointId"`
	// The attempt's ID
	Id string `json:"id"`
	Msg *MessageOut `json:"msg,omitempty"`
	// The msg's ID
	MsgId string `json:"msgId"`
	Response string `json:"response"`
	ResponseStatusCode int32 `json:"responseStatusCode"`
	Status MessageStatus `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	TriggerType MessageAttemptTriggerType `json:"triggerType"`
	Url string `json:"url"`
}

type _MessageAttemptEndpointOut MessageAttemptEndpointOut

// NewMessageAttemptEndpointOut instantiates a new MessageAttemptEndpointOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttemptEndpointOut(endpointId string, id string, msgId string, response string, responseStatusCode int32, status MessageStatus, timestamp time.Time, triggerType MessageAttemptTriggerType, url string) *MessageAttemptEndpointOut {
	this := MessageAttemptEndpointOut{}
	this.EndpointId = endpointId
	this.Id = id
	this.MsgId = msgId
	this.Response = response
	this.ResponseStatusCode = responseStatusCode
	this.Status = status
	this.Timestamp = timestamp
	this.TriggerType = triggerType
	this.Url = url
	return &this
}

// NewMessageAttemptEndpointOutWithDefaults instantiates a new MessageAttemptEndpointOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttemptEndpointOutWithDefaults() *MessageAttemptEndpointOut {
	this := MessageAttemptEndpointOut{}
	return &this
}

// GetEndpointId returns the EndpointId field value
func (o *MessageAttemptEndpointOut) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *MessageAttemptEndpointOut) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetId returns the Id field value
func (o *MessageAttemptEndpointOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageAttemptEndpointOut) SetId(v string) {
	o.Id = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *MessageAttemptEndpointOut) GetMsg() MessageOut {
	if o == nil || IsNil(o.Msg) {
		var ret MessageOut
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetMsgOk() (*MessageOut, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *MessageAttemptEndpointOut) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given MessageOut and assigns it to the Msg field.
func (o *MessageAttemptEndpointOut) SetMsg(v MessageOut) {
	o.Msg = &v
}

// GetMsgId returns the MsgId field value
func (o *MessageAttemptEndpointOut) GetMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetMsgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsgId, true
}

// SetMsgId sets field value
func (o *MessageAttemptEndpointOut) SetMsgId(v string) {
	o.MsgId = v
}

// GetResponse returns the Response field value
func (o *MessageAttemptEndpointOut) GetResponse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *MessageAttemptEndpointOut) SetResponse(v string) {
	o.Response = v
}

// GetResponseStatusCode returns the ResponseStatusCode field value
func (o *MessageAttemptEndpointOut) GetResponseStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseStatusCode
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetResponseStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseStatusCode, true
}

// SetResponseStatusCode sets field value
func (o *MessageAttemptEndpointOut) SetResponseStatusCode(v int32) {
	o.ResponseStatusCode = v
}

// GetStatus returns the Status field value
func (o *MessageAttemptEndpointOut) GetStatus() MessageStatus {
	if o == nil {
		var ret MessageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetStatusOk() (*MessageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MessageAttemptEndpointOut) SetStatus(v MessageStatus) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
func (o *MessageAttemptEndpointOut) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MessageAttemptEndpointOut) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTriggerType returns the TriggerType field value
func (o *MessageAttemptEndpointOut) GetTriggerType() MessageAttemptTriggerType {
	if o == nil {
		var ret MessageAttemptTriggerType
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetTriggerTypeOk() (*MessageAttemptTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *MessageAttemptEndpointOut) SetTriggerType(v MessageAttemptTriggerType) {
	o.TriggerType = v
}

// GetUrl returns the Url field value
func (o *MessageAttemptEndpointOut) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptEndpointOut) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MessageAttemptEndpointOut) SetUrl(v string) {
	o.Url = v
}

func (o MessageAttemptEndpointOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAttemptEndpointOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpointId"] = o.EndpointId
	toSerialize["id"] = o.Id
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	toSerialize["msgId"] = o.MsgId
	toSerialize["response"] = o.Response
	toSerialize["responseStatusCode"] = o.ResponseStatusCode
	toSerialize["status"] = o.Status
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["triggerType"] = o.TriggerType
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *MessageAttemptEndpointOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endpointId",
		"id",
		"msgId",
		"response",
		"responseStatusCode",
		"status",
		"timestamp",
		"triggerType",
		"url",
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

	varMessageAttemptEndpointOut := _MessageAttemptEndpointOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageAttemptEndpointOut)

	if err != nil {
		return err
	}

	*o = MessageAttemptEndpointOut(varMessageAttemptEndpointOut)

	return err
}

type NullableMessageAttemptEndpointOut struct {
	value *MessageAttemptEndpointOut
	isSet bool
}

func (v NullableMessageAttemptEndpointOut) Get() *MessageAttemptEndpointOut {
	return v.value
}

func (v *NullableMessageAttemptEndpointOut) Set(val *MessageAttemptEndpointOut) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttemptEndpointOut) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttemptEndpointOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttemptEndpointOut(val *MessageAttemptEndpointOut) *NullableMessageAttemptEndpointOut {
	return &NullableMessageAttemptEndpointOut{value: val, isSet: true}
}

func (v NullableMessageAttemptEndpointOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttemptEndpointOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


