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
	"strings"
)


type TaskIdentityLinkAPI interface {

	/*
	AddIdentityLink Add

	Adds an identity link to a task by id. Can be used to link any user or group to a task
and specify a relation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the task to add a link to.
	@return ApiAddIdentityLinkRequest
	*/
	AddIdentityLink(ctx context.Context, id string) ApiAddIdentityLinkRequest

	// AddIdentityLinkExecute executes the request
	AddIdentityLinkExecute(r ApiAddIdentityLinkRequest) (*http.Response, error)

	/*
	DeleteIdentityLink Delete

	Removes an identity link from a task by id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the task to remove a link from.
	@return ApiDeleteIdentityLinkRequest
	*/
	DeleteIdentityLink(ctx context.Context, id string) ApiDeleteIdentityLinkRequest

	// DeleteIdentityLinkExecute executes the request
	DeleteIdentityLinkExecute(r ApiDeleteIdentityLinkRequest) (*http.Response, error)

	/*
	GetIdentityLinks Get List

	Gets the identity links for a task by id, which are the users and groups that are in
*some* relation to it (including assignee and owner).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the task to retrieve the identity links for.
	@return ApiGetIdentityLinksRequest
	*/
	GetIdentityLinks(ctx context.Context, id string) ApiGetIdentityLinksRequest

	// GetIdentityLinksExecute executes the request
	//  @return []IdentityLinkDto
	GetIdentityLinksExecute(r ApiGetIdentityLinksRequest) ([]IdentityLinkDto, *http.Response, error)
}

// TaskIdentityLinkAPIService TaskIdentityLinkAPI service
type TaskIdentityLinkAPIService service

type ApiAddIdentityLinkRequest struct {
	ctx context.Context
	ApiService TaskIdentityLinkAPI
	id string
	identityLinkDto *IdentityLinkDto
}

func (r ApiAddIdentityLinkRequest) IdentityLinkDto(identityLinkDto IdentityLinkDto) ApiAddIdentityLinkRequest {
	r.identityLinkDto = &identityLinkDto
	return r
}

func (r ApiAddIdentityLinkRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddIdentityLinkExecute(r)
}

/*
AddIdentityLink Add

Adds an identity link to a task by id. Can be used to link any user or group to a task
and specify a relation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the task to add a link to.
 @return ApiAddIdentityLinkRequest
*/
func (a *TaskIdentityLinkAPIService) AddIdentityLink(ctx context.Context, id string) ApiAddIdentityLinkRequest {
	return ApiAddIdentityLinkRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TaskIdentityLinkAPIService) AddIdentityLinkExecute(r ApiAddIdentityLinkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskIdentityLinkAPIService.AddIdentityLink")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{id}/identity-links"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.identityLinkDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExceptionDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteIdentityLinkRequest struct {
	ctx context.Context
	ApiService TaskIdentityLinkAPI
	id string
	identityLinkDto *IdentityLinkDto
}

func (r ApiDeleteIdentityLinkRequest) IdentityLinkDto(identityLinkDto IdentityLinkDto) ApiDeleteIdentityLinkRequest {
	r.identityLinkDto = &identityLinkDto
	return r
}

func (r ApiDeleteIdentityLinkRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteIdentityLinkExecute(r)
}

/*
DeleteIdentityLink Delete

Removes an identity link from a task by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the task to remove a link from.
 @return ApiDeleteIdentityLinkRequest
*/
func (a *TaskIdentityLinkAPIService) DeleteIdentityLink(ctx context.Context, id string) ApiDeleteIdentityLinkRequest {
	return ApiDeleteIdentityLinkRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TaskIdentityLinkAPIService) DeleteIdentityLinkExecute(r ApiDeleteIdentityLinkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskIdentityLinkAPIService.DeleteIdentityLink")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{id}/identity-links/delete"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.identityLinkDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExceptionDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetIdentityLinksRequest struct {
	ctx context.Context
	ApiService TaskIdentityLinkAPI
	id string
	type_ *string
}

// Filter by the type of links to include.
func (r ApiGetIdentityLinksRequest) Type_(type_ string) ApiGetIdentityLinksRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetIdentityLinksRequest) Execute() ([]IdentityLinkDto, *http.Response, error) {
	return r.ApiService.GetIdentityLinksExecute(r)
}

/*
GetIdentityLinks Get List

Gets the identity links for a task by id, which are the users and groups that are in
*some* relation to it (including assignee and owner).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the task to retrieve the identity links for.
 @return ApiGetIdentityLinksRequest
*/
func (a *TaskIdentityLinkAPIService) GetIdentityLinks(ctx context.Context, id string) ApiGetIdentityLinksRequest {
	return ApiGetIdentityLinksRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []IdentityLinkDto
func (a *TaskIdentityLinkAPIService) GetIdentityLinksExecute(r ApiGetIdentityLinksRequest) ([]IdentityLinkDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []IdentityLinkDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskIdentityLinkAPIService.GetIdentityLinks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{id}/identity-links"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExceptionDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
