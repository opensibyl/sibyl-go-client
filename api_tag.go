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

// TagApiService TagApi service
type TagApiService service

type ApiApiV1TagFuncGetRequest struct {
	ctx        context.Context
	ApiService *TagApiService
	repo       *string
	rev        *string
	tag        *string
}

// repo
func (r ApiApiV1TagFuncGetRequest) Repo(repo string) ApiApiV1TagFuncGetRequest {
	r.repo = &repo
	return r
}

// rev
func (r ApiApiV1TagFuncGetRequest) Rev(rev string) ApiApiV1TagFuncGetRequest {
	r.rev = &rev
	return r
}

// tag
func (r ApiApiV1TagFuncGetRequest) Tag(tag string) ApiApiV1TagFuncGetRequest {
	r.tag = &tag
	return r
}

func (r ApiApiV1TagFuncGetRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ApiV1TagFuncGetExecute(r)
}

/*
ApiV1TagFuncGet query func by tag

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1TagFuncGetRequest
*/
func (a *TagApiService) ApiV1TagFuncGet(ctx context.Context) ApiApiV1TagFuncGetRequest {
	return ApiApiV1TagFuncGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []string
func (a *TagApiService) ApiV1TagFuncGetExecute(r ApiApiV1TagFuncGetRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagApiService.ApiV1TagFuncGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tag/func"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repo == nil {
		return localVarReturnValue, nil, reportError("repo is required and must be specified")
	}
	if r.rev == nil {
		return localVarReturnValue, nil, reportError("rev is required and must be specified")
	}
	if r.tag == nil {
		return localVarReturnValue, nil, reportError("tag is required and must be specified")
	}

	localVarQueryParams.Add("repo", parameterToString(*r.repo, ""))
	localVarQueryParams.Add("rev", parameterToString(*r.rev, ""))
	localVarQueryParams.Add("tag", parameterToString(*r.tag, ""))
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

type ApiApiV1TagFuncPostRequest struct {
	ctx        context.Context
	ApiService *TagApiService
	payload    *ServiceTagUpload
}

// tag upload payload
func (r ApiApiV1TagFuncPostRequest) Payload(payload ServiceTagUpload) ApiApiV1TagFuncPostRequest {
	r.payload = &payload
	return r
}

func (r ApiApiV1TagFuncPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1TagFuncPostExecute(r)
}

/*
ApiV1TagFuncPost create func tag

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1TagFuncPostRequest
*/
func (a *TagApiService) ApiV1TagFuncPost(ctx context.Context) ApiApiV1TagFuncPostRequest {
	return ApiApiV1TagFuncPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *TagApiService) ApiV1TagFuncPostExecute(r ApiApiV1TagFuncPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagApiService.ApiV1TagFuncPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tag/func"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.payload == nil {
		return nil, reportError("payload is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.payload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
