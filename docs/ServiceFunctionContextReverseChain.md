# ServiceFunctionContextReverseChain

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
**ReverseCallChains** | Pointer to [**ServiceContextTree**](ServiceContextTree.md) |  | [optional] 
**ReverseCalls** | Pointer to **[]string** |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

## Methods

### NewServiceFunctionContextReverseChain

`func NewServiceFunctionContextReverseChain() *ServiceFunctionContextReverseChain`

NewServiceFunctionContextReverseChain instantiates a new ServiceFunctionContextReverseChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceFunctionContextReverseChainWithDefaults

`func NewServiceFunctionContextReverseChainWithDefaults() *ServiceFunctionContextReverseChain`

NewServiceFunctionContextReverseChainWithDefaults instantiates a new ServiceFunctionContextReverseChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodySpan

`func (o *ServiceFunctionContextReverseChain) GetBodySpan() CoreSpan`

GetBodySpan returns the BodySpan field if non-nil, zero value otherwise.

### GetBodySpanOk

`func (o *ServiceFunctionContextReverseChain) GetBodySpanOk() (*CoreSpan, bool)`

GetBodySpanOk returns a tuple with the BodySpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySpan

`func (o *ServiceFunctionContextReverseChain) SetBodySpan(v CoreSpan)`

SetBodySpan sets BodySpan field to given value.

### HasBodySpan

`func (o *ServiceFunctionContextReverseChain) HasBodySpan() bool`

HasBodySpan returns a boolean if a field has been set.

### GetCalls

`func (o *ServiceFunctionContextReverseChain) GetCalls() []string`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ServiceFunctionContextReverseChain) GetCallsOk() (*[]string, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ServiceFunctionContextReverseChain) SetCalls(v []string)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *ServiceFunctionContextReverseChain) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetExtras

`func (o *ServiceFunctionContextReverseChain) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ServiceFunctionContextReverseChain) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ServiceFunctionContextReverseChain) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ServiceFunctionContextReverseChain) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *ServiceFunctionContextReverseChain) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ServiceFunctionContextReverseChain) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ServiceFunctionContextReverseChain) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ServiceFunctionContextReverseChain) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetName

`func (o *ServiceFunctionContextReverseChain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceFunctionContextReverseChain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceFunctionContextReverseChain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceFunctionContextReverseChain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ServiceFunctionContextReverseChain) GetParameters() []ObjectValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ServiceFunctionContextReverseChain) GetParametersOk() (*[]ObjectValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ServiceFunctionContextReverseChain) SetParameters(v []ObjectValueUnit)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ServiceFunctionContextReverseChain) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *ServiceFunctionContextReverseChain) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ServiceFunctionContextReverseChain) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ServiceFunctionContextReverseChain) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ServiceFunctionContextReverseChain) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceiver

`func (o *ServiceFunctionContextReverseChain) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *ServiceFunctionContextReverseChain) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *ServiceFunctionContextReverseChain) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *ServiceFunctionContextReverseChain) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReturns

`func (o *ServiceFunctionContextReverseChain) GetReturns() []ObjectValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ServiceFunctionContextReverseChain) GetReturnsOk() (*[]ObjectValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ServiceFunctionContextReverseChain) SetReturns(v []ObjectValueUnit)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *ServiceFunctionContextReverseChain) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetReverseCallChains

`func (o *ServiceFunctionContextReverseChain) GetReverseCallChains() ServiceContextTree`

GetReverseCallChains returns the ReverseCallChains field if non-nil, zero value otherwise.

### GetReverseCallChainsOk

`func (o *ServiceFunctionContextReverseChain) GetReverseCallChainsOk() (*ServiceContextTree, bool)`

GetReverseCallChainsOk returns a tuple with the ReverseCallChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseCallChains

`func (o *ServiceFunctionContextReverseChain) SetReverseCallChains(v ServiceContextTree)`

SetReverseCallChains sets ReverseCallChains field to given value.

### HasReverseCallChains

`func (o *ServiceFunctionContextReverseChain) HasReverseCallChains() bool`

HasReverseCallChains returns a boolean if a field has been set.

### GetReverseCalls

`func (o *ServiceFunctionContextReverseChain) GetReverseCalls() []string`

GetReverseCalls returns the ReverseCalls field if non-nil, zero value otherwise.

### GetReverseCallsOk

`func (o *ServiceFunctionContextReverseChain) GetReverseCallsOk() (*[]string, bool)`

GetReverseCallsOk returns a tuple with the ReverseCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseCalls

`func (o *ServiceFunctionContextReverseChain) SetReverseCalls(v []string)`

SetReverseCalls sets ReverseCalls field to given value.

### HasReverseCalls

`func (o *ServiceFunctionContextReverseChain) HasReverseCalls() bool`

HasReverseCalls returns a boolean if a field has been set.

### GetSpan

`func (o *ServiceFunctionContextReverseChain) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ServiceFunctionContextReverseChain) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ServiceFunctionContextReverseChain) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ServiceFunctionContextReverseChain) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


