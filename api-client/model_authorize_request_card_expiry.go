/*
Payment Gateway API

Payment gateway API for merchants usage.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// AuthorizeRequestCardExpiry struct for AuthorizeRequestCardExpiry
type AuthorizeRequestCardExpiry struct {
	Month *int32 `json:"month,omitempty"`
	Year  *int32 `json:"year,omitempty"`
}

// NewAuthorizeRequestCardExpiry instantiates a new AuthorizeRequestCardExpiry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeRequestCardExpiry() *AuthorizeRequestCardExpiry {
	this := AuthorizeRequestCardExpiry{}
	return &this
}

// NewAuthorizeRequestCardExpiryWithDefaults instantiates a new AuthorizeRequestCardExpiry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeRequestCardExpiryWithDefaults() *AuthorizeRequestCardExpiry {
	this := AuthorizeRequestCardExpiry{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *AuthorizeRequestCardExpiry) GetMonth() int32 {
	if o == nil || o.Month == nil {
		var ret int32
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeRequestCardExpiry) GetMonthOk() (*int32, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *AuthorizeRequestCardExpiry) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int32 and assigns it to the Month field.
func (o *AuthorizeRequestCardExpiry) SetMonth(v int32) {
	o.Month = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *AuthorizeRequestCardExpiry) GetYear() int32 {
	if o == nil || o.Year == nil {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeRequestCardExpiry) GetYearOk() (*int32, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *AuthorizeRequestCardExpiry) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *AuthorizeRequestCardExpiry) SetYear(v int32) {
	o.Year = &v
}

func (o AuthorizeRequestCardExpiry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.Year != nil {
		toSerialize["year"] = o.Year
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizeRequestCardExpiry struct {
	value *AuthorizeRequestCardExpiry
	isSet bool
}

func (v NullableAuthorizeRequestCardExpiry) Get() *AuthorizeRequestCardExpiry {
	return v.value
}

func (v *NullableAuthorizeRequestCardExpiry) Set(val *AuthorizeRequestCardExpiry) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeRequestCardExpiry) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeRequestCardExpiry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeRequestCardExpiry(val *AuthorizeRequestCardExpiry) *NullableAuthorizeRequestCardExpiry {
	return &NullableAuthorizeRequestCardExpiry{value: val, isSet: true}
}

func (v NullableAuthorizeRequestCardExpiry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeRequestCardExpiry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
