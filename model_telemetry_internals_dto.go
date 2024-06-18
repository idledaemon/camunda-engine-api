/*
Camunda Platform REST API

OpenApi Spec for Camunda Platform REST API.

API version: 7.21.2-ee
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package camundarestgo

import (
	"encoding/json"
	"time"
)

// checks if the TelemetryInternalsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryInternalsDto{}

// TelemetryInternalsDto struct for TelemetryInternalsDto
type TelemetryInternalsDto struct {
	Database *AbstractVendorVersionInformationDto `json:"database,omitempty"`
	ApplicationServer *AbstractVendorVersionInformationDto `json:"application-server,omitempty"`
	LicenseKey *TelemetryLicenseKeyDto `json:"license-key,omitempty"`
	// List of Camunda integrations used (e.g., Camunda Spring Boot Starter, Camunda Run, WildFly/JBoss subsystem, Camunda EJB).
	CamundaIntegration []string `json:"camunda-integration,omitempty"`
	// The count of executed commands after the last retrieved data.
	Commands map[string]TelemetryCountDto `json:"commands,omitempty"`
	// The collected metrics are the number of root process instance executions started, the number of activity instances started or also known as flow node instances, and the number of executed decision instances and elements.
	Metrics map[string]TelemetryCountDto `json:"metrics,omitempty"`
	// The webapps enabled in this installation of Camunda.
	Webapps []string `json:"webapps,omitempty"`
	Jdk *AbstractVendorVersionInformationDto `json:"jdk,omitempty"`
	// The date when the engine started to collect dynamic data, such as command executions and metrics. If telemetry sending is enabled, dynamic data resets on sending the data to Camunda. Dynamic data and the date returned by this method are reset in three cases: engine startup, after engine start when sending telemetry data to Camunda is enabled via API, after sending telemetry data to Camunda (only when this was enabled) The date is in the format <code>YYYY-MM-DD'T'HH:mm:ss.SSSZ</code>.
	DataCollectionStartDate *time.Time `json:"data-collection-start-date,omitempty"`
}

// NewTelemetryInternalsDto instantiates a new TelemetryInternalsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryInternalsDto() *TelemetryInternalsDto {
	this := TelemetryInternalsDto{}
	return &this
}

// NewTelemetryInternalsDtoWithDefaults instantiates a new TelemetryInternalsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryInternalsDtoWithDefaults() *TelemetryInternalsDto {
	this := TelemetryInternalsDto{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *TelemetryInternalsDto) GetDatabase() AbstractVendorVersionInformationDto {
	if o == nil || IsNil(o.Database) {
		var ret AbstractVendorVersionInformationDto
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryInternalsDto) GetDatabaseOk() (*AbstractVendorVersionInformationDto, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given AbstractVendorVersionInformationDto and assigns it to the Database field.
func (o *TelemetryInternalsDto) SetDatabase(v AbstractVendorVersionInformationDto) {
	o.Database = &v
}

// GetApplicationServer returns the ApplicationServer field value if set, zero value otherwise.
func (o *TelemetryInternalsDto) GetApplicationServer() AbstractVendorVersionInformationDto {
	if o == nil || IsNil(o.ApplicationServer) {
		var ret AbstractVendorVersionInformationDto
		return ret
	}
	return *o.ApplicationServer
}

// GetApplicationServerOk returns a tuple with the ApplicationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryInternalsDto) GetApplicationServerOk() (*AbstractVendorVersionInformationDto, bool) {
	if o == nil || IsNil(o.ApplicationServer) {
		return nil, false
	}
	return o.ApplicationServer, true
}

// HasApplicationServer returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasApplicationServer() bool {
	if o != nil && !IsNil(o.ApplicationServer) {
		return true
	}

	return false
}

// SetApplicationServer gets a reference to the given AbstractVendorVersionInformationDto and assigns it to the ApplicationServer field.
func (o *TelemetryInternalsDto) SetApplicationServer(v AbstractVendorVersionInformationDto) {
	o.ApplicationServer = &v
}

// GetLicenseKey returns the LicenseKey field value if set, zero value otherwise.
func (o *TelemetryInternalsDto) GetLicenseKey() TelemetryLicenseKeyDto {
	if o == nil || IsNil(o.LicenseKey) {
		var ret TelemetryLicenseKeyDto
		return ret
	}
	return *o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryInternalsDto) GetLicenseKeyOk() (*TelemetryLicenseKeyDto, bool) {
	if o == nil || IsNil(o.LicenseKey) {
		return nil, false
	}
	return o.LicenseKey, true
}

// HasLicenseKey returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasLicenseKey() bool {
	if o != nil && !IsNil(o.LicenseKey) {
		return true
	}

	return false
}

// SetLicenseKey gets a reference to the given TelemetryLicenseKeyDto and assigns it to the LicenseKey field.
func (o *TelemetryInternalsDto) SetLicenseKey(v TelemetryLicenseKeyDto) {
	o.LicenseKey = &v
}

// GetCamundaIntegration returns the CamundaIntegration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TelemetryInternalsDto) GetCamundaIntegration() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CamundaIntegration
}

// GetCamundaIntegrationOk returns a tuple with the CamundaIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TelemetryInternalsDto) GetCamundaIntegrationOk() ([]string, bool) {
	if o == nil || IsNil(o.CamundaIntegration) {
		return nil, false
	}
	return o.CamundaIntegration, true
}

// HasCamundaIntegration returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasCamundaIntegration() bool {
	if o != nil && !IsNil(o.CamundaIntegration) {
		return true
	}

	return false
}

// SetCamundaIntegration gets a reference to the given []string and assigns it to the CamundaIntegration field.
func (o *TelemetryInternalsDto) SetCamundaIntegration(v []string) {
	o.CamundaIntegration = v
}

// GetCommands returns the Commands field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TelemetryInternalsDto) GetCommands() map[string]TelemetryCountDto {
	if o == nil {
		var ret map[string]TelemetryCountDto
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TelemetryInternalsDto) GetCommandsOk() (*map[string]TelemetryCountDto, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return &o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasCommands() bool {
	if o != nil && !IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given map[string]TelemetryCountDto and assigns it to the Commands field.
func (o *TelemetryInternalsDto) SetCommands(v map[string]TelemetryCountDto) {
	o.Commands = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TelemetryInternalsDto) GetMetrics() map[string]TelemetryCountDto {
	if o == nil {
		var ret map[string]TelemetryCountDto
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TelemetryInternalsDto) GetMetricsOk() (*map[string]TelemetryCountDto, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given map[string]TelemetryCountDto and assigns it to the Metrics field.
func (o *TelemetryInternalsDto) SetMetrics(v map[string]TelemetryCountDto) {
	o.Metrics = v
}

// GetWebapps returns the Webapps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TelemetryInternalsDto) GetWebapps() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Webapps
}

// GetWebappsOk returns a tuple with the Webapps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TelemetryInternalsDto) GetWebappsOk() ([]string, bool) {
	if o == nil || IsNil(o.Webapps) {
		return nil, false
	}
	return o.Webapps, true
}

// HasWebapps returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasWebapps() bool {
	if o != nil && !IsNil(o.Webapps) {
		return true
	}

	return false
}

// SetWebapps gets a reference to the given []string and assigns it to the Webapps field.
func (o *TelemetryInternalsDto) SetWebapps(v []string) {
	o.Webapps = v
}

// GetJdk returns the Jdk field value if set, zero value otherwise.
func (o *TelemetryInternalsDto) GetJdk() AbstractVendorVersionInformationDto {
	if o == nil || IsNil(o.Jdk) {
		var ret AbstractVendorVersionInformationDto
		return ret
	}
	return *o.Jdk
}

// GetJdkOk returns a tuple with the Jdk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryInternalsDto) GetJdkOk() (*AbstractVendorVersionInformationDto, bool) {
	if o == nil || IsNil(o.Jdk) {
		return nil, false
	}
	return o.Jdk, true
}

// HasJdk returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasJdk() bool {
	if o != nil && !IsNil(o.Jdk) {
		return true
	}

	return false
}

// SetJdk gets a reference to the given AbstractVendorVersionInformationDto and assigns it to the Jdk field.
func (o *TelemetryInternalsDto) SetJdk(v AbstractVendorVersionInformationDto) {
	o.Jdk = &v
}

// GetDataCollectionStartDate returns the DataCollectionStartDate field value if set, zero value otherwise.
func (o *TelemetryInternalsDto) GetDataCollectionStartDate() time.Time {
	if o == nil || IsNil(o.DataCollectionStartDate) {
		var ret time.Time
		return ret
	}
	return *o.DataCollectionStartDate
}

// GetDataCollectionStartDateOk returns a tuple with the DataCollectionStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryInternalsDto) GetDataCollectionStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DataCollectionStartDate) {
		return nil, false
	}
	return o.DataCollectionStartDate, true
}

// HasDataCollectionStartDate returns a boolean if a field has been set.
func (o *TelemetryInternalsDto) HasDataCollectionStartDate() bool {
	if o != nil && !IsNil(o.DataCollectionStartDate) {
		return true
	}

	return false
}

// SetDataCollectionStartDate gets a reference to the given time.Time and assigns it to the DataCollectionStartDate field.
func (o *TelemetryInternalsDto) SetDataCollectionStartDate(v time.Time) {
	o.DataCollectionStartDate = &v
}

func (o TelemetryInternalsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryInternalsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.ApplicationServer) {
		toSerialize["application-server"] = o.ApplicationServer
	}
	if !IsNil(o.LicenseKey) {
		toSerialize["license-key"] = o.LicenseKey
	}
	if o.CamundaIntegration != nil {
		toSerialize["camunda-integration"] = o.CamundaIntegration
	}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Webapps != nil {
		toSerialize["webapps"] = o.Webapps
	}
	if !IsNil(o.Jdk) {
		toSerialize["jdk"] = o.Jdk
	}
	if !IsNil(o.DataCollectionStartDate) {
		toSerialize["data-collection-start-date"] = o.DataCollectionStartDate
	}
	return toSerialize, nil
}

type NullableTelemetryInternalsDto struct {
	value *TelemetryInternalsDto
	isSet bool
}

func (v NullableTelemetryInternalsDto) Get() *TelemetryInternalsDto {
	return v.value
}

func (v *NullableTelemetryInternalsDto) Set(val *TelemetryInternalsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryInternalsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryInternalsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryInternalsDto(val *TelemetryInternalsDto) *NullableTelemetryInternalsDto {
	return &NullableTelemetryInternalsDto{value: val, isSet: true}
}

func (v NullableTelemetryInternalsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryInternalsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

