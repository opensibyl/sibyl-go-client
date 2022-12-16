# CoreSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to [**CorePoint**](CorePoint.md) |  | [optional] 
**Start** | Pointer to [**CorePoint**](CorePoint.md) |  | [optional] 

## Methods

### NewCoreSpan

`func NewCoreSpan() *CoreSpan`

NewCoreSpan instantiates a new CoreSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreSpanWithDefaults

`func NewCoreSpanWithDefaults() *CoreSpan`

NewCoreSpanWithDefaults instantiates a new CoreSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *CoreSpan) GetEnd() CorePoint`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CoreSpan) GetEndOk() (*CorePoint, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CoreSpan) SetEnd(v CorePoint)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CoreSpan) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *CoreSpan) GetStart() CorePoint`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CoreSpan) GetStartOk() (*CorePoint, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CoreSpan) SetStart(v CorePoint)`

SetStart sets Start field to given value.

### HasStart

`func (o *CoreSpan) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


