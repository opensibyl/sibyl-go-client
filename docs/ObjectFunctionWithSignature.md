# ObjectFunctionWithSignature

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
**Signature** | Pointer to **string** |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewObjectFunctionWithSignature

`func NewObjectFunctionWithSignature() *ObjectFunctionWithSignature`

NewObjectFunctionWithSignature instantiates a new ObjectFunctionWithSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectFunctionWithSignatureWithDefaults

`func NewObjectFunctionWithSignatureWithDefaults() *ObjectFunctionWithSignature`

NewObjectFunctionWithSignatureWithDefaults instantiates a new ObjectFunctionWithSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodySpan

`func (o *ObjectFunctionWithSignature) GetBodySpan() CoreSpan`

GetBodySpan returns the BodySpan field if non-nil, zero value otherwise.

### GetBodySpanOk

`func (o *ObjectFunctionWithSignature) GetBodySpanOk() (*CoreSpan, bool)`

GetBodySpanOk returns a tuple with the BodySpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySpan

`func (o *ObjectFunctionWithSignature) SetBodySpan(v CoreSpan)`

SetBodySpan sets BodySpan field to given value.

### HasBodySpan

`func (o *ObjectFunctionWithSignature) HasBodySpan() bool`

HasBodySpan returns a boolean if a field has been set.

### GetExtras

`func (o *ObjectFunctionWithSignature) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ObjectFunctionWithSignature) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ObjectFunctionWithSignature) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ObjectFunctionWithSignature) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *ObjectFunctionWithSignature) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ObjectFunctionWithSignature) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ObjectFunctionWithSignature) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ObjectFunctionWithSignature) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetName

`func (o *ObjectFunctionWithSignature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectFunctionWithSignature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectFunctionWithSignature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectFunctionWithSignature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ObjectFunctionWithSignature) GetParameters() []ObjectValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ObjectFunctionWithSignature) GetParametersOk() (*[]ObjectValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ObjectFunctionWithSignature) SetParameters(v []ObjectValueUnit)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ObjectFunctionWithSignature) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *ObjectFunctionWithSignature) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ObjectFunctionWithSignature) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ObjectFunctionWithSignature) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ObjectFunctionWithSignature) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceiver

`func (o *ObjectFunctionWithSignature) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *ObjectFunctionWithSignature) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *ObjectFunctionWithSignature) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *ObjectFunctionWithSignature) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReturns

`func (o *ObjectFunctionWithSignature) GetReturns() []ObjectValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ObjectFunctionWithSignature) GetReturnsOk() (*[]ObjectValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ObjectFunctionWithSignature) SetReturns(v []ObjectValueUnit)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *ObjectFunctionWithSignature) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetSignature

`func (o *ObjectFunctionWithSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ObjectFunctionWithSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ObjectFunctionWithSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ObjectFunctionWithSignature) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSpan

`func (o *ObjectFunctionWithSignature) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ObjectFunctionWithSignature) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ObjectFunctionWithSignature) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ObjectFunctionWithSignature) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetTags

`func (o *ObjectFunctionWithSignature) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ObjectFunctionWithSignature) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ObjectFunctionWithSignature) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ObjectFunctionWithSignature) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


