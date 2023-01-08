# \EXPERIMENTALApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ClazzWithRuleGet**](EXPERIMENTALApi.md#ApiV1ClazzWithRuleGet) | **Get** /api/v1/clazz/with/rule | clazz query
[**ApiV1FuncSignatureGet**](EXPERIMENTALApi.md#ApiV1FuncSignatureGet) | **Get** /api/v1/func/signature | func query
[**ApiV1FuncWithRuleGet**](EXPERIMENTALApi.md#ApiV1FuncWithRuleGet) | **Get** /api/v1/func/with/rule | func query
[**ApiV1FuncWithSignatureGet**](EXPERIMENTALApi.md#ApiV1FuncWithSignatureGet) | **Get** /api/v1/func/with/signature | func query
[**ApiV1FuncctxWithRuleGet**](EXPERIMENTALApi.md#ApiV1FuncctxWithRuleGet) | **Get** /api/v1/funcctx/with/rule | func ctx query



## ApiV1ClazzWithRuleGet

> []Sibyl2ClazzWithPath ApiV1ClazzWithRuleGet(ctx).Repo(repo).Rev(rev).Rule(rule).Execute()

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
    rule := "rule_example" // string | rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1ClazzWithRuleGet(context.Background()).Repo(repo).Rev(rev).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1ClazzWithRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1ClazzWithRuleGet`: []Sibyl2ClazzWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1ClazzWithRuleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ClazzWithRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **rule** | **string** | rule | 

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


## ApiV1FuncWithRuleGet

> []Sibyl2FunctionWithPath ApiV1FuncWithRuleGet(ctx).Repo(repo).Rev(rev).Rule(rule).Execute()

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
    rule := "rule_example" // string | rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1FuncWithRuleGet(context.Background()).Repo(repo).Rev(rev).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1FuncWithRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncWithRuleGet`: []Sibyl2FunctionWithPath
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1FuncWithRuleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncWithRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **rule** | **string** | rule | 

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


## ApiV1FuncctxWithRuleGet

> []Sibyl2FunctionContext ApiV1FuncctxWithRuleGet(ctx).Repo(repo).Rev(rev).Rule(rule).Execute()

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
    rule := "rule_example" // string | rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EXPERIMENTALApi.ApiV1FuncctxWithRuleGet(context.Background()).Repo(repo).Rev(rev).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EXPERIMENTALApi.ApiV1FuncctxWithRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1FuncctxWithRuleGet`: []Sibyl2FunctionContext
    fmt.Fprintf(os.Stdout, "Response from `EXPERIMENTALApi.ApiV1FuncctxWithRuleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FuncctxWithRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **rule** | **string** | rule | 

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

