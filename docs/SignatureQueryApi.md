# \SignatureQueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1SignatureFuncGet**](SignatureQueryApi.md#ApiV1SignatureFuncGet) | **Get** /api/v1/signature/func | func query
[**ApiV1SignatureFuncctxChainGet**](SignatureQueryApi.md#ApiV1SignatureFuncctxChainGet) | **Get** /api/v1/signature/funcctx/chain | funcctx chain query
[**ApiV1SignatureFuncctxGet**](SignatureQueryApi.md#ApiV1SignatureFuncctxGet) | **Get** /api/v1/signature/funcctx | funcctx query
[**ApiV1SignatureFuncctxRchainGet**](SignatureQueryApi.md#ApiV1SignatureFuncctxRchainGet) | **Get** /api/v1/signature/funcctx/rchain | funcctx reverse chain query
[**ApiV1SignatureRegexFuncGet**](SignatureQueryApi.md#ApiV1SignatureRegexFuncGet) | **Get** /api/v1/signature/regex/func | func query



## ApiV1SignatureFuncGet

> ObjectFunctionServiceDTO ApiV1SignatureFuncGet(ctx).Repo(repo).Rev(rev).Signature(signature).Execute()

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
    resp, r, err := apiClient.SignatureQueryApi.ApiV1SignatureFuncGet(context.Background()).Repo(repo).Rev(rev).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignatureQueryApi.ApiV1SignatureFuncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1SignatureFuncGet`: ObjectFunctionServiceDTO
    fmt.Fprintf(os.Stdout, "Response from `SignatureQueryApi.ApiV1SignatureFuncGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SignatureFuncGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **signature** | **string** | signature | 

### Return type

[**ObjectFunctionServiceDTO**](ObjectFunctionServiceDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1SignatureFuncctxChainGet

> ServiceFunctionContextChain ApiV1SignatureFuncctxChainGet(ctx).Repo(repo).Rev(rev).Signature(signature).Depth(depth).Execute()

funcctx chain query

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
    depth := int32(56) // int32 | depth

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignatureQueryApi.ApiV1SignatureFuncctxChainGet(context.Background()).Repo(repo).Rev(rev).Signature(signature).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignatureQueryApi.ApiV1SignatureFuncctxChainGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1SignatureFuncctxChainGet`: ServiceFunctionContextChain
    fmt.Fprintf(os.Stdout, "Response from `SignatureQueryApi.ApiV1SignatureFuncctxChainGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SignatureFuncctxChainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **signature** | **string** | signature | 
 **depth** | **int32** | depth | 

### Return type

[**ServiceFunctionContextChain**](ServiceFunctionContextChain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1SignatureFuncctxGet

> ObjectFuncCtxServiceDTO ApiV1SignatureFuncctxGet(ctx).Repo(repo).Rev(rev).Signature(signature).Execute()

funcctx query

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
    resp, r, err := apiClient.SignatureQueryApi.ApiV1SignatureFuncctxGet(context.Background()).Repo(repo).Rev(rev).Signature(signature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignatureQueryApi.ApiV1SignatureFuncctxGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1SignatureFuncctxGet`: ObjectFuncCtxServiceDTO
    fmt.Fprintf(os.Stdout, "Response from `SignatureQueryApi.ApiV1SignatureFuncctxGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SignatureFuncctxGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **signature** | **string** | signature | 

### Return type

[**ObjectFuncCtxServiceDTO**](ObjectFuncCtxServiceDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1SignatureFuncctxRchainGet

> ServiceFunctionContextChain ApiV1SignatureFuncctxRchainGet(ctx).Repo(repo).Rev(rev).Signature(signature).Depth(depth).Execute()

funcctx reverse chain query

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
    depth := int32(56) // int32 | depth

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignatureQueryApi.ApiV1SignatureFuncctxRchainGet(context.Background()).Repo(repo).Rev(rev).Signature(signature).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignatureQueryApi.ApiV1SignatureFuncctxRchainGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1SignatureFuncctxRchainGet`: ServiceFunctionContextChain
    fmt.Fprintf(os.Stdout, "Response from `SignatureQueryApi.ApiV1SignatureFuncctxRchainGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SignatureFuncctxRchainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | repo | 
 **rev** | **string** | rev | 
 **signature** | **string** | signature | 
 **depth** | **int32** | depth | 

### Return type

[**ServiceFunctionContextChain**](ServiceFunctionContextChain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1SignatureRegexFuncGet

> []string ApiV1SignatureRegexFuncGet(ctx).Repo(repo).Rev(rev).Regex(regex).Execute()

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
    resp, r, err := apiClient.SignatureQueryApi.ApiV1SignatureRegexFuncGet(context.Background()).Repo(repo).Rev(rev).Regex(regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignatureQueryApi.ApiV1SignatureRegexFuncGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1SignatureRegexFuncGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `SignatureQueryApi.ApiV1SignatureRegexFuncGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1SignatureRegexFuncGetRequest struct via the builder pattern


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

