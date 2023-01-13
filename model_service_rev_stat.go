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

// ServiceRevStat struct for ServiceRevStat
type ServiceRevStat struct {
	FileCount     *int32         `json:"fileCount,omitempty"`
	FunctionCount *int32         `json:"functionCount,omitempty"`
	Info          *ObjectRevInfo `json:"info,omitempty"`
}

// NewServiceRevStat instantiates a new ServiceRevStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRevStat() *ServiceRevStat {
	this := ServiceRevStat{}
	return &this
}

// NewServiceRevStatWithDefaults instantiates a new ServiceRevStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRevStatWithDefaults() *ServiceRevStat {
	this := ServiceRevStat{}
	return &this
}

// GetFileCount returns the FileCount field value if set, zero value otherwise.
func (o *ServiceRevStat) GetFileCount() int32 {
	if o == nil || isNil(o.FileCount) {
		var ret int32
		return ret
	}
	return *o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRevStat) GetFileCountOk() (*int32, bool) {
	if o == nil || isNil(o.FileCount) {
		return nil, false
	}
	return o.FileCount, true
}

// HasFileCount returns a boolean if a field has been set.
func (o *ServiceRevStat) HasFileCount() bool {
	if o != nil && !isNil(o.FileCount) {
		return true
	}

	return false
}

// SetFileCount gets a reference to the given int32 and assigns it to the FileCount field.
func (o *ServiceRevStat) SetFileCount(v int32) {
	o.FileCount = &v
}

// GetFunctionCount returns the FunctionCount field value if set, zero value otherwise.
func (o *ServiceRevStat) GetFunctionCount() int32 {
	if o == nil || isNil(o.FunctionCount) {
		var ret int32
		return ret
	}
	return *o.FunctionCount
}

// GetFunctionCountOk returns a tuple with the FunctionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRevStat) GetFunctionCountOk() (*int32, bool) {
	if o == nil || isNil(o.FunctionCount) {
		return nil, false
	}
	return o.FunctionCount, true
}

// HasFunctionCount returns a boolean if a field has been set.
func (o *ServiceRevStat) HasFunctionCount() bool {
	if o != nil && !isNil(o.FunctionCount) {
		return true
	}

	return false
}

// SetFunctionCount gets a reference to the given int32 and assigns it to the FunctionCount field.
func (o *ServiceRevStat) SetFunctionCount(v int32) {
	o.FunctionCount = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ServiceRevStat) GetInfo() ObjectRevInfo {
	if o == nil || isNil(o.Info) {
		var ret ObjectRevInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRevStat) GetInfoOk() (*ObjectRevInfo, bool) {
	if o == nil || isNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ServiceRevStat) HasInfo() bool {
	if o != nil && !isNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given ObjectRevInfo and assigns it to the Info field.
func (o *ServiceRevStat) SetInfo(v ObjectRevInfo) {
	o.Info = &v
}

func (o ServiceRevStat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FileCount) {
		toSerialize["fileCount"] = o.FileCount
	}
	if !isNil(o.FunctionCount) {
		toSerialize["functionCount"] = o.FunctionCount
	}
	if !isNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	return json.Marshal(toSerialize)
}

type NullableServiceRevStat struct {
	value *ServiceRevStat
	isSet bool
}

func (v NullableServiceRevStat) Get() *ServiceRevStat {
	return v.value
}

func (v *NullableServiceRevStat) Set(val *ServiceRevStat) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRevStat) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRevStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRevStat(val *ServiceRevStat) *NullableServiceRevStat {
	return &NullableServiceRevStat{value: val, isSet: true}
}

func (v NullableServiceRevStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRevStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}