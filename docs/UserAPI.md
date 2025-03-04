# \UserAPI

All URIs are relative to *https://certificationapi.oshwa.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersSignupPost**](UserAPI.md#UsersSignupPost) | **Post** /users/signup | Register as user and receive an api token to authorize API requests.



## UsersSignupPost

> UsersSignupPost201Response UsersSignupPost(ctx).UsersSignupPostRequest(usersSignupPostRequest).Execute()

Register as user and receive an api token to authorize API requests.

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
	usersSignupPostRequest := *openapiclient.NewUsersSignupPostRequest("FirstName_example", "LastName_example", "Email_example") // UsersSignupPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UsersSignupPost(context.Background()).UsersSignupPostRequest(usersSignupPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UsersSignupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersSignupPost`: UsersSignupPost201Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UsersSignupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersSignupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersSignupPostRequest** | [**UsersSignupPostRequest**](UsersSignupPostRequest.md) |  | 

### Return type

[**UsersSignupPost201Response**](UsersSignupPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

