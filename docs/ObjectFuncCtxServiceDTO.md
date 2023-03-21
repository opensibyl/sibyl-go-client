# ObjectFuncCtxServiceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodySpan** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**Calls** | Pointer to **[]string** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Lang** | Pointer to **string** | language | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**Returns** | Pointer to [**[]ObjectValueUnit**](ObjectValueUnit.md) |  | [optional] 
**ReverseCalls** | Pointer to **[]string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

## Methods

### NewObjectFuncCtxServiceDTO

`func NewObjectFuncCtxServiceDTO() *ObjectFuncCtxServiceDTO`

NewObjectFuncCtxServiceDTO instantiates a new ObjectFuncCtxServiceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectFuncCtxServiceDTOWithDefaults

`func NewObjectFuncCtxServiceDTOWithDefaults() *ObjectFuncCtxServiceDTO`

NewObjectFuncCtxServiceDTOWithDefaults instantiates a new ObjectFuncCtxServiceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodySpan

`func (o *ObjectFuncCtxServiceDTO) GetBodySpan() CoreSpan`

GetBodySpan returns the BodySpan field if non-nil, zero value otherwise.

### GetBodySpanOk

`func (o *ObjectFuncCtxServiceDTO) GetBodySpanOk() (*CoreSpan, bool)`

GetBodySpanOk returns a tuple with the BodySpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySpan

`func (o *ObjectFuncCtxServiceDTO) SetBodySpan(v CoreSpan)`

SetBodySpan sets BodySpan field to given value.

### HasBodySpan

`func (o *ObjectFuncCtxServiceDTO) HasBodySpan() bool`

HasBodySpan returns a boolean if a field has been set.

### GetCalls

`func (o *ObjectFuncCtxServiceDTO) GetCalls() []string`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ObjectFuncCtxServiceDTO) GetCallsOk() (*[]string, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ObjectFuncCtxServiceDTO) SetCalls(v []string)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *ObjectFuncCtxServiceDTO) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetExtras

`func (o *ObjectFuncCtxServiceDTO) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ObjectFuncCtxServiceDTO) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ObjectFuncCtxServiceDTO) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ObjectFuncCtxServiceDTO) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *ObjectFuncCtxServiceDTO) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ObjectFuncCtxServiceDTO) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ObjectFuncCtxServiceDTO) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ObjectFuncCtxServiceDTO) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetName

`func (o *ObjectFuncCtxServiceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectFuncCtxServiceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectFuncCtxServiceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectFuncCtxServiceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ObjectFuncCtxServiceDTO) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ObjectFuncCtxServiceDTO) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ObjectFuncCtxServiceDTO) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ObjectFuncCtxServiceDTO) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetParameters

`func (o *ObjectFuncCtxServiceDTO) GetParameters() []ObjectValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ObjectFuncCtxServiceDTO) GetParametersOk() (*[]ObjectValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ObjectFuncCtxServiceDTO) SetParameters(v []ObjectValueUnit)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ObjectFuncCtxServiceDTO) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPath

`func (o *ObjectFuncCtxServiceDTO) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ObjectFuncCtxServiceDTO) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ObjectFuncCtxServiceDTO) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ObjectFuncCtxServiceDTO) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReceiver

`func (o *ObjectFuncCtxServiceDTO) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *ObjectFuncCtxServiceDTO) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *ObjectFuncCtxServiceDTO) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *ObjectFuncCtxServiceDTO) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReturns

`func (o *ObjectFuncCtxServiceDTO) GetReturns() []ObjectValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ObjectFuncCtxServiceDTO) GetReturnsOk() (*[]ObjectValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ObjectFuncCtxServiceDTO) SetReturns(v []ObjectValueUnit)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *ObjectFuncCtxServiceDTO) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetReverseCalls

`func (o *ObjectFuncCtxServiceDTO) GetReverseCalls() []string`

GetReverseCalls returns the ReverseCalls field if non-nil, zero value otherwise.

### GetReverseCallsOk

`func (o *ObjectFuncCtxServiceDTO) GetReverseCallsOk() (*[]string, bool)`

GetReverseCallsOk returns a tuple with the ReverseCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseCalls

`func (o *ObjectFuncCtxServiceDTO) SetReverseCalls(v []string)`

SetReverseCalls sets ReverseCalls field to given value.

### HasReverseCalls

`func (o *ObjectFuncCtxServiceDTO) HasReverseCalls() bool`

HasReverseCalls returns a boolean if a field has been set.

### GetSignature

`func (o *ObjectFuncCtxServiceDTO) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ObjectFuncCtxServiceDTO) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ObjectFuncCtxServiceDTO) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ObjectFuncCtxServiceDTO) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSpan

`func (o *ObjectFuncCtxServiceDTO) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ObjectFuncCtxServiceDTO) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ObjectFuncCtxServiceDTO) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ObjectFuncCtxServiceDTO) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


