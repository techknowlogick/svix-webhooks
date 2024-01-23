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

// checks if the ListResponseApplicationOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseApplicationOut{}

// ListResponseApplicationOut struct for ListResponseApplicationOut
type ListResponseApplicationOut struct {
	Data []ApplicationOut `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator,omitempty"`
	PrevIterator NullableString `json:"prevIterator,omitempty"`
}

type _ListResponseApplicationOut ListResponseApplicationOut

// NewListResponseApplicationOut instantiates a new ListResponseApplicationOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseApplicationOut(data []ApplicationOut, done bool) *ListResponseApplicationOut {
	this := ListResponseApplicationOut{}
	this.Data = data
	this.Done = done
	return &this
}

// NewListResponseApplicationOutWithDefaults instantiates a new ListResponseApplicationOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseApplicationOutWithDefaults() *ListResponseApplicationOut {
	this := ListResponseApplicationOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseApplicationOut) GetData() []ApplicationOut {
	if o == nil {
		var ret []ApplicationOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseApplicationOut) GetDataOk() ([]ApplicationOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseApplicationOut) SetData(v []ApplicationOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseApplicationOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseApplicationOut) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseApplicationOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseApplicationOut) GetIterator() string {
	if o == nil || IsNil(o.Iterator.Get()) {
		var ret string
		return ret
	}
	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseApplicationOut) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// HasIterator returns a boolean if a field has been set.
func (o *ListResponseApplicationOut) HasIterator() bool {
	if o != nil && o.Iterator.IsSet() {
		return true
	}

	return false
}

// SetIterator gets a reference to the given NullableString and assigns it to the Iterator field.
func (o *ListResponseApplicationOut) SetIterator(v string) {
	o.Iterator.Set(&v)
}
// SetIteratorNil sets the value for Iterator to be an explicit nil
func (o *ListResponseApplicationOut) SetIteratorNil() {
	o.Iterator.Set(nil)
}

// UnsetIterator ensures that no value is present for Iterator, not even an explicit nil
func (o *ListResponseApplicationOut) UnsetIterator() {
	o.Iterator.Unset()
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseApplicationOut) GetPrevIterator() string {
	if o == nil || IsNil(o.PrevIterator.Get()) {
		var ret string
		return ret
	}
	return *o.PrevIterator.Get()
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseApplicationOut) GetPrevIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevIterator.Get(), o.PrevIterator.IsSet()
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseApplicationOut) HasPrevIterator() bool {
	if o != nil && o.PrevIterator.IsSet() {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given NullableString and assigns it to the PrevIterator field.
func (o *ListResponseApplicationOut) SetPrevIterator(v string) {
	o.PrevIterator.Set(&v)
}
// SetPrevIteratorNil sets the value for PrevIterator to be an explicit nil
func (o *ListResponseApplicationOut) SetPrevIteratorNil() {
	o.PrevIterator.Set(nil)
}

// UnsetPrevIterator ensures that no value is present for PrevIterator, not even an explicit nil
func (o *ListResponseApplicationOut) UnsetPrevIterator() {
	o.PrevIterator.Unset()
}

func (o ListResponseApplicationOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseApplicationOut) ToMap() (map[string]interface{}, error) {
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

func (o *ListResponseApplicationOut) UnmarshalJSON(data []byte) (err error) {
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

	varListResponseApplicationOut := _ListResponseApplicationOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseApplicationOut)

	if err != nil {
		return err
	}

	*o = ListResponseApplicationOut(varListResponseApplicationOut)

	return err
}

type NullableListResponseApplicationOut struct {
	value *ListResponseApplicationOut
	isSet bool
}

func (v NullableListResponseApplicationOut) Get() *ListResponseApplicationOut {
	return v.value
}

func (v *NullableListResponseApplicationOut) Set(val *ListResponseApplicationOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseApplicationOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseApplicationOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseApplicationOut(val *ListResponseApplicationOut) *NullableListResponseApplicationOut {
	return &NullableListResponseApplicationOut{value: val, isSet: true}
}

func (v NullableListResponseApplicationOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseApplicationOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


