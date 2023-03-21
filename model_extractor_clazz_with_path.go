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

// ExtractorClazzWithPath struct for ExtractorClazzWithPath
type ExtractorClazzWithPath struct {
	// which contains language-specific contents
	Extras map[string]interface{} `json:"extras,omitempty"`
	// language
	Lang   *string   `json:"lang,omitempty"`
	Module *string   `json:"module,omitempty"`
	Name   *string   `json:"name,omitempty"`
	Path   *string   `json:"path,omitempty"`
	Span   *CoreSpan `json:"span,omitempty"`
}

// NewExtractorClazzWithPath instantiates a new ExtractorClazzWithPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractorClazzWithPath() *ExtractorClazzWithPath {
	this := ExtractorClazzWithPath{}
	return &this
}

// NewExtractorClazzWithPathWithDefaults instantiates a new ExtractorClazzWithPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractorClazzWithPathWithDefaults() *ExtractorClazzWithPath {
	this := ExtractorClazzWithPath{}
	return &this
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *ExtractorClazzWithPath) GetExtras() map[string]interface{} {
	if o == nil || isNil(o.Extras) {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractorClazzWithPath) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Extras) {
		return map[string]interface{}{}, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *ExtractorClazzWithPath) HasExtras() bool {
	if o != nil && !isNil(o.Extras) {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *ExtractorClazzWithPath) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *ExtractorClazzWithPath) GetLang() string {
	if o == nil || isNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractorClazzWithPath) GetLangOk() (*string, bool) {
	if o == nil || isNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *ExtractorClazzWithPath) HasLang() bool {
	if o != nil && !isNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *ExtractorClazzWithPath) SetLang(v string) {
	o.Lang = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *ExtractorClazzWithPath) GetModule() string {
	if o == nil || isNil(o.Module) {
		var ret string
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractorClazzWithPath) GetModuleOk() (*string, bool) {
	if o == nil || isNil(o.Module) {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *ExtractorClazzWithPath) HasModule() bool {
	if o != nil && !isNil(o.Module) {
		return true
	}

	return false
}

// SetModule gets a reference to the given string and assigns it to the Module field.
func (o *ExtractorClazzWithPath) SetModule(v string) {
	o.Module = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtractorClazzWithPath) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractorClazzWithPath) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtractorClazzWithPath) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtractorClazzWithPath) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ExtractorClazzWithPath) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractorClazzWithPath) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ExtractorClazzWithPath) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ExtractorClazzWithPath) SetPath(v string) {
	o.Path = &v
}

// GetSpan returns the Span field value if set, zero value otherwise.
func (o *ExtractorClazzWithPath) GetSpan() CoreSpan {
	if o == nil || isNil(o.Span) {
		var ret CoreSpan
		return ret
	}
	return *o.Span
}

// GetSpanOk returns a tuple with the Span field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractorClazzWithPath) GetSpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.Span) {
		return nil, false
	}
	return o.Span, true
}

// HasSpan returns a boolean if a field has been set.
func (o *ExtractorClazzWithPath) HasSpan() bool {
	if o != nil && !isNil(o.Span) {
		return true
	}

	return false
}

// SetSpan gets a reference to the given CoreSpan and assigns it to the Span field.
func (o *ExtractorClazzWithPath) SetSpan(v CoreSpan) {
	o.Span = &v
}

func (o ExtractorClazzWithPath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Extras) {
		toSerialize["extras"] = o.Extras
	}
	if !isNil(o.Lang) {
		toSerialize["lang"] = o.Lang
	}
	if !isNil(o.Module) {
		toSerialize["module"] = o.Module
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Span) {
		toSerialize["span"] = o.Span
	}
	return json.Marshal(toSerialize)
}

type NullableExtractorClazzWithPath struct {
	value *ExtractorClazzWithPath
	isSet bool
}

func (v NullableExtractorClazzWithPath) Get() *ExtractorClazzWithPath {
	return v.value
}

func (v *NullableExtractorClazzWithPath) Set(val *ExtractorClazzWithPath) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractorClazzWithPath) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractorClazzWithPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractorClazzWithPath(val *ExtractorClazzWithPath) *NullableExtractorClazzWithPath {
	return &NullableExtractorClazzWithPath{value: val, isSet: true}
}

func (v NullableExtractorClazzWithPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractorClazzWithPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}