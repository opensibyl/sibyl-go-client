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

// Sibyl2FunctionContext struct for Sibyl2FunctionContext
type Sibyl2FunctionContext struct {
	BodySpan *CoreSpan `json:"bodySpan,omitempty"`
	Calls []Sibyl2FunctionWithPath `json:"calls,omitempty"`
	// which contains language-specific contents
	Extras map[string]interface{} `json:"extras,omitempty"`
	Language *string `json:"language,omitempty"`
	Name *string `json:"name,omitempty"`
	Parameters []ExtractorValueUnit `json:"parameters,omitempty"`
	Path *string `json:"path,omitempty"`
	Receiver *string `json:"receiver,omitempty"`
	Returns []ExtractorValueUnit `json:"returns,omitempty"`
	ReverseCalls []Sibyl2FunctionWithPath `json:"reverseCalls,omitempty"`
	Span *CoreSpan `json:"span,omitempty"`
}

// NewSibyl2FunctionContext instantiates a new Sibyl2FunctionContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSibyl2FunctionContext() *Sibyl2FunctionContext {
	this := Sibyl2FunctionContext{}
	return &this
}

// NewSibyl2FunctionContextWithDefaults instantiates a new Sibyl2FunctionContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSibyl2FunctionContextWithDefaults() *Sibyl2FunctionContext {
	this := Sibyl2FunctionContext{}
	return &this
}

// GetBodySpan returns the BodySpan field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetBodySpan() CoreSpan {
	if o == nil || isNil(o.BodySpan) {
		var ret CoreSpan
		return ret
	}
	return *o.BodySpan
}

// GetBodySpanOk returns a tuple with the BodySpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetBodySpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.BodySpan) {
    return nil, false
	}
	return o.BodySpan, true
}

// HasBodySpan returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasBodySpan() bool {
	if o != nil && !isNil(o.BodySpan) {
		return true
	}

	return false
}

// SetBodySpan gets a reference to the given CoreSpan and assigns it to the BodySpan field.
func (o *Sibyl2FunctionContext) SetBodySpan(v CoreSpan) {
	o.BodySpan = &v
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetCalls() []Sibyl2FunctionWithPath {
	if o == nil || isNil(o.Calls) {
		var ret []Sibyl2FunctionWithPath
		return ret
	}
	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetCallsOk() ([]Sibyl2FunctionWithPath, bool) {
	if o == nil || isNil(o.Calls) {
    return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasCalls() bool {
	if o != nil && !isNil(o.Calls) {
		return true
	}

	return false
}

// SetCalls gets a reference to the given []Sibyl2FunctionWithPath and assigns it to the Calls field.
func (o *Sibyl2FunctionContext) SetCalls(v []Sibyl2FunctionWithPath) {
	o.Calls = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetExtras() map[string]interface{} {
	if o == nil || isNil(o.Extras) {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Extras) {
    return map[string]interface{}{}, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasExtras() bool {
	if o != nil && !isNil(o.Extras) {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Sibyl2FunctionContext) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetLanguage() string {
	if o == nil || isNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetLanguageOk() (*string, bool) {
	if o == nil || isNil(o.Language) {
    return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasLanguage() bool {
	if o != nil && !isNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Sibyl2FunctionContext) SetLanguage(v string) {
	o.Language = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Sibyl2FunctionContext) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetParameters() []ExtractorValueUnit {
	if o == nil || isNil(o.Parameters) {
		var ret []ExtractorValueUnit
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetParametersOk() ([]ExtractorValueUnit, bool) {
	if o == nil || isNil(o.Parameters) {
    return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ExtractorValueUnit and assigns it to the Parameters field.
func (o *Sibyl2FunctionContext) SetParameters(v []ExtractorValueUnit) {
	o.Parameters = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
    return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Sibyl2FunctionContext) SetPath(v string) {
	o.Path = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetReceiver() string {
	if o == nil || isNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetReceiverOk() (*string, bool) {
	if o == nil || isNil(o.Receiver) {
    return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasReceiver() bool {
	if o != nil && !isNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *Sibyl2FunctionContext) SetReceiver(v string) {
	o.Receiver = &v
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetReturns() []ExtractorValueUnit {
	if o == nil || isNil(o.Returns) {
		var ret []ExtractorValueUnit
		return ret
	}
	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetReturnsOk() ([]ExtractorValueUnit, bool) {
	if o == nil || isNil(o.Returns) {
    return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasReturns() bool {
	if o != nil && !isNil(o.Returns) {
		return true
	}

	return false
}

// SetReturns gets a reference to the given []ExtractorValueUnit and assigns it to the Returns field.
func (o *Sibyl2FunctionContext) SetReturns(v []ExtractorValueUnit) {
	o.Returns = v
}

// GetReverseCalls returns the ReverseCalls field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetReverseCalls() []Sibyl2FunctionWithPath {
	if o == nil || isNil(o.ReverseCalls) {
		var ret []Sibyl2FunctionWithPath
		return ret
	}
	return o.ReverseCalls
}

// GetReverseCallsOk returns a tuple with the ReverseCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetReverseCallsOk() ([]Sibyl2FunctionWithPath, bool) {
	if o == nil || isNil(o.ReverseCalls) {
    return nil, false
	}
	return o.ReverseCalls, true
}

// HasReverseCalls returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasReverseCalls() bool {
	if o != nil && !isNil(o.ReverseCalls) {
		return true
	}

	return false
}

// SetReverseCalls gets a reference to the given []Sibyl2FunctionWithPath and assigns it to the ReverseCalls field.
func (o *Sibyl2FunctionContext) SetReverseCalls(v []Sibyl2FunctionWithPath) {
	o.ReverseCalls = v
}

// GetSpan returns the Span field value if set, zero value otherwise.
func (o *Sibyl2FunctionContext) GetSpan() CoreSpan {
	if o == nil || isNil(o.Span) {
		var ret CoreSpan
		return ret
	}
	return *o.Span
}

// GetSpanOk returns a tuple with the Span field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sibyl2FunctionContext) GetSpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.Span) {
    return nil, false
	}
	return o.Span, true
}

// HasSpan returns a boolean if a field has been set.
func (o *Sibyl2FunctionContext) HasSpan() bool {
	if o != nil && !isNil(o.Span) {
		return true
	}

	return false
}

// SetSpan gets a reference to the given CoreSpan and assigns it to the Span field.
func (o *Sibyl2FunctionContext) SetSpan(v CoreSpan) {
	o.Span = &v
}

func (o Sibyl2FunctionContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BodySpan) {
		toSerialize["bodySpan"] = o.BodySpan
	}
	if !isNil(o.Calls) {
		toSerialize["calls"] = o.Calls
	}
	if !isNil(o.Extras) {
		toSerialize["extras"] = o.Extras
	}
	if !isNil(o.Language) {
		toSerialize["language"] = o.Language
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
	if !isNil(o.ReverseCalls) {
		toSerialize["reverseCalls"] = o.ReverseCalls
	}
	if !isNil(o.Span) {
		toSerialize["span"] = o.Span
	}
	return json.Marshal(toSerialize)
}

type NullableSibyl2FunctionContext struct {
	value *Sibyl2FunctionContext
	isSet bool
}

func (v NullableSibyl2FunctionContext) Get() *Sibyl2FunctionContext {
	return v.value
}

func (v *NullableSibyl2FunctionContext) Set(val *Sibyl2FunctionContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSibyl2FunctionContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSibyl2FunctionContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSibyl2FunctionContext(val *Sibyl2FunctionContext) *NullableSibyl2FunctionContext {
	return &NullableSibyl2FunctionContext{value: val, isSet: true}
}

func (v NullableSibyl2FunctionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSibyl2FunctionContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


