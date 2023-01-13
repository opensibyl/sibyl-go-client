# \ReferenceQueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1FuncctxWithReferenceCountGet**](ReferenceQueryApi.md#ApiV1FuncctxWithReferenceCountGet) | **Get** /api/v1/funcctx/with/reference/count | funcctx query by ref
[**ApiV1FuncctxWithReferencedCountGet**](ReferenceQueryApi.md#ApiV1FuncctxWithReferencedCountGet) | **Get** /api/v1/funcctx/with/referenced/count | funcctx query by referenced



## ApiV1FuncctxWithReferenceCountGet

> []Sibyl2FunctionContextSlim ApiV1FuncctxWithReferenceCountGet(ctx).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()

funcctx query by ref

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repo := "repo_example" // string | repo
    rev := "rev_example" // string | rev
    moreThan := int32(56) // int32 | moreThan
    lessThan := int32(56) // int32 | lessThan

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReferenceQueryApi.ApiV1FuncctxWithReferenceCountGet(context.Background()).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceQueryApi.ApiV1FuncctxWithReferenceCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxWithReferenceCountGet`: []Sibyl2FunctionContextSlim
    fmt.Fprintf(os.Stdout, "Response from `ReferenceQueryApi.ApiV1FuncctxWithReferenceCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncctxWithReferenceCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **moreThan** | **int32** | moreThan | 
 **lessThan** | **int32** | lessThan | 

### Return type

[**[]Sibyl2FunctionContextSlim**](Sibyl2FunctionContextSlim.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncctxWithReferencedCountGet

> []Sibyl2FunctionContextSlim ApiV1FuncctxWithReferencedCountGet(ctx).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()

funcctx query by referenced

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repo := "repo_example" // string | repo
    rev := "rev_example" // string | rev
    moreThan := int32(56) // int32 | moreThan
    lessThan := int32(56) // int32 | lessThan

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReferenceQueryApi.ApiV1FuncctxWithReferencedCountGet(context.Background()).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceQueryApi.ApiV1FuncctxWithReferencedCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxWithReferencedCountGet`: []Sibyl2FunctionContextSlim
    fmt.Fprintf(os.Stdout, "Response from `ReferenceQueryApi.ApiV1FuncctxWithReferencedCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncctxWithReferencedCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **moreThan** | **int32** | moreThan | 
 **lessThan** | **int32** | lessThan | 

### Return type

[**[]Sibyl2FunctionContextSlim**](Sibyl2FunctionContextSlim.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

