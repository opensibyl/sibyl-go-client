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

// ServiceFunctionContextChain struct for ServiceFunctionContextChain
type ServiceFunctionContextChain struct {
	BodySpan   *CoreSpan           `json:"bodySpan,omitempty"`
	CallChains *ServiceContextTree `json:"callChains,omitempty"`
	Calls      []string            `json:"calls,omitempty"`
	// which contains language-specific contents
	Extras map[string]interface{} `json:"extras,omitempty"`
	// language
	Lang              *string             `json:"lang,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Parameters        []ObjectValueUnit   `json:"parameters,omitempty"`
	Path              *string             `json:"path,omitempty"`
	Receiver          *string             `json:"receiver,omitempty"`
	Returns           []ObjectValueUnit   `json:"returns,omitempty"`
	ReverseCallChains *ServiceContextTree `json:"reverseCallChains,omitempty"`
	ReverseCalls      []string            `json:"reverseCalls,omitempty"`
	Span              *CoreSpan           `json:"span,omitempty"`
}

// NewServiceFunctionContextChain instantiates a new ServiceFunctionContextChain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceFunctionContextChain() *ServiceFunctionContextChain {
	this := ServiceFunctionContextChain{}
	return &this
}

// NewServiceFunctionContextChainWithDefaults instantiates a new ServiceFunctionContextChain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceFunctionContextChainWithDefaults() *ServiceFunctionContextChain {
	this := ServiceFunctionContextChain{}
	return &this
}

// GetBodySpan returns the BodySpan field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetBodySpan() CoreSpan {
	if o == nil || isNil(o.BodySpan) {
		var ret CoreSpan
		return ret
	}
	return *o.BodySpan
}

// GetBodySpanOk returns a tuple with the BodySpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetBodySpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.BodySpan) {
		return nil, false
	}
	return o.BodySpan, true
}

// HasBodySpan returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasBodySpan() bool {
	if o != nil && !isNil(o.BodySpan) {
		return true
	}

	return false
}

// SetBodySpan gets a reference to the given CoreSpan and assigns it to the BodySpan field.
func (o *ServiceFunctionContextChain) SetBodySpan(v CoreSpan) {
	o.BodySpan = &v
}

// GetCallChains returns the CallChains field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetCallChains() ServiceContextTree {
	if o == nil || isNil(o.CallChains) {
		var ret ServiceContextTree
		return ret
	}
	return *o.CallChains
}

// GetCallChainsOk returns a tuple with the CallChains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetCallChainsOk() (*ServiceContextTree, bool) {
	if o == nil || isNil(o.CallChains) {
		return nil, false
	}
	return o.CallChains, true
}

// HasCallChains returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasCallChains() bool {
	if o != nil && !isNil(o.CallChains) {
		return true
	}

	return false
}

// SetCallChains gets a reference to the given ServiceContextTree and assigns it to the CallChains field.
func (o *ServiceFunctionContextChain) SetCallChains(v ServiceContextTree) {
	o.CallChains = &v
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetCalls() []string {
	if o == nil || isNil(o.Calls) {
		var ret []string
		return ret
	}
	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetCallsOk() ([]string, bool) {
	if o == nil || isNil(o.Calls) {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasCalls() bool {
	if o != nil && !isNil(o.Calls) {
		return true
	}

	return false
}

// SetCalls gets a reference to the given []string and assigns it to the Calls field.
func (o *ServiceFunctionContextChain) SetCalls(v []string) {
	o.Calls = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetExtras() map[string]interface{} {
	if o == nil || isNil(o.Extras) {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Extras) {
		return map[string]interface{}{}, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasExtras() bool {
	if o != nil && !isNil(o.Extras) {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *ServiceFunctionContextChain) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetLang() string {
	if o == nil || isNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetLangOk() (*string, bool) {
	if o == nil || isNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasLang() bool {
	if o != nil && !isNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *ServiceFunctionContextChain) SetLang(v string) {
	o.Lang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceFunctionContextChain) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetParameters() []ObjectValueUnit {
	if o == nil || isNil(o.Parameters) {
		var ret []ObjectValueUnit
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetParametersOk() ([]ObjectValueUnit, bool) {
	if o == nil || isNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ObjectValueUnit and assigns it to the Parameters field.
func (o *ServiceFunctionContextChain) SetParameters(v []ObjectValueUnit) {
	o.Parameters = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetPath() string {
	if o == nil || isNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetPathOk() (*string, bool) {
	if o == nil || isNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasPath() bool {
	if o != nil && !isNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ServiceFunctionContextChain) SetPath(v string) {
	o.Path = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetReceiver() string {
	if o == nil || isNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetReceiverOk() (*string, bool) {
	if o == nil || isNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasReceiver() bool {
	if o != nil && !isNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *ServiceFunctionContextChain) SetReceiver(v string) {
	o.Receiver = &v
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetReturns() []ObjectValueUnit {
	if o == nil || isNil(o.Returns) {
		var ret []ObjectValueUnit
		return ret
	}
	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetReturnsOk() ([]ObjectValueUnit, bool) {
	if o == nil || isNil(o.Returns) {
		return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasReturns() bool {
	if o != nil && !isNil(o.Returns) {
		return true
	}

	return false
}

// SetReturns gets a reference to the given []ObjectValueUnit and assigns it to the Returns field.
func (o *ServiceFunctionContextChain) SetReturns(v []ObjectValueUnit) {
	o.Returns = v
}

// GetReverseCallChains returns the ReverseCallChains field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetReverseCallChains() ServiceContextTree {
	if o == nil || isNil(o.ReverseCallChains) {
		var ret ServiceContextTree
		return ret
	}
	return *o.ReverseCallChains
}

// GetReverseCallChainsOk returns a tuple with the ReverseCallChains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetReverseCallChainsOk() (*ServiceContextTree, bool) {
	if o == nil || isNil(o.ReverseCallChains) {
		return nil, false
	}
	return o.ReverseCallChains, true
}

// HasReverseCallChains returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasReverseCallChains() bool {
	if o != nil && !isNil(o.ReverseCallChains) {
		return true
	}

	return false
}

// SetReverseCallChains gets a reference to the given ServiceContextTree and assigns it to the ReverseCallChains field.
func (o *ServiceFunctionContextChain) SetReverseCallChains(v ServiceContextTree) {
	o.ReverseCallChains = &v
}

// GetReverseCalls returns the ReverseCalls field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetReverseCalls() []string {
	if o == nil || isNil(o.ReverseCalls) {
		var ret []string
		return ret
	}
	return o.ReverseCalls
}

// GetReverseCallsOk returns a tuple with the ReverseCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetReverseCallsOk() ([]string, bool) {
	if o == nil || isNil(o.ReverseCalls) {
		return nil, false
	}
	return o.ReverseCalls, true
}

// HasReverseCalls returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasReverseCalls() bool {
	if o != nil && !isNil(o.ReverseCalls) {
		return true
	}

	return false
}

// SetReverseCalls gets a reference to the given []string and assigns it to the ReverseCalls field.
func (o *ServiceFunctionContextChain) SetReverseCalls(v []string) {
	o.ReverseCalls = v
}

// GetSpan returns the Span field value if set, zero value otherwise.
func (o *ServiceFunctionContextChain) GetSpan() CoreSpan {
	if o == nil || isNil(o.Span) {
		var ret CoreSpan
		return ret
	}
	return *o.Span
}

// GetSpanOk returns a tuple with the Span field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFunctionContextChain) GetSpanOk() (*CoreSpan, bool) {
	if o == nil || isNil(o.Span) {
		return nil, false
	}
	return o.Span, true
}

// HasSpan returns a boolean if a field has been set.
func (o *ServiceFunctionContextChain) HasSpan() bool {
	if o != nil && !isNil(o.Span) {
		return true
	}

	return false
}

// SetSpan gets a reference to the given CoreSpan and assigns it to the Span field.
func (o *ServiceFunctionContextChain) SetSpan(v CoreSpan) {
	o.Span = &v
}

func (o ServiceFunctionContextChain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BodySpan) {
		toSerialize["bodySpan"] = o.BodySpan
	}
	if !isNil(o.CallChains) {
		toSerialize["callChains"] = o.CallChains
	}
	if !isNil(o.Calls) {
		toSerialize["calls"] = o.Calls
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
	if !isNil(o.ReverseCallChains) {
		toSerialize["reverseCallChains"] = o.ReverseCallChains
	}
	if !isNil(o.ReverseCalls) {
		toSerialize["reverseCalls"] = o.ReverseCalls
	}
	if !isNil(o.Span) {
		toSerialize["span"] = o.Span
	}
	return json.Marshal(toSerialize)
}

type NullableServiceFunctionContextChain struct {
	value *ServiceFunctionContextChain
	isSet bool
}

func (v NullableServiceFunctionContextChain) Get() *ServiceFunctionContextChain {
	return v.value
}

func (v *NullableServiceFunctionContextChain) Set(val *ServiceFunctionContextChain) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceFunctionContextChain) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceFunctionContextChain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceFunctionContextChain(val *ServiceFunctionContextChain) *NullableServiceFunctionContextChain {
	return &NullableServiceFunctionContextChain{value: val, isSet: true}
}

func (v NullableServiceFunctionContextChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceFunctionContextChain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
