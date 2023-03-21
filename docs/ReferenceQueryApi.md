# \ReferenceQueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ReferenceCountFuncctxGet**](ReferenceQueryApi.md#ApiV1ReferenceCountFuncctxGet) | **Get** /api/v1/reference/count/funcctx | funcctx query by ref
[**ApiV1ReferenceCountFuncctxReverseGet**](ReferenceQueryApi.md#ApiV1ReferenceCountFuncctxReverseGet) | **Get** /api/v1/reference/count/funcctx/reverse | funcctx query by referenced



## ApiV1ReferenceCountFuncctxGet

> []ObjectFuncCtxServiceDTO ApiV1ReferenceCountFuncctxGet(ctx).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()

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
    resp, r, err := apiClient.ReferenceQueryApi.ApiV1ReferenceCountFuncctxGet(context.Background()).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceQueryApi.ApiV1ReferenceCountFuncctxGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1ReferenceCountFuncctxGet`: []ObjectFuncCtxServiceDTO
    fmt.Fprintf(os.Stdout, "Response from `ReferenceQueryApi.ApiV1ReferenceCountFuncctxGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ReferenceCountFuncctxGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **moreThan** | **int32** | moreThan | 
 **lessThan** | **int32** | lessThan | 

### Return type

[**[]ObjectFuncCtxServiceDTO**](ObjectFuncCtxServiceDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ReferenceCountFuncctxReverseGet

> []ObjectFuncCtxServiceDTO ApiV1ReferenceCountFuncctxReverseGet(ctx).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()

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
    resp, r, err := apiClient.ReferenceQueryApi.ApiV1ReferenceCountFuncctxReverseGet(context.Background()).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceQueryApi.ApiV1ReferenceCountFuncctxReverseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1ReferenceCountFuncctxReverseGet`: []ObjectFuncCtxServiceDTO
    fmt.Fprintf(os.Stdout, "Response from `ReferenceQueryApi.ApiV1ReferenceCountFuncctxReverseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ReferenceCountFuncctxReverseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **moreThan** | **int32** | moreThan | 
 **lessThan** | **int32** | lessThan | 

### Return type

[**[]ObjectFuncCtxServiceDTO**](ObjectFuncCtxServiceDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

