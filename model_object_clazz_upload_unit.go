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

// ObjectClazzUploadUnit struct for ObjectClazzUploadUnit
type ObjectClazzUploadUnit struct {
	ClazzFileResult *ExtractorClazzFileResult `json:"clazzFileResult,omitempty"`
	Workspace       *ObjectWorkspaceConfig    `json:"workspace,omitempty"`
}

// NewObjectClazzUploadUnit instantiates a new ObjectClazzUploadUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectClazzUploadUnit() *ObjectClazzUploadUnit {
	this := ObjectClazzUploadUnit{}
	return &this
}

// NewObjectClazzUploadUnitWithDefaults instantiates a new ObjectClazzUploadUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectClazzUploadUnitWithDefaults() *ObjectClazzUploadUnit {
	this := ObjectClazzUploadUnit{}
	return &this
}

// GetClazzFileResult returns the ClazzFileResult field value if set, zero value otherwise.
func (o *ObjectClazzUploadUnit) GetClazzFileResult() ExtractorClazzFileResult {
	if o == nil || isNil(o.ClazzFileResult) {
		var ret ExtractorClazzFileResult
		return ret
	}
	return *o.ClazzFileResult
}

// GetClazzFileResultOk returns a tuple with the ClazzFileResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectClazzUploadUnit) GetClazzFileResultOk() (*ExtractorClazzFileResult, bool) {
	if o == nil || isNil(o.ClazzFileResult) {
		return nil, false
	}
	return o.ClazzFileResult, true
}

// HasClazzFileResult returns a boolean if a field has been set.
func (o *ObjectClazzUploadUnit) HasClazzFileResult() bool {
	if o != nil && !isNil(o.ClazzFileResult) {
		return true
	}

	return false
}

// SetClazzFileResult gets a reference to the given ExtractorClazzFileResult and assigns it to the ClazzFileResult field.
func (o *ObjectClazzUploadUnit) SetClazzFileResult(v ExtractorClazzFileResult) {
	o.ClazzFileResult = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *ObjectClazzUploadUnit) GetWorkspace() ObjectWorkspaceConfig {
	if o == nil || isNil(o.Workspace) {
		var ret ObjectWorkspaceConfig
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectClazzUploadUnit) GetWorkspaceOk() (*ObjectWorkspaceConfig, bool) {
	if o == nil || isNil(o.Workspace) {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *ObjectClazzUploadUnit) HasWorkspace() bool {
	if o != nil && !isNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given ObjectWorkspaceConfig and assigns it to the Workspace field.
func (o *ObjectClazzUploadUnit) SetWorkspace(v ObjectWorkspaceConfig) {
	o.Workspace = &v
}

func (o ObjectClazzUploadUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClazzFileResult) {
		toSerialize["clazzFileResult"] = o.ClazzFileResult
	}
	if !isNil(o.Workspace) {
		toSerialize["workspace"] = o.Workspace
	}
	return json.Marshal(toSerialize)
}

type NullableObjectClazzUploadUnit struct {
	value *ObjectClazzUploadUnit
	isSet bool
}

func (v NullableObjectClazzUploadUnit) Get() *ObjectClazzUploadUnit {
	return v.value
}

func (v *NullableObjectClazzUploadUnit) Set(val *ObjectClazzUploadUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectClazzUploadUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectClazzUploadUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectClazzUploadUnit(val *ObjectClazzUploadUnit) *NullableObjectClazzUploadUnit {
	return &NullableObjectClazzUploadUnit{value: val, isSet: true}
}

func (v NullableObjectClazzUploadUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectClazzUploadUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}