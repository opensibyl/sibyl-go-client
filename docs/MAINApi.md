# \MAINApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ClazzGet**](MAINApi.md#ApiV1ClazzGet) | **Get** /api/v1/clazz | class query
[**ApiV1FileGet**](MAINApi.md#ApiV1FileGet) | **Get** /api/v1/file | file query
[**ApiV1FuncGet**](MAINApi.md#ApiV1FuncGet) | **Get** /api/v1/func | func query
[**ApiV1FuncctxGet**](MAINApi.md#ApiV1FuncctxGet) | **Get** /api/v1/funcctx | func ctx query
[**ApiV1RepoGet**](MAINApi.md#ApiV1RepoGet) | **Get** /api/v1/repo | repo query
[**ApiV1RevGet**](MAINApi.md#ApiV1RevGet) | **Get** /api/v1/rev | rev query



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
    resp, r, err := apiClient.MAINApi.ApiV1ClazzGet(context.Background()).Repo(repo).Rev(rev).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MAINApi.ApiV1ClazzGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1ClazzGet`: []Sibyl2ClazzWithPath
    fmt.Fprintf(os.Stdout, "Response from `MAINApi.ApiV1ClazzGet`: %v\n", resp)
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


## ApiV1FileGet

> []string ApiV1FileGet(ctx).Repo(repo).Rev(rev).Execute()

file query

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MAINApi.ApiV1FileGet(context.Background()).Repo(repo).Rev(rev).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MAINApi.ApiV1FileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FileGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `MAINApi.ApiV1FileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 

### Return type

**[]string**

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
    resp, r, err := apiClient.MAINApi.ApiV1FuncGet(context.Background()).Repo(repo).Rev(rev).File(file).Lines(lines).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MAINApi.ApiV1FuncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncGet`: []ObjectFunctionWithSignature
    fmt.Fprintf(os.Stdout, "Response from `MAINApi.ApiV1FuncGet`: %v\n", resp)
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

> []Sibyl2FunctionContext ApiV1FuncctxGet(ctx).Repo(repo).Rev(rev).File(file).Lines(lines).Execute()

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
    resp, r, err := apiClient.MAINApi.ApiV1FuncctxGet(context.Background()).Repo(repo).Rev(rev).File(file).Lines(lines).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MAINApi.ApiV1FuncctxGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxGet`: []Sibyl2FunctionContext
    fmt.Fprintf(os.Stdout, "Response from `MAINApi.ApiV1FuncctxGet`: %v\n", resp)
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

[**[]Sibyl2FunctionContext**](Sibyl2FunctionContext.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RepoGet

> []string ApiV1RepoGet(ctx).Execute()

repo query

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MAINApi.ApiV1RepoGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MAINApi.ApiV1RepoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RepoGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `MAINApi.ApiV1RepoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RepoGetRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RevGet

> []string ApiV1RevGet(ctx).Repo(repo).Execute()

rev query

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
    repo := "repo_example" // string | rev search by repo

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MAINApi.ApiV1RevGet(context.Background()).Repo(repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MAINApi.ApiV1RevGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RevGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `MAINApi.ApiV1RevGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RevGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | rev search by repo | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

