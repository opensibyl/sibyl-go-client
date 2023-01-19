# ServiceContextTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]ServiceContextTree**](ServiceContextTree.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceContextTree

`func NewServiceContextTree() *ServiceContextTree`

NewServiceContextTree instantiates a new ServiceContextTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceContextTreeWithDefaults

`func NewServiceContextTreeWithDefaults() *ServiceContextTree`

NewServiceContextTreeWithDefaults instantiates a new ServiceContextTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ServiceContextTree) GetChildren() []ServiceContextTree`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ServiceContextTree) GetChildrenOk() (*[]ServiceContextTree, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ServiceContextTree) SetChildren(v []ServiceContextTree)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ServiceContextTree) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetContent

`func (o *ServiceContextTree) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ServiceContextTree) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ServiceContextTree) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ServiceContextTree) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


