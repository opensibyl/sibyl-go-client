/*
openapi for sibyl2 server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ObjectValueUnit struct for ObjectValueUnit
type ObjectValueUnit struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewObjectValueUnit instantiates a new ObjectValueUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectValueUnit() *ObjectValueUnit {
	this := ObjectValueUnit{}
	return &this
}

// NewObjectValueUnitWithDefaults instantiates a new ObjectValueUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectValueUnitWithDefaults() *ObjectValueUnit {
	this := ObjectValueUnit{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ObjectValueUnit) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectValueUnit) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ObjectValueUnit) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ObjectValueUnit) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ObjectValueUnit) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectValueUnit) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ObjectValueUnit) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ObjectValueUnit) SetType(v string) {
	o.Type = &v
}

func (o ObjectValueUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableObjectValueUnit struct {
	value *ObjectValueUnit
	isSet bool
}

func (v NullableObjectValueUnit) Get() *ObjectValueUnit {
	return v.value
}

func (v *NullableObjectValueUnit) Set(val *ObjectValueUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectValueUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectValueUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectValueUnit(val *ObjectValueUnit) *NullableObjectValueUnit {
	return &NullableObjectValueUnit{value: val, isSet: true}
}

func (v NullableObjectValueUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectValueUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
