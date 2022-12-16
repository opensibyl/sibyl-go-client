# CorePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **int32** |  | [optional] 
**Row** | Pointer to **int32** | NOTICE: INDEX, NOT REAL LINE NUMBER | [optional] 

## Methods

### NewCorePoint

`func NewCorePoint() *CorePoint`

NewCorePoint instantiates a new CorePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorePointWithDefaults

`func NewCorePointWithDefaults() *CorePoint`

NewCorePointWithDefaults instantiates a new CorePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *CorePoint) GetColumn() int32`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *CorePoint) GetColumnOk() (*int32, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *CorePoint) SetColumn(v int32)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *CorePoint) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetRow

`func (o *CorePoint) GetRow() int32`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *CorePoint) GetRowOk() (*int32, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *CorePoint) SetRow(v int32)`

SetRow sets Row field to given value.

### HasRow

`func (o *CorePoint) HasRow() bool`

HasRow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


