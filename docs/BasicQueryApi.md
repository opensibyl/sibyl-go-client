# \BasicQueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ClazzGet**](BasicQueryApi.md#ApiV1ClazzGet) | **Get** /api/v1/clazz | class query
[**ApiV1FuncGet**](BasicQueryApi.md#ApiV1FuncGet) | **Get** /api/v1/func | func query
[**ApiV1FuncctxGet**](BasicQueryApi.md#ApiV1FuncctxGet) | **Get** /api/v1/funcctx | func ctx query



## ApiV1ClazzGet

> []Sibyl2ClazzWithPath ApiV1ClazzGet(ctx).Repo(repo).Rev(rev).File(file).Execute()

class query

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
    file := "file_example" // string | file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicQueryApi.ApiV1ClazzGet(context.Background()).Repo(repo).Rev(rev).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicQueryApi.ApiV1ClazzGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1ClazzGet`: []Sibyl2ClazzWithPath
    fmt.Fprintf(os.Stdout, "Response from `BasicQueryApi.ApiV1ClazzGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ClazzGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **file** | **string** | file | 

### Return type

[**[]Sibyl2ClazzWithPath**](Sibyl2ClazzWithPath.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncGet

> []ObjectFunctionWithSignature ApiV1FuncGet(ctx).Repo(repo).Rev(rev).File(file).Lines(lines).Execute()

func query

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
    file := "file_example" // string | file
    lines := "lines_example" // string | specific lines (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicQueryApi.ApiV1FuncGet(context.Background()).Repo(repo).Rev(rev).File(file).Lines(lines).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicQueryApi.ApiV1FuncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncGet`: []ObjectFunctionWithSignature
    fmt.Fprintf(os.Stdout, "Response from `BasicQueryApi.ApiV1FuncGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **file** | **string** | file | 
 **lines** | **string** | specific lines | 

### Return type

[**[]ObjectFunctionWithSignature**](ObjectFunctionWithSignature.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncctxGet

> []Sibyl2FunctionContextSlim ApiV1FuncctxGet(ctx).Repo(repo).Rev(rev).File(file).Lines(lines).Execute()

func ctx query

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
    file := "file_example" // string | file
    lines := "lines_example" // string | specific lines (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BasicQueryApi.ApiV1FuncctxGet(context.Background()).Repo(repo).Rev(rev).File(file).Lines(lines).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BasicQueryApi.ApiV1FuncctxGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxGet`: []Sibyl2FunctionContextSlim
    fmt.Fprintf(os.Stdout, "Response from `BasicQueryApi.ApiV1FuncctxGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncctxGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **file** | **string** | file | 
 **lines** | **string** | specific lines | 

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

