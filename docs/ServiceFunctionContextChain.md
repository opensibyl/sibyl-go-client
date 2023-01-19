# ServiceFunctionContextChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodySpan** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**CallChains** | Pointer to [**ServiceContextTree**](ServiceContextTree.md) |  | [optional] 
**Calls** | Pointer to **[]string** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Lang** | Pointer to **string** | language | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**Returns** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**ReverseCallChains** | Pointer to [**ServiceContextTree**](ServiceContextTree.md) |  | [optional] 
**ReverseCalls** | Pointer to **[]string** |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

## Methods

### NewServiceFunctionContextChain

`func NewServiceFunctionContextChain() *ServiceFunctionContextChain`

NewServiceFunctionContextChain instantiates a new ServiceFunctionContextChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceFunctionContextChainWithDefaults

`func NewServiceFunctionContextChainWithDefaults() *ServiceFunctionContextChain`

NewServiceFunctionContextChainWithDefaults instantiates a new ServiceFunctionContextChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodySpan

`func (o *ServiceFunctionContextChain) GetBodySpan() CoreSpan`

GetBodySpan returns the BodySpan field if non-nil, zero value otherwise.

### GetBodySpanOk

`func (o *ServiceFunctionContextChain) GetBodySpanOk() (*CoreSpan, bool)`

GetBodySpanOk returns a tuple with the BodySpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySpan

`func (o *ServiceFunctionContextChain) SetBodySpan(v CoreSpan)`

SetBodySpan sets BodySpan field to given value.

### HasBodySpan

`func (o *ServiceFunctionContextChain) HasBodySpan() bool`

HasBodySpan returns a boolean if a field has been set.

### GetCallChains

`func (o *ServiceFunctionContextChain) GetCallChains() ServiceContextTree`

GetCallChains returns the CallChains field if non-nil, zero value otherwise.

### GetCallChainsOk

`func (o *ServiceFunctionContextChain) GetCallChainsOk() (*ServiceContextTree, bool)`

GetCallChainsOk returns a tuple with the CallChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallChains

`func (o *ServiceFunctionContextChain) SetCallChains(v ServiceContextTree)`

SetCallChains sets CallChains field to given value.

### HasCallChains

`func (o *ServiceFunctionContextChain) HasCallChains() bool`

HasCallChains returns a boolean if a field has been set.

### GetCalls

`func (o *ServiceFunctionContextChain) GetCalls() []string`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ServiceFunctionContextChain) GetCallsOk() (*[]string, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ServiceFunctionContextChain) SetCalls(v []string)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *ServiceFunctionContextChain) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetExtras

`func (o *ServiceFunctionContextChain) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ServiceFunctionContextChain) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ServiceFunctionContextChain) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ServiceFunctionContextChain) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *ServiceFunctionContextChain) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ServiceFunctionContextChain) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ServiceFunctionContextChain) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ServiceFunctionContextChain) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetName

`func (o *ServiceFunctionContextChain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceFunctionContextChain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceFunctionContextChain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceFunctionContextChain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ServiceFunctionContextChain) GetParameters() []ObjectValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ServiceFunctionContextChain) GetParametersOk() (*[]ObjectValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ServiceFunctionContextChain) SetParameters(v []ObjectValueUnit)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ServiceFunctionContextChain) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *ServiceFunctionContextChain) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ServiceFunctionContextChain) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ServiceFunctionContextChain) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ServiceFunctionContextChain) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceiver

`func (o *ServiceFunctionContextChain) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *ServiceFunctionContextChain) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *ServiceFunctionContextChain) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *ServiceFunctionContextChain) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReturns

`func (o *ServiceFunctionContextChain) GetReturns() []ObjectValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ServiceFunctionContextChain) GetReturnsOk() (*[]ObjectValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ServiceFunctionContextChain) SetReturns(v []ObjectValueUnit)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *ServiceFunctionContextChain) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetReverseCallChains

`func (o *ServiceFunctionContextChain) GetReverseCallChains() ServiceContextTree`

GetReverseCallChains returns the ReverseCallChains field if non-nil, zero value otherwise.

### GetReverseCallChainsOk

`func (o *ServiceFunctionContextChain) GetReverseCallChainsOk() (*ServiceContextTree, bool)`

GetReverseCallChainsOk returns a tuple with the ReverseCallChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseCallChains

`func (o *ServiceFunctionContextChain) SetReverseCallChains(v ServiceContextTree)`

SetReverseCallChains sets ReverseCallChains field to given value.

### HasReverseCallChains

`func (o *ServiceFunctionContextChain) HasReverseCallChains() bool`

HasReverseCallChains returns a boolean if a field has been set.

### GetReverseCalls

`func (o *ServiceFunctionContextChain) GetReverseCalls() []string`

GetReverseCalls returns the ReverseCalls field if non-nil, zero value otherwise.

### GetReverseCallsOk

`func (o *ServiceFunctionContextChain) GetReverseCallsOk() (*[]string, bool)`

GetReverseCallsOk returns a tuple with the ReverseCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseCalls

`func (o *ServiceFunctionContextChain) SetReverseCalls(v []string)`

SetReverseCalls sets ReverseCalls field to given value.

### HasReverseCalls

`func (o *ServiceFunctionContextChain) HasReverseCalls() bool`

HasReverseCalls returns a boolean if a field has been set.

### GetSpan

`func (o *ServiceFunctionContextChain) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ServiceFunctionContextChain) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ServiceFunctionContextChain) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ServiceFunctionContextChain) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


