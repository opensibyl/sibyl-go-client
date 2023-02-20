# Sibyl2FunctionWithTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodySpan** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Lang** | Pointer to **string** | language | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**Returns** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSibyl2FunctionWithTag

`func NewSibyl2FunctionWithTag() *Sibyl2FunctionWithTag`

NewSibyl2FunctionWithTag instantiates a new Sibyl2FunctionWithTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSibyl2FunctionWithTagWithDefaults

`func NewSibyl2FunctionWithTagWithDefaults() *Sibyl2FunctionWithTag`

NewSibyl2FunctionWithTagWithDefaults instantiates a new Sibyl2FunctionWithTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodySpan

`func (o *Sibyl2FunctionWithTag) GetBodySpan() CoreSpan`

GetBodySpan returns the BodySpan field if non-nil, zero value otherwise.

### GetBodySpanOk

`func (o *Sibyl2FunctionWithTag) GetBodySpanOk() (*CoreSpan, bool)`

GetBodySpanOk returns a tuple with the BodySpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySpan

`func (o *Sibyl2FunctionWithTag) SetBodySpan(v CoreSpan)`

SetBodySpan sets BodySpan field to given value.

### HasBodySpan

`func (o *Sibyl2FunctionWithTag) HasBodySpan() bool`

HasBodySpan returns a boolean if a field has been set.

### GetExtras

`func (o *Sibyl2FunctionWithTag) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Sibyl2FunctionWithTag) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Sibyl2FunctionWithTag) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Sibyl2FunctionWithTag) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *Sibyl2FunctionWithTag) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Sibyl2FunctionWithTag) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Sibyl2FunctionWithTag) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Sibyl2FunctionWithTag) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetName

`func (o *Sibyl2FunctionWithTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sibyl2FunctionWithTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sibyl2FunctionWithTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sibyl2FunctionWithTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *Sibyl2FunctionWithTag) GetParameters() []ObjectValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Sibyl2FunctionWithTag) GetParametersOk() (*[]ObjectValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Sibyl2FunctionWithTag) SetParameters(v []ObjectValueUnit)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Sibyl2FunctionWithTag) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *Sibyl2FunctionWithTag) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Sibyl2FunctionWithTag) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Sibyl2FunctionWithTag) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Sibyl2FunctionWithTag) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceiver

`func (o *Sibyl2FunctionWithTag) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *Sibyl2FunctionWithTag) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *Sibyl2FunctionWithTag) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *Sibyl2FunctionWithTag) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReturns

`func (o *Sibyl2FunctionWithTag) GetReturns() []ObjectValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *Sibyl2FunctionWithTag) GetReturnsOk() (*[]ObjectValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *Sibyl2FunctionWithTag) SetReturns(v []ObjectValueUnit)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *Sibyl2FunctionWithTag) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetSpan

`func (o *Sibyl2FunctionWithTag) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *Sibyl2FunctionWithTag) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *Sibyl2FunctionWithTag) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *Sibyl2FunctionWithTag) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetTags

`func (o *Sibyl2FunctionWithTag) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Sibyl2FunctionWithTag) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Sibyl2FunctionWithTag) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Sibyl2FunctionWithTag) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


