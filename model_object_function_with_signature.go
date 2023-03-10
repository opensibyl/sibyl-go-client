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

// ObjectFunctionWithSignature struct for ObjectFunctionWithSignature
type ObjectFunctionWithSignature struct {
	BodySpan *CoreSpan `json:"bodySpan,omitempty"`
	// which contains language-specific contents
	Extras map[string]interface{} `json:"extras,omitempty"`
	// language
	Lang       *string           `json:"lang,omitempty"`
	Name       *string           `json:"name,omitempty"`
	Parameters []ObjectValueUnit `json:"parameters,omitempty"`
	Path       *string           `json:"path,omitempty"`
	Receiver   *string           `json:"receiver,omitempty"`
	Returns    []ObjectValueUnit `json:"returns,omitempty"`
	Signature  *string           `json:"signature,omitempty"`
	Span       *CoreSpan         `json:"span,omitempty"`
	Tags       []string          `json:"tags,omitempty"`
}

// NewObjectFunctionWithSignature instantiates a new ObjectFunctionWithSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectFunctionWithSignature() *ObjectFunctionWithSignature {
	this := ObjectFunctionWithSignature{}
	return &this
}

// NewObjectFunctionWithSignatureWithDefaults instantiates a new ObjectFunctionWithSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectFunctionWithSignatureWithDefaults() *ObjectFunctionWithSignature {
	this := ObjectFunctionWithSignature{}
	return &this
}

// GetBodySpan returns the BodySpan field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetBodySpan() CoreSpan {
	if o == nil || isNil(o.BodySpan) {
		var ret CoreSpan
		return ret
	}
	return *o.BodySpan
}

// GetBodySpanOk returns a tuple with the BodySpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetBodySpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.BodySpan) {
		return nil, false
	}
	return o.BodySpan, true
}

// HasBodySpan returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasBodySpan() bool {
	if o != nil && !isNil(o.BodySpan) {
		return true
	}

	return false
}

// SetBodySpan gets a reference to the given CoreSpan and assigns it to the BodySpan field.
func (o *ObjectFunctionWithSignature) SetBodySpan(v CoreSpan) {
	o.BodySpan = &v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetExtras() map[string]interface{} {
	if o == nil || isNil(o.Extras) {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Extras) {
		return map[string]interface{}{}, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasExtras() bool {
	if o != nil && !isNil(o.Extras) {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *ObjectFunctionWithSignature) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetLang() string {
	if o == nil || isNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetLangOk() (*string, bool) {
	if o == nil || isNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasLang() bool {
	if o != nil && !isNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *ObjectFunctionWithSignature) SetLang(v string) {
	o.Lang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ObjectFunctionWithSignature) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetParameters() []ObjectValueUnit {
	if o == nil || isNil(o.Parameters) {
		var ret []ObjectValueUnit
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetParametersOk() ([]ObjectValueUnit, bool) {
	if o == nil || isNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ObjectValueUnit and assigns it to the Parameters field.
func (o *ObjectFunctionWithSignature) SetParameters(v []ObjectValueUnit) {
	o.Parameters = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ObjectFunctionWithSignature) SetPath(v string) {
	o.Path = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetReceiver() string {
	if o == nil || isNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetReceiverOk() (*string, bool) {
	if o == nil || isNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasReceiver() bool {
	if o != nil && !isNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *ObjectFunctionWithSignature) SetReceiver(v string) {
	o.Receiver = &v
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetReturns() []ObjectValueUnit {
	if o == nil || isNil(o.Returns) {
		var ret []ObjectValueUnit
		return ret
	}
	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetReturnsOk() ([]ObjectValueUnit, bool) {
	if o == nil || isNil(o.Returns) {
		return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasReturns() bool {
	if o != nil && !isNil(o.Returns) {
		return true
	}

	return false
}

// SetReturns gets a reference to the given []ObjectValueUnit and assigns it to the Returns field.
func (o *ObjectFunctionWithSignature) SetReturns(v []ObjectValueUnit) {
	o.Returns = v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetSignature() string {
	if o == nil || isNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetSignatureOk() (*string, bool) {
	if o == nil || isNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasSignature() bool {
	if o != nil && !isNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *ObjectFunctionWithSignature) SetSignature(v string) {
	o.Signature = &v
}

// GetSpan returns the Span field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetSpan() CoreSpan {
	if o == nil || isNil(o.Span) {
		var ret CoreSpan
		return ret
	}
	return *o.Span
}

// GetSpanOk returns a tuple with the Span field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetSpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.Span) {
		return nil, false
	}
	return o.Span, true
}

// HasSpan returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasSpan() bool {
	if o != nil && !isNil(o.Span) {
		return true
	}

	return false
}

// SetSpan gets a reference to the given CoreSpan and assigns it to the Span field.
func (o *ObjectFunctionWithSignature) SetSpan(v CoreSpan) {
	o.Span = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ObjectFunctionWithSignature) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectFunctionWithSignature) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ObjectFunctionWithSignature) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ObjectFunctionWithSignature) SetTags(v []string) {
	o.Tags = v
}

func (o ObjectFunctionWithSignature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BodySpan) {
		toSerialize["bodySpan"] = o.BodySpan
	}
	if !isNil(o.Extras) {
		toSerialize["extras"] = o.Extras
	}
	if !isNil(o.Lang) {
		toSerialize["lang"] = o.Lang
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !isNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Receiver) {
		toSerialize["receiver"] = o.Receiver
	}
	if !isNil(o.Returns) {
		toSerialize["returns"] = o.Returns
	}
	if !isNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !isNil(o.Span) {
		toSerialize["span"] = o.Span
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableObjectFunctionWithSignature struct {
	value *ObjectFunctionWithSignature
	isSet bool
}

func (v NullableObjectFunctionWithSignature) Get() *ObjectFunctionWithSignature {
	return v.value
}

func (v *NullableObjectFunctionWithSignature) Set(val *ObjectFunctionWithSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectFunctionWithSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectFunctionWithSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectFunctionWithSignature(val *ObjectFunctionWithSignature) *NullableObjectFunctionWithSignature {
	return &NullableObjectFunctionWithSignature{value: val, isSet: true}
}

func (v NullableObjectFunctionWithSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectFunctionWithSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
