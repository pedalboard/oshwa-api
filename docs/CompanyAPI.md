# \CompanyAPI

All URIs are relative to *https://certificationapi.oshwa.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCompaniesNameGet**](CompanyAPI.md#ApiCompaniesNameGet) | **Get** /api/companies/{name} | Search by company name



## ApiCompaniesNameGet

> PublicProjects ApiCompaniesNameGet(ctx, name).Limit(limit).Offset(offset).Execute()

Search by company name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pedalboard/oshwa-api"
)

func main() {
	name := "name_example" // string | 
	limit := int32(56) // int32 | The number of items to return. Default is 100, maximum is 1000 (optional)
	offset := int32(56) // int32 | The number of records to skip when retrieving results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyAPI.ApiCompaniesNameGet(context.Background(), name).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyAPI.ApiCompaniesNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiCompaniesNameGet`: PublicProjects
	fmt.Fprintf(os.Stdout, "Response from `CompanyAPI.ApiCompaniesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiCompaniesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of items to return. Default is 100, maximum is 1000 | 
 **offset** | **int32** | The number of records to skip when retrieving results | 

### Return type

[**PublicProjects**](PublicProjects.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

