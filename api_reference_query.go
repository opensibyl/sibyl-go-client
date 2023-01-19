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

// ReferenceQueryApiService ReferenceQueryApi service
type ReferenceQueryApiService service

type ApiApiV1ReferenceCountFuncctxGetRequest struct {
	ctx        context.Context
	ApiService *ReferenceQueryApiService
	repo       *string
	rev        *string
	moreThan   *int32
	lessThan   *int32
}

// repo
func (r ApiApiV1ReferenceCountFuncctxGetRequest) Repo(repo string) ApiApiV1ReferenceCountFuncctxGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1ReferenceCountFuncctxGetRequest) Rev(rev string) ApiApiV1ReferenceCountFuncctxGetRequest {
	r.rev = &rev
	return r
}

// moreThan
func (r ApiApiV1ReferenceCountFuncctxGetRequest) MoreThan(moreThan int32) ApiApiV1ReferenceCountFuncctxGetRequest {
	r.moreThan = &moreThan
	return r
}

// lessThan
func (r ApiApiV1ReferenceCountFuncctxGetRequest) LessThan(lessThan int32) ApiApiV1ReferenceCountFuncctxGetRequest {
	r.lessThan = &lessThan
	return r
}

func (r ApiApiV1ReferenceCountFuncctxGetRequest) Execute() ([]Sibyl2FunctionContextSlim, *http.Response, error) {
	return r.ApiService.ApiV1ReferenceCountFuncctxGetExecute(r)
}

/*
ApiV1ReferenceCountFuncctxGet funcctx query by ref

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1ReferenceCountFuncctxGetRequest
*/
func (a *ReferenceQueryApiService) ApiV1ReferenceCountFuncctxGet(ctx context.Context) ApiApiV1ReferenceCountFuncctxGetRequest {
	return ApiApiV1ReferenceCountFuncctxGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Sibyl2FunctionContextSlim
func (a *ReferenceQueryApiService) ApiV1ReferenceCountFuncctxGetExecute(r ApiApiV1ReferenceCountFuncctxGetRequest) ([]Sibyl2FunctionContextSlim, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Sibyl2FunctionContextSlim
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceQueryApiService.ApiV1ReferenceCountFuncctxGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/reference/count/funcctx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.moreThan == nil {
		return localVarReturnValue, nil, reportError("moreThan is required and must be specified")
	}
	if r.lessThan == nil {
		return localVarReturnValue, nil, reportError("lessThan is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("moreThan", parameterToString(*r.moreThan, ""))
	localVarQueryParams.Add("lessThan", parameterToString(*r.lessThan, ""))
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

type ApiApiV1ReferenceCountFuncctxReverseGetRequest struct {
	ctx        context.Context
	ApiService *ReferenceQueryApiService
	repo       *string
	rev        *string
	moreThan   *int32
	lessThan   *int32
}

// repo
func (r ApiApiV1ReferenceCountFuncctxReverseGetRequest) Repo(repo string) ApiApiV1ReferenceCountFuncctxReverseGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1ReferenceCountFuncctxReverseGetRequest) Rev(rev string) ApiApiV1ReferenceCountFuncctxReverseGetRequest {
	r.rev = &rev
	return r
}

// moreThan
func (r ApiApiV1ReferenceCountFuncctxReverseGetRequest) MoreThan(moreThan int32) ApiApiV1ReferenceCountFuncctxReverseGetRequest {
	r.moreThan = &moreThan
	return r
}

// lessThan
func (r ApiApiV1ReferenceCountFuncctxReverseGetRequest) LessThan(lessThan int32) ApiApiV1ReferenceCountFuncctxReverseGetRequest {
	r.lessThan = &lessThan
	return r
}

func (r ApiApiV1ReferenceCountFuncctxReverseGetRequest) Execute() ([]Sibyl2FunctionContextSlim, *http.Response, error) {
	return r.ApiService.ApiV1ReferenceCountFuncctxReverseGetExecute(r)
}

/*
ApiV1ReferenceCountFuncctxReverseGet funcctx query by referenced

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1ReferenceCountFuncctxReverseGetRequest
*/
func (a *ReferenceQueryApiService) ApiV1ReferenceCountFuncctxReverseGet(ctx context.Context) ApiApiV1ReferenceCountFuncctxReverseGetRequest {
	return ApiApiV1ReferenceCountFuncctxReverseGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []Sibyl2FunctionContextSlim
func (a *ReferenceQueryApiService) ApiV1ReferenceCountFuncctxReverseGetExecute(r ApiApiV1ReferenceCountFuncctxReverseGetRequest) ([]Sibyl2FunctionContextSlim, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []Sibyl2FunctionContextSlim
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReferenceQueryApiService.ApiV1ReferenceCountFuncctxReverseGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/reference/count/funcctx/reverse"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.moreThan == nil {
		return localVarReturnValue, nil, reportError("moreThan is required and must be specified")
	}
	if r.lessThan == nil {
		return localVarReturnValue, nil, reportError("lessThan is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("moreThan", parameterToString(*r.moreThan, ""))
	localVarQueryParams.Add("lessThan", parameterToString(*r.lessThan, ""))
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
