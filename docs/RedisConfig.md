# RedisConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewRedisConfig

`func NewRedisConfig() *RedisConfig`

NewRedisConfig instantiates a new RedisConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedisConfigWithDefaults

`func NewRedisConfigWithDefaults() *RedisConfig`

NewRedisConfigWithDefaults instantiates a new RedisConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RedisConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RedisConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RedisConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RedisConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *RedisConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RedisConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RedisConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *RedisConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPassword

`func (o *RedisConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RedisConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RedisConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RedisConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


