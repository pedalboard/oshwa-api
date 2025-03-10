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
	"bytes"
	"fmt"
)

// checks if the UsersSignupPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersSignupPostRequest{}

// UsersSignupPostRequest struct for UsersSignupPostRequest
type UsersSignupPostRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}

type _UsersSignupPostRequest UsersSignupPostRequest

// NewUsersSignupPostRequest instantiates a new UsersSignupPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersSignupPostRequest(firstName string, lastName string, email string) *UsersSignupPostRequest {
	this := UsersSignupPostRequest{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	return &this
}

// NewUsersSignupPostRequestWithDefaults instantiates a new UsersSignupPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersSignupPostRequestWithDefaults() *UsersSignupPostRequest {
	this := UsersSignupPostRequest{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *UsersSignupPostRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UsersSignupPostRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UsersSignupPostRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UsersSignupPostRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UsersSignupPostRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UsersSignupPostRequest) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *UsersSignupPostRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UsersSignupPostRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UsersSignupPostRequest) SetEmail(v string) {
	o.Email = v
}

func (o UsersSignupPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersSignupPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *UsersSignupPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"firstName",
		"lastName",
		"email",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUsersSignupPostRequest := _UsersSignupPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsersSignupPostRequest)

	if err != nil {
		return err
	}

	*o = UsersSignupPostRequest(varUsersSignupPostRequest)

	return err
}

type NullableUsersSignupPostRequest struct {
	value *UsersSignupPostRequest
	isSet bool
}

func (v NullableUsersSignupPostRequest) Get() *UsersSignupPostRequest {
	return v.value
}

func (v *NullableUsersSignupPostRequest) Set(val *UsersSignupPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersSignupPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersSignupPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersSignupPostRequest(val *UsersSignupPostRequest) *NullableUsersSignupPostRequest {
	return &NullableUsersSignupPostRequest{value: val, isSet: true}
}

func (v NullableUsersSignupPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersSignupPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


