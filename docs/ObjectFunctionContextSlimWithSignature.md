# ObjectFunctionContextSlimWithSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodySpan** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**Calls** | Pointer to **[]string** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Lang** | Pointer to **string** | language | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**Returns** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**ReverseCalls** | Pointer to **[]string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

## Methods

### NewObjectFunctionContextSlimWithSignature

`func NewObjectFunctionContextSlimWithSignature() *ObjectFunctionContextSlimWithSignature`

NewObjectFunctionContextSlimWithSignature instantiates a new ObjectFunctionContextSlimWithSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectFunctionContextSlimWithSignatureWithDefaults

`func NewObjectFunctionContextSlimWithSignatureWithDefaults() *ObjectFunctionContextSlimWithSignature`

NewObjectFunctionContextSlimWithSignatureWithDefaults instantiates a new ObjectFunctionContextSlimWithSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodySpan

`func (o *ObjectFunctionContextSlimWithSignature) GetBodySpan() CoreSpan`

GetBodySpan returns the BodySpan field if non-nil, zero value otherwise.

### GetBodySpanOk

`func (o *ObjectFunctionContextSlimWithSignature) GetBodySpanOk() (*CoreSpan, bool)`

GetBodySpanOk returns a tuple with the BodySpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySpan

`func (o *ObjectFunctionContextSlimWithSignature) SetBodySpan(v CoreSpan)`

SetBodySpan sets BodySpan field to given value.

### HasBodySpan

`func (o *ObjectFunctionContextSlimWithSignature) HasBodySpan() bool`

HasBodySpan returns a boolean if a field has been set.

### GetCalls

`func (o *ObjectFunctionContextSlimWithSignature) GetCalls() []string`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ObjectFunctionContextSlimWithSignature) GetCallsOk() (*[]string, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ObjectFunctionContextSlimWithSignature) SetCalls(v []string)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *ObjectFunctionContextSlimWithSignature) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetExtras

`func (o *ObjectFunctionContextSlimWithSignature) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ObjectFunctionContextSlimWithSignature) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ObjectFunctionContextSlimWithSignature) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ObjectFunctionContextSlimWithSignature) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *ObjectFunctionContextSlimWithSignature) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ObjectFunctionContextSlimWithSignature) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ObjectFunctionContextSlimWithSignature) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ObjectFunctionContextSlimWithSignature) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetName

`func (o *ObjectFunctionContextSlimWithSignature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectFunctionContextSlimWithSignature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectFunctionContextSlimWithSignature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectFunctionContextSlimWithSignature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ObjectFunctionContextSlimWithSignature) GetParameters() []ObjectValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ObjectFunctionContextSlimWithSignature) GetParametersOk() (*[]ObjectValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ObjectFunctionContextSlimWithSignature) SetParameters(v []ObjectValueUnit)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ObjectFunctionContextSlimWithSignature) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *ObjectFunctionContextSlimWithSignature) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ObjectFunctionContextSlimWithSignature) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ObjectFunctionContextSlimWithSignature) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ObjectFunctionContextSlimWithSignature) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceiver

`func (o *ObjectFunctionContextSlimWithSignature) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *ObjectFunctionContextSlimWithSignature) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *ObjectFunctionContextSlimWithSignature) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *ObjectFunctionContextSlimWithSignature) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReturns

`func (o *ObjectFunctionContextSlimWithSignature) GetReturns() []ObjectValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ObjectFunctionContextSlimWithSignature) GetReturnsOk() (*[]ObjectValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ObjectFunctionContextSlimWithSignature) SetReturns(v []ObjectValueUnit)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *ObjectFunctionContextSlimWithSignature) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetReverseCalls

`func (o *ObjectFunctionContextSlimWithSignature) GetReverseCalls() []string`

GetReverseCalls returns the ReverseCalls field if non-nil, zero value otherwise.

### GetReverseCallsOk

`func (o *ObjectFunctionContextSlimWithSignature) GetReverseCallsOk() (*[]string, bool)`

GetReverseCallsOk returns a tuple with the ReverseCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseCalls

`func (o *ObjectFunctionContextSlimWithSignature) SetReverseCalls(v []string)`

SetReverseCalls sets ReverseCalls field to given value.

### HasReverseCalls

`func (o *ObjectFunctionContextSlimWithSignature) HasReverseCalls() bool`

HasReverseCalls returns a boolean if a field has been set.

### GetSignature

`func (o *ObjectFunctionContextSlimWithSignature) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ObjectFunctionContextSlimWithSignature) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ObjectFunctionContextSlimWithSignature) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ObjectFunctionContextSlimWithSignature) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSpan

`func (o *ObjectFunctionContextSlimWithSignature) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ObjectFunctionContextSlimWithSignature) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ObjectFunctionContextSlimWithSignature) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ObjectFunctionContextSlimWithSignature) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


