# \SCOPEApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1FileGet**](SCOPEApi.md#ApiV1FileGet) | **Get** /api/v1/file | file query
[**ApiV1RepoDelete**](SCOPEApi.md#ApiV1RepoDelete) | **Delete** /api/v1/repo | repo delete
[**ApiV1RepoGet**](SCOPEApi.md#ApiV1RepoGet) | **Get** /api/v1/repo | repo query
[**ApiV1RevDelete**](SCOPEApi.md#ApiV1RevDelete) | **Delete** /api/v1/rev | rev delte
[**ApiV1RevGet**](SCOPEApi.md#ApiV1RevGet) | **Get** /api/v1/rev | rev query



## ApiV1FileGet

> []string ApiV1FileGet(ctx).Repo(repo).Rev(rev).IncludeRegex(includeRegex).Execute()

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
    includeRegex := "includeRegex_example" // string | includeRegex (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCOPEApi.ApiV1FileGet(context.Background()).Repo(repo).Rev(rev).IncludeRegex(includeRegex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCOPEApi.ApiV1FileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FileGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `SCOPEApi.ApiV1FileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **includeRegex** | **string** | includeRegex | 

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


## ApiV1RepoDelete

> ApiV1RepoDelete(ctx).Repo(repo).Execute()

repo delete

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
    repo := "repo_example" // string | rev delete by repo

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCOPEApi.ApiV1RepoDelete(context.Background()).Repo(repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCOPEApi.ApiV1RepoDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RepoDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | rev delete by repo | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
    resp, r, err := apiClient.SCOPEApi.ApiV1RepoGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCOPEApi.ApiV1RepoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RepoGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `SCOPEApi.ApiV1RepoGet`: %v\n", resp)
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


## ApiV1RevDelete

> ApiV1RevDelete(ctx).Repo(repo).Rev(rev).Execute()

rev delte

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
    resp, r, err := apiClient.SCOPEApi.ApiV1RevDelete(context.Background()).Repo(repo).Rev(rev).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCOPEApi.ApiV1RevDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RevDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
    resp, r, err := apiClient.SCOPEApi.ApiV1RevGet(context.Background()).Repo(repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCOPEApi.ApiV1RevGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RevGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `SCOPEApi.ApiV1RevGet`: %v\n", resp)
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

