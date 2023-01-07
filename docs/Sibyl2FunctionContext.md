# Sibyl2FunctionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodySpan** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**Calls** | Pointer to [**[]Sibyl2FunctionWithPath**](Sibyl2FunctionWithPath.md) |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Lang** | Pointer to **string** | language | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**Returns** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**ReverseCalls** | Pointer to [**[]Sibyl2FunctionWithPath**](Sibyl2FunctionWithPath.md) |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

## Methods

### NewSibyl2FunctionContext

`func NewSibyl2FunctionContext() *Sibyl2FunctionContext`

NewSibyl2FunctionContext instantiates a new Sibyl2FunctionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSibyl2FunctionContextWithDefaults

`func NewSibyl2FunctionContextWithDefaults() *Sibyl2FunctionContext`

NewSibyl2FunctionContextWithDefaults instantiates a new Sibyl2FunctionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodySpan

`func (o *Sibyl2FunctionContext) GetBodySpan() CoreSpan`

GetBodySpan returns the BodySpan field if non-nil, zero value otherwise.

### GetBodySpanOk

`func (o *Sibyl2FunctionContext) GetBodySpanOk() (*CoreSpan, bool)`

GetBodySpanOk returns a tuple with the BodySpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySpan

`func (o *Sibyl2FunctionContext) SetBodySpan(v CoreSpan)`

SetBodySpan sets BodySpan field to given value.

### HasBodySpan

`func (o *Sibyl2FunctionContext) HasBodySpan() bool`

HasBodySpan returns a boolean if a field has been set.

### GetCalls

`func (o *Sibyl2FunctionContext) GetCalls() []Sibyl2FunctionWithPath`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *Sibyl2FunctionContext) GetCallsOk() (*[]Sibyl2FunctionWithPath, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *Sibyl2FunctionContext) SetCalls(v []Sibyl2FunctionWithPath)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *Sibyl2FunctionContext) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetExtras

`func (o *Sibyl2FunctionContext) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Sibyl2FunctionContext) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Sibyl2FunctionContext) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Sibyl2FunctionContext) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *Sibyl2FunctionContext) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Sibyl2FunctionContext) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Sibyl2FunctionContext) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Sibyl2FunctionContext) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetName

`func (o *Sibyl2FunctionContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sibyl2FunctionContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sibyl2FunctionContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sibyl2FunctionContext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *Sibyl2FunctionContext) GetParameters() []ObjectValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Sibyl2FunctionContext) GetParametersOk() (*[]ObjectValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Sibyl2FunctionContext) SetParameters(v []ObjectValueUnit)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Sibyl2FunctionContext) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *Sibyl2FunctionContext) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Sibyl2FunctionContext) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Sibyl2FunctionContext) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Sibyl2FunctionContext) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceiver

`func (o *Sibyl2FunctionContext) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *Sibyl2FunctionContext) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *Sibyl2FunctionContext) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *Sibyl2FunctionContext) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReturns

`func (o *Sibyl2FunctionContext) GetReturns() []ObjectValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *Sibyl2FunctionContext) GetReturnsOk() (*[]ObjectValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *Sibyl2FunctionContext) SetReturns(v []ObjectValueUnit)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *Sibyl2FunctionContext) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetReverseCalls

`func (o *Sibyl2FunctionContext) GetReverseCalls() []Sibyl2FunctionWithPath`

GetReverseCalls returns the ReverseCalls field if non-nil, zero value otherwise.

### GetReverseCallsOk

`func (o *Sibyl2FunctionContext) GetReverseCallsOk() (*[]Sibyl2FunctionWithPath, bool)`

GetReverseCallsOk returns a tuple with the ReverseCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseCalls

`func (o *Sibyl2FunctionContext) SetReverseCalls(v []Sibyl2FunctionWithPath)`

SetReverseCalls sets ReverseCalls field to given value.

### HasReverseCalls

`func (o *Sibyl2FunctionContext) HasReverseCalls() bool`

HasReverseCalls returns a boolean if a field has been set.

### GetSpan

`func (o *Sibyl2FunctionContext) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *Sibyl2FunctionContext) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *Sibyl2FunctionContext) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *Sibyl2FunctionContext) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


