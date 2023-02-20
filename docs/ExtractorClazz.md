# ExtractorClazz

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extras** | Pointer to **map[string]interface{}** | which contains language-specific contents | [optional] 
**Lang** | Pointer to **string** | language | [optional] 
**Module** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Span** | Pointer to [**CoreSpan**](CoreSpan.md) |  | [optional] 

## Methods

### NewExtractorClazz

`func NewExtractorClazz() *ExtractorClazz`

NewExtractorClazz instantiates a new ExtractorClazz object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractorClazzWithDefaults

`func NewExtractorClazzWithDefaults() *ExtractorClazz`

NewExtractorClazzWithDefaults instantiates a new ExtractorClazz object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtras

`func (o *ExtractorClazz) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ExtractorClazz) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ExtractorClazz) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ExtractorClazz) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetLang

`func (o *ExtractorClazz) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ExtractorClazz) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ExtractorClazz) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ExtractorClazz) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetModule

`func (o *ExtractorClazz) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ExtractorClazz) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ExtractorClazz) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ExtractorClazz) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetName

`func (o *ExtractorClazz) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtractorClazz) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtractorClazz) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtractorClazz) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpan

`func (o *ExtractorClazz) GetSpan() CoreSpan`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ExtractorClazz) GetSpanOk() (*CoreSpan, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ExtractorClazz) SetSpan(v CoreSpan)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ExtractorClazz) HasSpan() bool`

HasSpan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


