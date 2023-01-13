# \StatQueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1RevStatGet**](StatQueryApi.md#ApiV1RevStatGet) | **Get** /api/v1/rev/stat | rev stat



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
    resp, r, err := apiClient.StatQueryApi.ApiV1RevStatGet(context.Background()).Repo(repo).Rev(rev).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatQueryApi.ApiV1RevStatGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1RevStatGet`: ServiceRevStat
    fmt.Fprintf(os.Stdout, "Response from `StatQueryApi.ApiV1RevStatGet`: %v\n", resp)
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

