# ServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** |  | [optional] 
**JwtSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewServerConfig

`func NewServerConfig() *ServerConfig`

NewServerConfig instantiates a new ServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigWithDefaults

`func NewServerConfigWithDefaults() *ServerConfig`

NewServerConfigWithDefaults instantiates a new ServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ServerConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServerConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServerConfig) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServerConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetJwtSecret

`func (o *ServerConfig) GetJwtSecret() string`

GetJwtSecret returns the JwtSecret field if non-nil, zero value otherwise.

### GetJwtSecretOk

`func (o *ServerConfig) GetJwtSecretOk() (*string, bool)`

GetJwtSecretOk returns a tuple with the JwtSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSecret

`func (o *ServerConfig) SetJwtSecret(v string)`

SetJwtSecret sets JwtSecret field to given value.

### HasJwtSecret

`func (o *ServerConfig) HasJwtSecret() bool`

HasJwtSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


