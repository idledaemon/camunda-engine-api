# \TelemetryAPI

All URIs are relative to *http://localhost:8080/engine-rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureTelemetry**](TelemetryAPI.md#ConfigureTelemetry) | **Post** /telemetry/configuration | Configure Telemetry
[**GetTelemetryConfiguration**](TelemetryAPI.md#GetTelemetryConfiguration) | **Get** /telemetry/configuration | Fetch Telemetry Configuration
[**GetTelemetryData**](TelemetryAPI.md#GetTelemetryData) | **Get** /telemetry/data | Fetch Telemetry Data



## ConfigureTelemetry

> ConfigureTelemetry(ctx).TelemetryConfigurationDto(telemetryConfigurationDto).Execute()

Configure Telemetry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	telemetryConfigurationDto := *openapiclient.NewTelemetryConfigurationDto() // TelemetryConfigurationDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TelemetryAPI.ConfigureTelemetry(context.Background()).TelemetryConfigurationDto(telemetryConfigurationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelemetryAPI.ConfigureTelemetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureTelemetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryConfigurationDto** | [**TelemetryConfigurationDto**](TelemetryConfigurationDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryConfiguration

> TelemetryConfigurationDto GetTelemetryConfiguration(ctx).Execute()

Fetch Telemetry Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TelemetryAPI.GetTelemetryConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelemetryAPI.GetTelemetryConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryConfiguration`: TelemetryConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `TelemetryAPI.GetTelemetryConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryConfigurationRequest struct via the builder pattern


### Return type

[**TelemetryConfigurationDto**](TelemetryConfigurationDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelemetryData

> TelemetryDataDto GetTelemetryData(ctx).Execute()

Fetch Telemetry Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TelemetryAPI.GetTelemetryData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelemetryAPI.GetTelemetryData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelemetryData`: TelemetryDataDto
	fmt.Fprintf(os.Stdout, "Response from `TelemetryAPI.GetTelemetryData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryDataRequest struct via the builder pattern


### Return type

[**TelemetryDataDto**](TelemetryDataDto.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

