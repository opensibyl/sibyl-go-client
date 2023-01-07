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

// Sibyl2FunctionWithPath struct for Sibyl2FunctionWithPath
type Sibyl2FunctionWithPath struct {
	BodySpan *CoreSpan `json:"bodySpan,omitempty"`
	// which contains language-specific contents
	Extras map[string]interface{} `json:"extras,omitempty"`
	// language
	Lang *string `json:"lang,omitempty"`
	Name *string `json:"name,omitempty"`
	Parameters []ObjectValueUnit `json:"parameters,omitempty"`
	Path *string `json:"path,omitempty"`
	Receiver *string `json:"receiver,omitempty"`
	Returns []ObjectValueUnit `json:"returns,omitempty"`
	Span *CoreSpan `json:"span,omitempty"`
}

// NewSibyl2FunctionWithPath instantiates a new Sibyl2FunctionWithPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSibyl2FunctionWithPath() *Sibyl2FunctionWithPath {
	this := Sibyl2FunctionWithPath{}
	return &this
}

// NewSibyl2FunctionWithPathWithDefaults instantiates a new Sibyl2FunctionWithPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSibyl2FunctionWithPathWithDefaults() *Sibyl2FunctionWithPath {
	this := Sibyl2FunctionWithPath{}
	return &this
}

// GetBodySpan returns the BodySpan field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetBodySpan() CoreSpan {
	if o == nil || isNil(o.BodySpan) {
		var ret CoreSpan
		return ret
	}
	return *o.BodySpan
}

// GetBodySpanOk returns a tuple with the BodySpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetBodySpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.BodySpan) {
    return nil, false
	}
	return o.BodySpan, true
}

// HasBodySpan returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasBodySpan() bool {
	if o != nil && !isNil(o.BodySpan) {
		return true
	}

	return false
}

// SetBodySpan gets a reference to the given CoreSpan and assigns it to the BodySpan field.
func (o *Sibyl2FunctionWithPath) SetBodySpan(v CoreSpan) {
	o.BodySpan = &v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetExtras() map[string]interface{} {
	if o == nil || isNil(o.Extras) {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Extras) {
    return map[string]interface{}{}, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasExtras() bool {
	if o != nil && !isNil(o.Extras) {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Sibyl2FunctionWithPath) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetLang() string {
	if o == nil || isNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetLangOk() (*string, bool) {
	if o == nil || isNil(o.Lang) {
    return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasLang() bool {
	if o != nil && !isNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *Sibyl2FunctionWithPath) SetLang(v string) {
	o.Lang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Sibyl2FunctionWithPath) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetParameters() []ObjectValueUnit {
	if o == nil || isNil(o.Parameters) {
		var ret []ObjectValueUnit
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetParametersOk() ([]ObjectValueUnit, bool) {
	if o == nil || isNil(o.Parameters) {
    return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ObjectValueUnit and assigns it to the Parameters field.
func (o *Sibyl2FunctionWithPath) SetParameters(v []ObjectValueUnit) {
	o.Parameters = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Sibyl2FunctionWithPath) SetPath(v string) {
	o.Path = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetReceiver() string {
	if o == nil || isNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetReceiverOk() (*string, bool) {
	if o == nil || isNil(o.Receiver) {
    return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasReceiver() bool {
	if o != nil && !isNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *Sibyl2FunctionWithPath) SetReceiver(v string) {
	o.Receiver = &v
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetReturns() []ObjectValueUnit {
	if o == nil || isNil(o.Returns) {
		var ret []ObjectValueUnit
		return ret
	}
	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetReturnsOk() ([]ObjectValueUnit, bool) {
	if o == nil || isNil(o.Returns) {
    return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasReturns() bool {
	if o != nil && !isNil(o.Returns) {
		return true
	}

	return false
}

// SetReturns gets a reference to the given []ObjectValueUnit and assigns it to the Returns field.
func (o *Sibyl2FunctionWithPath) SetReturns(v []ObjectValueUnit) {
	o.Returns = v
}

// GetSpan returns the Span field value if set, zero value otherwise.
func (o *Sibyl2FunctionWithPath) GetSpan() CoreSpan {
	if o == nil || isNil(o.Span) {
		var ret CoreSpan
		return ret
	}
	return *o.Span
}

// GetSpanOk returns a tuple with the Span field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionWithPath) GetSpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.Span) {
    return nil, false
	}
	return o.Span, true
}

// HasSpan returns a boolean if a field has been set.
func (o *Sibyl2FunctionWithPath) HasSpan() bool {
	if o != nil && !isNil(o.Span) {
		return true
	}

	return false
}

// SetSpan gets a reference to the given CoreSpan and assigns it to the Span field.
func (o *Sibyl2FunctionWithPath) SetSpan(v CoreSpan) {
	o.Span = &v
}

func (o Sibyl2FunctionWithPath) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Span) {
		toSerialize["span"] = o.Span
	}
	return json.Marshal(toSerialize)
}

type NullableSibyl2FunctionWithPath struct {
	value *Sibyl2FunctionWithPath
	isSet bool
}

func (v NullableSibyl2FunctionWithPath) Get() *Sibyl2FunctionWithPath {
	return v.value
}

func (v *NullableSibyl2FunctionWithPath) Set(val *Sibyl2FunctionWithPath) {
	v.value = val
	v.isSet = true
}

func (v NullableSibyl2FunctionWithPath) IsSet() bool {
	return v.isSet
}

func (v *NullableSibyl2FunctionWithPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSibyl2FunctionWithPath(val *Sibyl2FunctionWithPath) *NullableSibyl2FunctionWithPath {
	return &NullableSibyl2FunctionWithPath{value: val, isSet: true}
}

func (v NullableSibyl2FunctionWithPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSibyl2FunctionWithPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


