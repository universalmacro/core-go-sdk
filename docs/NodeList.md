# NodeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Node**](Node.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewNodeList

`func NewNodeList(items []Node, pagination Pagination, ) *NodeList`

NewNodeList instantiates a new NodeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeListWithDefaults

`func NewNodeListWithDefaults() *NodeList`

NewNodeListWithDefaults instantiates a new NodeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *NodeList) GetItems() []Node`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NodeList) GetItemsOk() (*[]Node, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NodeList) SetItems(v []Node)`

SetItems sets Items field to given value.


### GetPagination

`func (o *NodeList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NodeList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NodeList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


