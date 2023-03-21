# ObjectClazzServiceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Lang** | Pointer to **string** | language | [optional] 
**Module** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

## Methods

### NewObjectClazzServiceDTO

`func NewObjectClazzServiceDTO() *ObjectClazzServiceDTO`

NewObjectClazzServiceDTO instantiates a new ObjectClazzServiceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectClazzServiceDTOWithDefaults

`func NewObjectClazzServiceDTOWithDefaults() *ObjectClazzServiceDTO`

NewObjectClazzServiceDTOWithDefaults instantiates a new ObjectClazzServiceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtras

`func (o *ObjectClazzServiceDTO) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ObjectClazzServiceDTO) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ObjectClazzServiceDTO) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ObjectClazzServiceDTO) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *ObjectClazzServiceDTO) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ObjectClazzServiceDTO) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ObjectClazzServiceDTO) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ObjectClazzServiceDTO) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetModule

`func (o *ObjectClazzServiceDTO) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ObjectClazzServiceDTO) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ObjectClazzServiceDTO) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ObjectClazzServiceDTO) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetName

`func (o *ObjectClazzServiceDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectClazzServiceDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectClazzServiceDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectClazzServiceDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ObjectClazzServiceDTO) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ObjectClazzServiceDTO) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ObjectClazzServiceDTO) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ObjectClazzServiceDTO) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSignature

`func (o *ObjectClazzServiceDTO) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ObjectClazzServiceDTO) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ObjectClazzServiceDTO) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ObjectClazzServiceDTO) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSpan

`func (o *ObjectClazzServiceDTO) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ObjectClazzServiceDTO) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ObjectClazzServiceDTO) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ObjectClazzServiceDTO) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


