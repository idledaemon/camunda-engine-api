# TelemetryInternalsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to [**AbstractVendorVersionInformationDto**](AbstractVendorVersionInformationDto.md) |  | [optional] 
**ApplicationServer** | Pointer to [**AbstractVendorVersionInformationDto**](AbstractVendorVersionInformationDto.md) |  | [optional] 
**LicenseKey** | Pointer to [**TelemetryLicenseKeyDto**](TelemetryLicenseKeyDto.md) |  | [optional] 
**CamundaIntegration** | Pointer to **[]string** | List of Camunda integrations used (e.g., Camunda Spring Boot Starter, Camunda Run, WildFly/JBoss subsystem, Camunda EJB). | [optional] 
**Commands** | Pointer to [**map[string]TelemetryCountDto**](TelemetryCountDto.md) | The count of executed commands after the last retrieved data. | [optional] 
**Metrics** | Pointer to [**map[string]TelemetryCountDto**](TelemetryCountDto.md) | The collected metrics are the number of root process instance executions started, the number of activity instances started or also known as flow node instances, and the number of executed decision instances and elements. | [optional] 
**Webapps** | Pointer to **[]string** | The webapps enabled in this installation of Camunda. | [optional] 
**Jdk** | Pointer to [**AbstractVendorVersionInformationDto**](AbstractVendorVersionInformationDto.md) |  | [optional] 
**DataCollectionStartDate** | Pointer to **time.Time** | The date when the engine started to collect dynamic data, such as command executions and metrics. If telemetry sending is enabled, dynamic data resets on sending the data to Camunda. Dynamic data and the date returned by this method are reset in three cases: engine startup, after engine start when sending telemetry data to Camunda is enabled via API, after sending telemetry data to Camunda (only when this was enabled) The date is in the format &lt;code&gt;YYYY-MM-DD&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt;. | [optional] 

## Methods

### NewTelemetryInternalsDto

`func NewTelemetryInternalsDto() *TelemetryInternalsDto`

NewTelemetryInternalsDto instantiates a new TelemetryInternalsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryInternalsDtoWithDefaults

`func NewTelemetryInternalsDtoWithDefaults() *TelemetryInternalsDto`

NewTelemetryInternalsDtoWithDefaults instantiates a new TelemetryInternalsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *TelemetryInternalsDto) GetDatabase() AbstractVendorVersionInformationDto`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *TelemetryInternalsDto) GetDatabaseOk() (*AbstractVendorVersionInformationDto, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *TelemetryInternalsDto) SetDatabase(v AbstractVendorVersionInformationDto)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *TelemetryInternalsDto) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetApplicationServer

`func (o *TelemetryInternalsDto) GetApplicationServer() AbstractVendorVersionInformationDto`

GetApplicationServer returns the ApplicationServer field if non-nil, zero value otherwise.

### GetApplicationServerOk

`func (o *TelemetryInternalsDto) GetApplicationServerOk() (*AbstractVendorVersionInformationDto, bool)`

GetApplicationServerOk returns a tuple with the ApplicationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationServer

`func (o *TelemetryInternalsDto) SetApplicationServer(v AbstractVendorVersionInformationDto)`

SetApplicationServer sets ApplicationServer field to given value.

### HasApplicationServer

`func (o *TelemetryInternalsDto) HasApplicationServer() bool`

HasApplicationServer returns a boolean if a field has been set.

### GetLicenseKey

`func (o *TelemetryInternalsDto) GetLicenseKey() TelemetryLicenseKeyDto`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *TelemetryInternalsDto) GetLicenseKeyOk() (*TelemetryLicenseKeyDto, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *TelemetryInternalsDto) SetLicenseKey(v TelemetryLicenseKeyDto)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *TelemetryInternalsDto) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetCamundaIntegration

`func (o *TelemetryInternalsDto) GetCamundaIntegration() []string`

GetCamundaIntegration returns the CamundaIntegration field if non-nil, zero value otherwise.

### GetCamundaIntegrationOk

`func (o *TelemetryInternalsDto) GetCamundaIntegrationOk() (*[]string, bool)`

GetCamundaIntegrationOk returns a tuple with the CamundaIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCamundaIntegration

`func (o *TelemetryInternalsDto) SetCamundaIntegration(v []string)`

SetCamundaIntegration sets CamundaIntegration field to given value.

### HasCamundaIntegration

`func (o *TelemetryInternalsDto) HasCamundaIntegration() bool`

HasCamundaIntegration returns a boolean if a field has been set.

### SetCamundaIntegrationNil

`func (o *TelemetryInternalsDto) SetCamundaIntegrationNil(b bool)`

 SetCamundaIntegrationNil sets the value for CamundaIntegration to be an explicit nil

### UnsetCamundaIntegration
`func (o *TelemetryInternalsDto) UnsetCamundaIntegration()`

UnsetCamundaIntegration ensures that no value is present for CamundaIntegration, not even an explicit nil
### GetCommands

`func (o *TelemetryInternalsDto) GetCommands() map[string]TelemetryCountDto`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *TelemetryInternalsDto) GetCommandsOk() (*map[string]TelemetryCountDto, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *TelemetryInternalsDto) SetCommands(v map[string]TelemetryCountDto)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *TelemetryInternalsDto) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### SetCommandsNil

`func (o *TelemetryInternalsDto) SetCommandsNil(b bool)`

 SetCommandsNil sets the value for Commands to be an explicit nil

### UnsetCommands
`func (o *TelemetryInternalsDto) UnsetCommands()`

UnsetCommands ensures that no value is present for Commands, not even an explicit nil
### GetMetrics

`func (o *TelemetryInternalsDto) GetMetrics() map[string]TelemetryCountDto`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TelemetryInternalsDto) GetMetricsOk() (*map[string]TelemetryCountDto, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TelemetryInternalsDto) SetMetrics(v map[string]TelemetryCountDto)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *TelemetryInternalsDto) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetricsNil

`func (o *TelemetryInternalsDto) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *TelemetryInternalsDto) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil
### GetWebapps

`func (o *TelemetryInternalsDto) GetWebapps() []string`

GetWebapps returns the Webapps field if non-nil, zero value otherwise.

### GetWebappsOk

`func (o *TelemetryInternalsDto) GetWebappsOk() (*[]string, bool)`

GetWebappsOk returns a tuple with the Webapps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebapps

`func (o *TelemetryInternalsDto) SetWebapps(v []string)`

SetWebapps sets Webapps field to given value.

### HasWebapps

`func (o *TelemetryInternalsDto) HasWebapps() bool`

HasWebapps returns a boolean if a field has been set.

### SetWebappsNil

`func (o *TelemetryInternalsDto) SetWebappsNil(b bool)`

 SetWebappsNil sets the value for Webapps to be an explicit nil

### UnsetWebapps
`func (o *TelemetryInternalsDto) UnsetWebapps()`

UnsetWebapps ensures that no value is present for Webapps, not even an explicit nil
### GetJdk

`func (o *TelemetryInternalsDto) GetJdk() AbstractVendorVersionInformationDto`

GetJdk returns the Jdk field if non-nil, zero value otherwise.

### GetJdkOk

`func (o *TelemetryInternalsDto) GetJdkOk() (*AbstractVendorVersionInformationDto, bool)`

GetJdkOk returns a tuple with the Jdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdk

`func (o *TelemetryInternalsDto) SetJdk(v AbstractVendorVersionInformationDto)`

SetJdk sets Jdk field to given value.

### HasJdk

`func (o *TelemetryInternalsDto) HasJdk() bool`

HasJdk returns a boolean if a field has been set.

### GetDataCollectionStartDate

`func (o *TelemetryInternalsDto) GetDataCollectionStartDate() time.Time`

GetDataCollectionStartDate returns the DataCollectionStartDate field if non-nil, zero value otherwise.

### GetDataCollectionStartDateOk

`func (o *TelemetryInternalsDto) GetDataCollectionStartDateOk() (*time.Time, bool)`

GetDataCollectionStartDateOk returns a tuple with the DataCollectionStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectionStartDate

`func (o *TelemetryInternalsDto) SetDataCollectionStartDate(v time.Time)`

SetDataCollectionStartDate sets DataCollectionStartDate field to given value.

### HasDataCollectionStartDate

`func (o *TelemetryInternalsDto) HasDataCollectionStartDate() bool`

HasDataCollectionStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


