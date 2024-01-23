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

// checks if the ListResponseEndpointMessageOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseEndpointMessageOut{}

// ListResponseEndpointMessageOut struct for ListResponseEndpointMessageOut
type ListResponseEndpointMessageOut struct {
	Data []EndpointMessageOut `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator,omitempty"`
	PrevIterator NullableString `json:"prevIterator,omitempty"`
}

type _ListResponseEndpointMessageOut ListResponseEndpointMessageOut

// NewListResponseEndpointMessageOut instantiates a new ListResponseEndpointMessageOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseEndpointMessageOut(data []EndpointMessageOut, done bool) *ListResponseEndpointMessageOut {
	this := ListResponseEndpointMessageOut{}
	this.Data = data
	this.Done = done
	return &this
}

// NewListResponseEndpointMessageOutWithDefaults instantiates a new ListResponseEndpointMessageOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseEndpointMessageOutWithDefaults() *ListResponseEndpointMessageOut {
	this := ListResponseEndpointMessageOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseEndpointMessageOut) GetData() []EndpointMessageOut {
	if o == nil {
		var ret []EndpointMessageOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseEndpointMessageOut) GetDataOk() ([]EndpointMessageOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseEndpointMessageOut) SetData(v []EndpointMessageOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseEndpointMessageOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseEndpointMessageOut) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseEndpointMessageOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseEndpointMessageOut) GetIterator() string {
	if o == nil || IsNil(o.Iterator.Get()) {
		var ret string
		return ret
	}
	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseEndpointMessageOut) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// HasIterator returns a boolean if a field has been set.
func (o *ListResponseEndpointMessageOut) HasIterator() bool {
	if o != nil && o.Iterator.IsSet() {
		return true
	}

	return false
}

// SetIterator gets a reference to the given NullableString and assigns it to the Iterator field.
func (o *ListResponseEndpointMessageOut) SetIterator(v string) {
	o.Iterator.Set(&v)
}
// SetIteratorNil sets the value for Iterator to be an explicit nil
func (o *ListResponseEndpointMessageOut) SetIteratorNil() {
	o.Iterator.Set(nil)
}

// UnsetIterator ensures that no value is present for Iterator, not even an explicit nil
func (o *ListResponseEndpointMessageOut) UnsetIterator() {
	o.Iterator.Unset()
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseEndpointMessageOut) GetPrevIterator() string {
	if o == nil || IsNil(o.PrevIterator.Get()) {
		var ret string
		return ret
	}
	return *o.PrevIterator.Get()
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseEndpointMessageOut) GetPrevIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevIterator.Get(), o.PrevIterator.IsSet()
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseEndpointMessageOut) HasPrevIterator() bool {
	if o != nil && o.PrevIterator.IsSet() {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given NullableString and assigns it to the PrevIterator field.
func (o *ListResponseEndpointMessageOut) SetPrevIterator(v string) {
	o.PrevIterator.Set(&v)
}
// SetPrevIteratorNil sets the value for PrevIterator to be an explicit nil
func (o *ListResponseEndpointMessageOut) SetPrevIteratorNil() {
	o.PrevIterator.Set(nil)
}

// UnsetPrevIterator ensures that no value is present for PrevIterator, not even an explicit nil
func (o *ListResponseEndpointMessageOut) UnsetPrevIterator() {
	o.PrevIterator.Unset()
}

func (o ListResponseEndpointMessageOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseEndpointMessageOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["done"] = o.Done
	if o.Iterator.IsSet() {
		toSerialize["iterator"] = o.Iterator.Get()
	}
	if o.PrevIterator.IsSet() {
		toSerialize["prevIterator"] = o.PrevIterator.Get()
	}
	return toSerialize, nil
}

func (o *ListResponseEndpointMessageOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"done",
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

	varListResponseEndpointMessageOut := _ListResponseEndpointMessageOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseEndpointMessageOut)

	if err != nil {
		return err
	}

	*o = ListResponseEndpointMessageOut(varListResponseEndpointMessageOut)

	return err
}

type NullableListResponseEndpointMessageOut struct {
	value *ListResponseEndpointMessageOut
	isSet bool
}

func (v NullableListResponseEndpointMessageOut) Get() *ListResponseEndpointMessageOut {
	return v.value
}

func (v *NullableListResponseEndpointMessageOut) Set(val *ListResponseEndpointMessageOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseEndpointMessageOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseEndpointMessageOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseEndpointMessageOut(val *ListResponseEndpointMessageOut) *NullableListResponseEndpointMessageOut {
	return &NullableListResponseEndpointMessageOut{value: val, isSet: true}
}

func (v NullableListResponseEndpointMessageOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseEndpointMessageOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


