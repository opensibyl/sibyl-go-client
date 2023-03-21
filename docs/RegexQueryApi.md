# \RegexQueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1RegexClazzGet**](RegexQueryApi.md#ApiV1RegexClazzGet) | **Get** /api/v1/regex/clazz | clazz query
[**ApiV1RegexFuncGet**](RegexQueryApi.md#ApiV1RegexFuncGet) | **Get** /api/v1/regex/func | func query
[**ApiV1RegexFuncctxGet**](RegexQueryApi.md#ApiV1RegexFuncctxGet) | **Get** /api/v1/regex/funcctx | func ctx query



## ApiV1RegexClazzGet

> []ObjectClazzServiceDTO ApiV1RegexClazzGet(ctx).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()

clazz query

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
    field := "field_example" // string | field
    regex := "regex_example" // string | regex

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegexQueryApi.ApiV1RegexClazzGet(context.Background()).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegexQueryApi.ApiV1RegexClazzGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RegexClazzGet`: []ObjectClazzServiceDTO
    fmt.Fprintf(os.Stdout, "Response from `RegexQueryApi.ApiV1RegexClazzGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RegexClazzGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **field** | **string** | field | 
 **regex** | **string** | regex | 

### Return type

[**[]ObjectClazzServiceDTO**](ObjectClazzServiceDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RegexFuncGet

> []ObjectFunctionServiceDTO ApiV1RegexFuncGet(ctx).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()

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
    field := "field_example" // string | field
    regex := "regex_example" // string | regex

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegexQueryApi.ApiV1RegexFuncGet(context.Background()).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegexQueryApi.ApiV1RegexFuncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RegexFuncGet`: []ObjectFunctionServiceDTO
    fmt.Fprintf(os.Stdout, "Response from `RegexQueryApi.ApiV1RegexFuncGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RegexFuncGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **field** | **string** | field | 
 **regex** | **string** | regex | 

### Return type

[**[]ObjectFunctionServiceDTO**](ObjectFunctionServiceDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RegexFuncctxGet

> []ObjectFuncCtxServiceDTO ApiV1RegexFuncctxGet(ctx).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()

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
    field := "field_example" // string | field
    regex := "regex_example" // string | regex

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegexQueryApi.ApiV1RegexFuncctxGet(context.Background()).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegexQueryApi.ApiV1RegexFuncctxGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RegexFuncctxGet`: []ObjectFuncCtxServiceDTO
    fmt.Fprintf(os.Stdout, "Response from `RegexQueryApi.ApiV1RegexFuncctxGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RegexFuncctxGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **field** | **string** | field | 
 **regex** | **string** | regex | 

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

