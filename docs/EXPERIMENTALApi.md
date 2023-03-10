# \EXPERIMENTALApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ClazzWithRegexGet**](EXPERIMENTALApi.md#ApiV1ClazzWithRegexGet) | **Get** /api/v1/clazz/with/regex | clazz query
[**ApiV1FuncSignatureGet**](EXPERIMENTALApi.md#ApiV1FuncSignatureGet) | **Get** /api/v1/func/signature | func query
[**ApiV1FuncWithRegexGet**](EXPERIMENTALApi.md#ApiV1FuncWithRegexGet) | **Get** /api/v1/func/with/regex | func query
[**ApiV1FuncWithSignatureGet**](EXPERIMENTALApi.md#ApiV1FuncWithSignatureGet) | **Get** /api/v1/func/with/signature | func query
[**ApiV1FuncctxWithReferenceCountGet**](EXPERIMENTALApi.md#ApiV1FuncctxWithReferenceCountGet) | **Get** /api/v1/funcctx/with/reference/count | funcctx query by ref
[**ApiV1FuncctxWithReferencedCountGet**](EXPERIMENTALApi.md#ApiV1FuncctxWithReferencedCountGet) | **Get** /api/v1/funcctx/with/referenced/count | funcctx query by referenced
[**ApiV1FuncctxWithRegexGet**](EXPERIMENTALApi.md#ApiV1FuncctxWithRegexGet) | **Get** /api/v1/funcctx/with/regex | func ctx query
[**ApiV1RevStatGet**](EXPERIMENTALApi.md#ApiV1RevStatGet) | **Get** /api/v1/rev/stat | rev stat



## ApiV1ClazzWithRegexGet

> []Sibyl2ClazzWithPath ApiV1ClazzWithRegexGet(ctx).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()

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
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1ClazzWithRegexGet(context.Background()).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1ClazzWithRegexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1ClazzWithRegexGet`: []Sibyl2ClazzWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1ClazzWithRegexGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ClazzWithRegexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **field** | **string** | field | 
 **regex** | **string** | regex | 

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


## ApiV1FuncSignatureGet

> []string ApiV1FuncSignatureGet(ctx).Repo(repo).Rev(rev).Regex(regex).Execute()

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
    regex := "regex_example" // string | regex

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1FuncSignatureGet(context.Background()).Repo(repo).Rev(rev).Regex(regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1FuncSignatureGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncSignatureGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1FuncSignatureGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncSignatureGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **regex** | **string** | regex | 

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


## ApiV1FuncWithRegexGet

> []Sibyl2FunctionWithPath ApiV1FuncWithRegexGet(ctx).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()

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
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1FuncWithRegexGet(context.Background()).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1FuncWithRegexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncWithRegexGet`: []Sibyl2FunctionWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1FuncWithRegexGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncWithRegexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **field** | **string** | field | 
 **regex** | **string** | regex | 

### Return type

[**[]Sibyl2FunctionWithPath**](Sibyl2FunctionWithPath.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncWithSignatureGet

> Sibyl2FunctionWithPath ApiV1FuncWithSignatureGet(ctx).Repo(repo).Rev(rev).Signature(signature).Execute()

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
    signature := "signature_example" // string | signature

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1FuncWithSignatureGet(context.Background()).Repo(repo).Rev(rev).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1FuncWithSignatureGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncWithSignatureGet`: Sibyl2FunctionWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1FuncWithSignatureGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncWithSignatureGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **signature** | **string** | signature | 

### Return type

[**Sibyl2FunctionWithPath**](Sibyl2FunctionWithPath.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncctxWithReferenceCountGet

> []Sibyl2FunctionWithPath ApiV1FuncctxWithReferenceCountGet(ctx).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()

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
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1FuncctxWithReferenceCountGet(context.Background()).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1FuncctxWithReferenceCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxWithReferenceCountGet`: []Sibyl2FunctionWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1FuncctxWithReferenceCountGet`: %v\n", resp)
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

[**[]Sibyl2FunctionWithPath**](Sibyl2FunctionWithPath.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncctxWithReferencedCountGet

> []Sibyl2FunctionWithPath ApiV1FuncctxWithReferencedCountGet(ctx).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()

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
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1FuncctxWithReferencedCountGet(context.Background()).Repo(repo).Rev(rev).MoreThan(moreThan).LessThan(lessThan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1FuncctxWithReferencedCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxWithReferencedCountGet`: []Sibyl2FunctionWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1FuncctxWithReferencedCountGet`: %v\n", resp)
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

[**[]Sibyl2FunctionWithPath**](Sibyl2FunctionWithPath.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FuncctxWithRegexGet

> []Sibyl2FunctionContext ApiV1FuncctxWithRegexGet(ctx).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()

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
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1FuncctxWithRegexGet(context.Background()).Repo(repo).Rev(rev).Field(field).Regex(regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1FuncctxWithRegexGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxWithRegexGet`: []Sibyl2FunctionContext
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1FuncctxWithRegexGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncctxWithRegexGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **field** | **string** | field | 
 **regex** | **string** | regex | 

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


## ApiV1RevStatGet

> ServiceRevStat ApiV1RevStatGet(ctx).Repo(repo).Rev(rev).Execute()

rev stat

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
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1RevStatGet(context.Background()).Repo(repo).Rev(rev).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1RevStatGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RevStatGet`: ServiceRevStat
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1RevStatGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RevStatGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 

### Return type

[**ServiceRevStat**](ServiceRevStat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

