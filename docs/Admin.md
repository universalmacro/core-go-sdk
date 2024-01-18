# Admin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Account** | **string** |  | 
**PhoneNumber** | Pointer to [**PhoneNumber**](PhoneNumber.md) |  | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewAdmin

`func NewAdmin(id string, account string, ) *Admin`

NewAdmin instantiates a new Admin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminWithDefaults

`func NewAdminWithDefaults() *Admin`

NewAdminWithDefaults instantiates a new Admin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Admin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Admin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Admin) SetId(v string)`

SetId sets Id field to given value.


### GetAccount

`func (o *Admin) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Admin) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Admin) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetPhoneNumber

`func (o *Admin) GetPhoneNumber() PhoneNumber`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Admin) GetPhoneNumberOk() (*PhoneNumber, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Admin) SetPhoneNumber(v PhoneNumber)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Admin) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRole

`func (o *Admin) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Admin) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Admin) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *Admin) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Admin) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Admin) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Admin) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Admin) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


