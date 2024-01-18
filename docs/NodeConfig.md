# NodeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to [**ApiConfig**](ApiConfig.md) |  | [optional] 
**Server** | Pointer to [**ServerConfig**](ServerConfig.md) |  | [optional] 
**Database** | Pointer to [**DBConfig**](DBConfig.md) |  | [optional] 
**Redis** | Pointer to [**RedisConfig**](RedisConfig.md) |  | [optional] 

## Methods

### NewNodeConfig

`func NewNodeConfig() *NodeConfig`

NewNodeConfig instantiates a new NodeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeConfigWithDefaults

`func NewNodeConfigWithDefaults() *NodeConfig`

NewNodeConfigWithDefaults instantiates a new NodeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *NodeConfig) GetApi() ApiConfig`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *NodeConfig) GetApiOk() (*ApiConfig, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *NodeConfig) SetApi(v ApiConfig)`

SetApi sets Api field to given value.

### HasApi

`func (o *NodeConfig) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetServer

`func (o *NodeConfig) GetServer() ServerConfig`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *NodeConfig) GetServerOk() (*ServerConfig, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *NodeConfig) SetServer(v ServerConfig)`

SetServer sets Server field to given value.

### HasServer

`func (o *NodeConfig) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetDatabase

`func (o *NodeConfig) GetDatabase() DBConfig`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *NodeConfig) GetDatabaseOk() (*DBConfig, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *NodeConfig) SetDatabase(v DBConfig)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *NodeConfig) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetRedis

`func (o *NodeConfig) GetRedis() RedisConfig`

GetRedis returns the Redis field if non-nil, zero value otherwise.

### GetRedisOk

`func (o *NodeConfig) GetRedisOk() (*RedisConfig, bool)`

GetRedisOk returns a tuple with the Redis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis

`func (o *NodeConfig) SetRedis(v RedisConfig)`

SetRedis sets Redis field to given value.

### HasRedis

`func (o *NodeConfig) HasRedis() bool`

HasRedis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


