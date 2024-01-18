# \MerchantAPI

All URIs are relative to *https://uat.api.universalmacro.com/core*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMerchantToNode**](MerchantAPI.md#AddMerchantToNode) | **Post** /nodes/{id}/merchants | 
[**ListNodeMerchants**](MerchantAPI.md#ListNodeMerchants) | **Get** /nodes/{id}/merchants | 



## AddMerchantToNode

> Merchant AddMerchantToNode(ctx, id).CreateMerchantRequest(createMerchantRequest).Execute()





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
	createMerchantRequest := *openapiclient.NewCreateMerchantRequest("Account_example", "Password_example") // CreateMerchantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAPI.AddMerchantToNode(context.Background(), id).CreateMerchantRequest(createMerchantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAPI.AddMerchantToNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMerchantToNode`: Merchant
	fmt.Fprintf(os.Stdout, "Response from `MerchantAPI.AddMerchantToNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMerchantToNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMerchantRequest** | [**CreateMerchantRequest**](CreateMerchantRequest.md) |  | 

### Return type

[**Merchant**](Merchant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodeMerchants

> MerchantList ListNodeMerchants(ctx).Index(index).Limit(limit).Execute()





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
	index := int32(56) // int32 | Current page index (optional) (default to 0)
	limit := int32(56) // int32 | Total pages (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAPI.ListNodeMerchants(context.Background()).Index(index).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAPI.ListNodeMerchants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodeMerchants`: MerchantList
	fmt.Fprintf(os.Stdout, "Response from `MerchantAPI.ListNodeMerchants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNodeMerchantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **int32** | Current page index | [default to 0]
 **limit** | **int32** | Total pages | [default to 1]

### Return type

[**MerchantList**](MerchantList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

