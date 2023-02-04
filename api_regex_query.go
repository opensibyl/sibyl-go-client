/*
openapi for sibyl2 server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// RegexQueryApiService RegexQueryApi service
type RegexQueryApiService service

type ApiApiV1RegexClazzGetRequest struct {
	ctx        context.Context
	ApiService *RegexQueryApiService
	repo       *string
	rev        *string
	field      *string
	regex      *string
}

// repo
func (r ApiApiV1RegexClazzGetRequest) Repo(repo string) ApiApiV1RegexClazzGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1RegexClazzGetRequest) Rev(rev string) ApiApiV1RegexClazzGetRequest {
	r.rev = &rev
	return r
}

// field
func (r ApiApiV1RegexClazzGetRequest) Field(field string) ApiApiV1RegexClazzGetRequest {
	r.field = &field
	return r
}

// regex
func (r ApiApiV1RegexClazzGetRequest) Regex(regex string) ApiApiV1RegexClazzGetRequest {
	r.regex = &regex
	return r
}

func (r ApiApiV1RegexClazzGetRequest) Execute() ([]Sibyl2ClazzWithPath, *http.Response, error) {
	return r.ApiService.ApiV1RegexClazzGetExecute(r)
}

/*
ApiV1RegexClazzGet clazz query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1RegexClazzGetRequest
*/
func (a *RegexQueryApiService) ApiV1RegexClazzGet(ctx context.Context) ApiApiV1RegexClazzGetRequest {
	return ApiApiV1RegexClazzGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Sibyl2ClazzWithPath
func (a *RegexQueryApiService) ApiV1RegexClazzGetExecute(r ApiApiV1RegexClazzGetRequest) ([]Sibyl2ClazzWithPath, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Sibyl2ClazzWithPath
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegexQueryApiService.ApiV1RegexClazzGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/regex/clazz"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.field == nil {
		return localVarReturnValue, nil, reportError("field is required and must be specified")
	}
	if r.regex == nil {
		return localVarReturnValue, nil, reportError("regex is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("field", parameterToString(*r.field, ""))
	localVarQueryParams.Add("regex", parameterToString(*r.regex, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV1RegexFuncGetRequest struct {
	ctx        context.Context
	ApiService *RegexQueryApiService
	repo       *string
	rev        *string
	field      *string
	regex      *string
}

// repo
func (r ApiApiV1RegexFuncGetRequest) Repo(repo string) ApiApiV1RegexFuncGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1RegexFuncGetRequest) Rev(rev string) ApiApiV1RegexFuncGetRequest {
	r.rev = &rev
	return r
}

// field
func (r ApiApiV1RegexFuncGetRequest) Field(field string) ApiApiV1RegexFuncGetRequest {
	r.field = &field
	return r
}

// regex
func (r ApiApiV1RegexFuncGetRequest) Regex(regex string) ApiApiV1RegexFuncGetRequest {
	r.regex = &regex
	return r
}

func (r ApiApiV1RegexFuncGetRequest) Execute() ([]Sibyl2FunctionWithPath, *http.Response, error) {
	return r.ApiService.ApiV1RegexFuncGetExecute(r)
}

/*
ApiV1RegexFuncGet func query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1RegexFuncGetRequest
*/
func (a *RegexQueryApiService) ApiV1RegexFuncGet(ctx context.Context) ApiApiV1RegexFuncGetRequest {
	return ApiApiV1RegexFuncGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Sibyl2FunctionWithPath
func (a *RegexQueryApiService) ApiV1RegexFuncGetExecute(r ApiApiV1RegexFuncGetRequest) ([]Sibyl2FunctionWithPath, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Sibyl2FunctionWithPath
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegexQueryApiService.ApiV1RegexFuncGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/regex/func"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.field == nil {
		return localVarReturnValue, nil, reportError("field is required and must be specified")
	}
	if r.regex == nil {
		return localVarReturnValue, nil, reportError("regex is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("field", parameterToString(*r.field, ""))
	localVarQueryParams.Add("regex", parameterToString(*r.regex, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV1RegexFuncctxGetRequest struct {
	ctx        context.Context
	ApiService *RegexQueryApiService
	repo       *string
	rev        *string
	field      *string
	regex      *string
}

// repo
func (r ApiApiV1RegexFuncctxGetRequest) Repo(repo string) ApiApiV1RegexFuncctxGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1RegexFuncctxGetRequest) Rev(rev string) ApiApiV1RegexFuncctxGetRequest {
	r.rev = &rev
	return r
}

// field
func (r ApiApiV1RegexFuncctxGetRequest) Field(field string) ApiApiV1RegexFuncctxGetRequest {
	r.field = &field
	return r
}

// regex
func (r ApiApiV1RegexFuncctxGetRequest) Regex(regex string) ApiApiV1RegexFuncctxGetRequest {
	r.regex = &regex
	return r
}

func (r ApiApiV1RegexFuncctxGetRequest) Execute() ([]Sibyl2FunctionContextSlim, *http.Response, error) {
	return r.ApiService.ApiV1RegexFuncctxGetExecute(r)
}

/*
ApiV1RegexFuncctxGet func ctx query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1RegexFuncctxGetRequest
*/
func (a *RegexQueryApiService) ApiV1RegexFuncctxGet(ctx context.Context) ApiApiV1RegexFuncctxGetRequest {
	return ApiApiV1RegexFuncctxGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Sibyl2FunctionContextSlim
func (a *RegexQueryApiService) ApiV1RegexFuncctxGetExecute(r ApiApiV1RegexFuncctxGetRequest) ([]Sibyl2FunctionContextSlim, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Sibyl2FunctionContextSlim
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegexQueryApiService.ApiV1RegexFuncctxGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/regex/funcctx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.field == nil {
		return localVarReturnValue, nil, reportError("field is required and must be specified")
	}
	if r.regex == nil {
		return localVarReturnValue, nil, reportError("regex is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("field", parameterToString(*r.field, ""))
	localVarQueryParams.Add("regex", parameterToString(*r.regex, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}