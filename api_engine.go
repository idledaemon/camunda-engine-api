/*
Camunda Platform REST API

OpenApi Spec for Camunda Platform REST API.

API version: 7.21.2-ee
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package camundarestgo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type EngineAPI interface {

	/*
	GetProcessEngineNames Get List

	Retrieves the names of all process engines available on your platform.
**Note**: You cannot prepend `/engine/{name}` to this method.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetProcessEngineNamesRequest
	*/
	GetProcessEngineNames(ctx context.Context) ApiGetProcessEngineNamesRequest

	// GetProcessEngineNamesExecute executes the request
	//  @return []ProcessEngineDto
	GetProcessEngineNamesExecute(r ApiGetProcessEngineNamesRequest) ([]ProcessEngineDto, *http.Response, error)
}

// EngineAPIService EngineAPI service
type EngineAPIService service

type ApiGetProcessEngineNamesRequest struct {
	ctx context.Context
	ApiService EngineAPI
}

func (r ApiGetProcessEngineNamesRequest) Execute() ([]ProcessEngineDto, *http.Response, error) {
	return r.ApiService.GetProcessEngineNamesExecute(r)
}

/*
GetProcessEngineNames Get List

Retrieves the names of all process engines available on your platform.
**Note**: You cannot prepend `/engine/{name}` to this method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProcessEngineNamesRequest
*/
func (a *EngineAPIService) GetProcessEngineNames(ctx context.Context) ApiGetProcessEngineNamesRequest {
	return ApiGetProcessEngineNamesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ProcessEngineDto
func (a *EngineAPIService) GetProcessEngineNamesExecute(r ApiGetProcessEngineNamesRequest) ([]ProcessEngineDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ProcessEngineDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EngineAPIService.GetProcessEngineNames")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/engine"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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