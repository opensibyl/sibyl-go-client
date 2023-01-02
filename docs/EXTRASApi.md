# \EXTRASApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1FuncDiffGet**](EXTRASApi.md#ApiV1FuncDiffGet) | **Get** /api/v1/func/diff | func diff query



## ApiV1FuncDiffGet

> map[string][]ObjectFunctionWithSignature ApiV1FuncDiffGet(ctx).Repo(repo).Rev(rev).Diff(diff).Execute()

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
    // response from `ApiV1FuncDiffGet`: map[string][]ObjectFunctionWithSignature
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

[**map[string][]ObjectFunctionWithSignature**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

