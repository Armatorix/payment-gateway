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

// AuthorizeResponse struct for AuthorizeResponse
type AuthorizeResponse struct {
	AuthorizationId *string `json:"authorization_id,omitempty"`
	// Amount of money in a form of least significant monetary unit.
	AmountAvailable *int32  `json:"amount_available,omitempty"`
	Currency        *string `json:"currency,omitempty"`
	Status          *string `json:"status,omitempty"`
}

// NewAuthorizeResponse instantiates a new AuthorizeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeResponse() *AuthorizeResponse {
	this := AuthorizeResponse{}
	return &this
}

// NewAuthorizeResponseWithDefaults instantiates a new AuthorizeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeResponseWithDefaults() *AuthorizeResponse {
	this := AuthorizeResponse{}
	return &this
}

// GetAuthorizationId returns the AuthorizationId field value if set, zero value otherwise.
func (o *AuthorizeResponse) GetAuthorizationId() string {
	if o == nil || o.AuthorizationId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationId
}

// GetAuthorizationIdOk returns a tuple with the AuthorizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeResponse) GetAuthorizationIdOk() (*string, bool) {
	if o == nil || o.AuthorizationId == nil {
		return nil, false
	}
	return o.AuthorizationId, true
}

// HasAuthorizationId returns a boolean if a field has been set.
func (o *AuthorizeResponse) HasAuthorizationId() bool {
	if o != nil && o.AuthorizationId != nil {
		return true
	}

	return false
}

// SetAuthorizationId gets a reference to the given string and assigns it to the AuthorizationId field.
func (o *AuthorizeResponse) SetAuthorizationId(v string) {
	o.AuthorizationId = &v
}

// GetAmountAvailable returns the AmountAvailable field value if set, zero value otherwise.
func (o *AuthorizeResponse) GetAmountAvailable() int32 {
	if o == nil || o.AmountAvailable == nil {
		var ret int32
		return ret
	}
	return *o.AmountAvailable
}

// GetAmountAvailableOk returns a tuple with the AmountAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeResponse) GetAmountAvailableOk() (*int32, bool) {
	if o == nil || o.AmountAvailable == nil {
		return nil, false
	}
	return o.AmountAvailable, true
}

// HasAmountAvailable returns a boolean if a field has been set.
func (o *AuthorizeResponse) HasAmountAvailable() bool {
	if o != nil && o.AmountAvailable != nil {
		return true
	}

	return false
}

// SetAmountAvailable gets a reference to the given int32 and assigns it to the AmountAvailable field.
func (o *AuthorizeResponse) SetAmountAvailable(v int32) {
	o.AmountAvailable = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AuthorizeResponse) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AuthorizeResponse) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AuthorizeResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthorizeResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthorizeResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthorizeResponse) SetStatus(v string) {
	o.Status = &v
}

func (o AuthorizeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationId != nil {
		toSerialize["authorization_id"] = o.AuthorizationId
	}
	if o.AmountAvailable != nil {
		toSerialize["amount_available"] = o.AmountAvailable
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizeResponse struct {
	value *AuthorizeResponse
	isSet bool
}

func (v NullableAuthorizeResponse) Get() *AuthorizeResponse {
	return v.value
}

func (v *NullableAuthorizeResponse) Set(val *AuthorizeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeResponse(val *AuthorizeResponse) *NullableAuthorizeResponse {
	return &NullableAuthorizeResponse{value: val, isSet: true}
}

func (v NullableAuthorizeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}