# \EXTRASApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ClazzDiffGet**](EXTRASApi.md#ApiV1ClazzDiffGet) | **Get** /api/v1/clazz/diff | clazz diff query
[**ApiV1FuncDiffGet**](EXTRASApi.md#ApiV1FuncDiffGet) | **Get** /api/v1/func/diff | func diff query
[**ApiV1FuncctxDiffGet**](EXTRASApi.md#ApiV1FuncctxDiffGet) | **Get** /api/v1/funcctx/diff | func ctx diff query



## ApiV1ClazzDiffGet

> map[string][]Sibyl2ClazzWithPath ApiV1ClazzDiffGet(ctx).Repo(repo).Rev(rev).Diff(diff).Execute()

clazz diff query

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
    diff := "diff_example" // string | unified diff

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EXTRASApi.ApiV1ClazzDiffGet(context.Background()).Repo(repo).Rev(rev).Diff(diff).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXTRASApi.ApiV1ClazzDiffGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1ClazzDiffGet`: map[string][]Sibyl2ClazzWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXTRASApi.ApiV1ClazzDiffGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ClazzDiffGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **diff** | **string** | unified diff | 

### Return type

[**map[string][]Sibyl2ClazzWithPath**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncDiffGet

> map[string][]Sibyl2FunctionWithPath ApiV1FuncDiffGet(ctx).Repo(repo).Rev(rev).Diff(diff).Execute()

func diff query

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
    diff := "diff_example" // string | unified diff

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EXTRASApi.ApiV1FuncDiffGet(context.Background()).Repo(repo).Rev(rev).Diff(diff).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXTRASApi.ApiV1FuncDiffGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncDiffGet`: map[string][]Sibyl2FunctionWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXTRASApi.ApiV1FuncDiffGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncDiffGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **diff** | **string** | unified diff | 

### Return type

[**map[string][]Sibyl2FunctionWithPath**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncctxDiffGet

> map[string][]Sibyl2FunctionContext ApiV1FuncctxDiffGet(ctx).Repo(repo).Rev(rev).Diff(diff).Execute()

func ctx diff query

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
    diff := "diff_example" // string | unified diff

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EXTRASApi.ApiV1FuncctxDiffGet(context.Background()).Repo(repo).Rev(rev).Diff(diff).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXTRASApi.ApiV1FuncctxDiffGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxDiffGet`: map[string][]Sibyl2FunctionContext
    fmt.Fprintf(os.Stdout, "Response from `EXTRASApi.ApiV1FuncctxDiffGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncctxDiffGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **diff** | **string** | unified diff | 

### Return type

[**map[string][]Sibyl2FunctionContext**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

