# MerchantList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Merchant**](Merchant.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewMerchantList

`func NewMerchantList(items []Merchant, pagination Pagination, ) *MerchantList`

NewMerchantList instantiates a new MerchantList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantListWithDefaults

`func NewMerchantListWithDefaults() *MerchantList`

NewMerchantListWithDefaults instantiates a new MerchantList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MerchantList) GetItems() []Merchant`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MerchantList) GetItemsOk() (*[]Merchant, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MerchantList) SetItems(v []Merchant)`

SetItems sets Items field to given value.


### GetPagination

`func (o *MerchantList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *MerchantList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *MerchantList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


