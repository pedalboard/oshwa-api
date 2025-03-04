# \ProjectAPI

All URIs are relative to *https://certificationapi.oshwa.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProjectsGet**](ProjectAPI.md#ApiProjectsGet) | **Get** /api/projects | Get all OSHWA certified projects
[**ApiProjectsIdGet**](ProjectAPI.md#ApiProjectsIdGet) | **Get** /api/projects/{id} | Get an OSHWA certified project by ID.
[**ApiProjectsPost**](ProjectAPI.md#ApiProjectsPost) | **Post** /api/projects | Create a project for OSHWA certification



## ApiProjectsGet

> PublicProjects ApiProjectsGet(ctx).Q(q).Types(types).OshwaUid(oshwaUid).ProjectName(projectName).Country(country).PrimaryType(primaryType).AdditionalType(additionalType).ProjectKeywords(projectKeywords).HardwareLicense(hardwareLicense).SoftWareLicense(softWareLicense).DocumentationLicense(documentationLicense).Limit(limit).Offset(offset).Execute()

Get all OSHWA certified projects

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
	q := "q_example" // string | Full text search across project text fields (optional)
	types := "types_example" // string | Returns matches in primaryType, additionalType, or projectKeywords. Multiple types can be searched with comma-separated values, i.e. `arts,education,electronics` (optional)
	oshwaUid := "oshwaUid_example" // string | OSHWA UID (optional)
	projectName := "projectName_example" // string | Project Name (optional)
	country := "country_example" // string | Country (optional)
	primaryType := "primaryType_example" // string | Primary Type (optional)
	additionalType := []string{"Inner_example"} // []string | Additional Type (optional)
	projectKeywords := []string{"Inner_example"} // []string | Project Keywords (optional)
	hardwareLicense := "hardwareLicense_example" // string | Hardware License (optional)
	softWareLicense := "softWareLicense_example" // string | Software License (optional)
	documentationLicense := "documentationLicense_example" // string | Documentation License (optional)
	limit := int32(56) // int32 | The number of items to return. Default is 100, maximum is 1000 (optional)
	offset := int32(56) // int32 | The number of records to skip when retrieving results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ApiProjectsGet(context.Background()).Q(q).Types(types).OshwaUid(oshwaUid).ProjectName(projectName).Country(country).PrimaryType(primaryType).AdditionalType(additionalType).ProjectKeywords(projectKeywords).HardwareLicense(hardwareLicense).SoftWareLicense(softWareLicense).DocumentationLicense(documentationLicense).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ApiProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiProjectsGet`: PublicProjects
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ApiProjectsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Full text search across project text fields | 
 **types** | **string** | Returns matches in primaryType, additionalType, or projectKeywords. Multiple types can be searched with comma-separated values, i.e. &#x60;arts,education,electronics&#x60; | 
 **oshwaUid** | **string** | OSHWA UID | 
 **projectName** | **string** | Project Name | 
 **country** | **string** | Country | 
 **primaryType** | **string** | Primary Type | 
 **additionalType** | **[]string** | Additional Type | 
 **projectKeywords** | **[]string** | Project Keywords | 
 **hardwareLicense** | **string** | Hardware License | 
 **softWareLicense** | **string** | Software License | 
 **documentationLicense** | **string** | Documentation License | 
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


## ApiProjectsIdGet

> PublicProject ApiProjectsIdGet(ctx, id).Execute()

Get an OSHWA certified project by ID.

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ApiProjectsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ApiProjectsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiProjectsIdGet`: PublicProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ApiProjectsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicProject**](PublicProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProjectsPost

> ApiProjectsPost200Response ApiProjectsPost(ctx).Project(project).Execute()

Create a project for OSHWA certification

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
	project := *openapiclient.NewProject("ResponsiblePartyType_example", "ResponsibleParty_example", "BindingParty_example", "Country_example", "ProjectName_example", "PrimaryType_example", "HardwareLicense_example", "SoftwareLicense_example", "DocumentationLicense_example", "NoCommercialRestriction_example", "ExplanationNcr_example", false, "ExplanationNdr_example", false, "ExplanationOhwc_example", false, "ExplanationCcr_example", false, "ExplanationNur_example", false, "ExplanationRwr_example", false, "ExplanationNsp_example", false, "ExplanationNor_example", false, "ExplanationTn_example", *openapiclient.NewProjectCertificationMarkTerms(), "ExplanationCertificationTerms_example", "Relationship_example", false, "ParentName_example") // Project | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ApiProjectsPost(context.Background()).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ApiProjectsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiProjectsPost`: ApiProjectsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ApiProjectsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) |  | 

### Return type

[**ApiProjectsPost200Response**](ApiProjectsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

