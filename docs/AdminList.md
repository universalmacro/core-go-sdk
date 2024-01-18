# AdminList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Admin**](Admin.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewAdminList

`func NewAdminList(items []Admin, pagination Pagination, ) *AdminList`

NewAdminList instantiates a new AdminList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminListWithDefaults

`func NewAdminListWithDefaults() *AdminList`

NewAdminListWithDefaults instantiates a new AdminList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AdminList) GetItems() []Admin`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AdminList) GetItemsOk() (*[]Admin, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AdminList) SetItems(v []Admin)`

SetItems sets Items field to given value.


### GetPagination

`func (o *AdminList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AdminList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AdminList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


