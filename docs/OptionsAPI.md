# \OptionsAPI

All URIs are relative to *https://certificationapi.oshwa.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOptionsGet**](OptionsAPI.md#ApiOptionsGet) | **Get** /api/options | Get a list of valid project values for country, responsibleParty, primaryType, additionalType, hardwareLicense,softwareLicense, and documentationLicense fields.



## ApiOptionsGet

> ApiOptionsGet200Response ApiOptionsGet(ctx).Execute()

Get a list of valid project values for country, responsibleParty, primaryType, additionalType, hardwareLicense,softwareLicense, and documentationLicense fields.

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
	resp, r, err := apiClient.OptionsAPI.ApiOptionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.ApiOptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiOptionsGet`: ApiOptionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.ApiOptionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiOptionsGetRequest struct via the builder pattern


### Return type

[**ApiOptionsGet200Response**](ApiOptionsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

