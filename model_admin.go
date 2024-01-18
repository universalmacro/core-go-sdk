/*
Core APIs

universalmacro

API version: 0.0.2
Contact: chenyunda218@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Admin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Admin{}

// Admin struct for Admin
type Admin struct {
	Id string `json:"id"`
	Account string `json:"account"`
	PhoneNumber *PhoneNumber `json:"phoneNumber,omitempty"`
	Role *Role `json:"role,omitempty"`
	CreatedAt *int64 `json:"createdAt,omitempty"`
}

type _Admin Admin

// NewAdmin instantiates a new Admin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdmin(id string, account string) *Admin {
	this := Admin{}
	this.Id = id
	this.Account = account
	return &this
}

// NewAdminWithDefaults instantiates a new Admin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminWithDefaults() *Admin {
	this := Admin{}
	return &this
}

// GetId returns the Id field value
func (o *Admin) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Admin) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Admin) SetId(v string) {
	o.Id = v
}

// GetAccount returns the Account field value
func (o *Admin) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *Admin) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *Admin) SetAccount(v string) {
	o.Account = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Admin) GetPhoneNumber() PhoneNumber {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret PhoneNumber
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Admin) GetPhoneNumberOk() (*PhoneNumber, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Admin) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given PhoneNumber and assigns it to the PhoneNumber field.
func (o *Admin) SetPhoneNumber(v PhoneNumber) {
	o.PhoneNumber = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Admin) GetRole() Role {
	if o == nil || IsNil(o.Role) {
		var ret Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Admin) GetRoleOk() (*Role, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Admin) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given Role and assigns it to the Role field.
func (o *Admin) SetRole(v Role) {
	o.Role = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Admin) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Admin) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Admin) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Admin) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

func (o Admin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Admin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["account"] = o.Account
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

func (o *Admin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"account",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAdmin := _Admin{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdmin)

	if err != nil {
		return err
	}

	*o = Admin(varAdmin)

	return err
}

type NullableAdmin struct {
	value *Admin
	isSet bool
}

func (v NullableAdmin) Get() *Admin {
	return v.value
}

func (v *NullableAdmin) Set(val *Admin) {
	v.value = val
	v.isSet = true
}

func (v NullableAdmin) IsSet() bool {
	return v.isSet
}

func (v *NullableAdmin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdmin(val *Admin) *NullableAdmin {
	return &NullableAdmin{value: val, isSet: true}
}

func (v NullableAdmin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdmin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

