# ObjectRevInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int32** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewObjectRevInfo

`func NewObjectRevInfo() *ObjectRevInfo`

NewObjectRevInfo instantiates a new ObjectRevInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRevInfoWithDefaults

`func NewObjectRevInfoWithDefaults() *ObjectRevInfo`

NewObjectRevInfoWithDefaults instantiates a new ObjectRevInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *ObjectRevInfo) GetCreateTime() int32`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ObjectRevInfo) GetCreateTimeOk() (*int32, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ObjectRevInfo) SetCreateTime(v int32)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ObjectRevInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExtras

`func (o *ObjectRevInfo) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ObjectRevInfo) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ObjectRevInfo) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ObjectRevInfo) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetHash

`func (o *ObjectRevInfo) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ObjectRevInfo) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ObjectRevInfo) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ObjectRevInfo) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


