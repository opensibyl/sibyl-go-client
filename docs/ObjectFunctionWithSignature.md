# ObjectFunctionWithSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodySpan** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]ExtractorValueUnit**](ExtractorValueUnit.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to **string** |  | [optional] 
**Returns** | Pointer to [**[]ExtractorValueUnit**](ExtractorValueUnit.md) |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

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

### GetLanguage

`func (o *ObjectFunctionWithSignature) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ObjectFunctionWithSignature) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ObjectFunctionWithSignature) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ObjectFunctionWithSignature) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

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

`func (o *ObjectFunctionWithSignature) GetParameters() []ExtractorValueUnit`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ObjectFunctionWithSignature) GetParametersOk() (*[]ExtractorValueUnit, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ObjectFunctionWithSignature) SetParameters(v []ExtractorValueUnit)`

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

`func (o *ObjectFunctionWithSignature) GetReturns() []ExtractorValueUnit`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *ObjectFunctionWithSignature) GetReturnsOk() (*[]ExtractorValueUnit, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *ObjectFunctionWithSignature) SetReturns(v []ExtractorValueUnit)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


