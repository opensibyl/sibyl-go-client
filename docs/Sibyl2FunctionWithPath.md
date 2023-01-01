# Sibyl2FunctionWithPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodySpan** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**Returns** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

## Methods

### NewSibyl2FunctionWithPath

`func NewSibyl2FunctionWithPath() *Sibyl2FunctionWithPath`

NewSibyl2FunctionWithPath instantiates a new Sibyl2FunctionWithPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSibyl2FunctionWithPathWithDefaults

`func NewSibyl2FunctionWithPathWithDefaults() *Sibyl2FunctionWithPath`

NewSibyl2FunctionWithPathWithDefaults instantiates a new Sibyl2FunctionWithPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodySpan

`func (o *Sibyl2FunctionWithPath) GetBodySpan() CoreSpan`

GetBodySpan returns the BodySpan field if non-nil, zero value otherwise.

### GetBodySpanOk

`func (o *Sibyl2FunctionWithPath) GetBodySpanOk() (*CoreSpan, bool)`

GetBodySpanOk returns a tuple with the BodySpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySpan

`func (o *Sibyl2FunctionWithPath) SetBodySpan(v CoreSpan)`

SetBodySpan sets BodySpan field to given value.

### HasBodySpan

`func (o *Sibyl2FunctionWithPath) HasBodySpan() bool`

HasBodySpan returns a boolean if a field has been set.

### GetExtras

`func (o *Sibyl2FunctionWithPath) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Sibyl2FunctionWithPath) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Sibyl2FunctionWithPath) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Sibyl2FunctionWithPath) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLanguage

`func (o *Sibyl2FunctionWithPath) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Sibyl2FunctionWithPath) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Sibyl2FunctionWithPath) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Sibyl2FunctionWithPath) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetName

`func (o *Sibyl2FunctionWithPath) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sibyl2FunctionWithPath) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sibyl2FunctionWithPath) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sibyl2FunctionWithPath) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *Sibyl2FunctionWithPath) GetParameters() []ObjectValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Sibyl2FunctionWithPath) GetParametersOk() (*[]ObjectValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Sibyl2FunctionWithPath) SetParameters(v []ObjectValueUnit)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Sibyl2FunctionWithPath) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *Sibyl2FunctionWithPath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Sibyl2FunctionWithPath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Sibyl2FunctionWithPath) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Sibyl2FunctionWithPath) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceiver

`func (o *Sibyl2FunctionWithPath) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *Sibyl2FunctionWithPath) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *Sibyl2FunctionWithPath) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *Sibyl2FunctionWithPath) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReturns

`func (o *Sibyl2FunctionWithPath) GetReturns() []ObjectValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *Sibyl2FunctionWithPath) GetReturnsOk() (*[]ObjectValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *Sibyl2FunctionWithPath) SetReturns(v []ObjectValueUnit)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *Sibyl2FunctionWithPath) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetSpan

`func (o *Sibyl2FunctionWithPath) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *Sibyl2FunctionWithPath) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *Sibyl2FunctionWithPath) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *Sibyl2FunctionWithPath) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


