/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EndpointUpdatedEventData Sent when an endpoint is created, updated, or deleted
type EndpointUpdatedEventData struct {
	// The app's ID
	AppId string `json:"appId"`
	// The app's UID
	AppUid NullableString `json:"appUid,omitempty"`
	// The ep's ID
	EndpointId string `json:"endpointId"`
	// The ep's UID
	EndpointUid NullableString `json:"endpointUid,omitempty"`
}

// NewEndpointUpdatedEventData instantiates a new EndpointUpdatedEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointUpdatedEventData(appId string, endpointId string) *EndpointUpdatedEventData {
	this := EndpointUpdatedEventData{}
	this.AppId = appId
	this.EndpointId = endpointId
	return &this
}

// NewEndpointUpdatedEventDataWithDefaults instantiates a new EndpointUpdatedEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointUpdatedEventDataWithDefaults() *EndpointUpdatedEventData {
	this := EndpointUpdatedEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *EndpointUpdatedEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *EndpointUpdatedEventData) GetAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *EndpointUpdatedEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointUpdatedEventData) GetAppUid() string {
	if o == nil || o.AppUid.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppUid.Get()
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointUpdatedEventData) GetAppUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppUid.Get(), o.AppUid.IsSet()
}

// HasAppUid returns a boolean if a field has been set.
func (o *EndpointUpdatedEventData) HasAppUid() bool {
	if o != nil && o.AppUid.IsSet() {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given NullableString and assigns it to the AppUid field.
func (o *EndpointUpdatedEventData) SetAppUid(v string) {
	o.AppUid.Set(&v)
}
// SetAppUidNil sets the value for AppUid to be an explicit nil
func (o *EndpointUpdatedEventData) SetAppUidNil() {
	o.AppUid.Set(nil)
}

// UnsetAppUid ensures that no value is present for AppUid, not even an explicit nil
func (o *EndpointUpdatedEventData) UnsetAppUid() {
	o.AppUid.Unset()
}

// GetEndpointId returns the EndpointId field value
func (o *EndpointUpdatedEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *EndpointUpdatedEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *EndpointUpdatedEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetEndpointUid returns the EndpointUid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointUpdatedEventData) GetEndpointUid() string {
	if o == nil || o.EndpointUid.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndpointUid.Get()
}

// GetEndpointUidOk returns a tuple with the EndpointUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointUpdatedEventData) GetEndpointUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndpointUid.Get(), o.EndpointUid.IsSet()
}

// HasEndpointUid returns a boolean if a field has been set.
func (o *EndpointUpdatedEventData) HasEndpointUid() bool {
	if o != nil && o.EndpointUid.IsSet() {
		return true
	}

	return false
}

// SetEndpointUid gets a reference to the given NullableString and assigns it to the EndpointUid field.
func (o *EndpointUpdatedEventData) SetEndpointUid(v string) {
	o.EndpointUid.Set(&v)
}
// SetEndpointUidNil sets the value for EndpointUid to be an explicit nil
func (o *EndpointUpdatedEventData) SetEndpointUidNil() {
	o.EndpointUid.Set(nil)
}

// UnsetEndpointUid ensures that no value is present for EndpointUid, not even an explicit nil
func (o *EndpointUpdatedEventData) UnsetEndpointUid() {
	o.EndpointUid.Unset()
}

func (o EndpointUpdatedEventData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if o.AppUid.IsSet() {
		toSerialize["appUid"] = o.AppUid.Get()
	}
	if true {
		toSerialize["endpointId"] = o.EndpointId
	}
	if o.EndpointUid.IsSet() {
		toSerialize["endpointUid"] = o.EndpointUid.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointUpdatedEventData struct {
	value *EndpointUpdatedEventData
	isSet bool
}

func (v NullableEndpointUpdatedEventData) Get() *EndpointUpdatedEventData {
	return v.value
}

func (v *NullableEndpointUpdatedEventData) Set(val *EndpointUpdatedEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointUpdatedEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointUpdatedEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointUpdatedEventData(val *EndpointUpdatedEventData) *NullableEndpointUpdatedEventData {
	return &NullableEndpointUpdatedEventData{value: val, isSet: true}
}

func (v NullableEndpointUpdatedEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointUpdatedEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


