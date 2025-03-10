/*
OSHWA API

# Introduction  Welcome to the OSHWA Open Source Hardware Certification API. We hope that you will use it to build on top of the OSHWA certification program. This documentation provides information on how to use the OSHWA REST API to view, create, and search OSHWA certified open source hardware projects. You can learn more about OSHWA [here](https://www.oshwa.org/about/) and more about OSHWA’s free open source hardware certification program [here](https://certification.oshwa.org/).  This API supports both read and write functions. You can use the read functions to pull information about certified hardware from the [directory](https://certification.oshwa.org/list.html) in order to explore and present it in new ways. We encourage you to use it to find new ways to understand and visualize the world of open source hardware!  You can use the write functions to make it easier to submit registration applications to the program. Originally, the only way to submit hardware for registration was through OSHWA’s [application form](https://application.oshwa.org/apply). We hope that the write functionality will make it easier to integrate certification applications into existing workflows, and to connect platforms that already host open source hardware to the certification program.    If you have questions or comments about the API, please email us at info@oshwa.org.  We would also love to know how you use the API! We encourage you to contact us at info@oshwa.org, or let us know on twitter at @OHSummit.  # Tools This API is documented in **OpenAPI format**. It is built with [Swagger](http://swagger.io) and [ReDoc](https://github.com/Redocly/redoc).  # Using the API In order to use the API, you must register for an API key. You can get your own API key [here](https://certificationapi.oshwa.org/). If your token is not included or is invalid, the API will return an error. If you'd like to test the endpoints, check out our [Swagger](/endpoints) implementation, where you will be able to request a key, add the key to your requests, and explore OSHWA certified projects. You will also find code examples in this documentation.  # Pagination Project and Company results are returned in a wrapper object that contains total, limit, and offset values, which are useful for paginating over results.   ``` {   \"total\": 200, // Total number of matching items   \"offset\": 0, // Number of items skipped in request   \"limit\": 100, // Max number of items in request   \"items\": [ {...} ] // List of items } ```  The default `limit` for requests is 100, and the maximum `limit` is 1000. The default `offset` is 0. The above request returns the first 100 items. To get the next 100 items, the `offset` would be changed to 100. These parameters can be used to loop through the api and retrieve all items.  All items are returned in alphabetical order by `projectName`.   # Authentication  All OSHWA API endpoints require Bearer Token Authentication.  You can get an API key [here](https://certificationapi.oshwa.org/).  # Errors  The OSHWA API uses HTTP response status codes to indicate whether the response was successful. Detailed information about the response and how any errors might be resolved can be found in the body of the error message.  ``` {     \"error\": {         \"statusCode\": 422,         \"errorCode\": \"Unprocessable Entity\",         \"message\": \"Validation Error: Input validation failed\",         \"details\": [             {                 \"msg\": \"Responsible Party Type is required\",                 \"param\": \"responsiblePartyType\",                 \"location\": \"body\"             },         ]     } } ``` 

API version: 1.0.0
Contact: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiOptionsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiOptionsGet200Response{}

// ApiOptionsGet200Response struct for ApiOptionsGet200Response
type ApiOptionsGet200Response struct {
	CountryOptions []string `json:"countryOptions,omitempty"`
	ResponsiblePartyTypeOptions []string `json:"responsiblePartyTypeOptions,omitempty"`
	PrimaryTypeOptions []string `json:"primaryTypeOptions,omitempty"`
	AdditionalTypeOptions []string `json:"additionalTypeOptions,omitempty"`
	HardwareLicenseOptions []string `json:"hardwareLicenseOptions,omitempty"`
	SoftwareLicenseOptions []string `json:"softwareLicenseOptions,omitempty"`
	DocumentationLicenseOptions []string `json:"documentationLicenseOptions,omitempty"`
}

// NewApiOptionsGet200Response instantiates a new ApiOptionsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiOptionsGet200Response() *ApiOptionsGet200Response {
	this := ApiOptionsGet200Response{}
	return &this
}

// NewApiOptionsGet200ResponseWithDefaults instantiates a new ApiOptionsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiOptionsGet200ResponseWithDefaults() *ApiOptionsGet200Response {
	this := ApiOptionsGet200Response{}
	return &this
}

// GetCountryOptions returns the CountryOptions field value if set, zero value otherwise.
func (o *ApiOptionsGet200Response) GetCountryOptions() []string {
	if o == nil || IsNil(o.CountryOptions) {
		var ret []string
		return ret
	}
	return o.CountryOptions
}

// GetCountryOptionsOk returns a tuple with the CountryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptionsGet200Response) GetCountryOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryOptions) {
		return nil, false
	}
	return o.CountryOptions, true
}

// HasCountryOptions returns a boolean if a field has been set.
func (o *ApiOptionsGet200Response) HasCountryOptions() bool {
	if o != nil && !IsNil(o.CountryOptions) {
		return true
	}

	return false
}

// SetCountryOptions gets a reference to the given []string and assigns it to the CountryOptions field.
func (o *ApiOptionsGet200Response) SetCountryOptions(v []string) {
	o.CountryOptions = v
}

// GetResponsiblePartyTypeOptions returns the ResponsiblePartyTypeOptions field value if set, zero value otherwise.
func (o *ApiOptionsGet200Response) GetResponsiblePartyTypeOptions() []string {
	if o == nil || IsNil(o.ResponsiblePartyTypeOptions) {
		var ret []string
		return ret
	}
	return o.ResponsiblePartyTypeOptions
}

// GetResponsiblePartyTypeOptionsOk returns a tuple with the ResponsiblePartyTypeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptionsGet200Response) GetResponsiblePartyTypeOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponsiblePartyTypeOptions) {
		return nil, false
	}
	return o.ResponsiblePartyTypeOptions, true
}

// HasResponsiblePartyTypeOptions returns a boolean if a field has been set.
func (o *ApiOptionsGet200Response) HasResponsiblePartyTypeOptions() bool {
	if o != nil && !IsNil(o.ResponsiblePartyTypeOptions) {
		return true
	}

	return false
}

// SetResponsiblePartyTypeOptions gets a reference to the given []string and assigns it to the ResponsiblePartyTypeOptions field.
func (o *ApiOptionsGet200Response) SetResponsiblePartyTypeOptions(v []string) {
	o.ResponsiblePartyTypeOptions = v
}

// GetPrimaryTypeOptions returns the PrimaryTypeOptions field value if set, zero value otherwise.
func (o *ApiOptionsGet200Response) GetPrimaryTypeOptions() []string {
	if o == nil || IsNil(o.PrimaryTypeOptions) {
		var ret []string
		return ret
	}
	return o.PrimaryTypeOptions
}

// GetPrimaryTypeOptionsOk returns a tuple with the PrimaryTypeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptionsGet200Response) GetPrimaryTypeOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.PrimaryTypeOptions) {
		return nil, false
	}
	return o.PrimaryTypeOptions, true
}

// HasPrimaryTypeOptions returns a boolean if a field has been set.
func (o *ApiOptionsGet200Response) HasPrimaryTypeOptions() bool {
	if o != nil && !IsNil(o.PrimaryTypeOptions) {
		return true
	}

	return false
}

// SetPrimaryTypeOptions gets a reference to the given []string and assigns it to the PrimaryTypeOptions field.
func (o *ApiOptionsGet200Response) SetPrimaryTypeOptions(v []string) {
	o.PrimaryTypeOptions = v
}

// GetAdditionalTypeOptions returns the AdditionalTypeOptions field value if set, zero value otherwise.
func (o *ApiOptionsGet200Response) GetAdditionalTypeOptions() []string {
	if o == nil || IsNil(o.AdditionalTypeOptions) {
		var ret []string
		return ret
	}
	return o.AdditionalTypeOptions
}

// GetAdditionalTypeOptionsOk returns a tuple with the AdditionalTypeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptionsGet200Response) GetAdditionalTypeOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalTypeOptions) {
		return nil, false
	}
	return o.AdditionalTypeOptions, true
}

// HasAdditionalTypeOptions returns a boolean if a field has been set.
func (o *ApiOptionsGet200Response) HasAdditionalTypeOptions() bool {
	if o != nil && !IsNil(o.AdditionalTypeOptions) {
		return true
	}

	return false
}

// SetAdditionalTypeOptions gets a reference to the given []string and assigns it to the AdditionalTypeOptions field.
func (o *ApiOptionsGet200Response) SetAdditionalTypeOptions(v []string) {
	o.AdditionalTypeOptions = v
}

// GetHardwareLicenseOptions returns the HardwareLicenseOptions field value if set, zero value otherwise.
func (o *ApiOptionsGet200Response) GetHardwareLicenseOptions() []string {
	if o == nil || IsNil(o.HardwareLicenseOptions) {
		var ret []string
		return ret
	}
	return o.HardwareLicenseOptions
}

// GetHardwareLicenseOptionsOk returns a tuple with the HardwareLicenseOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptionsGet200Response) GetHardwareLicenseOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.HardwareLicenseOptions) {
		return nil, false
	}
	return o.HardwareLicenseOptions, true
}

// HasHardwareLicenseOptions returns a boolean if a field has been set.
func (o *ApiOptionsGet200Response) HasHardwareLicenseOptions() bool {
	if o != nil && !IsNil(o.HardwareLicenseOptions) {
		return true
	}

	return false
}

// SetHardwareLicenseOptions gets a reference to the given []string and assigns it to the HardwareLicenseOptions field.
func (o *ApiOptionsGet200Response) SetHardwareLicenseOptions(v []string) {
	o.HardwareLicenseOptions = v
}

// GetSoftwareLicenseOptions returns the SoftwareLicenseOptions field value if set, zero value otherwise.
func (o *ApiOptionsGet200Response) GetSoftwareLicenseOptions() []string {
	if o == nil || IsNil(o.SoftwareLicenseOptions) {
		var ret []string
		return ret
	}
	return o.SoftwareLicenseOptions
}

// GetSoftwareLicenseOptionsOk returns a tuple with the SoftwareLicenseOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptionsGet200Response) GetSoftwareLicenseOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.SoftwareLicenseOptions) {
		return nil, false
	}
	return o.SoftwareLicenseOptions, true
}

// HasSoftwareLicenseOptions returns a boolean if a field has been set.
func (o *ApiOptionsGet200Response) HasSoftwareLicenseOptions() bool {
	if o != nil && !IsNil(o.SoftwareLicenseOptions) {
		return true
	}

	return false
}

// SetSoftwareLicenseOptions gets a reference to the given []string and assigns it to the SoftwareLicenseOptions field.
func (o *ApiOptionsGet200Response) SetSoftwareLicenseOptions(v []string) {
	o.SoftwareLicenseOptions = v
}

// GetDocumentationLicenseOptions returns the DocumentationLicenseOptions field value if set, zero value otherwise.
func (o *ApiOptionsGet200Response) GetDocumentationLicenseOptions() []string {
	if o == nil || IsNil(o.DocumentationLicenseOptions) {
		var ret []string
		return ret
	}
	return o.DocumentationLicenseOptions
}

// GetDocumentationLicenseOptionsOk returns a tuple with the DocumentationLicenseOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiOptionsGet200Response) GetDocumentationLicenseOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.DocumentationLicenseOptions) {
		return nil, false
	}
	return o.DocumentationLicenseOptions, true
}

// HasDocumentationLicenseOptions returns a boolean if a field has been set.
func (o *ApiOptionsGet200Response) HasDocumentationLicenseOptions() bool {
	if o != nil && !IsNil(o.DocumentationLicenseOptions) {
		return true
	}

	return false
}

// SetDocumentationLicenseOptions gets a reference to the given []string and assigns it to the DocumentationLicenseOptions field.
func (o *ApiOptionsGet200Response) SetDocumentationLicenseOptions(v []string) {
	o.DocumentationLicenseOptions = v
}

func (o ApiOptionsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiOptionsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryOptions) {
		toSerialize["countryOptions"] = o.CountryOptions
	}
	if !IsNil(o.ResponsiblePartyTypeOptions) {
		toSerialize["responsiblePartyTypeOptions"] = o.ResponsiblePartyTypeOptions
	}
	if !IsNil(o.PrimaryTypeOptions) {
		toSerialize["primaryTypeOptions"] = o.PrimaryTypeOptions
	}
	if !IsNil(o.AdditionalTypeOptions) {
		toSerialize["additionalTypeOptions"] = o.AdditionalTypeOptions
	}
	if !IsNil(o.HardwareLicenseOptions) {
		toSerialize["hardwareLicenseOptions"] = o.HardwareLicenseOptions
	}
	if !IsNil(o.SoftwareLicenseOptions) {
		toSerialize["softwareLicenseOptions"] = o.SoftwareLicenseOptions
	}
	if !IsNil(o.DocumentationLicenseOptions) {
		toSerialize["documentationLicenseOptions"] = o.DocumentationLicenseOptions
	}
	return toSerialize, nil
}

type NullableApiOptionsGet200Response struct {
	value *ApiOptionsGet200Response
	isSet bool
}

func (v NullableApiOptionsGet200Response) Get() *ApiOptionsGet200Response {
	return v.value
}

func (v *NullableApiOptionsGet200Response) Set(val *ApiOptionsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiOptionsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiOptionsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiOptionsGet200Response(val *ApiOptionsGet200Response) *NullableApiOptionsGet200Response {
	return &NullableApiOptionsGet200Response{value: val, isSet: true}
}

func (v NullableApiOptionsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiOptionsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


