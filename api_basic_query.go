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

// BasicQueryApiService BasicQueryApi service
type BasicQueryApiService service

type ApiApiV1ClazzGetRequest struct {
	ctx        context.Context
	ApiService *BasicQueryApiService
	repo       *string
	rev        *string
	file       *string
}

// repo
func (r ApiApiV1ClazzGetRequest) Repo(repo string) ApiApiV1ClazzGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1ClazzGetRequest) Rev(rev string) ApiApiV1ClazzGetRequest {
	r.rev = &rev
	return r
}

// file
func (r ApiApiV1ClazzGetRequest) File(file string) ApiApiV1ClazzGetRequest {
	r.file = &file
	return r
}

func (r ApiApiV1ClazzGetRequest) Execute() ([]ObjectClazzServiceDTO, *http.Response, error) {
	return r.ApiService.ApiV1ClazzGetExecute(r)
}

/*
ApiV1ClazzGet class query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1ClazzGetRequest
*/
func (a *BasicQueryApiService) ApiV1ClazzGet(ctx context.Context) ApiApiV1ClazzGetRequest {
	return ApiApiV1ClazzGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ObjectClazzServiceDTO
func (a *BasicQueryApiService) ApiV1ClazzGetExecute(r ApiApiV1ClazzGetRequest) ([]ObjectClazzServiceDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ObjectClazzServiceDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicQueryApiService.ApiV1ClazzGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clazz"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("file", parameterToString(*r.file, ""))
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

type ApiApiV1FuncGetRequest struct {
	ctx        context.Context
	ApiService *BasicQueryApiService
	repo       *string
	rev        *string
	file       *string
	lines      *string
}

// repo
func (r ApiApiV1FuncGetRequest) Repo(repo string) ApiApiV1FuncGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1FuncGetRequest) Rev(rev string) ApiApiV1FuncGetRequest {
	r.rev = &rev
	return r
}

// file
func (r ApiApiV1FuncGetRequest) File(file string) ApiApiV1FuncGetRequest {
	r.file = &file
	return r
}

// specific lines
func (r ApiApiV1FuncGetRequest) Lines(lines string) ApiApiV1FuncGetRequest {
	r.lines = &lines
	return r
}

func (r ApiApiV1FuncGetRequest) Execute() ([]ObjectFunctionServiceDTO, *http.Response, error) {
	return r.ApiService.ApiV1FuncGetExecute(r)
}

/*
ApiV1FuncGet func query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1FuncGetRequest
*/
func (a *BasicQueryApiService) ApiV1FuncGet(ctx context.Context) ApiApiV1FuncGetRequest {
	return ApiApiV1FuncGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ObjectFunctionServiceDTO
func (a *BasicQueryApiService) ApiV1FuncGetExecute(r ApiApiV1FuncGetRequest) ([]ObjectFunctionServiceDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ObjectFunctionServiceDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicQueryApiService.ApiV1FuncGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/func"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("file", parameterToString(*r.file, ""))
	if r.lines != nil {
		localVarQueryParams.Add("lines", parameterToString(*r.lines, ""))
	}
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

type ApiApiV1FuncctxGetRequest struct {
	ctx        context.Context
	ApiService *BasicQueryApiService
	repo       *string
	rev        *string
	file       *string
	lines      *string
}

// repo
func (r ApiApiV1FuncctxGetRequest) Repo(repo string) ApiApiV1FuncctxGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1FuncctxGetRequest) Rev(rev string) ApiApiV1FuncctxGetRequest {
	r.rev = &rev
	return r
}

// file
func (r ApiApiV1FuncctxGetRequest) File(file string) ApiApiV1FuncctxGetRequest {
	r.file = &file
	return r
}

// specific lines
func (r ApiApiV1FuncctxGetRequest) Lines(lines string) ApiApiV1FuncctxGetRequest {
	r.lines = &lines
	return r
}

func (r ApiApiV1FuncctxGetRequest) Execute() ([]ObjectFuncCtxServiceDTO, *http.Response, error) {
	return r.ApiService.ApiV1FuncctxGetExecute(r)
}

/*
ApiV1FuncctxGet func ctx query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1FuncctxGetRequest
*/
func (a *BasicQueryApiService) ApiV1FuncctxGet(ctx context.Context) ApiApiV1FuncctxGetRequest {
	return ApiApiV1FuncctxGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ObjectFuncCtxServiceDTO
func (a *BasicQueryApiService) ApiV1FuncctxGetExecute(r ApiApiV1FuncctxGetRequest) ([]ObjectFuncCtxServiceDTO, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ObjectFuncCtxServiceDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BasicQueryApiService.ApiV1FuncctxGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/funcctx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("file", parameterToString(*r.file, ""))
	if r.lines != nil {
		localVarQueryParams.Add("lines", parameterToString(*r.lines, ""))
	}
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
