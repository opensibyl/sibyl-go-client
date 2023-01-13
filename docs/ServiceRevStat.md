# ServiceRevStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileCount** | Pointer to **int32** |  | [optional] 
**FunctionCount** | Pointer to **int32** |  | [optional] 
**Info** | Pointer to [**ObjectRevInfo**](ObjectRevInfo.md) |  | [optional] 

## Methods

### NewServiceRevStat

`func NewServiceRevStat() *ServiceRevStat`

NewServiceRevStat instantiates a new ServiceRevStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRevStatWithDefaults

`func NewServiceRevStatWithDefaults() *ServiceRevStat`

NewServiceRevStatWithDefaults instantiates a new ServiceRevStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileCount

`func (o *ServiceRevStat) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *ServiceRevStat) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *ServiceRevStat) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *ServiceRevStat) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetFunctionCount

`func (o *ServiceRevStat) GetFunctionCount() int32`

GetFunctionCount returns the FunctionCount field if non-nil, zero value otherwise.

### GetFunctionCountOk

`func (o *ServiceRevStat) GetFunctionCountOk() (*int32, bool)`

GetFunctionCountOk returns a tuple with the FunctionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCount

`func (o *ServiceRevStat) SetFunctionCount(v int32)`

SetFunctionCount sets FunctionCount field to given value.

### HasFunctionCount

`func (o *ServiceRevStat) HasFunctionCount() bool`

HasFunctionCount returns a boolean if a field has been set.

### GetInfo

`func (o *ServiceRevStat) GetInfo() ObjectRevInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceRevStat) GetInfoOk() (*ObjectRevInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceRevStat) SetInfo(v ObjectRevInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceRevStat) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


