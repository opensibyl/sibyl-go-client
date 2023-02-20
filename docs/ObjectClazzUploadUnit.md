# ObjectClazzUploadUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClazzFileResult** | Pointer to [**ExtractorClazzFileResult**](ExtractorClazzFileResult.md) |  | [optional] 
**Workspace** | Pointer to [**ObjectWorkspaceConfig**](ObjectWorkspaceConfig.md) |  | [optional] 

## Methods

### NewObjectClazzUploadUnit

`func NewObjectClazzUploadUnit() *ObjectClazzUploadUnit`

NewObjectClazzUploadUnit instantiates a new ObjectClazzUploadUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectClazzUploadUnitWithDefaults

`func NewObjectClazzUploadUnitWithDefaults() *ObjectClazzUploadUnit`

NewObjectClazzUploadUnitWithDefaults instantiates a new ObjectClazzUploadUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClazzFileResult

`func (o *ObjectClazzUploadUnit) GetClazzFileResult() ExtractorClazzFileResult`

GetClazzFileResult returns the ClazzFileResult field if non-nil, zero value otherwise.

### GetClazzFileResultOk

`func (o *ObjectClazzUploadUnit) GetClazzFileResultOk() (*ExtractorClazzFileResult, bool)`

GetClazzFileResultOk returns a tuple with the ClazzFileResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClazzFileResult

`func (o *ObjectClazzUploadUnit) SetClazzFileResult(v ExtractorClazzFileResult)`

SetClazzFileResult sets ClazzFileResult field to given value.

### HasClazzFileResult

`func (o *ObjectClazzUploadUnit) HasClazzFileResult() bool`

HasClazzFileResult returns a boolean if a field has been set.

### GetWorkspace

`func (o *ObjectClazzUploadUnit) GetWorkspace() ObjectWorkspaceConfig`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *ObjectClazzUploadUnit) GetWorkspaceOk() (*ObjectWorkspaceConfig, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *ObjectClazzUploadUnit) SetWorkspace(v ObjectWorkspaceConfig)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *ObjectClazzUploadUnit) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


