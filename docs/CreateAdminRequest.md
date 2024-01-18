# CreateAdminRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** |  | 
**Password** | **string** |  | 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 

## Methods

### NewCreateAdminRequest

`func NewCreateAdminRequest(account string, password string, ) *CreateAdminRequest`

NewCreateAdminRequest instantiates a new CreateAdminRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdminRequestWithDefaults

`func NewCreateAdminRequestWithDefaults() *CreateAdminRequest`

NewCreateAdminRequestWithDefaults instantiates a new CreateAdminRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateAdminRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateAdminRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateAdminRequest) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetPassword

`func (o *CreateAdminRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateAdminRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateAdminRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRole

`func (o *CreateAdminRequest) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateAdminRequest) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateAdminRequest) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateAdminRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


