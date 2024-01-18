# DBConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Port** | **string** |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 
**Database** | [**DatabaseType**](DatabaseType.md) |  | 

## Methods

### NewDBConfig

`func NewDBConfig(host string, port string, username string, password string, database DatabaseType, ) *DBConfig`

NewDBConfig instantiates a new DBConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBConfigWithDefaults

`func NewDBConfigWithDefaults() *DBConfig`

NewDBConfigWithDefaults instantiates a new DBConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *DBConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DBConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DBConfig) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DBConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DBConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DBConfig) SetPort(v string)`

SetPort sets Port field to given value.


### GetUsername

`func (o *DBConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DBConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DBConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *DBConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DBConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DBConfig) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDatabase

`func (o *DBConfig) GetDatabase() DatabaseType`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DBConfig) GetDatabaseOk() (*DatabaseType, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DBConfig) SetDatabase(v DatabaseType)`

SetDatabase sets Database field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


