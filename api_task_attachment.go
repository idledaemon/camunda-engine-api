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
	"os"
)


type TaskAttachmentAPI interface {

	/*
	AddAttachment Create

	Creates an attachment for a task.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the task to add the attachment to.
	@return ApiAddAttachmentRequest
	*/
	AddAttachment(ctx context.Context, id string) ApiAddAttachmentRequest

	// AddAttachmentExecute executes the request
	//  @return AttachmentDto
	AddAttachmentExecute(r ApiAddAttachmentRequest) (*AttachmentDto, *http.Response, error)

	/*
	DeleteAttachment Delete

	Removes an attachment from a task by id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the task.
	@param attachmentId The id of the attachment to be removed.
	@return ApiDeleteAttachmentRequest
	*/
	DeleteAttachment(ctx context.Context, id string, attachmentId string) ApiDeleteAttachmentRequest

	// DeleteAttachmentExecute executes the request
	DeleteAttachmentExecute(r ApiDeleteAttachmentRequest) (*http.Response, error)

	/*
	GetAttachment Get

	Retrieves a task attachment by task id and attachment id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the task.
	@param attachmentId The id of the attachment to be retrieved.
	@return ApiGetAttachmentRequest
	*/
	GetAttachment(ctx context.Context, id string, attachmentId string) ApiGetAttachmentRequest

	// GetAttachmentExecute executes the request
	//  @return AttachmentDto
	GetAttachmentExecute(r ApiGetAttachmentRequest) (*AttachmentDto, *http.Response, error)

	/*
	GetAttachmentData Get (Binary)

	Retrieves the binary content of a task attachment by task id and attachment id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the task.
	@param attachmentId The id of the attachment to be retrieved.
	@return ApiGetAttachmentDataRequest
	*/
	GetAttachmentData(ctx context.Context, id string, attachmentId string) ApiGetAttachmentDataRequest

	// GetAttachmentDataExecute executes the request
	//  @return *os.File
	GetAttachmentDataExecute(r ApiGetAttachmentDataRequest) (*os.File, *http.Response, error)

	/*
	GetAttachments Get List

	Gets the attachments for a task.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the task to retrieve the attachments for.
	@return ApiGetAttachmentsRequest
	*/
	GetAttachments(ctx context.Context, id string) ApiGetAttachmentsRequest

	// GetAttachmentsExecute executes the request
	//  @return []AttachmentDto
	GetAttachmentsExecute(r ApiGetAttachmentsRequest) ([]AttachmentDto, *http.Response, error)
}

// TaskAttachmentAPIService TaskAttachmentAPI service
type TaskAttachmentAPIService service

type ApiAddAttachmentRequest struct {
	ctx context.Context
	ApiService TaskAttachmentAPI
	id string
	attachmentName *string
	attachmentDescription *string
	attachmentType *string
	url *string
	content *os.File
}

// The name of the attachment.
func (r ApiAddAttachmentRequest) AttachmentName(attachmentName string) ApiAddAttachmentRequest {
	r.attachmentName = &attachmentName
	return r
}

// The description of the attachment.
func (r ApiAddAttachmentRequest) AttachmentDescription(attachmentDescription string) ApiAddAttachmentRequest {
	r.attachmentDescription = &attachmentDescription
	return r
}

// The type of the attachment.
func (r ApiAddAttachmentRequest) AttachmentType(attachmentType string) ApiAddAttachmentRequest {
	r.attachmentType = &attachmentType
	return r
}

// The url to the remote content of the attachment.
func (r ApiAddAttachmentRequest) Url(url string) ApiAddAttachmentRequest {
	r.url = &url
	return r
}

// The content of the attachment.
func (r ApiAddAttachmentRequest) Content(content *os.File) ApiAddAttachmentRequest {
	r.content = content
	return r
}

func (r ApiAddAttachmentRequest) Execute() (*AttachmentDto, *http.Response, error) {
	return r.ApiService.AddAttachmentExecute(r)
}

/*
AddAttachment Create

Creates an attachment for a task.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the task to add the attachment to.
 @return ApiAddAttachmentRequest
*/
func (a *TaskAttachmentAPIService) AddAttachment(ctx context.Context, id string) ApiAddAttachmentRequest {
	return ApiAddAttachmentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AttachmentDto
func (a *TaskAttachmentAPIService) AddAttachmentExecute(r ApiAddAttachmentRequest) (*AttachmentDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AttachmentDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAttachmentAPIService.AddAttachment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{id}/attachment/create"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.attachmentName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment-name", r.attachmentName, "")
	}
	if r.attachmentDescription != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment-description", r.attachmentDescription, "")
	}
	if r.attachmentType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment-type", r.attachmentType, "")
	}
	if r.url != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "url", r.url, "")
	}
	var contentLocalVarFormFileName string
	var contentLocalVarFileName     string
	var contentLocalVarFileBytes    []byte

	contentLocalVarFormFileName = "content"
	contentLocalVarFile := r.content

	if contentLocalVarFile != nil {
		fbs, _ := io.ReadAll(contentLocalVarFile)

		contentLocalVarFileBytes = fbs
		contentLocalVarFileName = contentLocalVarFile.Name()
		contentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: contentLocalVarFileBytes, fileName: contentLocalVarFileName, formFileName: contentLocalVarFormFileName})
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationExceptionDto
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

type ApiDeleteAttachmentRequest struct {
	ctx context.Context
	ApiService TaskAttachmentAPI
	id string
	attachmentId string
}

func (r ApiDeleteAttachmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAttachmentExecute(r)
}

/*
DeleteAttachment Delete

Removes an attachment from a task by id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the task.
 @param attachmentId The id of the attachment to be removed.
 @return ApiDeleteAttachmentRequest
*/
func (a *TaskAttachmentAPIService) DeleteAttachment(ctx context.Context, id string, attachmentId string) ApiDeleteAttachmentRequest {
	return ApiDeleteAttachmentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		attachmentId: attachmentId,
	}
}

// Execute executes the request
func (a *TaskAttachmentAPIService) DeleteAttachmentExecute(r ApiDeleteAttachmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAttachmentAPIService.DeleteAttachment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{id}/attachment/{attachmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attachmentId"+"}", url.PathEscape(parameterValueToString(r.attachmentId, "attachmentId")), -1)

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
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationExceptionDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetAttachmentRequest struct {
	ctx context.Context
	ApiService TaskAttachmentAPI
	id string
	attachmentId string
}

func (r ApiGetAttachmentRequest) Execute() (*AttachmentDto, *http.Response, error) {
	return r.ApiService.GetAttachmentExecute(r)
}

/*
GetAttachment Get

Retrieves a task attachment by task id and attachment id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the task.
 @param attachmentId The id of the attachment to be retrieved.
 @return ApiGetAttachmentRequest
*/
func (a *TaskAttachmentAPIService) GetAttachment(ctx context.Context, id string, attachmentId string) ApiGetAttachmentRequest {
	return ApiGetAttachmentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		attachmentId: attachmentId,
	}
}

// Execute executes the request
//  @return AttachmentDto
func (a *TaskAttachmentAPIService) GetAttachmentExecute(r ApiGetAttachmentRequest) (*AttachmentDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AttachmentDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAttachmentAPIService.GetAttachment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{id}/attachment/{attachmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attachmentId"+"}", url.PathEscape(parameterValueToString(r.attachmentId, "attachmentId")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetAttachmentDataRequest struct {
	ctx context.Context
	ApiService TaskAttachmentAPI
	id string
	attachmentId string
}

func (r ApiGetAttachmentDataRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetAttachmentDataExecute(r)
}

/*
GetAttachmentData Get (Binary)

Retrieves the binary content of a task attachment by task id and attachment id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the task.
 @param attachmentId The id of the attachment to be retrieved.
 @return ApiGetAttachmentDataRequest
*/
func (a *TaskAttachmentAPIService) GetAttachmentData(ctx context.Context, id string, attachmentId string) ApiGetAttachmentDataRequest {
	return ApiGetAttachmentDataRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		attachmentId: attachmentId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *TaskAttachmentAPIService) GetAttachmentDataExecute(r ApiGetAttachmentDataRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAttachmentAPIService.GetAttachmentData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{id}/attachment/{attachmentId}/data"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attachmentId"+"}", url.PathEscape(parameterValueToString(r.attachmentId, "attachmentId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "text/plain", "application/json"}

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetAttachmentsRequest struct {
	ctx context.Context
	ApiService TaskAttachmentAPI
	id string
}

func (r ApiGetAttachmentsRequest) Execute() ([]AttachmentDto, *http.Response, error) {
	return r.ApiService.GetAttachmentsExecute(r)
}

/*
GetAttachments Get List

Gets the attachments for a task.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the task to retrieve the attachments for.
 @return ApiGetAttachmentsRequest
*/
func (a *TaskAttachmentAPIService) GetAttachments(ctx context.Context, id string) ApiGetAttachmentsRequest {
	return ApiGetAttachmentsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []AttachmentDto
func (a *TaskAttachmentAPIService) GetAttachmentsExecute(r ApiGetAttachmentsRequest) ([]AttachmentDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AttachmentDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAttachmentAPIService.GetAttachments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/task/{id}/attachment"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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
